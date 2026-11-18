// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListDocsRequest struct {
	ListDocsQuery *ListDocsQuery `json:"listDocsQuery,omitempty"`
}

func (s ListDocsRequest) String() string {
	return utils.Beautify(s)
}

func (s ListDocsRequest) GoString() string {
	return s.String()
}

func (s ListDocsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListDocsRequest) SetListDocsQuery(v *ListDocsQuery) *ListDocsRequest {
	s.ListDocsQuery = v
	return s
}

type ListDocsRequestBuilder struct {
	s *ListDocsRequest
}

func NewListDocsRequestBuilder() *ListDocsRequestBuilder {
	s := &ListDocsRequest{}
	b := &ListDocsRequestBuilder{s: s}
	return b
}

func (b *ListDocsRequestBuilder) ListDocsQuery(v *ListDocsQuery) *ListDocsRequestBuilder {
	b.s.ListDocsQuery = v
	return b
}

func (b *ListDocsRequestBuilder) Build() *ListDocsRequest {
	return b.s
}
