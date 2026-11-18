// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CancelOrderRequest struct {
	CancelOrderBody *CancelOrderBody `json:"cancelOrderBody,omitempty"`
}

func (s CancelOrderRequest) String() string {
	return utils.Beautify(s)
}

func (s CancelOrderRequest) GoString() string {
	return s.String()
}

func (s CancelOrderRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CancelOrderRequest) SetCancelOrderBody(v *CancelOrderBody) *CancelOrderRequest {
	s.CancelOrderBody = v
	return s
}

type CancelOrderRequestBuilder struct {
	s *CancelOrderRequest
}

func NewCancelOrderRequestBuilder() *CancelOrderRequestBuilder {
	s := &CancelOrderRequest{}
	b := &CancelOrderRequestBuilder{s: s}
	return b
}

func (b *CancelOrderRequestBuilder) CancelOrderBody(v *CancelOrderBody) *CancelOrderRequestBuilder {
	b.s.CancelOrderBody = v
	return b
}

func (b *CancelOrderRequestBuilder) Build() *CancelOrderRequest {
	return b.s
}
