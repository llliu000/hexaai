// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetDocumentResponseBodyDocumentFormEnum string

// List of DocumentForm
const (
	GetDocumentResponseBodyDocumentFormEnumQa   GetDocumentResponseBodyDocumentFormEnum = "QA"
	GetDocumentResponseBodyDocumentFormEnumText GetDocumentResponseBodyDocumentFormEnum = "TEXT"
)

type GetDocumentResponseBodyIndexingStatusEnum string

// List of IndexingStatus
const (
	GetDocumentResponseBodyIndexingStatusEnumWaiting          GetDocumentResponseBodyIndexingStatusEnum = "WAITING"
	GetDocumentResponseBodyIndexingStatusEnumQueuing          GetDocumentResponseBodyIndexingStatusEnum = "QUEUING"
	GetDocumentResponseBodyIndexingStatusEnumParsing          GetDocumentResponseBodyIndexingStatusEnum = "PARSING"
	GetDocumentResponseBodyIndexingStatusEnumCleaning         GetDocumentResponseBodyIndexingStatusEnum = "CLEANING"
	GetDocumentResponseBodyIndexingStatusEnumSplitting        GetDocumentResponseBodyIndexingStatusEnum = "SPLITTING"
	GetDocumentResponseBodyIndexingStatusEnumIndexing         GetDocumentResponseBodyIndexingStatusEnum = "INDEXING"
	GetDocumentResponseBodyIndexingStatusEnumCompleted        GetDocumentResponseBodyIndexingStatusEnum = "COMPLETED"
	GetDocumentResponseBodyIndexingStatusEnumExceeded         GetDocumentResponseBodyIndexingStatusEnum = "EXCEEDED"
	GetDocumentResponseBodyIndexingStatusEnumFilenamenotmatch GetDocumentResponseBodyIndexingStatusEnum = "FILENAMENOTMATCH"
	GetDocumentResponseBodyIndexingStatusEnumEmpty            GetDocumentResponseBodyIndexingStatusEnum = "EMPTY"
	GetDocumentResponseBodyIndexingStatusEnumError            GetDocumentResponseBodyIndexingStatusEnum = "ERROR"
	GetDocumentResponseBodyIndexingStatusEnumDeleting         GetDocumentResponseBodyIndexingStatusEnum = "DELETING"
	GetDocumentResponseBodyIndexingStatusEnumDeleted          GetDocumentResponseBodyIndexingStatusEnum = "DELETED"
)

type GetDocumentResponseBodyExtractModeEnum string

// List of ExtractMode
const (
	GetDocumentResponseBodyExtractModeEnumFast    GetDocumentResponseBodyExtractModeEnum = "FAST"
	GetDocumentResponseBodyExtractModeEnumPrecise GetDocumentResponseBodyExtractModeEnum = "PRECISE"
)

type GetDocumentResponseBodyProcessModeEnum string

// List of ProcessMode
const (
	GetDocumentResponseBodyProcessModeEnumAutomatic    GetDocumentResponseBodyProcessModeEnum = "AUTOMATIC"
	GetDocumentResponseBodyProcessModeEnumCustom       GetDocumentResponseBodyProcessModeEnum = "CUSTOM"
	GetDocumentResponseBodyProcessModeEnumByLine       GetDocumentResponseBodyProcessModeEnum = "BY_LINE"
	GetDocumentResponseBodyProcessModeEnumQaForm       GetDocumentResponseBodyProcessModeEnum = "QA_FORM"
	GetDocumentResponseBodyProcessModeEnumSemantic     GetDocumentResponseBodyProcessModeEnum = "SEMANTIC"
	GetDocumentResponseBodyProcessModeEnumSmart        GetDocumentResponseBodyProcessModeEnum = "SMART"
	GetDocumentResponseBodyProcessModeEnumHierarchical GetDocumentResponseBodyProcessModeEnum = "HIERARCHICAL"
	GetDocumentResponseBodyProcessModeEnumImageText    GetDocumentResponseBodyProcessModeEnum = "IMAGE_TEXT"
)

