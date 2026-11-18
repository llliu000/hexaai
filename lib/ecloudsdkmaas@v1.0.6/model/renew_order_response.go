// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewOrderResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 订单批次编号
	Body *string `json:"body,omitempty"`
}

func (s RenewOrderResponse) String() string {
	return utils.Beautify(s)
}

func (s RenewOrderResponse) GoString() string {
	return s.String()
}

func (s RenewOrderResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewOrderResponse) SetRequestId(v string) *RenewOrderResponse {
	s.RequestId = &v
	return s
}

func (s *RenewOrderResponse) SetErrorMessage(v string) *RenewOrderResponse {
	s.ErrorMessage = &v
	return s
}

func (s *RenewOrderResponse) SetErrorCode(v string) *RenewOrderResponse {
	s.ErrorCode = &v
	return s
}

func (s *RenewOrderResponse) SetState(v string) *RenewOrderResponse {
	s.State = &v
	return s
}

func (s *RenewOrderResponse) SetBody(v string) *RenewOrderResponse {
	s.Body = &v
	return s
}

type RenewOrderResponseBuilder struct {
	s *RenewOrderResponse
}

func NewRenewOrderResponseBuilder() *RenewOrderResponseBuilder {
	s := &RenewOrderResponse{}
	b := &RenewOrderResponseBuilder{s: s}
	return b
}

func (b *RenewOrderResponseBuilder) RequestId(v string) *RenewOrderResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *RenewOrderResponseBuilder) ErrorMessage(v string) *RenewOrderResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *RenewOrderResponseBuilder) ErrorCode(v string) *RenewOrderResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *RenewOrderResponseBuilder) State(v string) *RenewOrderResponseBuilder {
	b.s.State = &v
	return b
}

func (b *RenewOrderResponseBuilder) Body(v string) *RenewOrderResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *RenewOrderResponseBuilder) Build() *RenewOrderResponse {
	return b.s
}
