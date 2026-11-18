// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateSessionResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *CreateSessionResponseBody `json:"body,omitempty"`
}

func (s CreateSessionResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateSessionResponse) GoString() string {
	return s.String()
}

func (s CreateSessionResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSessionResponse) SetRequestId(v string) *CreateSessionResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSessionResponse) SetErrorMessage(v string) *CreateSessionResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSessionResponse) SetErrorCode(v string) *CreateSessionResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateSessionResponse) SetState(v string) *CreateSessionResponse {
	s.State = &v
	return s
}

func (s *CreateSessionResponse) SetBody(v *CreateSessionResponseBody) *CreateSessionResponse {
	s.Body = v
	return s
}

type CreateSessionResponseBuilder struct {
	s *CreateSessionResponse
}

func NewCreateSessionResponseBuilder() *CreateSessionResponseBuilder {
	s := &CreateSessionResponse{}
	b := &CreateSessionResponseBuilder{s: s}
	return b
}

func (b *CreateSessionResponseBuilder) RequestId(v string) *CreateSessionResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateSessionResponseBuilder) ErrorMessage(v string) *CreateSessionResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateSessionResponseBuilder) ErrorCode(v string) *CreateSessionResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateSessionResponseBuilder) State(v string) *CreateSessionResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateSessionResponseBuilder) Body(v *CreateSessionResponseBody) *CreateSessionResponseBuilder {
	b.s.Body = v
	return b
}

func (b *CreateSessionResponseBuilder) Build() *CreateSessionResponse {
	return b.s
}
