// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListAssetsRequest struct {
	ListAssetsBody *ListAssetsBody `json:"listAssetsBody,omitempty"`
}

func (s ListAssetsRequest) String() string {
	return utils.Beautify(s)
}

func (s ListAssetsRequest) GoString() string {
	return s.String()
}

func (s ListAssetsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListAssetsRequest) SetListAssetsBody(v *ListAssetsBody) *ListAssetsRequest {
	s.ListAssetsBody = v
	return s
}

type ListAssetsRequestBuilder struct {
	s *ListAssetsRequest
}

func NewListAssetsRequestBuilder() *ListAssetsRequestBuilder {
	s := &ListAssetsRequest{}
	b := &ListAssetsRequestBuilder{s: s}
	return b
}

func (b *ListAssetsRequestBuilder) ListAssetsBody(v *ListAssetsBody) *ListAssetsRequestBuilder {
	b.s.ListAssetsBody = v
	return b
}

func (b *ListAssetsRequestBuilder) Build() *ListAssetsRequest {
	return b.s
}
