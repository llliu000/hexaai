package doubao

import (
	"fmt"
	"net/http"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/service"
	"github.com/pkg/errors"
)

const (
	ThirdAnyFast = "anyfast"
	Official     = "official"
	ThirdKWJM    = "kwjm"
)

var submitResultMapping = map[string]func(responseBody []byte) (string, *dto.TaskError){
	ThirdKWJM:    officialSubmit,
	Official:     officialSubmit,
	ThirdAnyFast: newApiSubmit,
}

func officialSubmit(responseBody []byte) (string, *dto.TaskError) {
	var dResp responsePayload
	if err := common.Unmarshal(responseBody, &dResp); err != nil {
		taskErr := service.TaskErrorWrapper(errors.Wrapf(err, "body: %s", responseBody), "unmarshal_response_body_failed", http.StatusInternalServerError)
		return "", taskErr
	}
	if dResp.ID == "" {
		taskErr := service.TaskErrorWrapper(fmt.Errorf("task_id is empty"), "invalid_response", http.StatusInternalServerError)
		return "", taskErr
	}
	return dResp.ID, nil
}

func newApiSubmit(responseBody []byte) (string, *dto.TaskError) {
	var nr dto.OpenAIVideo
	if err := common.Unmarshal(responseBody, &nr); nil != err {
		taskErr := service.TaskErrorWrapper(err, "unmarshal_response_failed", http.StatusInternalServerError)
		return "", taskErr
	}
	if nr.TaskID != "" {
		return nr.TaskID, nil
	}
	taskErr := &dto.TaskError{}
	if err := common.Unmarshal(responseBody, &taskErr); nil != err {
		taskErr = service.TaskErrorWrapper(err, "unmarshal_response_failed", http.StatusInternalServerError)
	}
	return "", taskErr
}

var fetchResultMapping = map[string]func(responseBody []byte) (*relaycommon.TaskInfo, error){
	Official:     officialFetch,
	ThirdAnyFast: newApiFetch,
	ThirdKWJM:    officialFetch,
}

// seedance-1.x获取任务详情响应结果处理
func officialFetch(responseBody []byte) (*relaycommon.TaskInfo, error) {
	resTask := responseTask{}
	if err := common.Unmarshal(responseBody, &resTask); err != nil {
		return nil, errors.Wrap(err, "unmarshal task result failed")
	}
	return taskInfo(resTask)
}

// anyfast seedance-2.x获取任务详情响应结果处理
func newApiFetch(responseBody []byte) (*relaycommon.TaskInfo, error) {
	var responseItems dto.TaskResponse[model.Task]
	if err := common.Unmarshal(responseBody, &responseItems); err != nil {
		return nil, errors.Wrap(err, "unmarshal task result failed")
	}
	if !responseItems.IsSuccess() {
		return nil, fmt.Errorf("task fetch failed:%s", responseItems.Message)
	}
	resTask := responseTask{}
	if err := common.Unmarshal(responseItems.Data.Data, &resTask); nil != err {
		return nil, err
	}
	return taskInfo(resTask)
}

func taskInfo(resTask responseTask) (*relaycommon.TaskInfo, error) {
	taskResult := relaycommon.TaskInfo{Code: 0}

	// Map Doubao status to internal status
	switch resTask.Status {
	case "pending", "queued":
		taskResult.Status = model.TaskStatusQueued
		taskResult.Progress = "10%"
	case "processing", "running":
		taskResult.Status = model.TaskStatusInProgress
		taskResult.Progress = "50%"
	case "succeeded":
		taskResult.Status = model.TaskStatusSuccess
		taskResult.Progress = "100%"
		taskResult.Url = resTask.Content.VideoURL
		// 解析 usage 信息用于按倍率计费
		taskResult.CompletionTokens = resTask.Usage.CompletionTokens
		taskResult.TotalTokens = resTask.Usage.TotalTokens
	case "failed":
		taskResult.Status = model.TaskStatusFailure
		taskResult.Progress = "100%"
		taskResult.Reason = "task failed"
	default:
		// Unknown status, treat as processing
		taskResult.Status = model.TaskStatusInProgress
		taskResult.Progress = "30%"
	}
	return &taskResult, nil
}
