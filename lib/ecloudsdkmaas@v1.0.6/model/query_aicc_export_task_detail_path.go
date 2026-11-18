// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryAICCExportTaskDetailPath struct {
	position.Path
	// 导出任务ID（创建任务时返回的taskId）
	TaskId *string `json:"taskId,omitempty"`
}

func (s QueryAICCExportTaskDetailPath) String() string {
	return utils.Beautify(s)
}

func (s QueryAICCExportTaskDetailPath) GoString() string {
	return s.String()
}

func (s QueryAICCExportTaskDetailPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryAICCExportTaskDetailPath) SetTaskId(v string) *QueryAICCExportTaskDetailPath {
	s.TaskId = &v
	return s
}

type QueryAICCExportTaskDetailPathBuilder struct {
	s *QueryAICCExportTaskDetailPath
}

func NewQueryAICCExportTaskDetailPathBuilder() *QueryAICCExportTaskDetailPathBuilder {
	s := &QueryAICCExportTaskDetailPath{}
	b := &QueryAICCExportTaskDetailPathBuilder{s: s}
	return b
}

func (b *QueryAICCExportTaskDetailPathBuilder) TaskId(v string) *QueryAICCExportTaskDetailPathBuilder {
	b.s.TaskId = &v
	return b
}

func (b *QueryAICCExportTaskDetailPathBuilder) Build() *QueryAICCExportTaskDetailPath {
	return b.s
}
