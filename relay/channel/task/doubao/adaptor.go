package doubao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/QuantumNous/new-api/common"
	"github.com/samber/lo"

	"github.com/QuantumNous/new-api/constant"
	taskdto "github.com/QuantumNous/new-api/dto"
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
	// 暂时用于同类型模型的不同渠道选择
	organization string
}

func (a *TaskAdaptor) Init(info *relaycommon.RelayInfo) {
	a.organization = info.Organization
	a.ChannelType = info.ChannelType
	a.baseURL = info.ChannelBaseUrl
	a.apiKey = info.ApiKey
}

// ValidateRequestAndSetAction parses body, validates fields and sets default action.
func (a *TaskAdaptor) ValidateRequestAndSetAction(c *gin.Context, info *relaycommon.RelayInfo) (taskErr *taskdto.TaskError) {
	// Accept only POST /v1/video/generations as "generate" action.
	return relaycommon.ValidateBasicTaskRequest(c, info, constant.TaskActionGenerate)
}

// BuildRequestURL constructs the upstream URL.
func (a *TaskAdaptor) BuildRequestURL(_ *relaycommon.RelayInfo) (string, error) {
	switch a.organization {
	case ThirdAnyFast:
		return fmt.Sprintf("%s/v1/video/generations", a.baseURL), nil
	case ThirdKWJM:
		return fmt.Sprintf("%s/v3/contents/generations/tasks", a.baseURL), nil
	case ThirdTokenMart:
		return fmt.Sprintf("%s/v1/video/generate", a.baseURL), nil
	default:
		return fmt.Sprintf("%s/api/v3/contents/generations/tasks", a.baseURL), nil
	}
}

// BuildRequestHeader sets required headers.
func (a *TaskAdaptor) BuildRequestHeader(_ *gin.Context, req *http.Request, _ *relaycommon.RelayInfo) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.apiKey)
	return nil
}

// EstimateBilling 根据请求 metadata 中的输出分辨率与是否包含视频输入，返回相对基准价的计费 OtherRatio。
func (a *TaskAdaptor) EstimateBilling(c *gin.Context, info *relaycommon.RelayInfo) map[string]float64 {
	req, err := relaycommon.GetTaskRequest(c)
	if err != nil {
		return nil
	}
	ratio, ok := GetVideoInputRatio(info.OriginModelName, req.Metadata)
	if !ok || ratio == 1.0 {
		return nil
	}
	return map[string]float64{"video_input": ratio}
}

// BuildRequestBody converts request into Doubao specific format.
func (a *TaskAdaptor) BuildRequestBody(c *gin.Context, info *relaycommon.RelayInfo) (io.Reader, error) {
	req, err := relaycommon.GetTaskRequest(c)
	if err != nil {
		return nil, err
	}

	body, err := a.convertToRequestPayload(&req)
	if err != nil {
		return nil, errors.Wrap(err, "convert request payload failed")
	}
	for i := range body.Content {
		if image := body.Content[i].ImageURL; image != nil {
			body.Content[i].ImageURL.URL = resolveAssetURL(info.ChannelId, image.URL)
		}
		if video := body.Content[i].VideoURL; video != nil {
			body.Content[i].VideoURL.URL = resolveAssetURL(info.ChannelId, video.URL)
		}
		if audio := body.Content[i].AudioURL; audio != nil {
			body.Content[i].AudioURL.URL = resolveAssetURL(info.ChannelId, audio.URL)
		}
	}
	if info.IsModelMapped {
		body.Model = info.UpstreamModelName
	} else {
		info.UpstreamModelName = body.Model
	}
	data, err := common.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

// DoRequest delegates to common helper.
func (a *TaskAdaptor) DoRequest(c *gin.Context, info *relaycommon.RelayInfo, requestBody io.Reader) (*http.Response, error) {
	return channel.DoTaskApiRequest(a, c, info, requestBody)
}

// DoResponse handles upstream response, returns taskID etc.
func (a *TaskAdaptor) DoResponse(c *gin.Context, resp *http.Response, info *relaycommon.RelayInfo) (taskID string, taskData []byte, taskErr *taskdto.TaskError) {
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		taskErr = service.TaskErrorWrapper(err, "read_response_body_failed", http.StatusInternalServerError)
		return
	}
	_ = resp.Body.Close()
	fn, ok := submitResultMapping[a.organization]
	if !ok {
		fn = submitResultMapping[Official]
	}
	upstreamTaskId, taskErr := fn(responseBody)
	if taskErr != nil {
		return
	}
	c.JSON(http.StatusOK, map[string]any{"id": info.PublicTaskID})
	return upstreamTaskId, responseBody, nil
}

