package mule

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	taskdto "github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/relay/channel"
	"github.com/QuantumNous/new-api/relay/channel/task/taskcommon"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/relaykit/dto"
	"github.com/QuantumNous/new-api/service"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// ============================
// Adaptor implementation
// ============================

type TaskAdaptor struct {
	taskcommon.BaseBilling
	ChannelType int
	apiKey      string
	baseURL     string
}

func (a *TaskAdaptor) Init(info *relaycommon.RelayInfo) {
	a.ChannelType = info.ChannelType
	a.baseURL = info.ChannelBaseUrl
	a.apiKey = info.ApiKey
}

func (a *TaskAdaptor) ValidateRequestAndSetAction(c *gin.Context, info *relaycommon.RelayInfo) (taskErr *taskdto.TaskError) {
	return relaycommon.ValidateBasicTaskRequest(c, info, constant.TaskActionGenerate)
}

func (a *TaskAdaptor) BuildRequestURL(info *relaycommon.RelayInfo) (string, error) {
	return fmt.Sprintf("%s/vendors/carrothub/v1/%s/generation", a.baseURL, info.OriginModelName), nil
}

// BuildRequestHeader sets required headers for Ali API
func (a *TaskAdaptor) BuildRequestHeader(c *gin.Context, req *http.Request, info *relaycommon.RelayInfo) error {
	req.Header.Set("Authorization", "Bearer "+a.apiKey)
	req.Header.Set("Content-Type", "application/json")
	return nil
}

func (a *TaskAdaptor) BuildRequestBody(c *gin.Context, info *relaycommon.RelayInfo) (io.Reader, error) {
	taskReq, err := relaycommon.GetTaskRequest(c)
	if err != nil {
		return nil, err
	}
	req, err := a.convertToAliRequest(info, taskReq)
	if err != nil {
		return nil, err
	}
	logger.LogJson(c, "ali video request body", req)

	bodyBytes, err := common.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "marshal_ali_request_failed")
	}
	return bytes.NewReader(bodyBytes), nil
}

func (a *TaskAdaptor) convertToAliRequest(info *relaycommon.RelayInfo, req relaycommon.TaskSubmitReq) (*MuleRequest, error) {
	var muleReq MuleRequest
	err := taskcommon.UnmarshalMetadata(req.Metadata, &muleReq)
	if err != nil {
		return nil, err
	}
	if muleReq.Duration == 0 || muleReq.Duration == -1 {
		muleReq.Duration = 5
	}
	if muleReq.Resolution == "" {
		muleReq.Resolution = "1080p"
	}
	if muleReq.Ratio == "" {
		muleReq.Ratio = "adaptive"
	}
	return &muleReq, err
}

// EstimateBilling 根据用户请求参数计算 OtherRatios（时长、分辨率等）。
// 在 ValidateRequestAndSetAction 之后、价格计算之前调用。
func (a *TaskAdaptor) EstimateBilling(c *gin.Context, info *relaycommon.RelayInfo) map[string]float64 {
	taskReq, err := relaycommon.GetTaskRequest(c)
	if err != nil {
		return nil
	}
	var muleReq MuleRequest
	_ = taskcommon.UnmarshalMetadata(taskReq.Metadata, &muleReq)

	duration := muleReq.Duration
	if duration == 0 || duration == -1 {
		duration = 5
	}

	resolution := strings.ToLower(muleReq.Resolution)
	if resolution == "" {
		resolution = "1080p"
	}

	otherRatios := map[string]float64{
		"seconds": float64(min(duration, relaycommon.MaxTaskDurationSeconds)),
	}

	if resolutionRatios, ok := muleResolutionPriceRatio[info.OriginModelName]; ok {
		if ratio, exists := resolutionRatios[resolution]; exists {
			otherRatios["resolution"] = ratio
		}
	}
	return otherRatios
}

func muleValidationError(code, message string) *taskdto.TaskError {
	return service.TaskErrorWrapperLocal(errors.New(message), code, http.StatusBadRequest)
}

// DoRequest delegates to common helper
func (a *TaskAdaptor) DoRequest(c *gin.Context, info *relaycommon.RelayInfo, requestBody io.Reader) (*http.Response, error) {
	return channel.DoTaskApiRequest(a, c, info, requestBody)
}

