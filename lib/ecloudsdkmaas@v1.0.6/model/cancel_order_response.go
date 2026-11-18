// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CancelOrderResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 是否退订成功
	Body *bool `json:"body,omitempty"`
}

func (s CancelOrderResponse) String() string {
	return utils.Beautify(s)
}

func (s CancelOrderResponse) GoString() string {
	return s.String()
}

func (s CancelOrderResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CancelOrderResponse) SetRequestId(v string) *CancelOrderResponse {
	s.RequestId = &v
	return s
}

func (s *CancelOrderResponse) SetErrorMessage(v string) *CancelOrderResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CancelOrderResponse) SetErrorCode(v string) *CancelOrderResponse {
	s.ErrorCode = &v
	return s
}

func (s *CancelOrderResponse) SetState(v string) *CancelOrderResponse {
	s.State = &v
	return s
}

func (s *CancelOrderResponse) SetBody(v bool) *CancelOrderResponse {
	s.Body = &v
	return s
}

type CancelOrderResponseBuilder struct {
	s *CancelOrderResponse
}

func NewCancelOrderResponseBuilder() *CancelOrderResponseBuilder {
	s := &CancelOrderResponse{}
	b := &CancelOrderResponseBuilder{s: s}
	return b
}

func (b *CancelOrderResponseBuilder) RequestId(v string) *CancelOrderResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CancelOrderResponseBuilder) ErrorMessage(v string) *CancelOrderResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CancelOrderResponseBuilder) ErrorCode(v string) *CancelOrderResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CancelOrderResponseBuilder) State(v string) *CancelOrderResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CancelOrderResponseBuilder) Body(v bool) *CancelOrderResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CancelOrderResponseBuilder) Build() *CancelOrderResponse {
	return b.s
}
