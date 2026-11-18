// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateExportTaskResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *CreateExportTaskResponseBody `json:"body,omitempty"`
}

func (s CreateExportTaskResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateExportTaskResponse) GoString() string {
	return s.String()
}

func (s CreateExportTaskResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateExportTaskResponse) SetRequestId(v string) *CreateExportTaskResponse {
	s.RequestId = &v
	return s
}

func (s *CreateExportTaskResponse) SetErrorMessage(v string) *CreateExportTaskResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateExportTaskResponse) SetErrorCode(v string) *CreateExportTaskResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateExportTaskResponse) SetState(v string) *CreateExportTaskResponse {
	s.State = &v
	return s
}

func (s *CreateExportTaskResponse) SetBody(v *CreateExportTaskResponseBody) *CreateExportTaskResponse {
	s.Body = v
	return s
}

type CreateExportTaskResponseBuilder struct {
	s *CreateExportTaskResponse
}

func NewCreateExportTaskResponseBuilder() *CreateExportTaskResponseBuilder {
	s := &CreateExportTaskResponse{}
	b := &CreateExportTaskResponseBuilder{s: s}
	return b
}

func (b *CreateExportTaskResponseBuilder) RequestId(v string) *CreateExportTaskResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateExportTaskResponseBuilder) ErrorMessage(v string) *CreateExportTaskResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateExportTaskResponseBuilder) ErrorCode(v string) *CreateExportTaskResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateExportTaskResponseBuilder) State(v string) *CreateExportTaskResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateExportTaskResponseBuilder) Body(v *CreateExportTaskResponseBody) *CreateExportTaskResponseBuilder {
	b.s.Body = v
	return b
}

func (b *CreateExportTaskResponseBuilder) Build() *CreateExportTaskResponse {
	return b.s
}