// FetchTask fetch task status
func (a *TaskAdaptor) FetchTask(baseUrl, key string, body map[string]any, proxy string) (*http.Response, error) {
	taskID, ok := body["task_id"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid task_id")
	}
	var api string
	switch a.organization {
	case ThirdAnyFast:
		api = fmt.Sprintf("%s/v1/video/generations/%s", baseUrl, taskID)
	case ThirdKWJM:
		api = fmt.Sprintf("%s/v3/contents/generations/tasks/%s", baseUrl, taskID)
	case ThirdTokenMart:
		api = fmt.Sprintf("%s/v1/video/tasks/%s", baseUrl, taskID)
	default:
		api = fmt.Sprintf("%s/api/v3/contents/generations/tasks/%s", baseUrl, taskID)
	}

	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	client, err := service.GetHttpClientWithProxy(proxy)
	if err != nil {
		return nil, fmt.Errorf("new proxy http client failed: %w", err)
	}
	return client.Do(req)
}

func (a *TaskAdaptor) GetModelList() []string {
	return ModelList
}

func (a *TaskAdaptor) GetChannelName() string {
	return ChannelName
}

func (a *TaskAdaptor) convertToRequestPayload(req *relaycommon.TaskSubmitReq) (*requestPayload, error) {
	r := requestPayload{
		Model:   req.Model,
		Content: []ContentItem{},
	}

	// Add images if present
	if req.HasImage() {
		for _, imgURL := range req.Images {
			r.Content = append(r.Content, ContentItem{
				Type: "image_url",
				ImageURL: &MediaURL{
					URL: imgURL,
				},
			})
		}
	}

	if sec, _ := strconv.Atoi(req.Seconds); sec > 0 {
		r.Duration = lo.ToPtr(dto.IntValue(sec))
	}

	r.Content = lo.Reject(r.Content, func(c ContentItem, _ int) bool { return c.Type == "text" })
	r.Content = append(r.Content, ContentItem{
		Type: "text",
		Text: req.Prompt,
	})

	metadata := req.Metadata
	if err := taskcommon.UnmarshalMetadata(metadata, &r); err != nil {
		return nil, errors.Wrap(err, "unmarshal metadata failed")
	}
	return &r, nil
}

func (a *TaskAdaptor) ParseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	fn, ok := fetchResultMapping[a.organization]
	if !ok {
		fn = fetchResultMapping[Official]
	}
	return fn(respBody)
}

func (a *TaskAdaptor) ConvertToOpenAIVideo(originTask *model.Task) ([]byte, error) {
	var dResp responseTask
	if err := common.Unmarshal(originTask.Data, &dResp); err != nil {
		return nil, errors.Wrap(err, "unmarshal doubao task data failed")
	}

	openAIVideo := dto.NewOpenAIVideo()
	openAIVideo.ID = originTask.TaskID
	openAIVideo.TaskID = originTask.TaskID
	openAIVideo.Status = originTask.Status.ToVideoStatus()
	openAIVideo.SetProgressStr(originTask.Progress)
	openAIVideo.SetMetadata("url", dResp.Content.VideoURL)
	openAIVideo.CreatedAt = originTask.CreatedAt
	openAIVideo.CompletedAt = originTask.UpdatedAt
	openAIVideo.Model = originTask.Properties.OriginModelName

	if dResp.Status == "failed" {
		openAIVideo.Error = &dto.OpenAIVideoError{
			Message: dResp.Error.Message,
			Code:    dResp.Error.Code,
		}
	}

	return common.Marshal(openAIVideo)
}

func (a *TaskAdaptor) ConvertToDoubaoVideo(originTask *model.Task) ([]byte, error) {
	var responseItems dto.TaskResponse[model.Task]
	_ = common.Unmarshal(originTask.Data, &responseItems)
	var data = originTask.Data
	if responseItems.Code != "" {
		data = responseItems.Data.Data
	}
	var sbm tokenMartResponse
	_ = common.Unmarshal(data, &sbm)
	if sbm.Task.Id != "" {
		var response = responseTask{
			ID:        originTask.TaskID,
			Status:    sbm.Task.Status,
			Model:     originTask.Properties.OriginModelName,
			Duration:  sbm.Task.DurationSeconds,
			CreatedAt: originTask.CreatedAt,
			UpdatedAt: originTask.UpdatedAt,
		}
		response.Usage.TotalTokens = sbm.Task.Usage.TotalTokens
		response.Usage.CompletionTokens = sbm.Task.Usage.CompletionTokens
		if len(sbm.Task.Outputs) > 0 {
			response.Content.VideoURL = sbm.Task.Outputs[0]
		}
		return common.Marshal(response)
	}

	var rt responseTask
	if err := json.Unmarshal(data, &rt); nil != err {
		return data.MarshalJSON()
	}
	rt.ID = originTask.TaskID
	rt.CreatedAt = originTask.CreatedAt
	rt.UpdatedAt = originTask.UpdatedAt
	rt.Model = originTask.Properties.OriginModelName
	if rt.Status == "" {
		rt.Status = "queued"
	}
	return json.Marshal(rt)
}
