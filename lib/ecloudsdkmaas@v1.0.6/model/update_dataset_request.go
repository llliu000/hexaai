// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateDatasetRequest struct {
	UpdateDatasetBody *UpdateDatasetBody `json:"updateDatasetBody,omitempty"`

	UpdateDatasetPath *UpdateDatasetPath `json:"updateDatasetPath,omitempty"`
}

func (s UpdateDatasetRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s UpdateDatasetRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateDatasetRequest) SetUpdateDatasetBody(v *UpdateDatasetBody) *UpdateDatasetRequest {
	s.UpdateDatasetBody = v
	return s
}

func (s *UpdateDatasetRequest) SetUpdateDatasetPath(v *UpdateDatasetPath) *UpdateDatasetRequest {
	s.UpdateDatasetPath = v
	return s
}

type UpdateDatasetRequestBuilder struct {
	s *UpdateDatasetRequest
}

func NewUpdateDatasetRequestBuilder() *UpdateDatasetRequestBuilder {
	s := &UpdateDatasetRequest{}
	b := &UpdateDatasetRequestBuilder{s: s}
	return b
}

func (b *UpdateDatasetRequestBuilder) UpdateDatasetBody(v *UpdateDatasetBody) *UpdateDatasetRequestBuilder {
	b.s.UpdateDatasetBody = v
	return b
}

func (b *UpdateDatasetRequestBuilder) UpdateDatasetPath(v *UpdateDatasetPath) *UpdateDatasetRequestBuilder {
	b.s.UpdateDatasetPath = v
	return b
}

func (b *UpdateDatasetRequestBuilder) Build() *UpdateDatasetRequest {
	return b.s
}
