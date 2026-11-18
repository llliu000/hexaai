// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAssetPath struct {
	position.Path
	// 素材ID（火山素材ID）
	AssetId *string `json:"assetId,omitempty"`
}

func (s GetAssetPath) String() string {
	return utils.Beautify(s)
}

func (s GetAssetPath) GoString() string {
	return s.String()
}

func (s GetAssetPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAssetPath) SetAssetId(v string) *GetAssetPath {
	s.AssetId = &v
	return s
}

type GetAssetPathBuilder struct {
	s *GetAssetPath
}

func NewGetAssetPathBuilder() *GetAssetPathBuilder {
	s := &GetAssetPath{}
	b := &GetAssetPathBuilder{s: s}
	return b
}

func (b *GetAssetPathBuilder) AssetId(v string) *GetAssetPathBuilder {
	b.s.AssetId = &v
	return b
}

func (b *GetAssetPathBuilder) Build() *GetAssetPath {
	return b.s
}
