// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateOrderResponseOrderExtInfoList struct {

	// 模型名称
	ChaGroupName *string `json:"chaGroupName,omitempty"`
	// 子订单ID
	OrderExtId *string `json:"orderExtId,omitempty"`
}

func (s CreateOrderResponseOrderExtInfoList) String() string {
	return utils.Beautify(s)
}

func (s CreateOrderResponseOrderExtInfoList) GoString() string {
	return s.String()
}

func (s CreateOrderResponseOrderExtInfoList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateOrderResponseOrderExtInfoList) SetChaGroupName(v string) *CreateOrderResponseOrderExtInfoList {
	s.ChaGroupName = &v
	return s
}

func (s *CreateOrderResponseOrderExtInfoList) SetOrderExtId(v string) *CreateOrderResponseOrderExtInfoList {
	s.OrderExtId = &v
	return s
}

type CreateOrderResponseOrderExtInfoListBuilder struct {
	s *CreateOrderResponseOrderExtInfoList
}

func NewCreateOrderResponseOrderExtInfoListBuilder() *CreateOrderResponseOrderExtInfoListBuilder {
	s := &CreateOrderResponseOrderExtInfoList{}
	b := &CreateOrderResponseOrderExtInfoListBuilder{s: s}
	return b
}

func (b *CreateOrderResponseOrderExtInfoListBuilder) ChaGroupName(v string) *CreateOrderResponseOrderExtInfoListBuilder {
	b.s.ChaGroupName = &v
	return b
}

func (b *CreateOrderResponseOrderExtInfoListBuilder) OrderExtId(v string) *CreateOrderResponseOrderExtInfoListBuilder {
	b.s.OrderExtId = &v
	return b
}

func (b *CreateOrderResponseOrderExtInfoListBuilder) Build() *CreateOrderResponseOrderExtInfoList {
	return b.s
}
