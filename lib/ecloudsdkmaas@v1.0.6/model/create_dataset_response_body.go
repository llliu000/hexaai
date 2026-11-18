// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateDatasetResponseBody struct {

	// 创建时间
	CreatedAt *string `json:"createdAt,omitempty"`
	// 知识库名称
	Name *string `json:"name,omitempty"`
	// 知识库ID
	DatasetId *string `json:"datasetId,omitempty"`
	// 知识库描述
	Description *string `json:"description,omitempty"`
	// 更新时间
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return utils.Beautify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s CreateDatasetResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateDatasetResponseBody) SetCreatedAt(v string) *CreateDatasetResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateDatasetResponseBody) SetName(v string) *CreateDatasetResponseBody {
	s.Name = &v
	return s
}

func (s *CreateDatasetResponseBody) SetDatasetId(v string) *CreateDatasetResponseBody {
	s.DatasetId = &v
	return s
}

func (s *CreateDatasetResponseBody) SetDescription(v string) *CreateDatasetResponseBody {
	s.Description = &v
	return s
}

func (s *CreateDatasetResponseBody) SetUpdatedAt(v string) *CreateDatasetResponseBody {
	s.UpdatedAt = &v
	return s
}

type CreateDatasetResponseBodyBuilder struct {
	s *CreateDatasetResponseBody
}

func NewCreateDatasetResponseBodyBuilder() *CreateDatasetResponseBodyBuilder {
	s := &CreateDatasetResponseBody{}
	b := &CreateDatasetResponseBodyBuilder{s: s}
	return b
}

func (b *CreateDatasetResponseBodyBuilder) CreatedAt(v string) *CreateDatasetResponseBodyBuilder {
	b.s.CreatedAt = &v
	return b
}

func (b *CreateDatasetResponseBodyBuilder) Name(v string) *CreateDatasetResponseBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *CreateDatasetResponseBodyBuilder) DatasetId(v string) *CreateDatasetResponseBodyBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *CreateDatasetResponseBodyBuilder) Description(v string) *CreateDatasetResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateDatasetResponseBodyBuilder) UpdatedAt(v string) *CreateDatasetResponseBodyBuilder {
	b.s.UpdatedAt = &v
	return b
}

func (b *CreateDatasetResponseBodyBuilder) Build() *CreateDatasetResponseBody {
	return b.s
}
