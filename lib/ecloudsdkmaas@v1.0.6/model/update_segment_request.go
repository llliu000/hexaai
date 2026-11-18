// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateSegmentRequest struct {
	UpdateSegmentBody *UpdateSegmentBody `json:"updateSegmentBody,omitempty"`

	UpdateSegmentPath *UpdateSegmentPath `json:"updateSegmentPath,omitempty"`
}

func (s UpdateSegmentRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateSegmentRequest) GoString() string {
	return s.String()
}

func (s UpdateSegmentRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateSegmentRequest) SetUpdateSegmentBody(v *UpdateSegmentBody) *UpdateSegmentRequest {
	s.UpdateSegmentBody = v
	return s
}

func (s *UpdateSegmentRequest) SetUpdateSegmentPath(v *UpdateSegmentPath) *UpdateSegmentRequest {
	s.UpdateSegmentPath = v
	return s
}

type UpdateSegmentRequestBuilder struct {
	s *UpdateSegmentRequest
}

func NewUpdateSegmentRequestBuilder() *UpdateSegmentRequestBuilder {
	s := &UpdateSegmentRequest{}
	b := &UpdateSegmentRequestBuilder{s: s}
	return b
}

func (b *UpdateSegmentRequestBuilder) UpdateSegmentBody(v *UpdateSegmentBody) *UpdateSegmentRequestBuilder {
	b.s.UpdateSegmentBody = v
	return b
}

func (b *UpdateSegmentRequestBuilder) UpdateSegmentPath(v *UpdateSegmentPath) *UpdateSegmentRequestBuilder {
	b.s.UpdateSegmentPath = v
	return b
}

func (b *UpdateSegmentRequestBuilder) Build() *UpdateSegmentRequest {
	return b.s
}
