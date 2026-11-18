// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetDocumentResponseSegmentation struct {

	// 分段重叠长度（设置分段之间的重叠长度可以保留分段之间的语义关系，提升召回效果。建议设置为最大分段长度的10%-25%）
	ChunkOverlap *int32 `json:"chunkOverlap,omitempty"`
	// 语义切分-标题增强。
	TitleEnhance *bool `json:"titleEnhance,omitempty"`
	// 分段最大长度
	MaxTokens *int32 `json:"maxTokens,omitempty"`
	// 分段标识符
	Separator *string `json:"separator,omitempty"`
	// 语义切分-断点阈值。
	BreakpointThreshold *int32 `json:"breakpointThreshold,omitempty"`
	// 语义切分-窗口大小
	BufferSize *int32 `json:"bufferSize,omitempty"`
}

func (s GetDocumentResponseSegmentation) String() string {
	return utils.Beautify(s)
}

func (s GetDocumentResponseSegmentation) GoString() string {
	return s.String()
}

func (s GetDocumentResponseSegmentation) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetDocumentResponseSegmentation) SetChunkOverlap(v int32) *GetDocumentResponseSegmentation {
	s.ChunkOverlap = &v
	return s
}

func (s *GetDocumentResponseSegmentation) SetTitleEnhance(v bool) *GetDocumentResponseSegmentation {
	s.TitleEnhance = &v
	return s
}

func (s *GetDocumentResponseSegmentation) SetMaxTokens(v int32) *GetDocumentResponseSegmentation {
	s.MaxTokens = &v
	return s
}

func (s *GetDocumentResponseSegmentation) SetSeparator(v string) *GetDocumentResponseSegmentation {
	s.Separator = &v
	return s
}

func (s *GetDocumentResponseSegmentation) SetBreakpointThreshold(v int32) *GetDocumentResponseSegmentation {
	s.BreakpointThreshold = &v
	return s
}

func (s *GetDocumentResponseSegmentation) SetBufferSize(v int32) *GetDocumentResponseSegmentation {
	s.BufferSize = &v
	return s
}

type GetDocumentResponseSegmentationBuilder struct {
	s *GetDocumentResponseSegmentation
}

func NewGetDocumentResponseSegmentationBuilder() *GetDocumentResponseSegmentationBuilder {
	s := &GetDocumentResponseSegmentation{}
	b := &GetDocumentResponseSegmentationBuilder{s: s}
	return b
}

func (b *GetDocumentResponseSegmentationBuilder) ChunkOverlap(v int32) *GetDocumentResponseSegmentationBuilder {
	b.s.ChunkOverlap = &v
	return b
}

func (b *GetDocumentResponseSegmentationBuilder) TitleEnhance(v bool) *GetDocumentResponseSegmentationBuilder {
	b.s.TitleEnhance = &v
	return b
}

func (b *GetDocumentResponseSegmentationBuilder) MaxTokens(v int32) *GetDocumentResponseSegmentationBuilder {
	b.s.MaxTokens = &v
	return b
}

func (b *GetDocumentResponseSegmentationBuilder) Separator(v string) *GetDocumentResponseSegmentationBuilder {
	b.s.Separator = &v
	return b
}

func (b *GetDocumentResponseSegmentationBuilder) BreakpointThreshold(v int32) *GetDocumentResponseSegmentationBuilder {
	b.s.BreakpointThreshold = &v
	return b
}

func (b *GetDocumentResponseSegmentationBuilder) BufferSize(v int32) *GetDocumentResponseSegmentationBuilder {
	b.s.BufferSize = &v
	return b
}

func (b *GetDocumentResponseSegmentationBuilder) Build() *GetDocumentResponseSegmentation {
	return b.s
}
