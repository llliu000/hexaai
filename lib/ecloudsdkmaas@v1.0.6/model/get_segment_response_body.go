// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetSegmentResponseBodyStatusEnum string

// List of Status
const (
	GetSegmentResponseBodyStatusEnumWaiting          GetSegmentResponseBodyStatusEnum = "WAITING"
	GetSegmentResponseBodyStatusEnumQueuing          GetSegmentResponseBodyStatusEnum = "QUEUING"
	GetSegmentResponseBodyStatusEnumParsing          GetSegmentResponseBodyStatusEnum = "PARSING"
	GetSegmentResponseBodyStatusEnumCleaning         GetSegmentResponseBodyStatusEnum = "CLEANING"
	GetSegmentResponseBodyStatusEnumSplitting        GetSegmentResponseBodyStatusEnum = "SPLITTING"
	GetSegmentResponseBodyStatusEnumIndexing         GetSegmentResponseBodyStatusEnum = "INDEXING"
	GetSegmentResponseBodyStatusEnumCompleted        GetSegmentResponseBodyStatusEnum = "COMPLETED"
	GetSegmentResponseBodyStatusEnumExceeded         GetSegmentResponseBodyStatusEnum = "EXCEEDED"
	GetSegmentResponseBodyStatusEnumFilenamenotmatch GetSegmentResponseBodyStatusEnum = "FILENAMENOTMATCH"
	GetSegmentResponseBodyStatusEnumEmpty            GetSegmentResponseBodyStatusEnum = "EMPTY"
	GetSegmentResponseBodyStatusEnumError            GetSegmentResponseBodyStatusEnum = "ERROR"
	GetSegmentResponseBodyStatusEnumDeleting         GetSegmentResponseBodyStatusEnum = "DELETING"
	GetSegmentResponseBodyStatusEnumDeleted          GetSegmentResponseBodyStatusEnum = "DELETED"
)

type GetSegmentResponseBody struct {

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
	Status *GetSegmentResponseBodyStatusEnum `json:"status,omitempty"`
	// 更新时间
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

func (s GetSegmentResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetSegmentResponseBody) GoString() string {
	return s.String()
}

func (s GetSegmentResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSegmentResponseBody) SetCompletedAt(v string) *GetSegmentResponseBody {
	s.CompletedAt = &v
	return s
}

func (s *GetSegmentResponseBody) SetHitCount(v int32) *GetSegmentResponseBody {
	s.HitCount = &v
	return s
}

func (s *GetSegmentResponseBody) SetUpdatedBy(v string) *GetSegmentResponseBody {
	s.UpdatedBy = &v
	return s
}

func (s *GetSegmentResponseBody) SetWordCount(v int32) *GetSegmentResponseBody {
	s.WordCount = &v
	return s
}

func (s *GetSegmentResponseBody) SetError(v string) *GetSegmentResponseBody {
	s.Error = &v
	return s
}

func (s *GetSegmentResponseBody) SetContent(v string) *GetSegmentResponseBody {
	s.Content = &v
	return s
}

func (s *GetSegmentResponseBody) SetEnabled(v bool) *GetSegmentResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetSegmentResponseBody) SetDisabledAt(v string) *GetSegmentResponseBody {
	s.DisabledAt = &v
	return s
}

func (s *GetSegmentResponseBody) SetContentHash(v string) *GetSegmentResponseBody {
	s.ContentHash = &v
	return s
}

func (s *GetSegmentResponseBody) SetCreatedAt(v string) *GetSegmentResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetSegmentResponseBody) SetIndexingAt(v string) *GetSegmentResponseBody {
	s.IndexingAt = &v
	return s
}

func (s *GetSegmentResponseBody) SetDisabledBy(v string) *GetSegmentResponseBody {
	s.DisabledBy = &v
	return s
}

func (s *GetSegmentResponseBody) SetStoppedAt(v string) *GetSegmentResponseBody {
	s.StoppedAt = &v
	return s
}

func (s *GetSegmentResponseBody) SetAnswer(v string) *GetSegmentResponseBody {
	s.Answer = &v
	return s
}

func (s *GetSegmentResponseBody) SetCreatedBy(v string) *GetSegmentResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *GetSegmentResponseBody) SetDatasetId(v string) *GetSegmentResponseBody {
	s.DatasetId = &v
	return s
}

func (s *GetSegmentResponseBody) SetDocumentId(v string) *GetSegmentResponseBody {
	s.DocumentId = &v
	return s
}

func (s *GetSegmentResponseBody) SetId(v string) *GetSegmentResponseBody {
	s.Id = &v
	return s
}

func (s *GetSegmentResponseBody) SetStatus(v GetSegmentResponseBodyStatusEnum) *GetSegmentResponseBody {
	s.Status = &v
	return s
}

func (s *GetSegmentResponseBody) SetUpdatedAt(v string) *GetSegmentResponseBody {
	s.UpdatedAt = &v
	return s
}

type GetSegmentResponseBodyBuilder struct {
	s *GetSegmentResponseBody
}

func NewGetSegmentResponseBodyBuilder() *GetSegmentResponseBodyBuilder {
	s := &GetSegmentResponseBody{}
	b := &GetSegmentResponseBodyBuilder{s: s}
	return b
}

func (b *GetSegmentResponseBodyBuilder) CompletedAt(v string) *GetSegmentResponseBodyBuilder {
	b.s.CompletedAt = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) HitCount(v int32) *GetSegmentResponseBodyBuilder {
	b.s.HitCount = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) UpdatedBy(v string) *GetSegmentResponseBodyBuilder {
	b.s.UpdatedBy = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) WordCount(v int32) *GetSegmentResponseBodyBuilder {
	b.s.WordCount = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) Error(v string) *GetSegmentResponseBodyBuilder {
	b.s.Error = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) Content(v string) *GetSegmentResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) Enabled(v bool) *GetSegmentResponseBodyBuilder {
	b.s.Enabled = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) DisabledAt(v string) *GetSegmentResponseBodyBuilder {
	b.s.DisabledAt = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) ContentHash(v string) *GetSegmentResponseBodyBuilder {
	b.s.ContentHash = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) CreatedAt(v string) *GetSegmentResponseBodyBuilder {
	b.s.CreatedAt = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) IndexingAt(v string) *GetSegmentResponseBodyBuilder {
	b.s.IndexingAt = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) DisabledBy(v string) *GetSegmentResponseBodyBuilder {
	b.s.DisabledBy = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) StoppedAt(v string) *GetSegmentResponseBodyBuilder {
	b.s.StoppedAt = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) Answer(v string) *GetSegmentResponseBodyBuilder {
	b.s.Answer = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) CreatedBy(v string) *GetSegmentResponseBodyBuilder {
	b.s.CreatedBy = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) DatasetId(v string) *GetSegmentResponseBodyBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) DocumentId(v string) *GetSegmentResponseBodyBuilder {
	b.s.DocumentId = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) Id(v string) *GetSegmentResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) Status(v GetSegmentResponseBodyStatusEnum) *GetSegmentResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) UpdatedAt(v string) *GetSegmentResponseBodyBuilder {
	b.s.UpdatedAt = &v
	return b
}

func (b *GetSegmentResponseBodyBuilder) Build() *GetSegmentResponseBody {
	return b.s
}
