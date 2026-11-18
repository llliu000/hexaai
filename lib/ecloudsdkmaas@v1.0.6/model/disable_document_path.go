// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DisableDocumentPath struct {
	position.Path
	// 知识库Id（URL路径参数）
	DatasetId *string `json:"datasetId,omitempty"`
	// 文档Id（URL路径参数）
	DocumentId *string `json:"documentId,omitempty"`
}

func (s DisableDocumentPath) String() string {
	return utils.Beautify(s)
}

func (s DisableDocumentPath) GoString() string {
	return s.String()
}

func (s DisableDocumentPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DisableDocumentPath) SetDatasetId(v string) *DisableDocumentPath {
	s.DatasetId = &v
	return s
}

func (s *DisableDocumentPath) SetDocumentId(v string) *DisableDocumentPath {
	s.DocumentId = &v
	return s
}

type DisableDocumentPathBuilder struct {
	s *DisableDocumentPath
}

func NewDisableDocumentPathBuilder() *DisableDocumentPathBuilder {
	s := &DisableDocumentPath{}
	b := &DisableDocumentPathBuilder{s: s}
	return b
}

func (b *DisableDocumentPathBuilder) DatasetId(v string) *DisableDocumentPathBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *DisableDocumentPathBuilder) DocumentId(v string) *DisableDocumentPathBuilder {
	b.s.DocumentId = &v
	return b
}

func (b *DisableDocumentPathBuilder) Build() *DisableDocumentPath {
	return b.s
}
