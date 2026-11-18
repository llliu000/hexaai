// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateSegmentResponseBodyStatusEnum string

// List of Status
const (
	UpdateSegmentResponseBodyStatusEnumWaiting          UpdateSegmentResponseBodyStatusEnum = "WAITING"
	UpdateSegmentResponseBodyStatusEnumQueuing          UpdateSegmentResponseBodyStatusEnum = "QUEUING"
	UpdateSegmentResponseBodyStatusEnumParsing          UpdateSegmentResponseBodyStatusEnum = "PARSING"
	UpdateSegmentResponseBodyStatusEnumCleaning         UpdateSegmentResponseBodyStatusEnum = "CLEANING"
	UpdateSegmentResponseBodyStatusEnumSplitting        UpdateSegmentResponseBodyStatusEnum = "SPLITTING"
	UpdateSegmentResponseBodyStatusEnumIndexing         UpdateSegmentResponseBodyStatusEnum = "INDEXING"
	UpdateSegmentResponseBodyStatusEnumCompleted        UpdateSegmentResponseBodyStatusEnum = "COMPLETED"
	UpdateSegmentResponseBodyStatusEnumExceeded         UpdateSegmentResponseBodyStatusEnum = "EXCEEDED"
	UpdateSegmentResponseBodyStatusEnumFilenamenotmatch UpdateSegmentResponseBodyStatusEnum = "FILENAMENOTMATCH"
	UpdateSegmentResponseBodyStatusEnumEmpty            UpdateSegmentResponseBodyStatusEnum = "EMPTY"
	UpdateSegmentResponseBodyStatusEnumError            UpdateSegmentResponseBodyStatusEnum = "ERROR"
	UpdateSegmentResponseBodyStatusEnumDeleting         UpdateSegmentResponseBodyStatusEnum = "DELETING"
	UpdateSegmentResponseBodyStatusEnumDeleted          UpdateSegmentResponseBodyStatusEnum = "DELETED"
)

type UpdateSegmentResponseBody struct {

	// 完成时间
	CompletedAt *string `json:"completedAt,omitempty"`
	// 召回命中次数
	HitCount *int32 `json:"hitCount,omitempty"`
	// 更新人
	UpdatedBy *string `json:"updatedBy,omitempty"`
	// 字数
	WordCount *int32 `json:"wordCount,omitempty"`
	// 如果发生错误，可能有错误提示
	Error *string `json:"error,omitempty"`
	// 文档内容
	Content *string `json:"content,omitempty"`
	// 是否启用，true启用，false禁用
	Enabled *bool `json:"enabled,omitempty"`
	// 禁用时间
	DisabledAt *string `json:"disabledAt,omitempty"`
	// 文档hash值
	ContentHash *string `json:"contentHash,omitempty"`
	// 创建时间
	CreatedAt *string `json:"createdAt,omitempty"`
	// 索引时间
	IndexingAt *string `json:"indexingAt,omitempty"`
	// 禁用操作人
	DisabledBy *string `json:"disabledBy,omitempty"`
	// 停止时间
	StoppedAt *string `json:"stoppedAt,omitempty"`
	// QA问答对的回答
	Answer *string `json:"answer,omitempty"`
	// 创建人
	CreatedBy *string `json:"createdBy,omitempty"`
	// 知识库id
	DatasetId *string `json:"datasetId,omitempty"`
	// 文档id
	DocumentId *string `json:"documentId,omitempty"`
	// 分段id
	Id *string `json:"id,omitempty"`
	// 文档状态，排队中：WAITING,QUEUING, 解析中：PARSING,CLEANING,SPLITTING,INDEXING, 解析成功：COMPLETED, 超过限额：EXCEEDED, 文件类型不支持：FILENAMENOTMATCH, 空文档：EMPTY, 内部错误：ERROR, 删除中：DELETING 已删除：DELETED
	Status *UpdateSegmentResponseBodyStatusEnum `json:"status,omitempty"`
	// 更新时间
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

func (s UpdateSegmentResponseBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateSegmentResponseBody) GoString() string {
	return s.String()
}

func (s UpdateSegmentResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateSegmentResponseBody) SetCompletedAt(v string) *UpdateSegmentResponseBody {
	s.CompletedAt = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetHitCount(v int32) *UpdateSegmentResponseBody {
	s.HitCount = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetUpdatedBy(v string) *UpdateSegmentResponseBody {
	s.UpdatedBy = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetWordCount(v int32) *UpdateSegmentResponseBody {
	s.WordCount = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetError(v string) *UpdateSegmentResponseBody {
	s.Error = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetContent(v string) *UpdateSegmentResponseBody {
	s.Content = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetEnabled(v bool) *UpdateSegmentResponseBody {
	s.Enabled = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetDisabledAt(v string) *UpdateSegmentResponseBody {
	s.DisabledAt = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetContentHash(v string) *UpdateSegmentResponseBody {
	s.ContentHash = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetCreatedAt(v string) *UpdateSegmentResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetIndexingAt(v string) *UpdateSegmentResponseBody {
	s.IndexingAt = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetDisabledBy(v string) *UpdateSegmentResponseBody {
	s.DisabledBy = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetStoppedAt(v string) *UpdateSegmentResponseBody {
	s.StoppedAt = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetAnswer(v string) *UpdateSegmentResponseBody {
	s.Answer = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetCreatedBy(v string) *UpdateSegmentResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetDatasetId(v string) *UpdateSegmentResponseBody {
	s.DatasetId = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetDocumentId(v string) *UpdateSegmentResponseBody {
	s.DocumentId = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetId(v string) *UpdateSegmentResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetStatus(v UpdateSegmentResponseBodyStatusEnum) *UpdateSegmentResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateSegmentResponseBody) SetUpdatedAt(v string) *UpdateSegmentResponseBody {
	s.UpdatedAt = &v
	return s
}

type UpdateSegmentResponseBodyBuilder struct {
	s *UpdateSegmentResponseBody
}

func NewUpdateSegmentResponseBodyBuilder() *UpdateSegmentResponseBodyBuilder {
	s := &UpdateSegmentResponseBody{}
	b := &UpdateSegmentResponseBodyBuilder{s: s}
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) CompletedAt(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.CompletedAt = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) HitCount(v int32) *UpdateSegmentResponseBodyBuilder {
	b.s.HitCount = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) UpdatedBy(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.UpdatedBy = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) WordCount(v int32) *UpdateSegmentResponseBodyBuilder {
	b.s.WordCount = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) Error(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.Error = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) Content(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) Enabled(v bool) *UpdateSegmentResponseBodyBuilder {
	b.s.Enabled = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) DisabledAt(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.DisabledAt = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) ContentHash(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.ContentHash = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) CreatedAt(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.CreatedAt = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) IndexingAt(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.IndexingAt = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) DisabledBy(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.DisabledBy = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) StoppedAt(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.StoppedAt = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) Answer(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.Answer = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) CreatedBy(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.CreatedBy = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) DatasetId(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) DocumentId(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.DocumentId = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) Id(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) Status(v UpdateSegmentResponseBodyStatusEnum) *UpdateSegmentResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) UpdatedAt(v string) *UpdateSegmentResponseBodyBuilder {
	b.s.UpdatedAt = &v
	return b
}

func (b *UpdateSegmentResponseBodyBuilder) Build() *UpdateSegmentResponseBody {
	return b.s
}
