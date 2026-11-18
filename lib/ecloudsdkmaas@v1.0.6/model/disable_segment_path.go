// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DisableSegmentPath struct {
	position.Path
	// 分段Id（URL路径参数）
	SegmentId *string `json:"segmentId,omitempty"`
	// 知识库Id（URL路径参数）
	DatasetId *string `json:"datasetId,omitempty"`
	// 文档Id（URL路径参数）
	DocumentId *string `json:"documentId,omitempty"`
}

func (s DisableSegmentPath) String() string {
	return utils.Beautify(s)
}

func (s DisableSegmentPath) GoString() string {
	return s.String()
}

func (s DisableSegmentPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DisableSegmentPath) SetSegmentId(v string) *DisableSegmentPath {
	s.SegmentId = &v
	return s
}

func (s *DisableSegmentPath) SetDatasetId(v string) *DisableSegmentPath {
	s.DatasetId = &v
	return s
}

func (s *DisableSegmentPath) SetDocumentId(v string) *DisableSegmentPath {
	s.DocumentId = &v
	return s
}

type DisableSegmentPathBuilder struct {
	s *DisableSegmentPath
}

func NewDisableSegmentPathBuilder() *DisableSegmentPathBuilder {
	s := &DisableSegmentPath{}
	b := &DisableSegmentPathBuilder{s: s}
	return b
}

func (b *DisableSegmentPathBuilder) SegmentId(v string) *DisableSegmentPathBuilder {
	b.s.SegmentId = &v
	return b
}

func (b *DisableSegmentPathBuilder) DatasetId(v string) *DisableSegmentPathBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *DisableSegmentPathBuilder) DocumentId(v string) *DisableSegmentPathBuilder {
	b.s.DocumentId = &v
	return b
}

func (b *DisableSegmentPathBuilder) Build() *DisableSegmentPath {
	return b.s
}
