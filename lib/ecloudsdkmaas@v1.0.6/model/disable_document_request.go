// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DisableDocumentRequest struct {
	DisableDocumentPath *DisableDocumentPath `json:"disableDocumentPath,omitempty"`
}

func (s DisableDocumentRequest) String() string {
	return utils.Beautify(s)
}

func (s DisableDocumentRequest) GoString() string {
	return s.String()
}

func (s DisableDocumentRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DisableDocumentRequest) SetDisableDocumentPath(v *DisableDocumentPath) *DisableDocumentRequest {
	s.DisableDocumentPath = v
	return s
}

type DisableDocumentRequestBuilder struct {
	s *DisableDocumentRequest
}

func NewDisableDocumentRequestBuilder() *DisableDocumentRequestBuilder {
	s := &DisableDocumentRequest{}
	b := &DisableDocumentRequestBuilder{s: s}
	return b
}

func (b *DisableDocumentRequestBuilder) DisableDocumentPath(v *DisableDocumentPath) *DisableDocumentRequestBuilder {
	b.s.DisableDocumentPath = v
	return b
}

func (b *DisableDocumentRequestBuilder) Build() *DisableDocumentRequest {
	return b.s
}
