// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewOrderBody struct {
	position.Body
	// 订购规格唯一标识信息
	InstanceIds []string `json:"instanceIds,omitempty"`
}

func (s RenewOrderBody) String() string {
	return utils.Beautify(s)
}

func (s RenewOrderBody) GoString() string {
	return s.String()
}

func (s RenewOrderBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewOrderBody) SetInstanceIds(v []string) *RenewOrderBody {
	s.InstanceIds = v
	return s
}

type RenewOrderBodyBuilder struct {
	s *RenewOrderBody
}

func NewRenewOrderBodyBuilder() *RenewOrderBodyBuilder {
	s := &RenewOrderBody{}
	b := &RenewOrderBodyBuilder{s: s}
	return b
}

func (b *RenewOrderBodyBuilder) InstanceIds(v []string) *RenewOrderBodyBuilder {
	b.s.InstanceIds = v
	return b
}

func (b *RenewOrderBodyBuilder) Build() *RenewOrderBody {
	return b.s
}
