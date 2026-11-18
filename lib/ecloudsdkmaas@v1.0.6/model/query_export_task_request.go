// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryExportTaskRequest struct {
	QueryExportTaskPath *QueryExportTaskPath `json:"queryExportTaskPath,omitempty"`
}

func (s QueryExportTaskRequest) String() string {
	return utils.Beautify(s)
}

func (s QueryExportTaskRequest) GoString() string {
	return s.String()
}

func (s QueryExportTaskRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryExportTaskRequest) SetQueryExportTaskPath(v *QueryExportTaskPath) *QueryExportTaskRequest {
	s.QueryExportTaskPath = v
	return s
}

type QueryExportTaskRequestBuilder struct {
	s *QueryExportTaskRequest
}

func NewQueryExportTaskRequestBuilder() *QueryExportTaskRequestBuilder {
	s := &QueryExportTaskRequest{}
	b := &QueryExportTaskRequestBuilder{s: s}
	return b
}

func (b *QueryExportTaskRequestBuilder) QueryExportTaskPath(v *QueryExportTaskPath) *QueryExportTaskRequestBuilder {
	b.s.QueryExportTaskPath = v
	return b
}

func (b *QueryExportTaskRequestBuilder) Build() *QueryExportTaskRequest {
	return b.s
}
