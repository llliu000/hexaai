// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateAssetGroupRequest struct {
	UpdateAssetGroupPath *UpdateAssetGroupPath `json:"updateAssetGroupPath,omitempty"`

	UpdateAssetGroupBody *UpdateAssetGroupBody `json:"updateAssetGroupBody,omitempty"`
}

func (s UpdateAssetGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateAssetGroupRequest) GoString() string {
	return s.String()
}

func (s UpdateAssetGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateAssetGroupRequest) SetUpdateAssetGroupPath(v *UpdateAssetGroupPath) *UpdateAssetGroupRequest {
	s.UpdateAssetGroupPath = v
	return s
}

func (s *UpdateAssetGroupRequest) SetUpdateAssetGroupBody(v *UpdateAssetGroupBody) *UpdateAssetGroupRequest {
	s.UpdateAssetGroupBody = v
	return s
}

type UpdateAssetGroupRequestBuilder struct {
	s *UpdateAssetGroupRequest
}

func NewUpdateAssetGroupRequestBuilder() *UpdateAssetGroupRequestBuilder {
	s := &UpdateAssetGroupRequest{}
	b := &UpdateAssetGroupRequestBuilder{s: s}
	return b
}

func (b *UpdateAssetGroupRequestBuilder) UpdateAssetGroupPath(v *UpdateAssetGroupPath) *UpdateAssetGroupRequestBuilder {
	b.s.UpdateAssetGroupPath = v
	return b
}

func (b *UpdateAssetGroupRequestBuilder) UpdateAssetGroupBody(v *UpdateAssetGroupBody) *UpdateAssetGroupRequestBuilder {
	b.s.UpdateAssetGroupBody = v
	return b
}

func (b *UpdateAssetGroupRequestBuilder) Build() *UpdateAssetGroupRequest {
	return b.s
}
