// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EnableDocumentRequest struct {
	EnableDocumentPath *EnableDocumentPath `json:"enableDocumentPath,omitempty"`
}

func (s EnableDocumentRequest) String() string {
	return utils.Beautify(s)
}

func (s EnableDocumentRequest) GoString() string {
	return s.String()
}

func (s EnableDocumentRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EnableDocumentRequest) SetEnableDocumentPath(v *EnableDocumentPath) *EnableDocumentRequest {
	s.EnableDocumentPath = v
	return s
}

type EnableDocumentRequestBuilder struct {
	s *EnableDocumentRequest
}

func NewEnableDocumentRequestBuilder() *EnableDocumentRequestBuilder {
	s := &EnableDocumentRequest{}
	b := &EnableDocumentRequestBuilder{s: s}
	return b
}

func (b *EnableDocumentRequestBuilder) EnableDocumentPath(v *EnableDocumentPath) *EnableDocumentRequestBuilder {
	b.s.EnableDocumentPath = v
	return b
}

func (b *EnableDocumentRequestBuilder) Build() *EnableDocumentRequest {
	return b.s
}
