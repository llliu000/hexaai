// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteSegmentRequest struct {
	DeleteSegmentPath *DeleteSegmentPath `json:"deleteSegmentPath,omitempty"`
}

func (s DeleteSegmentRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteSegmentRequest) GoString() string {
	return s.String()
}

func (s DeleteSegmentRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteSegmentRequest) SetDeleteSegmentPath(v *DeleteSegmentPath) *DeleteSegmentRequest {
	s.DeleteSegmentPath = v
	return s
}

type DeleteSegmentRequestBuilder struct {
	s *DeleteSegmentRequest
}

func NewDeleteSegmentRequestBuilder() *DeleteSegmentRequestBuilder {
	s := &DeleteSegmentRequest{}
	b := &DeleteSegmentRequestBuilder{s: s}
	return b
}

func (b *DeleteSegmentRequestBuilder) DeleteSegmentPath(v *DeleteSegmentPath) *DeleteSegmentRequestBuilder {
	b.s.DeleteSegmentPath = v
	return b
}

func (b *DeleteSegmentRequestBuilder) Build() *DeleteSegmentRequest {
	return b.s
}
