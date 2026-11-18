// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryDocsPageByTypeRequest struct {
	QueryDocsPageByTypePath *QueryDocsPageByTypePath `json:"queryDocsPageByTypePath,omitempty"`

	QueryDocsPageByTypeBody *QueryDocsPageByTypeBody `json:"queryDocsPageByTypeBody,omitempty"`
}

func (s QueryDocsPageByTypeRequest) String() string {
	return utils.Beautify(s)
}

func (s QueryDocsPageByTypeRequest) GoString() string {
	return s.String()
}

func (s QueryDocsPageByTypeRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryDocsPageByTypeRequest) SetQueryDocsPageByTypePath(v *QueryDocsPageByTypePath) *QueryDocsPageByTypeRequest {
	s.QueryDocsPageByTypePath = v
	return s
}

func (s *QueryDocsPageByTypeRequest) SetQueryDocsPageByTypeBody(v *QueryDocsPageByTypeBody) *QueryDocsPageByTypeRequest {
	s.QueryDocsPageByTypeBody = v
	return s
}

type QueryDocsPageByTypeRequestBuilder struct {
	s *QueryDocsPageByTypeRequest
}

func NewQueryDocsPageByTypeRequestBuilder() *QueryDocsPageByTypeRequestBuilder {
	s := &QueryDocsPageByTypeRequest{}
	b := &QueryDocsPageByTypeRequestBuilder{s: s}
	return b
}

func (b *QueryDocsPageByTypeRequestBuilder) QueryDocsPageByTypePath(v *QueryDocsPageByTypePath) *QueryDocsPageByTypeRequestBuilder {
	b.s.QueryDocsPageByTypePath = v
	return b
}

func (b *QueryDocsPageByTypeRequestBuilder) QueryDocsPageByTypeBody(v *QueryDocsPageByTypeBody) *QueryDocsPageByTypeRequestBuilder {
	b.s.QueryDocsPageByTypeBody = v
	return b
}

func (b *QueryDocsPageByTypeRequestBuilder) Build() *QueryDocsPageByTypeRequest {
	return b.s
}
