// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListAssetGroupsResponseBody struct {

	// 查询总条数
	Total *int64 `json:"total,omitempty"`
	// 分页查询列表
	Data *[]ListAssetGroupsResponseData `json:"data,omitempty"`
}

func (s ListAssetGroupsResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListAssetGroupsResponseBody) GoString() string {
	return s.String()
}

func (s ListAssetGroupsResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListAssetGroupsResponseBody) SetTotal(v int64) *ListAssetGroupsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAssetGroupsResponseBody) SetData(v []ListAssetGroupsResponseData) *ListAssetGroupsResponseBody {
	s.Data = &v
	return s
}

type ListAssetGroupsResponseBodyBuilder struct {
	s *ListAssetGroupsResponseBody
}

func NewListAssetGroupsResponseBodyBuilder() *ListAssetGroupsResponseBodyBuilder {
	s := &ListAssetGroupsResponseBody{}
	b := &ListAssetGroupsResponseBodyBuilder{s: s}
	return b
}

func (b *ListAssetGroupsResponseBodyBuilder) Total(v int64) *ListAssetGroupsResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListAssetGroupsResponseBodyBuilder) Data(v []ListAssetGroupsResponseData) *ListAssetGroupsResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *ListAssetGroupsResponseBodyBuilder) Build() *ListAssetGroupsResponseBody {
	return b.s
}
