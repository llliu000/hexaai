// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateSegmentPath struct {
	position.Path
	// 分段Id（URL路径参数）
	SegmentId *string `json:"segmentId,omitempty"`
	// 知识库Id（URL路径参数）
	DatasetId *string `json:"datasetId,omitempty"`
	// 文档Id（URL路径参数）
	DocumentId *string `json:"documentId,omitempty"`
}

func (s UpdateSegmentPath) String() string {
	return utils.Beautify(s)
}

func (s UpdateSegmentPath) GoString() string {
	return s.String()
}

func (s UpdateSegmentPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateSegmentPath) SetSegmentId(v string) *UpdateSegmentPath {
	s.SegmentId = &v
	return s
}

func (s *UpdateSegmentPath) SetDatasetId(v string) *UpdateSegmentPath {
	s.DatasetId = &v
	return s
}

func (s *UpdateSegmentPath) SetDocumentId(v string) *UpdateSegmentPath {
	s.DocumentId = &v
	return s
}

type UpdateSegmentPathBuilder struct {
	s *UpdateSegmentPath
}

func NewUpdateSegmentPathBuilder() *UpdateSegmentPathBuilder {
	s := &UpdateSegmentPath{}
	b := &UpdateSegmentPathBuilder{s: s}
	return b
}

func (b *UpdateSegmentPathBuilder) SegmentId(v string) *UpdateSegmentPathBuilder {
	b.s.SegmentId = &v
	return b
}

func (b *UpdateSegmentPathBuilder) DatasetId(v string) *UpdateSegmentPathBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *UpdateSegmentPathBuilder) DocumentId(v string) *UpdateSegmentPathBuilder {
	b.s.DocumentId = &v
	return b
}

func (b *UpdateSegmentPathBuilder) Build() *UpdateSegmentPath {
	return b.s
}
