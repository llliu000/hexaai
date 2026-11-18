// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListAssetGroupsRequest struct {
	ListAssetGroupsBody *ListAssetGroupsBody `json:"listAssetGroupsBody,omitempty"`
}

func (s ListAssetGroupsRequest) String() string {
	return utils.Beautify(s)
}

func (s ListAssetGroupsRequest) GoString() string {
	return s.String()
}

func (s ListAssetGroupsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListAssetGroupsRequest) SetListAssetGroupsBody(v *ListAssetGroupsBody) *ListAssetGroupsRequest {
	s.ListAssetGroupsBody = v
	return s
}

type ListAssetGroupsRequestBuilder struct {
	s *ListAssetGroupsRequest
}

func NewListAssetGroupsRequestBuilder() *ListAssetGroupsRequestBuilder {
	s := &ListAssetGroupsRequest{}
	b := &ListAssetGroupsRequestBuilder{s: s}
	return b
}

func (b *ListAssetGroupsRequestBuilder) ListAssetGroupsBody(v *ListAssetGroupsBody) *ListAssetGroupsRequestBuilder {
	b.s.ListAssetGroupsBody = v
	return b
}

func (b *ListAssetGroupsRequestBuilder) Build() *ListAssetGroupsRequest {
	return b.s
}
