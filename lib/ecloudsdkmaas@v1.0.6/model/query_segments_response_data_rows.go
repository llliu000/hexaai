// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QuerySegmentsResponseDataRowsStatusEnum string

// List of Status
const (
	QuerySegmentsResponseDataRowsStatusEnumWaiting   QuerySegmentsResponseDataRowsStatusEnum = "WAITING"
	QuerySegmentsResponseDataRowsStatusEnumQueuing   QuerySegmentsResponseDataRowsStatusEnum = "QUEUING"
	QuerySegmentsResponseDataRowsStatusEnumParsing   QuerySegmentsResponseDataRowsStatusEnum = "PARSING"
	QuerySegmentsResponseDataRowsStatusEnumCleaning  QuerySegmentsResponseDataRowsStatusEnum = "CLEANING"
	QuerySegmentsResponseDataRowsStatusEnumSplitting QuerySegmentsResponseDataRowsStatusEnum = "SPLITTING"
	QuerySegmentsResponseDataRowsStatusEnumIndexing  QuerySegmentsResponseDataRowsStatusEnum = "INDEXING"
	QuerySegmentsResponseDataRowsStatusEnumCompleted QuerySegmentsResponseDataRowsStatusEnum = "COMPLETED"
	QuerySegmentsResponseDataRowsStatusEnumExceeded  QuerySegmentsResponseDataRowsStatusEnum = "EXCEEDED"
	QuerySegmentsResponseDataRowsStatusEnumEmpty     QuerySegmentsResponseDataRowsStatusEnum = "EMPTY"
	QuerySegmentsResponseDataRowsStatusEnumError     QuerySegmentsResponseDataRowsStatusEnum = "ERROR"
	QuerySegmentsResponseDataRowsStatusEnumDeleting  QuerySegmentsResponseDataRowsStatusEnum = "DELETING"
	QuerySegmentsResponseDataRowsStatusEnumDeleted   QuerySegmentsResponseDataRowsStatusEnum = "DELETED"
)

type QuerySegmentsResponseDataRows struct {

	// 创建时间
	CreatedAt *string `json:"createdAt,omitempty"`
	// 命中次数
	HitCount *int32 `json:"hitCount,omitempty"`
	// QA模式  问答
	Answer *string `json:"answer,omitempty"`
	// 字数
	WordCount *int32 `json:"wordCount,omitempty"`
	// 创建人
	CreatedBy *string `json:"createdBy,omitempty"`
	// 知识库id
	DatasetId *string `json:"datasetId,omitempty"`
	// 文档id
	DocumentId *string `json:"documentId,omitempty"`
	// 分段id
	Id *string `json:"id,omitempty"`
	// 分段内容
	Content *string `json:"content,omitempty"`
	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	ContentHash *string `json:"contentHash,omitempty"`
	// 状态
	Status *QuerySegmentsResponseDataRowsStatusEnum `json:"status,omitempty"`
}

func (s QuerySegmentsResponseDataRows) String() string {
	return utils.Beautify(s)
}

func (s QuerySegmentsResponseDataRows) GoString() string {
	return s.String()
}

func (s QuerySegmentsResponseDataRows) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QuerySegmentsResponseDataRows) SetCreatedAt(v string) *QuerySegmentsResponseDataRows {
	s.CreatedAt = &v
	return s
}

func (s *QuerySegmentsResponseDataRows) SetHitCount(v int32) *QuerySegmentsResponseDataRows {
	s.HitCount = &v
	return s
}

func (s *QuerySegmentsResponseDataRows) SetAnswer(v string) *QuerySegmentsResponseDataRows {
	s.Answer = &v
	return s
}

func (s *QuerySegmentsResponseDataRows) SetWordCount(v int32) *QuerySegmentsResponseDataRows {
	s.WordCount = &v
	return s
}

func (s *QuerySegmentsResponseDataRows) SetCreatedBy(v string) *QuerySegmentsResponseDataRows {
	s.CreatedBy = &v
	return s
}

func (s *QuerySegmentsResponseDataRows) SetDatasetId(v string) *QuerySegmentsResponseDataRows {
	s.DatasetId = &v
	return s
}

func (s *QuerySegmentsResponseDataRows) SetDocumentId(v string) *QuerySegmentsResponseDataRows {
	s.DocumentId = &v
	return s
}

func (s *QuerySegmentsResponseDataRows) SetId(v string) *QuerySegmentsResponseDataRows {
	s.Id = &v
	return s
}

func (s *QuerySegmentsResponseDataRows) SetContent(v string) *QuerySegmentsResponseDataRows {
	s.Content = &v
	return s
}

func (s *QuerySegmentsResponseDataRows) SetEnabled(v bool) *QuerySegmentsResponseDataRows {
	s.Enabled = &v
	return s
}

func (s *QuerySegmentsResponseDataRows) SetContentHash(v string) *QuerySegmentsResponseDataRows {
	s.ContentHash = &v
	return s
}

func (s *QuerySegmentsResponseDataRows) SetStatus(v QuerySegmentsResponseDataRowsStatusEnum) *QuerySegmentsResponseDataRows {
	s.Status = &v
	return s
}

type QuerySegmentsResponseDataRowsBuilder struct {
	s *QuerySegmentsResponseDataRows
}

func NewQuerySegmentsResponseDataRowsBuilder() *QuerySegmentsResponseDataRowsBuilder {
	s := &QuerySegmentsResponseDataRows{}
	b := &QuerySegmentsResponseDataRowsBuilder{s: s}
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) CreatedAt(v string) *QuerySegmentsResponseDataRowsBuilder {
	b.s.CreatedAt = &v
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) HitCount(v int32) *QuerySegmentsResponseDataRowsBuilder {
	b.s.HitCount = &v
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) Answer(v string) *QuerySegmentsResponseDataRowsBuilder {
	b.s.Answer = &v
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) WordCount(v int32) *QuerySegmentsResponseDataRowsBuilder {
	b.s.WordCount = &v
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) CreatedBy(v string) *QuerySegmentsResponseDataRowsBuilder {
	b.s.CreatedBy = &v
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) DatasetId(v string) *QuerySegmentsResponseDataRowsBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) DocumentId(v string) *QuerySegmentsResponseDataRowsBuilder {
	b.s.DocumentId = &v
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) Id(v string) *QuerySegmentsResponseDataRowsBuilder {
	b.s.Id = &v
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) Content(v string) *QuerySegmentsResponseDataRowsBuilder {
	b.s.Content = &v
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) Enabled(v bool) *QuerySegmentsResponseDataRowsBuilder {
	b.s.Enabled = &v
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) ContentHash(v string) *QuerySegmentsResponseDataRowsBuilder {
	b.s.ContentHash = &v
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) Status(v QuerySegmentsResponseDataRowsStatusEnum) *QuerySegmentsResponseDataRowsBuilder {
	b.s.Status = &v
	return b
}

func (b *QuerySegmentsResponseDataRowsBuilder) Build() *QuerySegmentsResponseDataRows {
	return b.s
}
