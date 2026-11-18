// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QuerySegmentsRequest struct {
	QuerySegmentsPath *QuerySegmentsPath `json:"querySegmentsPath,omitempty"`

	QuerySegmentsBody *QuerySegmentsBody `json:"querySegmentsBody,omitempty"`
}

func (s QuerySegmentsRequest) String() string {
	return utils.Beautify(s)
}

func (s QuerySegmentsRequest) GoString() string {
	return s.String()
}

func (s QuerySegmentsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QuerySegmentsRequest) SetQuerySegmentsPath(v *QuerySegmentsPath) *QuerySegmentsRequest {
	s.QuerySegmentsPath = v
	return s
}

func (s *QuerySegmentsRequest) SetQuerySegmentsBody(v *QuerySegmentsBody) *QuerySegmentsRequest {
	s.QuerySegmentsBody = v
	return s
}

type QuerySegmentsRequestBuilder struct {
	s *QuerySegmentsRequest
}

func NewQuerySegmentsRequestBuilder() *QuerySegmentsRequestBuilder {
	s := &QuerySegmentsRequest{}
	b := &QuerySegmentsRequestBuilder{s: s}
	return b
}

func (b *QuerySegmentsRequestBuilder) QuerySegmentsPath(v *QuerySegmentsPath) *QuerySegmentsRequestBuilder {
	b.s.QuerySegmentsPath = v
	return b
}

func (b *QuerySegmentsRequestBuilder) QuerySegmentsBody(v *QuerySegmentsBody) *QuerySegmentsRequestBuilder {
	b.s.QuerySegmentsBody = v
	return b
}

func (b *QuerySegmentsRequestBuilder) Build() *QuerySegmentsRequest {
	return b.s
}
