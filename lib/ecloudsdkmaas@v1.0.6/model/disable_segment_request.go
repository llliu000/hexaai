// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DisableSegmentRequest struct {
	DisableSegmentPath *DisableSegmentPath `json:"disableSegmentPath,omitempty"`
}

func (s DisableSegmentRequest) String() string {
	return utils.Beautify(s)
}

func (s DisableSegmentRequest) GoString() string {
	return s.String()
}

func (s DisableSegmentRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DisableSegmentRequest) SetDisableSegmentPath(v *DisableSegmentPath) *DisableSegmentRequest {
	s.DisableSegmentPath = v
	return s
}

type DisableSegmentRequestBuilder struct {
	s *DisableSegmentRequest
}

func NewDisableSegmentRequestBuilder() *DisableSegmentRequestBuilder {
	s := &DisableSegmentRequest{}
	b := &DisableSegmentRequestBuilder{s: s}
	return b
}

func (b *DisableSegmentRequestBuilder) DisableSegmentPath(v *DisableSegmentPath) *DisableSegmentRequestBuilder {
	b.s.DisableSegmentPath = v
	return b
}

func (b *DisableSegmentRequestBuilder) Build() *DisableSegmentRequest {
	return b.s
}
