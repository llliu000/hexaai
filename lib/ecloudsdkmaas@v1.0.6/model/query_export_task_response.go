// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryExportTaskResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *QueryExportTaskResponseBody `json:"body,omitempty"`
}

func (s QueryExportTaskResponse) String() string {
	return utils.Beautify(s)
}

func (s QueryExportTaskResponse) GoString() string {
	return s.String()
}

func (s QueryExportTaskResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryExportTaskResponse) SetRequestId(v string) *QueryExportTaskResponse {
	s.RequestId = &v
	return s
}

func (s *QueryExportTaskResponse) SetErrorMessage(v string) *QueryExportTaskResponse {
	s.ErrorMessage = &v
	return s
}

func (s *QueryExportTaskResponse) SetErrorCode(v string) *QueryExportTaskResponse {
	s.ErrorCode = &v
	return s
}

func (s *QueryExportTaskResponse) SetState(v string) *QueryExportTaskResponse {
	s.State = &v
	return s
}

func (s *QueryExportTaskResponse) SetBody(v *QueryExportTaskResponseBody) *QueryExportTaskResponse {
	s.Body = v
	return s
}

type QueryExportTaskResponseBuilder struct {
	s *QueryExportTaskResponse
}

func NewQueryExportTaskResponseBuilder() *QueryExportTaskResponseBuilder {
	s := &QueryExportTaskResponse{}
	b := &QueryExportTaskResponseBuilder{s: s}
	return b
}

func (b *QueryExportTaskResponseBuilder) RequestId(v string) *QueryExportTaskResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *QueryExportTaskResponseBuilder) ErrorMessage(v string) *QueryExportTaskResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *QueryExportTaskResponseBuilder) ErrorCode(v string) *QueryExportTaskResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *QueryExportTaskResponseBuilder) State(v string) *QueryExportTaskResponseBuilder {
	b.s.State = &v
	return b
}

func (b *QueryExportTaskResponseBuilder) Body(v *QueryExportTaskResponseBody) *QueryExportTaskResponseBuilder {
	b.s.Body = v
	return b
}

func (b *QueryExportTaskResponseBuilder) Build() *QueryExportTaskResponse {
	return b.s
}
