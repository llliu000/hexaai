// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryUsageDetailResponseBody struct {

	// 查询总条数
	Total *int64 `json:"total,omitempty"`
	// 分页查询列表
	Data *[]QueryUsageDetailResponseData `json:"data,omitempty"`
}

func (s QueryUsageDetailResponseBody) String() string {
	return utils.Beautify(s)
}

func (s QueryUsageDetailResponseBody) GoString() string {
	return s.String()
}

func (s QueryUsageDetailResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryUsageDetailResponseBody) SetTotal(v int64) *QueryUsageDetailResponseBody {
	s.Total = &v
	return s
}

func (s *QueryUsageDetailResponseBody) SetData(v []QueryUsageDetailResponseData) *QueryUsageDetailResponseBody {
	s.Data = &v
	return s
}

type QueryUsageDetailResponseBodyBuilder struct {
	s *QueryUsageDetailResponseBody
}

func NewQueryUsageDetailResponseBodyBuilder() *QueryUsageDetailResponseBodyBuilder {
	s := &QueryUsageDetailResponseBody{}
	b := &QueryUsageDetailResponseBodyBuilder{s: s}
	return b
}

func (b *QueryUsageDetailResponseBodyBuilder) Total(v int64) *QueryUsageDetailResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *QueryUsageDetailResponseBodyBuilder) Data(v []QueryUsageDetailResponseData) *QueryUsageDetailResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *QueryUsageDetailResponseBodyBuilder) Build() *QueryUsageDetailResponseBody {
	return b.s
}
