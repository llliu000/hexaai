// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateOrderResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 返回体
	Body *CreateOrderResponseBody `json:"body,omitempty"`
}

func (s CreateOrderResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateOrderResponse) GoString() string {
	return s.String()
}

func (s CreateOrderResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateOrderResponse) SetRequestId(v string) *CreateOrderResponse {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponse) SetErrorMessage(v string) *CreateOrderResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateOrderResponse) SetErrorCode(v string) *CreateOrderResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateOrderResponse) SetState(v string) *CreateOrderResponse {
	s.State = &v
	return s
}

func (s *CreateOrderResponse) SetBody(v *CreateOrderResponseBody) *CreateOrderResponse {
	s.Body = v
	return s
}

type CreateOrderResponseBuilder struct {
	s *CreateOrderResponse
}

func NewCreateOrderResponseBuilder() *CreateOrderResponseBuilder {
	s := &CreateOrderResponse{}
	b := &CreateOrderResponseBuilder{s: s}
	return b
}

func (b *CreateOrderResponseBuilder) RequestId(v string) *CreateOrderResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateOrderResponseBuilder) ErrorMessage(v string) *CreateOrderResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateOrderResponseBuilder) ErrorCode(v string) *CreateOrderResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateOrderResponseBuilder) State(v string) *CreateOrderResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateOrderResponseBuilder) Body(v *CreateOrderResponseBody) *CreateOrderResponseBuilder {
	b.s.Body = v
	return b
}

func (b *CreateOrderResponseBuilder) Build() *CreateOrderResponse {
	return b.s
}
