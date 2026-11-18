// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListDatasetByUserIdResponseDataRows struct {

	// 创建时间
	CreatedAt *string `json:"createdAt,omitempty"`
	// 召回命中数
	HitCount *int32 `json:"hitCount,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updatedBy,omitempty"`
	// 字符数
	WordCount *int32 `json:"wordCount,omitempty"`
	// 创建者
	CreatedBy *string `json:"createdBy,omitempty"`
	// 知识库名称
	Name *string `json:"name,omitempty"`
	// 文档个数
	DocumentNum *int32 `json:"documentNum,omitempty"`
	// 知识库描述
	Description *string `json:"description,omitempty"`
	// 知识库id
	Id *string `json:"id,omitempty"`
	// 检索配置
	RetrievalConfig *ListDatasetByUserIdResponseRetrievalConfig `json:"retrievalConfig,omitempty"`
	// 更新时间
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

func (s ListDatasetByUserIdResponseDataRows) String() string {
	return utils.Beautify(s)
}

func (s ListDatasetByUserIdResponseDataRows) GoString() string {
	return s.String()
}

func (s ListDatasetByUserIdResponseDataRows) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListDatasetByUserIdResponseDataRows) SetCreatedAt(v string) *ListDatasetByUserIdResponseDataRows {
	s.CreatedAt = &v
	return s
}

func (s *ListDatasetByUserIdResponseDataRows) SetHitCount(v int32) *ListDatasetByUserIdResponseDataRows {
	s.HitCount = &v
	return s
}

func (s *ListDatasetByUserIdResponseDataRows) SetUpdatedBy(v string) *ListDatasetByUserIdResponseDataRows {
	s.UpdatedBy = &v
	return s
}

func (s *ListDatasetByUserIdResponseDataRows) SetWordCount(v int32) *ListDatasetByUserIdResponseDataRows {
	s.WordCount = &v
	return s
}

func (s *ListDatasetByUserIdResponseDataRows) SetCreatedBy(v string) *ListDatasetByUserIdResponseDataRows {
	s.CreatedBy = &v
	return s
}

func (s *ListDatasetByUserIdResponseDataRows) SetName(v string) *ListDatasetByUserIdResponseDataRows {
	s.Name = &v
	return s
}

func (s *ListDatasetByUserIdResponseDataRows) SetDocumentNum(v int32) *ListDatasetByUserIdResponseDataRows {
	s.DocumentNum = &v
	return s
}

func (s *ListDatasetByUserIdResponseDataRows) SetDescription(v string) *ListDatasetByUserIdResponseDataRows {
	s.Description = &v
	return s
}

func (s *ListDatasetByUserIdResponseDataRows) SetId(v string) *ListDatasetByUserIdResponseDataRows {
	s.Id = &v
	return s
}

func (s *ListDatasetByUserIdResponseDataRows) SetRetrievalConfig(v *ListDatasetByUserIdResponseRetrievalConfig) *ListDatasetByUserIdResponseDataRows {
	s.RetrievalConfig = v
	return s
}

func (s *ListDatasetByUserIdResponseDataRows) SetUpdatedAt(v string) *ListDatasetByUserIdResponseDataRows {
	s.UpdatedAt = &v
	return s
}

type ListDatasetByUserIdResponseDataRowsBuilder struct {
	s *ListDatasetByUserIdResponseDataRows
}

func NewListDatasetByUserIdResponseDataRowsBuilder() *ListDatasetByUserIdResponseDataRowsBuilder {
	s := &ListDatasetByUserIdResponseDataRows{}
	b := &ListDatasetByUserIdResponseDataRowsBuilder{s: s}
	return b
}

func (b *ListDatasetByUserIdResponseDataRowsBuilder) CreatedAt(v string) *ListDatasetByUserIdResponseDataRowsBuilder {
	b.s.CreatedAt = &v
	return b
}

func (b *ListDatasetByUserIdResponseDataRowsBuilder) HitCount(v int32) *ListDatasetByUserIdResponseDataRowsBuilder {
	b.s.HitCount = &v
	return b
}

func (b *ListDatasetByUserIdResponseDataRowsBuilder) UpdatedBy(v string) *ListDatasetByUserIdResponseDataRowsBuilder {
	b.s.UpdatedBy = &v
	return b
}

func (b *ListDatasetByUserIdResponseDataRowsBuilder) WordCount(v int32) *ListDatasetByUserIdResponseDataRowsBuilder {
	b.s.WordCount = &v
	return b
}

func (b *ListDatasetByUserIdResponseDataRowsBuilder) CreatedBy(v string) *ListDatasetByUserIdResponseDataRowsBuilder {
	b.s.CreatedBy = &v
	return b
}

func (b *ListDatasetByUserIdResponseDataRowsBuilder) Name(v string) *ListDatasetByUserIdResponseDataRowsBuilder {
	b.s.Name = &v
	return b
}

func (b *ListDatasetByUserIdResponseDataRowsBuilder) DocumentNum(v int32) *ListDatasetByUserIdResponseDataRowsBuilder {
	b.s.DocumentNum = &v
	return b
}

func (b *ListDatasetByUserIdResponseDataRowsBuilder) Description(v string) *ListDatasetByUserIdResponseDataRowsBuilder {
	b.s.Description = &v
	return b
}

func (b *ListDatasetByUserIdResponseDataRowsBuilder) Id(v string) *ListDatasetByUserIdResponseDataRowsBuilder {
	b.s.Id = &v
	return b
}

func (b *ListDatasetByUserIdResponseDataRowsBuilder) RetrievalConfig(v *ListDatasetByUserIdResponseRetrievalConfig) *ListDatasetByUserIdResponseDataRowsBuilder {
	b.s.RetrievalConfig = v
	return b
}

func (b *ListDatasetByUserIdResponseDataRowsBuilder) UpdatedAt(v string) *ListDatasetByUserIdResponseDataRowsBuilder {
	b.s.UpdatedAt = &v
	return b
}

func (b *ListDatasetByUserIdResponseDataRowsBuilder) Build() *ListDatasetByUserIdResponseDataRows {
	return b.s
}
