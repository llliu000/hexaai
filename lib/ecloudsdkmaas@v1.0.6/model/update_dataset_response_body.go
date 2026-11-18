// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateDatasetResponseBody struct {

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

func (s UpdateDatasetResponseBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s UpdateDatasetResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateDatasetResponseBody) SetCreatedAt(v string) *UpdateDatasetResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetName(v string) *UpdateDatasetResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetDatasetId(v string) *UpdateDatasetResponseBody {
	s.DatasetId = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetDescription(v string) *UpdateDatasetResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetUpdatedAt(v string) *UpdateDatasetResponseBody {
	s.UpdatedAt = &v
	return s
}

type UpdateDatasetResponseBodyBuilder struct {
	s *UpdateDatasetResponseBody
}

func NewUpdateDatasetResponseBodyBuilder() *UpdateDatasetResponseBodyBuilder {
	s := &UpdateDatasetResponseBody{}
	b := &UpdateDatasetResponseBodyBuilder{s: s}
	return b
}

func (b *UpdateDatasetResponseBodyBuilder) CreatedAt(v string) *UpdateDatasetResponseBodyBuilder {
	b.s.CreatedAt = &v
	return b
}

func (b *UpdateDatasetResponseBodyBuilder) Name(v string) *UpdateDatasetResponseBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *UpdateDatasetResponseBodyBuilder) DatasetId(v string) *UpdateDatasetResponseBodyBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *UpdateDatasetResponseBodyBuilder) Description(v string) *UpdateDatasetResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *UpdateDatasetResponseBodyBuilder) UpdatedAt(v string) *UpdateDatasetResponseBodyBuilder {
	b.s.UpdatedAt = &v
	return b
}

func (b *UpdateDatasetResponseBodyBuilder) Build() *UpdateDatasetResponseBody {
	return b.s
}