// DoResponse handles upstream response
func (a *TaskAdaptor) DoResponse(c *gin.Context, resp *http.Response, info *relaycommon.RelayInfo) (taskID string, taskData []byte, taskErr *taskdto.TaskError) {
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		taskErr = service.TaskErrorWrapper(err, "read_response_body_failed", http.StatusInternalServerError)
		return
	}
	_ = resp.Body.Close()
	// 解析阿里响应
	var aliResp MuleResponse
	if err = common.Unmarshal(responseBody, &aliResp); err != nil {
		taskErr = service.TaskErrorWrapper(errors.Wrapf(err, "body: %s", responseBody), "unmarshal_response_body_failed", http.StatusInternalServerError)
		return
	}
	// 检查错误
	if aliResp.Status != 0 {
		taskErr = service.TaskErrorWrapper(fmt.Errorf("%s", aliResp.Detail), "ali_api_error", resp.StatusCode)
		return
	}
	taskId := aliResp.TaskInfo.Id
	// 返回 OpenAI 格式
	aliResp.TaskInfo.Id = info.PublicTaskID
	c.JSON(http.StatusOK, aliResp)
	return taskId, responseBody, nil
}

// FetchTask 查询任务状态
func (a *TaskAdaptor) FetchTask(baseUrl, key string, body map[string]any, proxy string) (*http.Response, error) {
	taskID, ok := body["task_id"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid task_id")
	}
	m, ok := body["model"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid model")
	}
	uri := fmt.Sprintf("%s/vendors/carrothub/v1/%s/generation/%s", baseUrl, m, taskID)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+key)

	client, err := service.GetHttpClientWithProxy(proxy)
	if err != nil {
		return nil, fmt.Errorf("new proxy http client failed: %w", err)
	}
	return client.Do(req)
}

func (a *TaskAdaptor) GetModelList() []string {
	return []string{}
}

func (a *TaskAdaptor) GetChannelName() string {
	return ChannelName
}

// ParseTaskResult 解析任务结果
func (a *TaskAdaptor) ParseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	var resp MuleResponse
	if err := common.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}
	if resp.Status != 0 {
		return nil, fmt.Errorf("%s", resp.Detail)
	}

	taskResult := relaycommon.TaskInfo{
		Code: 0,
	}

	// 状态映射
	switch resp.TaskInfo.Status {
	case "pending":
		taskResult.Status = model.TaskStatusQueued
	case "processing":
		taskResult.Status = model.TaskStatusInProgress
	case "completed":
		taskResult.Status = model.TaskStatusSuccess
		if len(resp.Videos) > 0 {
			taskResult.Url = resp.Videos[0]
		}
	case "failed":
		taskResult.Status = model.TaskStatusFailure
		taskResult.Reason = fmt.Sprintf("%s", resp.Detail)
	default:
		taskResult.Status = model.TaskStatusQueued
	}

	return &taskResult, nil
}

func (a *TaskAdaptor) ConvertToOpenAIVideo(task *model.Task) ([]byte, error) {
	var aliResp MuleResponse
	if err := common.Unmarshal(task.Data, &aliResp); err != nil {
		return nil, errors.Wrap(err, "unmarshal ali response failed")
	}

	openAIResp := dto.NewOpenAIVideo()
	openAIResp.ID = task.TaskID
	openAIResp.Status = string(task.Status)
	openAIResp.Model = task.Properties.OriginModelName
	openAIResp.SetProgressStr(task.Progress)
	openAIResp.CreatedAt = task.CreatedAt
	openAIResp.CompletedAt = task.UpdatedAt

	// 设置视频URL（核心字段）
	openAIResp.SetMetadata("url", task.GetResultURL())

	// 错误处理
	if aliResp.Status != 0 {
		openAIResp.Error = &dto.OpenAIVideoError{
			Code:    strconv.Itoa(aliResp.Status),
			Message: aliResp.Detail,
		}
	} else if aliResp.TaskInfo.Error.Code != 0 {
		openAIResp.Error = &dto.OpenAIVideoError{
			Code:    strconv.Itoa(aliResp.TaskInfo.Error.Code),
			Message: aliResp.TaskInfo.Error.Detail,
		}
	}

	return common.Marshal(openAIResp)
}

func (a *TaskAdaptor) ConvertToMuleVideo(originTask *model.Task) ([]byte, error) {
	var mr MuleResponse
	if err := json.Unmarshal(originTask.Data, &mr); err != nil {
		return nil, err
	}
	mr.TaskInfo.Id = originTask.TaskID
	if mr.TaskInfo.Status == "" {
		mr.TaskInfo.Status = "pending"
	}
	if mr.TaskInfo.CreatedAt.IsZero() {
		mr.TaskInfo.CreatedAt = time.Unix(originTask.CreatedAt, 0)
		mr.TaskInfo.UpdatedAt = mr.TaskInfo.CreatedAt
	}
	return json.Marshal(mr)
}
