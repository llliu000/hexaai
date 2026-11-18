// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateOrderBody struct {
	position.Body
	// 模型名称集合，同一模型名称可对应多家服务商，请传入与模型广场 API 调用完全一致的模型名称，名称不一致将导致订购匹配异常。
	Models []string `json:"models,omitempty"`
}

func (s CreateOrderBody) String() string {
	return utils.Beautify(s)
}

func (s CreateOrderBody) GoString() string {
	return s.String()
}

func (s CreateOrderBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateOrderBody) SetModels(v []string) *CreateOrderBody {
	s.Models = v
	return s
}

type CreateOrderBodyBuilder struct {
	s *CreateOrderBody
}

func NewCreateOrderBodyBuilder() *CreateOrderBodyBuilder {
	s := &CreateOrderBody{}
	b := &CreateOrderBodyBuilder{s: s}
	return b
}

func (b *CreateOrderBodyBuilder) Models(v []string) *CreateOrderBodyBuilder {
	b.s.Models = v
	return b
}

func (b *CreateOrderBodyBuilder) Build() *CreateOrderBody {
	return b.s
}
