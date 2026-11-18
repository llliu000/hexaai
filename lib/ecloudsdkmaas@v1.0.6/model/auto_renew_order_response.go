// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AutoRenewOrderResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 是否成功开启或关闭自动续订
	Body *bool `json:"body,omitempty"`
}

func (s AutoRenewOrderResponse) String() string {
	return utils.Beautify(s)
}

func (s AutoRenewOrderResponse) GoString() string {
	return s.String()
}

func (s AutoRenewOrderResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AutoRenewOrderResponse) SetRequestId(v string) *AutoRenewOrderResponse {
	s.RequestId = &v
	return s
}

func (s *AutoRenewOrderResponse) SetErrorMessage(v string) *AutoRenewOrderResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AutoRenewOrderResponse) SetErrorCode(v string) *AutoRenewOrderResponse {
	s.ErrorCode = &v
	return s
}

func (s *AutoRenewOrderResponse) SetState(v string) *AutoRenewOrderResponse {
	s.State = &v
	return s
}

func (s *AutoRenewOrderResponse) SetBody(v bool) *AutoRenewOrderResponse {
	s.Body = &v
	return s
}

type AutoRenewOrderResponseBuilder struct {
	s *AutoRenewOrderResponse
}

func NewAutoRenewOrderResponseBuilder() *AutoRenewOrderResponseBuilder {
	s := &AutoRenewOrderResponse{}
	b := &AutoRenewOrderResponseBuilder{s: s}
	return b
}

func (b *AutoRenewOrderResponseBuilder) RequestId(v string) *AutoRenewOrderResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AutoRenewOrderResponseBuilder) ErrorMessage(v string) *AutoRenewOrderResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AutoRenewOrderResponseBuilder) ErrorCode(v string) *AutoRenewOrderResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AutoRenewOrderResponseBuilder) State(v string) *AutoRenewOrderResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AutoRenewOrderResponseBuilder) Body(v bool) *AutoRenewOrderResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AutoRenewOrderResponseBuilder) Build() *AutoRenewOrderResponse {
	return b.s
}
