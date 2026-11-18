// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListDatasetByUserIdRequest struct {
	ListDatasetByUserIdHeader *ListDatasetByUserIdHeader `json:"listDatasetByUserIdHeader,omitempty"`

	ListDatasetByUserIdBody *ListDatasetByUserIdBody `json:"listDatasetByUserIdBody,omitempty"`
}

func (s ListDatasetByUserIdRequest) String() string {
	return utils.Beautify(s)
}

func (s ListDatasetByUserIdRequest) GoString() string {
	return s.String()
}

func (s ListDatasetByUserIdRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListDatasetByUserIdRequest) SetListDatasetByUserIdHeader(v *ListDatasetByUserIdHeader) *ListDatasetByUserIdRequest {
	s.ListDatasetByUserIdHeader = v
	return s
}

func (s *ListDatasetByUserIdRequest) SetListDatasetByUserIdBody(v *ListDatasetByUserIdBody) *ListDatasetByUserIdRequest {
	s.ListDatasetByUserIdBody = v
	return s
}

type ListDatasetByUserIdRequestBuilder struct {
	s *ListDatasetByUserIdRequest
}

func NewListDatasetByUserIdRequestBuilder() *ListDatasetByUserIdRequestBuilder {
	s := &ListDatasetByUserIdRequest{}
	b := &ListDatasetByUserIdRequestBuilder{s: s}
	return b
}

func (b *ListDatasetByUserIdRequestBuilder) ListDatasetByUserIdHeader(v *ListDatasetByUserIdHeader) *ListDatasetByUserIdRequestBuilder {
	b.s.ListDatasetByUserIdHeader = v
	return b
}

func (b *ListDatasetByUserIdRequestBuilder) ListDatasetByUserIdBody(v *ListDatasetByUserIdBody) *ListDatasetByUserIdRequestBuilder {
	b.s.ListDatasetByUserIdBody = v
	return b
}

func (b *ListDatasetByUserIdRequestBuilder) Build() *ListDatasetByUserIdRequest {
	return b.s
}
