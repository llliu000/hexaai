// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AutoRenewOrderBody struct {
	position.Body
	// 自动续订设置列表
	AutoRenewList *[]AutoRenewOrderRequestAutoRenewList `json:"autoRenewList,omitempty"`
}

func (s AutoRenewOrderBody) String() string {
	return utils.Beautify(s)
}

func (s AutoRenewOrderBody) GoString() string {
	return s.String()
}

func (s AutoRenewOrderBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AutoRenewOrderBody) SetAutoRenewList(v []AutoRenewOrderRequestAutoRenewList) *AutoRenewOrderBody {
	s.AutoRenewList = &v
	return s
}

type AutoRenewOrderBodyBuilder struct {
	s *AutoRenewOrderBody
}

func NewAutoRenewOrderBodyBuilder() *AutoRenewOrderBodyBuilder {
	s := &AutoRenewOrderBody{}
	b := &AutoRenewOrderBodyBuilder{s: s}
	return b
}

func (b *AutoRenewOrderBodyBuilder) AutoRenewList(v []AutoRenewOrderRequestAutoRenewList) *AutoRenewOrderBodyBuilder {
	b.s.AutoRenewList = &v
	return b
}

func (b *AutoRenewOrderBodyBuilder) Build() *AutoRenewOrderBody {
	return b.s
}
