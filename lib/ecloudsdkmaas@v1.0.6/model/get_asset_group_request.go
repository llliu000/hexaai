// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAssetGroupRequest struct {
	GetAssetGroupPath *GetAssetGroupPath `json:"getAssetGroupPath,omitempty"`
}

func (s GetAssetGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s GetAssetGroupRequest) GoString() string {
	return s.String()
}

func (s GetAssetGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAssetGroupRequest) SetGetAssetGroupPath(v *GetAssetGroupPath) *GetAssetGroupRequest {
	s.GetAssetGroupPath = v
	return s
}

type GetAssetGroupRequestBuilder struct {
	s *GetAssetGroupRequest
}

func NewGetAssetGroupRequestBuilder() *GetAssetGroupRequestBuilder {
	s := &GetAssetGroupRequest{}
	b := &GetAssetGroupRequestBuilder{s: s}
	return b
}

func (b *GetAssetGroupRequestBuilder) GetAssetGroupPath(v *GetAssetGroupPath) *GetAssetGroupRequestBuilder {
	b.s.GetAssetGroupPath = v
	return b
}

func (b *GetAssetGroupRequestBuilder) Build() *GetAssetGroupRequest {
	return b.s
}
