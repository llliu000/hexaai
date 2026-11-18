// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EnableDocumentPath struct {
	position.Path
	// 知识库Id（URL路径参数）
	DatasetId *string `json:"datasetId,omitempty"`
	// 文档Id（URL路径参数）
	DocumentId *string `json:"documentId,omitempty"`
}

func (s EnableDocumentPath) String() string {
	return utils.Beautify(s)
}

func (s EnableDocumentPath) GoString() string {
	return s.String()
}

func (s EnableDocumentPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EnableDocumentPath) SetDatasetId(v string) *EnableDocumentPath {
	s.DatasetId = &v
	return s
}

func (s *EnableDocumentPath) SetDocumentId(v string) *EnableDocumentPath {
	s.DocumentId = &v
	return s
}

type EnableDocumentPathBuilder struct {
	s *EnableDocumentPath
}

func NewEnableDocumentPathBuilder() *EnableDocumentPathBuilder {
	s := &EnableDocumentPath{}
	b := &EnableDocumentPathBuilder{s: s}
	return b
}

func (b *EnableDocumentPathBuilder) DatasetId(v string) *EnableDocumentPathBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *EnableDocumentPathBuilder) DocumentId(v string) *EnableDocumentPathBuilder {
	b.s.DocumentId = &v
	return b
}

func (b *EnableDocumentPathBuilder) Build() *EnableDocumentPath {
	return b.s
}
