// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetDocumentPath struct {
	position.Path
	// 文档 Id
	DocumentId *string `json:"documentId,omitempty"`
}

func (s GetDocumentPath) String() string {
	return utils.Beautify(s)
}

func (s GetDocumentPath) GoString() string {
	return s.String()
}

func (s GetDocumentPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetDocumentPath) SetDocumentId(v string) *GetDocumentPath {
	s.DocumentId = &v
	return s
}

type GetDocumentPathBuilder struct {
	s *GetDocumentPath
}

func NewGetDocumentPathBuilder() *GetDocumentPathBuilder {
	s := &GetDocumentPath{}
	b := &GetDocumentPathBuilder{s: s}
	return b
}

func (b *GetDocumentPathBuilder) DocumentId(v string) *GetDocumentPathBuilder {
	b.s.DocumentId = &v
	return b
}

func (b *GetDocumentPathBuilder) Build() *GetDocumentPath {
	return b.s
}
