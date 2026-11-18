// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EnableSegmentRequest struct {
	EnableSegmentPath *EnableSegmentPath `json:"enableSegmentPath,omitempty"`
}

func (s EnableSegmentRequest) String() string {
	return utils.Beautify(s)
}

func (s EnableSegmentRequest) GoString() string {
	return s.String()
}

func (s EnableSegmentRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EnableSegmentRequest) SetEnableSegmentPath(v *EnableSegmentPath) *EnableSegmentRequest {
	s.EnableSegmentPath = v
	return s
}

type EnableSegmentRequestBuilder struct {
	s *EnableSegmentRequest
}

func NewEnableSegmentRequestBuilder() *EnableSegmentRequestBuilder {
	s := &EnableSegmentRequest{}
	b := &EnableSegmentRequestBuilder{s: s}
	return b
}

func (b *EnableSegmentRequestBuilder) EnableSegmentPath(v *EnableSegmentPath) *EnableSegmentRequestBuilder {
	b.s.EnableSegmentPath = v
	return b
}

func (b *EnableSegmentRequestBuilder) Build() *EnableSegmentRequest {
	return b.s
}