type GetDocumentResponseBody struct {

	// 清理完成时间
	CleaningCompletedAt *string `json:"cleaningCompletedAt,omitempty"`
	// 命中次数（文档下所有分段命中次数之和）
	HitCount *int32 `json:"hitCount,omitempty"`
	// 索引配置
	IndexConfig *GetDocumentResponseIndexConfig `json:"indexConfig,omitempty"`
	// 可对用户显示的错误详情
	Error *string `json:"error,omitempty"`
	// 禁用时间
	DisabledAt *string `json:"disabledAt,omitempty"`
	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`
	// 创建时间
	CreatedAt *string `json:"createdAt,omitempty"`
	// 文档形式（QA 拆分时传：QA，原始文档已为 QA 形式时不适用）
	DocumentForm *GetDocumentResponseBodyDocumentFormEnum `json:"documentForm,omitempty"`
	// 解析完成时间
	ParsingCompletedAt *string `json:"parsingCompletedAt,omitempty"`
	// 知识库id
	DatasetId *string `json:"datasetId,omitempty"`
	// 文档状态，排队中：WAITING,QUEUING, 解析中：PARSING,CLEANING,SPLITTING,INDEXING, 解析成功：COMPLETED, 超过限额：EXCEEDED, 文件类型不支持：FILENAMENOTMATCH, 空文档：EMPTY, 内部错误：ERROR, 删除中：DELETING 已删除：DELETED
	IndexingStatus *GetDocumentResponseBodyIndexingStatusEnum `json:"indexingStatus,omitempty"`
	// 文档id
	Id *string `json:"id,omitempty"`
	// 处理开始时间
	ProcessingStartedAt *string `json:"processingStartedAt,omitempty"`
	// 更新时间
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// 完成时间
	CompletedAt *string `json:"completedAt,omitempty"`
	// 更新人
	UpdatedBy *string `json:"updatedBy,omitempty"`
	// 字数（文档下所有分段字数之和）
	WordCount *int32 `json:"wordCount,omitempty"`
	// 解析规则，快速解析：FAST，精准解析：PRECISE
	ExtractMode *GetDocumentResponseBodyExtractModeEnum `json:"extractMode,omitempty"`
	// 分段数
	SegmentNum *int32 `json:"segmentNum,omitempty"`
	// 分割完成时间
	SplittingCompletedAt *string `json:"splittingCompletedAt,omitempty"`
	// 停止时间
	StoppedAt *string `json:"stoppedAt,omitempty"`
	// 禁用操作人
	DisabledBy *string `json:"disabledBy,omitempty"`
	// 创建人
	CreatedBy *string `json:"createdBy,omitempty"`
	// 文档名称
	Name *string `json:"name,omitempty"`
	// 文档的（分段索引）处理模式,默认切分：AUTOMATIC, 自定义切分：CUSTOM, 按行切分：BY_LINE, QA问答对：QA_FORM, 语义切分：SEMANTIC, 智能切分：SMART, 按层级切分：HIERARCHICAL, 多模态：IMAGE_TEXT
	ProcessMode *GetDocumentResponseBodyProcessModeEnum `json:"processMode,omitempty"`
}

func (s GetDocumentResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetDocumentResponseBody) GoString() string {
	return s.String()
}

func (s GetDocumentResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetDocumentResponseBody) SetCleaningCompletedAt(v string) *GetDocumentResponseBody {
	s.CleaningCompletedAt = &v
	return s
}

func (s *GetDocumentResponseBody) SetHitCount(v int32) *GetDocumentResponseBody {
	s.HitCount = &v
	return s
}

func (s *GetDocumentResponseBody) SetIndexConfig(v *GetDocumentResponseIndexConfig) *GetDocumentResponseBody {
	s.IndexConfig = v
	return s
}

func (s *GetDocumentResponseBody) SetError(v string) *GetDocumentResponseBody {
	s.Error = &v
	return s
}

func (s *GetDocumentResponseBody) SetDisabledAt(v string) *GetDocumentResponseBody {
	s.DisabledAt = &v
	return s
}

func (s *GetDocumentResponseBody) SetEnabled(v bool) *GetDocumentResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetDocumentResponseBody) SetCreatedAt(v string) *GetDocumentResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetDocumentResponseBody) SetDocumentForm(v GetDocumentResponseBodyDocumentFormEnum) *GetDocumentResponseBody {
	s.DocumentForm = &v
	return s
}

func (s *GetDocumentResponseBody) SetParsingCompletedAt(v string) *GetDocumentResponseBody {
	s.ParsingCompletedAt = &v
	return s
}

func (s *GetDocumentResponseBody) SetDatasetId(v string) *GetDocumentResponseBody {
	s.DatasetId = &v
	return s
}

func (s *GetDocumentResponseBody) SetIndexingStatus(v GetDocumentResponseBodyIndexingStatusEnum) *GetDocumentResponseBody {
	s.IndexingStatus = &v
	return s
}

func (s *GetDocumentResponseBody) SetId(v string) *GetDocumentResponseBody {
	s.Id = &v
	return s
}

func (s *GetDocumentResponseBody) SetProcessingStartedAt(v string) *GetDocumentResponseBody {
	s.ProcessingStartedAt = &v
	return s
}

func (s *GetDocumentResponseBody) SetUpdatedAt(v string) *GetDocumentResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetDocumentResponseBody) SetCompletedAt(v string) *GetDocumentResponseBody {
	s.CompletedAt = &v
	return s
}

func (s *GetDocumentResponseBody) SetUpdatedBy(v string) *GetDocumentResponseBody {
	s.UpdatedBy = &v
	return s
}

func (s *GetDocumentResponseBody) SetWordCount(v int32) *GetDocumentResponseBody {
	s.WordCount = &v
	return s
}

func (s *GetDocumentResponseBody) SetExtractMode(v GetDocumentResponseBodyExtractModeEnum) *GetDocumentResponseBody {
	s.ExtractMode = &v
	return s
}

func (s *GetDocumentResponseBody) SetSegmentNum(v int32) *GetDocumentResponseBody {
	s.SegmentNum = &v
	return s
}

func (s *GetDocumentResponseBody) SetSplittingCompletedAt(v string) *GetDocumentResponseBody {
	s.SplittingCompletedAt = &v
	return s
}

func (s *GetDocumentResponseBody) SetStoppedAt(v string) *GetDocumentResponseBody {
	s.StoppedAt = &v
	return s
}

func (s *GetDocumentResponseBody) SetDisabledBy(v string) *GetDocumentResponseBody {
	s.DisabledBy = &v
	return s
}

func (s *GetDocumentResponseBody) SetCreatedBy(v string) *GetDocumentResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *GetDocumentResponseBody) SetName(v string) *GetDocumentResponseBody {
	s.Name = &v
	return s
}

func (s *GetDocumentResponseBody) SetProcessMode(v GetDocumentResponseBodyProcessModeEnum) *GetDocumentResponseBody {
	s.ProcessMode = &v
	return s
}

type GetDocumentResponseBodyBuilder struct {
	s *GetDocumentResponseBody
}

func NewGetDocumentResponseBodyBuilder() *GetDocumentResponseBodyBuilder {
	s := &GetDocumentResponseBody{}
	b := &GetDocumentResponseBodyBuilder{s: s}
	return b
}

func (b *GetDocumentResponseBodyBuilder) CleaningCompletedAt(v string) *GetDocumentResponseBodyBuilder {
	b.s.CleaningCompletedAt = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) HitCount(v int32) *GetDocumentResponseBodyBuilder {
	b.s.HitCount = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) IndexConfig(v *GetDocumentResponseIndexConfig) *GetDocumentResponseBodyBuilder {
	b.s.IndexConfig = v
	return b
}

func (b *GetDocumentResponseBodyBuilder) Error(v string) *GetDocumentResponseBodyBuilder {
	b.s.Error = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) DisabledAt(v string) *GetDocumentResponseBodyBuilder {
	b.s.DisabledAt = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) Enabled(v bool) *GetDocumentResponseBodyBuilder {
	b.s.Enabled = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) CreatedAt(v string) *GetDocumentResponseBodyBuilder {
	b.s.CreatedAt = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) DocumentForm(v GetDocumentResponseBodyDocumentFormEnum) *GetDocumentResponseBodyBuilder {
	b.s.DocumentForm = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) ParsingCompletedAt(v string) *GetDocumentResponseBodyBuilder {
	b.s.ParsingCompletedAt = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) DatasetId(v string) *GetDocumentResponseBodyBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) IndexingStatus(v GetDocumentResponseBodyIndexingStatusEnum) *GetDocumentResponseBodyBuilder {
	b.s.IndexingStatus = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) Id(v string) *GetDocumentResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) ProcessingStartedAt(v string) *GetDocumentResponseBodyBuilder {
	b.s.ProcessingStartedAt = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) UpdatedAt(v string) *GetDocumentResponseBodyBuilder {
	b.s.UpdatedAt = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) CompletedAt(v string) *GetDocumentResponseBodyBuilder {
	b.s.CompletedAt = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) UpdatedBy(v string) *GetDocumentResponseBodyBuilder {
	b.s.UpdatedBy = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) WordCount(v int32) *GetDocumentResponseBodyBuilder {
	b.s.WordCount = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) ExtractMode(v GetDocumentResponseBodyExtractModeEnum) *GetDocumentResponseBodyBuilder {
	b.s.ExtractMode = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) SegmentNum(v int32) *GetDocumentResponseBodyBuilder {
	b.s.SegmentNum = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) SplittingCompletedAt(v string) *GetDocumentResponseBodyBuilder {
	b.s.SplittingCompletedAt = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) StoppedAt(v string) *GetDocumentResponseBodyBuilder {
	b.s.StoppedAt = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) DisabledBy(v string) *GetDocumentResponseBodyBuilder {
	b.s.DisabledBy = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) CreatedBy(v string) *GetDocumentResponseBodyBuilder {
	b.s.CreatedBy = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) Name(v string) *GetDocumentResponseBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) ProcessMode(v GetDocumentResponseBodyProcessModeEnum) *GetDocumentResponseBodyBuilder {
	b.s.ProcessMode = &v
	return b
}

func (b *GetDocumentResponseBodyBuilder) Build() *GetDocumentResponseBody {
	return b.s
}
