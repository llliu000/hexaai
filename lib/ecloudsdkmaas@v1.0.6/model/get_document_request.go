// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetDocumentRequest struct {
	GetDocumentPath *GetDocumentPath `json:"getDocumentPath,omitempty"`
}

func (s GetDocumentRequest) String() string {
	return utils.Beautify(s)
}

func (s GetDocumentRequest) GoString() string {
	return s.String()
}

func (s GetDocumentRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetDocumentRequest) SetGetDocumentPath(v *GetDocumentPath) *GetDocumentRequest {
	s.GetDocumentPath = v
	return s
}

type GetDocumentRequestBuilder struct {
	s *GetDocumentRequest
}

func NewGetDocumentRequestBuilder() *GetDocumentRequestBuilder {
	s := &GetDocumentRequest{}
	b := &GetDocumentRequestBuilder{s: s}
	return b
}

func (b *GetDocumentRequestBuilder) GetDocumentPath(v *GetDocumentPath) *GetDocumentRequestBuilder {
	b.s.GetDocumentPath = v
	return b
}

func (b *GetDocumentRequestBuilder) Build() *GetDocumentRequest {
	return b.s
}
