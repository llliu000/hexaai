// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetDocumentResponseIndexConfigDocumentFormEnum string

// List of DocumentForm
const (
	GetDocumentResponseIndexConfigDocumentFormEnumQa   GetDocumentResponseIndexConfigDocumentFormEnum = "QA"
	GetDocumentResponseIndexConfigDocumentFormEnumText GetDocumentResponseIndexConfigDocumentFormEnum = "TEXT"
)

type GetDocumentResponseIndexConfigExtractModeEnum string

// List of ExtractMode
const (
	GetDocumentResponseIndexConfigExtractModeEnumFast    GetDocumentResponseIndexConfigExtractModeEnum = "FAST"
	GetDocumentResponseIndexConfigExtractModeEnumPrecise GetDocumentResponseIndexConfigExtractModeEnum = "PRECISE"
)

type GetDocumentResponseIndexConfigProcessModeEnum string

// List of ProcessMode
const (
	GetDocumentResponseIndexConfigProcessModeEnumAutomatic    GetDocumentResponseIndexConfigProcessModeEnum = "AUTOMATIC"
	GetDocumentResponseIndexConfigProcessModeEnumCustom       GetDocumentResponseIndexConfigProcessModeEnum = "CUSTOM"
	GetDocumentResponseIndexConfigProcessModeEnumByLine       GetDocumentResponseIndexConfigProcessModeEnum = "BY_LINE"
	GetDocumentResponseIndexConfigProcessModeEnumQaForm       GetDocumentResponseIndexConfigProcessModeEnum = "QA_FORM"
	GetDocumentResponseIndexConfigProcessModeEnumSemantic     GetDocumentResponseIndexConfigProcessModeEnum = "SEMANTIC"
	GetDocumentResponseIndexConfigProcessModeEnumSmart        GetDocumentResponseIndexConfigProcessModeEnum = "SMART"
	GetDocumentResponseIndexConfigProcessModeEnumHierarchical GetDocumentResponseIndexConfigProcessModeEnum = "HIERARCHICAL"
	GetDocumentResponseIndexConfigProcessModeEnumImageText    GetDocumentResponseIndexConfigProcessModeEnum = "IMAGE_TEXT"
)

type GetDocumentResponseIndexConfig struct {

	// 文档形式（QA 拆分时传：QA，原始文档已为 QA 形式时不适用）
	DocumentForm *GetDocumentResponseIndexConfigDocumentFormEnum `json:"documentForm,omitempty"`
	// 解析规则，快速解析：FAST，精准解析：PRECISE
	ExtractMode *GetDocumentResponseIndexConfigExtractModeEnum `json:"extractMode,omitempty"`
	// 分段规则
	Segmentation *GetDocumentResponseSegmentation `json:"segmentation,omitempty"`
	// 文档的（分段索引）处理模式,默认切分：AUTOMATIC, 自定义切分：CUSTOM, 按行切分：BY_LINE, QA问答对：QA_FORM, 语义切分：SEMANTIC, 智能切分：SMART, 按层级切分：HIERARCHICAL, 多模态：IMAGE_TEXT
	ProcessMode *GetDocumentResponseIndexConfigProcessModeEnum `json:"processMode,omitempty"`
}

func (s GetDocumentResponseIndexConfig) String() string {
	return utils.Beautify(s)
}

func (s GetDocumentResponseIndexConfig) GoString() string {
	return s.String()
}

func (s GetDocumentResponseIndexConfig) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetDocumentResponseIndexConfig) SetDocumentForm(v GetDocumentResponseIndexConfigDocumentFormEnum) *GetDocumentResponseIndexConfig {
	s.DocumentForm = &v
	return s
}

func (s *GetDocumentResponseIndexConfig) SetExtractMode(v GetDocumentResponseIndexConfigExtractModeEnum) *GetDocumentResponseIndexConfig {
	s.ExtractMode = &v
	return s
}

func (s *GetDocumentResponseIndexConfig) SetSegmentation(v *GetDocumentResponseSegmentation) *GetDocumentResponseIndexConfig {
	s.Segmentation = v
	return s
}

func (s *GetDocumentResponseIndexConfig) SetProcessMode(v GetDocumentResponseIndexConfigProcessModeEnum) *GetDocumentResponseIndexConfig {
	s.ProcessMode = &v
	return s
}

type GetDocumentResponseIndexConfigBuilder struct {
	s *GetDocumentResponseIndexConfig
}

func NewGetDocumentResponseIndexConfigBuilder() *GetDocumentResponseIndexConfigBuilder {
	s := &GetDocumentResponseIndexConfig{}
	b := &GetDocumentResponseIndexConfigBuilder{s: s}
	return b
}

func (b *GetDocumentResponseIndexConfigBuilder) DocumentForm(v GetDocumentResponseIndexConfigDocumentFormEnum) *GetDocumentResponseIndexConfigBuilder {
	b.s.DocumentForm = &v
	return b
}

func (b *GetDocumentResponseIndexConfigBuilder) ExtractMode(v GetDocumentResponseIndexConfigExtractModeEnum) *GetDocumentResponseIndexConfigBuilder {
	b.s.ExtractMode = &v
	return b
}

func (b *GetDocumentResponseIndexConfigBuilder) Segmentation(v *GetDocumentResponseSegmentation) *GetDocumentResponseIndexConfigBuilder {
	b.s.Segmentation = v
	return b
}

func (b *GetDocumentResponseIndexConfigBuilder) ProcessMode(v GetDocumentResponseIndexConfigProcessModeEnum) *GetDocumentResponseIndexConfigBuilder {
	b.s.ProcessMode = &v
	return b
}

func (b *GetDocumentResponseIndexConfigBuilder) Build() *GetDocumentResponseIndexConfig {
	return b.s
}
