// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetSegmentPath struct {
	position.Path
	// 分段Id（URL路径参数）
	SegmentId *string `json:"segmentId,omitempty"`
	// 知识库Id（URL路径参数）
	DatasetId *string `json:"datasetId,omitempty"`
	// 文档Id（URL路径参数）
	DocumentId *string `json:"documentId,omitempty"`
}

func (s GetSegmentPath) String() string {
	return utils.Beautify(s)
}

func (s GetSegmentPath) GoString() string {
	return s.String()
}

func (s GetSegmentPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSegmentPath) SetSegmentId(v string) *GetSegmentPath {
	s.SegmentId = &v
	return s
}

func (s *GetSegmentPath) SetDatasetId(v string) *GetSegmentPath {
	s.DatasetId = &v
	return s
}

func (s *GetSegmentPath) SetDocumentId(v string) *GetSegmentPath {
	s.DocumentId = &v
	return s
}

type GetSegmentPathBuilder struct {
	s *GetSegmentPath
}

func NewGetSegmentPathBuilder() *GetSegmentPathBuilder {
	s := &GetSegmentPath{}
	b := &GetSegmentPathBuilder{s: s}
	return b
}

func (b *GetSegmentPathBuilder) SegmentId(v string) *GetSegmentPathBuilder {
	b.s.SegmentId = &v
	return b
}

func (b *GetSegmentPathBuilder) DatasetId(v string) *GetSegmentPathBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *GetSegmentPathBuilder) DocumentId(v string) *GetSegmentPathBuilder {
	b.s.DocumentId = &v
	return b
}

func (b *GetSegmentPathBuilder) Build() *GetSegmentPath {
	return b.s
}
