// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConfirmVerificationResultResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *ConfirmVerificationResultResponseBody `json:"body,omitempty"`
}

func (s ConfirmVerificationResultResponse) String() string {
	return utils.Beautify(s)
}

func (s ConfirmVerificationResultResponse) GoString() string {
	return s.String()
}

func (s ConfirmVerificationResultResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConfirmVerificationResultResponse) SetRequestId(v string) *ConfirmVerificationResultResponse {
	s.RequestId = &v
	return s
}

func (s *ConfirmVerificationResultResponse) SetErrorMessage(v string) *ConfirmVerificationResultResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ConfirmVerificationResultResponse) SetErrorCode(v string) *ConfirmVerificationResultResponse {
	s.ErrorCode = &v
	return s
}

func (s *ConfirmVerificationResultResponse) SetState(v string) *ConfirmVerificationResultResponse {
	s.State = &v
	return s
}

func (s *ConfirmVerificationResultResponse) SetBody(v *ConfirmVerificationResultResponseBody) *ConfirmVerificationResultResponse {
	s.Body = v
	return s
}

type ConfirmVerificationResultResponseBuilder struct {
	s *ConfirmVerificationResultResponse
}

func NewConfirmVerificationResultResponseBuilder() *ConfirmVerificationResultResponseBuilder {
	s := &ConfirmVerificationResultResponse{}
	b := &ConfirmVerificationResultResponseBuilder{s: s}
	return b
}

func (b *ConfirmVerificationResultResponseBuilder) RequestId(v string) *ConfirmVerificationResultResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ConfirmVerificationResultResponseBuilder) ErrorMessage(v string) *ConfirmVerificationResultResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ConfirmVerificationResultResponseBuilder) ErrorCode(v string) *ConfirmVerificationResultResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ConfirmVerificationResultResponseBuilder) State(v string) *ConfirmVerificationResultResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ConfirmVerificationResultResponseBuilder) Body(v *ConfirmVerificationResultResponseBody) *ConfirmVerificationResultResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ConfirmVerificationResultResponseBuilder) Build() *ConfirmVerificationResultResponse {
	return b.s
}
