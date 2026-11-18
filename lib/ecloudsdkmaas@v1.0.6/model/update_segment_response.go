// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateSegmentResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *UpdateSegmentResponseBody `json:"body,omitempty"`
}

func (s UpdateSegmentResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateSegmentResponse) GoString() string {
	return s.String()
}

func (s UpdateSegmentResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateSegmentResponse) SetRequestId(v string) *UpdateSegmentResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateSegmentResponse) SetErrorMessage(v string) *UpdateSegmentResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateSegmentResponse) SetErrorCode(v string) *UpdateSegmentResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateSegmentResponse) SetState(v string) *UpdateSegmentResponse {
	s.State = &v
	return s
}

func (s *UpdateSegmentResponse) SetBody(v *UpdateSegmentResponseBody) *UpdateSegmentResponse {
	s.Body = v
	return s
}

type UpdateSegmentResponseBuilder struct {
	s *UpdateSegmentResponse
}

func NewUpdateSegmentResponseBuilder() *UpdateSegmentResponseBuilder {
	s := &UpdateSegmentResponse{}
	b := &UpdateSegmentResponseBuilder{s: s}
	return b
}

func (b *UpdateSegmentResponseBuilder) RequestId(v string) *UpdateSegmentResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateSegmentResponseBuilder) ErrorMessage(v string) *UpdateSegmentResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateSegmentResponseBuilder) ErrorCode(v string) *UpdateSegmentResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateSegmentResponseBuilder) State(v string) *UpdateSegmentResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateSegmentResponseBuilder) Body(v *UpdateSegmentResponseBody) *UpdateSegmentResponseBuilder {
	b.s.Body = v
	return b
}

func (b *UpdateSegmentResponseBuilder) Build() *UpdateSegmentResponse {
	return b.s
}
