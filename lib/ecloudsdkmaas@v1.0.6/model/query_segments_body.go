// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QuerySegmentsBody struct {
	position.Body
	// 页码（从 1 开始）
	PageNo *int32 `json:"pageNo,omitempty"`
	// 每页行数
	PageSize *int32 `json:"pageSize,omitempty"`
	// 关键词
	Keyword *string `json:"keyword,omitempty"`
}

func (s QuerySegmentsBody) String() string {
	return utils.Beautify(s)
}

func (s QuerySegmentsBody) GoString() string {
	return s.String()
}

func (s QuerySegmentsBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QuerySegmentsBody) SetPageNo(v int32) *QuerySegmentsBody {
	s.PageNo = &v
	return s
}

func (s *QuerySegmentsBody) SetPageSize(v int32) *QuerySegmentsBody {
	s.PageSize = &v
	return s
}

func (s *QuerySegmentsBody) SetKeyword(v string) *QuerySegmentsBody {
	s.Keyword = &v
	return s
}

type QuerySegmentsBodyBuilder struct {
	s *QuerySegmentsBody
}

func NewQuerySegmentsBodyBuilder() *QuerySegmentsBodyBuilder {
	s := &QuerySegmentsBody{}
	b := &QuerySegmentsBodyBuilder{s: s}
	return b
}

func (b *QuerySegmentsBodyBuilder) PageNo(v int32) *QuerySegmentsBodyBuilder {
	b.s.PageNo = &v
	return b
}

func (b *QuerySegmentsBodyBuilder) PageSize(v int32) *QuerySegmentsBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *QuerySegmentsBodyBuilder) Keyword(v string) *QuerySegmentsBodyBuilder {
	b.s.Keyword = &v
	return b
}

func (b *QuerySegmentsBodyBuilder) Build() *QuerySegmentsBody {
	return b.s
}
