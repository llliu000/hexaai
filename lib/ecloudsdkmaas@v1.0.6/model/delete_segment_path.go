// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteSegmentPath struct {
	position.Path
	// 分段Id（URL路径参数）
	SegmentId *string `json:"segmentId,omitempty"`
	// 知识库Id（URL路径参数）
	DatasetId *string `json:"datasetId,omitempty"`
	// 文档Id（URL路径参数）
	DocumentId *string `json:"documentId,omitempty"`
}

func (s DeleteSegmentPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteSegmentPath) GoString() string {
	return s.String()
}

func (s DeleteSegmentPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteSegmentPath) SetSegmentId(v string) *DeleteSegmentPath {
	s.SegmentId = &v
	return s
}

func (s *DeleteSegmentPath) SetDatasetId(v string) *DeleteSegmentPath {
	s.DatasetId = &v
	return s
}

func (s *DeleteSegmentPath) SetDocumentId(v string) *DeleteSegmentPath {
	s.DocumentId = &v
	return s
}

type DeleteSegmentPathBuilder struct {
	s *DeleteSegmentPath
}

func NewDeleteSegmentPathBuilder() *DeleteSegmentPathBuilder {
	s := &DeleteSegmentPath{}
	b := &DeleteSegmentPathBuilder{s: s}
	return b
}

func (b *DeleteSegmentPathBuilder) SegmentId(v string) *DeleteSegmentPathBuilder {
	b.s.SegmentId = &v
	return b
}

func (b *DeleteSegmentPathBuilder) DatasetId(v string) *DeleteSegmentPathBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *DeleteSegmentPathBuilder) DocumentId(v string) *DeleteSegmentPathBuilder {
	b.s.DocumentId = &v
	return b
}

func (b *DeleteSegmentPathBuilder) Build() *DeleteSegmentPath {
	return b.s
}
