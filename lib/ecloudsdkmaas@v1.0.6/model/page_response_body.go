// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PageResponseBody struct {

	// 查询总条数
	Total *int64 `json:"total,omitempty"`
	// 分页查询列表
	Data *[]PageResponseData `json:"data,omitempty"`
}

func (s PageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s PageResponseBody) GoString() string {
	return s.String()
}

func (s PageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PageResponseBody) SetTotal(v int64) *PageResponseBody {
	s.Total = &v
	return s
}

func (s *PageResponseBody) SetData(v []PageResponseData) *PageResponseBody {
	s.Data = &v
	return s
}

type PageResponseBodyBuilder struct {
	s *PageResponseBody
}

func NewPageResponseBodyBuilder() *PageResponseBodyBuilder {
	s := &PageResponseBody{}
	b := &PageResponseBodyBuilder{s: s}
	return b
}

func (b *PageResponseBodyBuilder) Total(v int64) *PageResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *PageResponseBodyBuilder) Data(v []PageResponseData) *PageResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *PageResponseBodyBuilder) Build() *PageResponseBody {
	return b.s
}
