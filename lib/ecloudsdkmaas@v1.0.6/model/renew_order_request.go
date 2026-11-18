// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewOrderRequest struct {
	RenewOrderBody *RenewOrderBody `json:"renewOrderBody,omitempty"`
}

func (s RenewOrderRequest) String() string {
	return utils.Beautify(s)
}

func (s RenewOrderRequest) GoString() string {
	return s.String()
}

func (s RenewOrderRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewOrderRequest) SetRenewOrderBody(v *RenewOrderBody) *RenewOrderRequest {
	s.RenewOrderBody = v
	return s
}

type RenewOrderRequestBuilder struct {
	s *RenewOrderRequest
}

func NewRenewOrderRequestBuilder() *RenewOrderRequestBuilder {
	s := &RenewOrderRequest{}
	b := &RenewOrderRequestBuilder{s: s}
	return b
}

func (b *RenewOrderRequestBuilder) RenewOrderBody(v *RenewOrderBody) *RenewOrderRequestBuilder {
	b.s.RenewOrderBody = v
	return b
}

func (b *RenewOrderRequestBuilder) Build() *RenewOrderRequest {
	return b.s
}
