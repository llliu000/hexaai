// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetSegmentRequest struct {
	GetSegmentPath *GetSegmentPath `json:"getSegmentPath,omitempty"`
}

func (s GetSegmentRequest) String() string {
	return utils.Beautify(s)
}

func (s GetSegmentRequest) GoString() string {
	return s.String()
}

func (s GetSegmentRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSegmentRequest) SetGetSegmentPath(v *GetSegmentPath) *GetSegmentRequest {
	s.GetSegmentPath = v
	return s
}

type GetSegmentRequestBuilder struct {
	s *GetSegmentRequest
}

func NewGetSegmentRequestBuilder() *GetSegmentRequestBuilder {
	s := &GetSegmentRequest{}
	b := &GetSegmentRequestBuilder{s: s}
	return b
}

func (b *GetSegmentRequestBuilder) GetSegmentPath(v *GetSegmentPath) *GetSegmentRequestBuilder {
	b.s.GetSegmentPath = v
	return b
}

func (b *GetSegmentRequestBuilder) Build() *GetSegmentRequest {
	return b.s
}
