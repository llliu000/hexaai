// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryExportTaskPath struct {
	position.Path
	// 导出任务ID
	TaskId *string `json:"taskId,omitempty"`
}

func (s QueryExportTaskPath) String() string {
	return utils.Beautify(s)
}

func (s QueryExportTaskPath) GoString() string {
	return s.String()
}

func (s QueryExportTaskPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryExportTaskPath) SetTaskId(v string) *QueryExportTaskPath {
	s.TaskId = &v
	return s
}

type QueryExportTaskPathBuilder struct {
	s *QueryExportTaskPath
}

func NewQueryExportTaskPathBuilder() *QueryExportTaskPathBuilder {
	s := &QueryExportTaskPath{}
	b := &QueryExportTaskPathBuilder{s: s}
	return b
}

func (b *QueryExportTaskPathBuilder) TaskId(v string) *QueryExportTaskPathBuilder {
	b.s.TaskId = &v
	return b
}

func (b *QueryExportTaskPathBuilder) Build() *QueryExportTaskPath {
	return b.s
}
