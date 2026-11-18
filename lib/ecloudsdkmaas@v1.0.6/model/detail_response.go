// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DetailResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *DetailResponseBody `json:"body,omitempty"`
}

func (s DetailResponse) String() string {
	return utils.Beautify(s)
}

func (s DetailResponse) GoString() string {
	return s.String()
}

func (s DetailResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DetailResponse) SetRequestId(v string) *DetailResponse {
	s.RequestId = &v
	return s
}

func (s *DetailResponse) SetErrorMessage(v string) *DetailResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DetailResponse) SetErrorCode(v string) *DetailResponse {
	s.ErrorCode = &v
	return s
}

func (s *DetailResponse) SetState(v string) *DetailResponse {
	s.State = &v
	return s
}

func (s *DetailResponse) SetBody(v *DetailResponseBody) *DetailResponse {
	s.Body = v
	return s
}

type DetailResponseBuilder struct {
	s *DetailResponse
}

func NewDetailResponseBuilder() *DetailResponseBuilder {
	s := &DetailResponse{}
	b := &DetailResponseBuilder{s: s}
	return b
}

func (b *DetailResponseBuilder) RequestId(v string) *DetailResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DetailResponseBuilder) ErrorMessage(v string) *DetailResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DetailResponseBuilder) ErrorCode(v string) *DetailResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DetailResponseBuilder) State(v string) *DetailResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DetailResponseBuilder) Body(v *DetailResponseBody) *DetailResponseBuilder {
	b.s.Body = v
	return b
}

func (b *DetailResponseBuilder) Build() *DetailResponse {
	return b.s
}
