// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryUsageStatsRequest struct {
	QueryUsageStatsBody *QueryUsageStatsBody `json:"queryUsageStatsBody,omitempty"`
}

func (s QueryUsageStatsRequest) String() string {
	return utils.Beautify(s)
}

func (s QueryUsageStatsRequest) GoString() string {
	return s.String()
}

func (s QueryUsageStatsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryUsageStatsRequest) SetQueryUsageStatsBody(v *QueryUsageStatsBody) *QueryUsageStatsRequest {
	s.QueryUsageStatsBody = v
	return s
}

type QueryUsageStatsRequestBuilder struct {
	s *QueryUsageStatsRequest
}

func NewQueryUsageStatsRequestBuilder() *QueryUsageStatsRequestBuilder {
	s := &QueryUsageStatsRequest{}
	b := &QueryUsageStatsRequestBuilder{s: s}
	return b
}

func (b *QueryUsageStatsRequestBuilder) QueryUsageStatsBody(v *QueryUsageStatsBody) *QueryUsageStatsRequestBuilder {
	b.s.QueryUsageStatsBody = v
	return b
}

func (b *QueryUsageStatsRequestBuilder) Build() *QueryUsageStatsRequest {
	return b.s
}
