// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAssetGroupByBytedTokenRequest struct {
	GetAssetGroupByBytedTokenBody *GetAssetGroupByBytedTokenBody `json:"getAssetGroupByBytedTokenBody,omitempty"`
}

func (s GetAssetGroupByBytedTokenRequest) String() string {
	return utils.Beautify(s)
}

func (s GetAssetGroupByBytedTokenRequest) GoString() string {
	return s.String()
}

func (s GetAssetGroupByBytedTokenRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAssetGroupByBytedTokenRequest) SetGetAssetGroupByBytedTokenBody(v *GetAssetGroupByBytedTokenBody) *GetAssetGroupByBytedTokenRequest {
	s.GetAssetGroupByBytedTokenBody = v
	return s
}

type GetAssetGroupByBytedTokenRequestBuilder struct {
	s *GetAssetGroupByBytedTokenRequest
}

func NewGetAssetGroupByBytedTokenRequestBuilder() *GetAssetGroupByBytedTokenRequestBuilder {
	s := &GetAssetGroupByBytedTokenRequest{}
	b := &GetAssetGroupByBytedTokenRequestBuilder{s: s}
	return b
}

func (b *GetAssetGroupByBytedTokenRequestBuilder) GetAssetGroupByBytedTokenBody(v *GetAssetGroupByBytedTokenBody) *GetAssetGroupByBytedTokenRequestBuilder {
	b.s.GetAssetGroupByBytedTokenBody = v
	return b
}

func (b *GetAssetGroupByBytedTokenRequestBuilder) Build() *GetAssetGroupByBytedTokenRequest {
	return b.s
}
