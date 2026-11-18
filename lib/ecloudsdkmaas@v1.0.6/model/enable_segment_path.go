// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EnableSegmentPath struct {
	position.Path
	// 分段Id（URL路径参数）
	SegmentId *string `json:"segmentId,omitempty"`
	// 知识库Id（URL路径参数）
	DatasetId *string `json:"datasetId,omitempty"`
	// 文档Id（URL路径参数）
	DocumentId *string `json:"documentId,omitempty"`
}

func (s EnableSegmentPath) String() string {
	return utils.Beautify(s)
}

func (s EnableSegmentPath) GoString() string {
	return s.String()
}

func (s EnableSegmentPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EnableSegmentPath) SetSegmentId(v string) *EnableSegmentPath {
	s.SegmentId = &v
	return s
}

func (s *EnableSegmentPath) SetDatasetId(v string) *EnableSegmentPath {
	s.DatasetId = &v
	return s
}

func (s *EnableSegmentPath) SetDocumentId(v string) *EnableSegmentPath {
	s.DocumentId = &v
	return s
}

type EnableSegmentPathBuilder struct {
	s *EnableSegmentPath
}

func NewEnableSegmentPathBuilder() *EnableSegmentPathBuilder {
	s := &EnableSegmentPath{}
	b := &EnableSegmentPathBuilder{s: s}
	return b
}

func (b *EnableSegmentPathBuilder) SegmentId(v string) *EnableSegmentPathBuilder {
	b.s.SegmentId = &v
	return b
}

func (b *EnableSegmentPathBuilder) DatasetId(v string) *EnableSegmentPathBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *EnableSegmentPathBuilder) DocumentId(v string) *EnableSegmentPathBuilder {
	b.s.DocumentId = &v
	return b
}

func (b *EnableSegmentPathBuilder) Build() *EnableSegmentPath {
	return b.s
}
