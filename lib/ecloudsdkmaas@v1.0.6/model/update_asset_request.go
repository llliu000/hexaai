// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateAssetRequest struct {
	UpdateAssetBody *UpdateAssetBody `json:"updateAssetBody,omitempty"`

	UpdateAssetPath *UpdateAssetPath `json:"updateAssetPath,omitempty"`
}

func (s UpdateAssetRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateAssetRequest) GoString() string {
	return s.String()
}

func (s UpdateAssetRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateAssetRequest) SetUpdateAssetBody(v *UpdateAssetBody) *UpdateAssetRequest {
	s.UpdateAssetBody = v
	return s
}

func (s *UpdateAssetRequest) SetUpdateAssetPath(v *UpdateAssetPath) *UpdateAssetRequest {
	s.UpdateAssetPath = v
	return s
}

type UpdateAssetRequestBuilder struct {
	s *UpdateAssetRequest
}

func NewUpdateAssetRequestBuilder() *UpdateAssetRequestBuilder {
	s := &UpdateAssetRequest{}
	b := &UpdateAssetRequestBuilder{s: s}
	return b
}

func (b *UpdateAssetRequestBuilder) UpdateAssetBody(v *UpdateAssetBody) *UpdateAssetRequestBuilder {
	b.s.UpdateAssetBody = v
	return b
}

func (b *UpdateAssetRequestBuilder) UpdateAssetPath(v *UpdateAssetPath) *UpdateAssetRequestBuilder {
	b.s.UpdateAssetPath = v
	return b
}

func (b *UpdateAssetRequestBuilder) Build() *UpdateAssetRequest {
	return b.s
}
