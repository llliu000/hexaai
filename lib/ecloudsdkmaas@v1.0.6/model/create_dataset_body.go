// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateDatasetBody struct {
	position.Body
	// 知识库名称
	Name *string `json:"name,omitempty"`
	// 知识库描述
	Description *string `json:"description,omitempty"`
	// 检索设置
	RetrievalConfig *CreateDatasetRequestRetrievalConfig `json:"retrievalConfig,omitempty"`
}

func (s CreateDatasetBody) String() string {
	return utils.Beautify(s)
}

func (s CreateDatasetBody) GoString() string {
	return s.String()
}

func (s CreateDatasetBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateDatasetBody) SetName(v string) *CreateDatasetBody {
	s.Name = &v
	return s
}

func (s *CreateDatasetBody) SetDescription(v string) *CreateDatasetBody {
	s.Description = &v
	return s
}

func (s *CreateDatasetBody) SetRetrievalConfig(v *CreateDatasetRequestRetrievalConfig) *CreateDatasetBody {
	s.RetrievalConfig = v
	return s
}

type CreateDatasetBodyBuilder struct {
	s *CreateDatasetBody
}

func NewCreateDatasetBodyBuilder() *CreateDatasetBodyBuilder {
	s := &CreateDatasetBody{}
	b := &CreateDatasetBodyBuilder{s: s}
	return b
}

func (b *CreateDatasetBodyBuilder) Name(v string) *CreateDatasetBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *CreateDatasetBodyBuilder) Description(v string) *CreateDatasetBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateDatasetBodyBuilder) RetrievalConfig(v *CreateDatasetRequestRetrievalConfig) *CreateDatasetBodyBuilder {
	b.s.RetrievalConfig = v
	return b
}

func (b *CreateDatasetBodyBuilder) Build() *CreateDatasetBody {
	return b.s
}
