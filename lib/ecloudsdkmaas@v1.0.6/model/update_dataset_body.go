// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateDatasetBody struct {
	position.Body
	// 知识库名称
	Name *string `json:"name,omitempty"`
	// 知识库描述
	Description *string `json:"description,omitempty"`
	// 检索设置
	RetrievalConfig *UpdateDatasetRequestRetrievalConfig `json:"retrievalConfig,omitempty"`
}

func (s UpdateDatasetBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateDatasetBody) GoString() string {
	return s.String()
}

func (s UpdateDatasetBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateDatasetBody) SetName(v string) *UpdateDatasetBody {
	s.Name = &v
	return s
}

func (s *UpdateDatasetBody) SetDescription(v string) *UpdateDatasetBody {
	s.Description = &v
	return s
}

func (s *UpdateDatasetBody) SetRetrievalConfig(v *UpdateDatasetRequestRetrievalConfig) *UpdateDatasetBody {
	s.RetrievalConfig = v
	return s
}

type UpdateDatasetBodyBuilder struct {
	s *UpdateDatasetBody
}

func NewUpdateDatasetBodyBuilder() *UpdateDatasetBodyBuilder {
	s := &UpdateDatasetBody{}
	b := &UpdateDatasetBodyBuilder{s: s}
	return b
}

func (b *UpdateDatasetBodyBuilder) Name(v string) *UpdateDatasetBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateDatasetBodyBuilder) Description(v string) *UpdateDatasetBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *UpdateDatasetBodyBuilder) RetrievalConfig(v *UpdateDatasetRequestRetrievalConfig) *UpdateDatasetBodyBuilder {
	b.s.RetrievalConfig = v
	return b
}

func (b *UpdateDatasetBodyBuilder) Build() *UpdateDatasetBody {
	return b.s
}
