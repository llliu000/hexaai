// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAssetRequest struct {
	GetAssetPath *GetAssetPath `json:"getAssetPath,omitempty"`
}

func (s GetAssetRequest) String() string {
	return utils.Beautify(s)
}

func (s GetAssetRequest) GoString() string {
	return s.String()
}

func (s GetAssetRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAssetRequest) SetGetAssetPath(v *GetAssetPath) *GetAssetRequest {
	s.GetAssetPath = v
	return s
}

type GetAssetRequestBuilder struct {
	s *GetAssetRequest
}

func NewGetAssetRequestBuilder() *GetAssetRequestBuilder {
	s := &GetAssetRequest{}
	b := &GetAssetRequestBuilder{s: s}
	return b
}

func (b *GetAssetRequestBuilder) GetAssetPath(v *GetAssetPath) *GetAssetRequestBuilder {
	b.s.GetAssetPath = v
	return b
}

func (b *GetAssetRequestBuilder) Build() *GetAssetRequest {
	return b.s
}
