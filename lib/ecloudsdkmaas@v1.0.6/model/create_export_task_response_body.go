// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateExportTaskResponseBody struct {

	// 导出任务ID（用于查询任务状态）
	TaskId *string `json:"taskId,omitempty"`
}

func (s CreateExportTaskResponseBody) String() string {
	return utils.Beautify(s)
}

func (s CreateExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s CreateExportTaskResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateExportTaskResponseBody) SetTaskId(v string) *CreateExportTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateExportTaskResponseBodyBuilder struct {
	s *CreateExportTaskResponseBody
}

func NewCreateExportTaskResponseBodyBuilder() *CreateExportTaskResponseBodyBuilder {
	s := &CreateExportTaskResponseBody{}
	b := &CreateExportTaskResponseBodyBuilder{s: s}
	return b
}

func (b *CreateExportTaskResponseBodyBuilder) TaskId(v string) *CreateExportTaskResponseBodyBuilder {
	b.s.TaskId = &v
	return b
}

func (b *CreateExportTaskResponseBodyBuilder) Build() *CreateExportTaskResponseBody {
	return b.s
}
