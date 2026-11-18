// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryAICCExportTaskDetailRequest struct {
	QueryAICCExportTaskDetailPath *QueryAICCExportTaskDetailPath `json:"queryAICCExportTaskDetailPath,omitempty"`
}

func (s QueryAICCExportTaskDetailRequest) String() string {
	return utils.Beautify(s)
}

func (s QueryAICCExportTaskDetailRequest) GoString() string {
	return s.String()
}

func (s QueryAICCExportTaskDetailRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryAICCExportTaskDetailRequest) SetQueryAICCExportTaskDetailPath(v *QueryAICCExportTaskDetailPath) *QueryAICCExportTaskDetailRequest {
	s.QueryAICCExportTaskDetailPath = v
	return s
}

type QueryAICCExportTaskDetailRequestBuilder struct {
	s *QueryAICCExportTaskDetailRequest
}

func NewQueryAICCExportTaskDetailRequestBuilder() *QueryAICCExportTaskDetailRequestBuilder {
	s := &QueryAICCExportTaskDetailRequest{}
	b := &QueryAICCExportTaskDetailRequestBuilder{s: s}
	return b
}

func (b *QueryAICCExportTaskDetailRequestBuilder) QueryAICCExportTaskDetailPath(v *QueryAICCExportTaskDetailPath) *QueryAICCExportTaskDetailRequestBuilder {
	b.s.QueryAICCExportTaskDetailPath = v
	return b
}

func (b *QueryAICCExportTaskDetailRequestBuilder) Build() *QueryAICCExportTaskDetailRequest {
	return b.s
}
