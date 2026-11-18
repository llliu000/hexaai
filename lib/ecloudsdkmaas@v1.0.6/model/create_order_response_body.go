// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateOrderResponseBodyDurationTypeEnum string

// List of DurationType
const (
	CreateOrderResponseBodyDurationTypeEnumYear  CreateOrderResponseBodyDurationTypeEnum = "YEAR"
	CreateOrderResponseBodyDurationTypeEnumMonth CreateOrderResponseBodyDurationTypeEnum = "MONTH"
	CreateOrderResponseBodyDurationTypeEnumDay   CreateOrderResponseBodyDurationTypeEnum = "DAY"
)

type CreateOrderResponseBody struct {

	// 订购时长
	Duration *int32 `json:"duration,omitempty"`
	// 订单批次编号
	OrderId *string `json:"orderId,omitempty"`
	// MOP 返回的子订单信息列表
	OrderExtInfoList *[]CreateOrderResponseOrderExtInfoList `json:"orderExtInfoList,omitempty"`
	// 订购时长单位类型
	DurationType *CreateOrderResponseBodyDurationTypeEnum `json:"durationType,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return utils.Beautify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s CreateOrderResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateOrderResponseBody) SetDuration(v int32) *CreateOrderResponseBody {
	s.Duration = &v
	return s
}

func (s *CreateOrderResponseBody) SetOrderId(v string) *CreateOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateOrderResponseBody) SetOrderExtInfoList(v []CreateOrderResponseOrderExtInfoList) *CreateOrderResponseBody {
	s.OrderExtInfoList = &v
	return s
}

func (s *CreateOrderResponseBody) SetDurationType(v CreateOrderResponseBodyDurationTypeEnum) *CreateOrderResponseBody {
	s.DurationType = &v
	return s
}

type CreateOrderResponseBodyBuilder struct {
	s *CreateOrderResponseBody
}

func NewCreateOrderResponseBodyBuilder() *CreateOrderResponseBodyBuilder {
	s := &CreateOrderResponseBody{}
	b := &CreateOrderResponseBodyBuilder{s: s}
	return b
}

func (b *CreateOrderResponseBodyBuilder) Duration(v int32) *CreateOrderResponseBodyBuilder {
	b.s.Duration = &v
	return b
}

func (b *CreateOrderResponseBodyBuilder) OrderId(v string) *CreateOrderResponseBodyBuilder {
	b.s.OrderId = &v
	return b
}

func (b *CreateOrderResponseBodyBuilder) OrderExtInfoList(v []CreateOrderResponseOrderExtInfoList) *CreateOrderResponseBodyBuilder {
	b.s.OrderExtInfoList = &v
	return b
}

func (b *CreateOrderResponseBodyBuilder) DurationType(v CreateOrderResponseBodyDurationTypeEnum) *CreateOrderResponseBodyBuilder {
	b.s.DurationType = &v
	return b
}

func (b *CreateOrderResponseBodyBuilder) Build() *CreateOrderResponseBody {
	return b.s
}
