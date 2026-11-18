// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AutoRenewOrderRequest struct {
	AutoRenewOrderBody *AutoRenewOrderBody `json:"autoRenewOrderBody,omitempty"`
}

func (s AutoRenewOrderRequest) String() string {
	return utils.Beautify(s)
}

func (s AutoRenewOrderRequest) GoString() string {
	return s.String()
}

func (s AutoRenewOrderRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AutoRenewOrderRequest) SetAutoRenewOrderBody(v *AutoRenewOrderBody) *AutoRenewOrderRequest {
	s.AutoRenewOrderBody = v
	return s
}

type AutoRenewOrderRequestBuilder struct {
	s *AutoRenewOrderRequest
}

func NewAutoRenewOrderRequestBuilder() *AutoRenewOrderRequestBuilder {
	s := &AutoRenewOrderRequest{}
	b := &AutoRenewOrderRequestBuilder{s: s}
	return b
}

func (b *AutoRenewOrderRequestBuilder) AutoRenewOrderBody(v *AutoRenewOrderBody) *AutoRenewOrderRequestBuilder {
	b.s.AutoRenewOrderBody = v
	return b
}

func (b *AutoRenewOrderRequestBuilder) Build() *AutoRenewOrderRequest {
	return b.s
}
