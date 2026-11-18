// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryUsageStatsResponseBody struct {

	// 查询总条数
	Total *int64 `json:"total,omitempty"`
	// 分页查询列表
	Data *[]QueryUsageStatsResponseData `json:"data,omitempty"`
}

func (s QueryUsageStatsResponseBody) String() string {
	return utils.Beautify(s)
}

func (s QueryUsageStatsResponseBody) GoString() string {
	return s.String()
}

func (s QueryUsageStatsResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryUsageStatsResponseBody) SetTotal(v int64) *QueryUsageStatsResponseBody {
	s.Total = &v
	return s
}

func (s *QueryUsageStatsResponseBody) SetData(v []QueryUsageStatsResponseData) *QueryUsageStatsResponseBody {
	s.Data = &v
	return s
}

type QueryUsageStatsResponseBodyBuilder struct {
	s *QueryUsageStatsResponseBody
}

func NewQueryUsageStatsResponseBodyBuilder() *QueryUsageStatsResponseBodyBuilder {
	s := &QueryUsageStatsResponseBody{}
	b := &QueryUsageStatsResponseBodyBuilder{s: s}
	return b
}

func (b *QueryUsageStatsResponseBodyBuilder) Total(v int64) *QueryUsageStatsResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *QueryUsageStatsResponseBodyBuilder) Data(v []QueryUsageStatsResponseData) *QueryUsageStatsResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *QueryUsageStatsResponseBodyBuilder) Build() *QueryUsageStatsResponseBody {
	return b.s
}
