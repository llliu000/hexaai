// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListAssetsResponseBody struct {

	// 查询总条数
	Total *int64 `json:"total,omitempty"`
	// 分页查询列表
	Data *[]ListAssetsResponseData `json:"data,omitempty"`
}

func (s ListAssetsResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListAssetsResponseBody) GoString() string {
	return s.String()
}

func (s ListAssetsResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListAssetsResponseBody) SetTotal(v int64) *ListAssetsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAssetsResponseBody) SetData(v []ListAssetsResponseData) *ListAssetsResponseBody {
	s.Data = &v
	return s
}

type ListAssetsResponseBodyBuilder struct {
	s *ListAssetsResponseBody
}

func NewListAssetsResponseBodyBuilder() *ListAssetsResponseBodyBuilder {
	s := &ListAssetsResponseBody{}
	b := &ListAssetsResponseBodyBuilder{s: s}
	return b
}

func (b *ListAssetsResponseBodyBuilder) Total(v int64) *ListAssetsResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListAssetsResponseBodyBuilder) Data(v []ListAssetsResponseData) *ListAssetsResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *ListAssetsResponseBodyBuilder) Build() *ListAssetsResponseBody {
	return b.s
}
