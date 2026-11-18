// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QuerySegmentsPath struct {
	position.Path
	// 知识库Id（URL路径参数）
	DatasetId *string `json:"datasetId,omitempty"`
	// 文档Id（URL路径参数）
	DocumentId *string `json:"documentId,omitempty"`
}

func (s QuerySegmentsPath) String() string {
	return utils.Beautify(s)
}

func (s QuerySegmentsPath) GoString() string {
	return s.String()
}

func (s QuerySegmentsPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QuerySegmentsPath) SetDatasetId(v string) *QuerySegmentsPath {
	s.DatasetId = &v
	return s
}

func (s *QuerySegmentsPath) SetDocumentId(v string) *QuerySegmentsPath {
	s.DocumentId = &v
	return s
}

type QuerySegmentsPathBuilder struct {
	s *QuerySegmentsPath
}

func NewQuerySegmentsPathBuilder() *QuerySegmentsPathBuilder {
	s := &QuerySegmentsPath{}
	b := &QuerySegmentsPathBuilder{s: s}
	return b
}

func (b *QuerySegmentsPathBuilder) DatasetId(v string) *QuerySegmentsPathBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *QuerySegmentsPathBuilder) DocumentId(v string) *QuerySegmentsPathBuilder {
	b.s.DocumentId = &v
	return b
}

func (b *QuerySegmentsPathBuilder) Build() *QuerySegmentsPath {
	return b.s
}
