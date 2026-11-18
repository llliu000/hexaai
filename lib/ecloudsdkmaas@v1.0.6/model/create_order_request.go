// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateOrderRequest struct {
	CreateOrderBody *CreateOrderBody `json:"createOrderBody,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s CreateOrderRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateOrderRequest) SetCreateOrderBody(v *CreateOrderBody) *CreateOrderRequest {
	s.CreateOrderBody = v
	return s
}

type CreateOrderRequestBuilder struct {
	s *CreateOrderRequest
}

func NewCreateOrderRequestBuilder() *CreateOrderRequestBuilder {
	s := &CreateOrderRequest{}
	b := &CreateOrderRequestBuilder{s: s}
	return b
}

func (b *CreateOrderRequestBuilder) CreateOrderBody(v *CreateOrderBody) *CreateOrderRequestBuilder {
	b.s.CreateOrderBody = v
	return b
}

func (b *CreateOrderRequestBuilder) Build() *CreateOrderRequest {
	return b.s
}
