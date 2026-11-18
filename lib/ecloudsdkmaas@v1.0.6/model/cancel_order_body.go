// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CancelOrderBody struct {
	position.Body
	// 模型名称集合，同一模型名称可对应多家服务商，请传入与模型广场 API 调用完全一致的模型名称；与 instanceIds 二选一，若同时传入，以 models 为准。
	Models []string `json:"models,omitempty"`
	// 订购规格唯一标识 ID 集合；与 models 二选一，若同时传入，以 models 为准。
	InstanceIds []string `json:"instanceIds,omitempty"`
}

func (s CancelOrderBody) String() string {
	return utils.Beautify(s)
}

func (s CancelOrderBody) GoString() string {
	return s.String()
}

func (s CancelOrderBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CancelOrderBody) SetModels(v []string) *CancelOrderBody {
	s.Models = v
	return s
}

func (s *CancelOrderBody) SetInstanceIds(v []string) *CancelOrderBody {
	s.InstanceIds = v
	return s
}

type CancelOrderBodyBuilder struct {
	s *CancelOrderBody
}

func NewCancelOrderBodyBuilder() *CancelOrderBodyBuilder {
	s := &CancelOrderBody{}
	b := &CancelOrderBodyBuilder{s: s}
	return b
}

func (b *CancelOrderBodyBuilder) Models(v []string) *CancelOrderBodyBuilder {
	b.s.Models = v
	return b
}

func (b *CancelOrderBodyBuilder) InstanceIds(v []string) *CancelOrderBodyBuilder {
	b.s.InstanceIds = v
	return b
}

func (b *CancelOrderBodyBuilder) Build() *CancelOrderBody {
	return b.s
}
