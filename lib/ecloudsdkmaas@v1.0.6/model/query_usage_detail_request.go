// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryUsageDetailRequest struct {
	QueryUsageDetailBody *QueryUsageDetailBody `json:"queryUsageDetailBody,omitempty"`
}

func (s QueryUsageDetailRequest) String() string {
	return utils.Beautify(s)
}

func (s QueryUsageDetailRequest) GoString() string {
	return s.String()
}

func (s QueryUsageDetailRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryUsageDetailRequest) SetQueryUsageDetailBody(v *QueryUsageDetailBody) *QueryUsageDetailRequest {
	s.QueryUsageDetailBody = v
	return s
}

type QueryUsageDetailRequestBuilder struct {
	s *QueryUsageDetailRequest
}

func NewQueryUsageDetailRequestBuilder() *QueryUsageDetailRequestBuilder {
	s := &QueryUsageDetailRequest{}
	b := &QueryUsageDetailRequestBuilder{s: s}
	return b
}

func (b *QueryUsageDetailRequestBuilder) QueryUsageDetailBody(v *QueryUsageDetailBody) *QueryUsageDetailRequestBuilder {
	b.s.QueryUsageDetailBody = v
	return b
}

func (b *QueryUsageDetailRequestBuilder) Build() *QueryUsageDetailRequest {
	return b.s
}
