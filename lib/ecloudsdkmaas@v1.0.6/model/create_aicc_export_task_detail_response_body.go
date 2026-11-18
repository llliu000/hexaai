// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAICCExportTaskDetailResponseBody struct {

	// 导出任务ID（用于查询任务状态）
	TaskId *string `json:"taskId,omitempty"`
}

func (s CreateAICCExportTaskDetailResponseBody) String() string {
	return utils.Beautify(s)
}

func (s CreateAICCExportTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s CreateAICCExportTaskDetailResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAICCExportTaskDetailResponseBody) SetTaskId(v string) *CreateAICCExportTaskDetailResponseBody {
	s.TaskId = &v
	return s
}

type CreateAICCExportTaskDetailResponseBodyBuilder struct {
	s *CreateAICCExportTaskDetailResponseBody
}

func NewCreateAICCExportTaskDetailResponseBodyBuilder() *CreateAICCExportTaskDetailResponseBodyBuilder {
	s := &CreateAICCExportTaskDetailResponseBody{}
	b := &CreateAICCExportTaskDetailResponseBodyBuilder{s: s}
	return b
}

func (b *CreateAICCExportTaskDetailResponseBodyBuilder) TaskId(v string) *CreateAICCExportTaskDetailResponseBodyBuilder {
	b.s.TaskId = &v
	return b
}

func (b *CreateAICCExportTaskDetailResponseBodyBuilder) Build() *CreateAICCExportTaskDetailResponseBody {
	return b.s
}
