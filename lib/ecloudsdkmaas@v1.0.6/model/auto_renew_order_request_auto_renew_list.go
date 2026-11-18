// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AutoRenewOrderRequestAutoRenewList struct {

	// 订购规格唯一标识信息
	InstanceId *string `json:"instanceId,omitempty"`
	// 是否开启自动续订
	AutoRenew *bool `json:"autoRenew,omitempty"`
}

func (s AutoRenewOrderRequestAutoRenewList) String() string {
	return utils.Beautify(s)
}

func (s AutoRenewOrderRequestAutoRenewList) GoString() string {
	return s.String()
}

func (s AutoRenewOrderRequestAutoRenewList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AutoRenewOrderRequestAutoRenewList) SetInstanceId(v string) *AutoRenewOrderRequestAutoRenewList {
	s.InstanceId = &v
	return s
}

func (s *AutoRenewOrderRequestAutoRenewList) SetAutoRenew(v bool) *AutoRenewOrderRequestAutoRenewList {
	s.AutoRenew = &v
	return s
}

type AutoRenewOrderRequestAutoRenewListBuilder struct {
	s *AutoRenewOrderRequestAutoRenewList
}

func NewAutoRenewOrderRequestAutoRenewListBuilder() *AutoRenewOrderRequestAutoRenewListBuilder {
	s := &AutoRenewOrderRequestAutoRenewList{}
	b := &AutoRenewOrderRequestAutoRenewListBuilder{s: s}
	return b
}

func (b *AutoRenewOrderRequestAutoRenewListBuilder) InstanceId(v string) *AutoRenewOrderRequestAutoRenewListBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *AutoRenewOrderRequestAutoRenewListBuilder) AutoRenew(v bool) *AutoRenewOrderRequestAutoRenewListBuilder {
	b.s.AutoRenew = &v
	return b
}

func (b *AutoRenewOrderRequestAutoRenewListBuilder) Build() *AutoRenewOrderRequestAutoRenewList {
	return b.s
}
