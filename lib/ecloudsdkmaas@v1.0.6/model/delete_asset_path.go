// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteAssetPath struct {
	position.Path
	// 素材ID（火山素材ID）
	AssetId *string `json:"assetId,omitempty"`
}

func (s DeleteAssetPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteAssetPath) GoString() string {
	return s.String()
}

func (s DeleteAssetPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteAssetPath) SetAssetId(v string) *DeleteAssetPath {
	s.AssetId = &v
	return s
}

type DeleteAssetPathBuilder struct {
	s *DeleteAssetPath
}

func NewDeleteAssetPathBuilder() *DeleteAssetPathBuilder {
	s := &DeleteAssetPath{}
	b := &DeleteAssetPathBuilder{s: s}
	return b
}

func (b *DeleteAssetPathBuilder) AssetId(v string) *DeleteAssetPathBuilder {
	b.s.AssetId = &v
	return b
}

func (b *DeleteAssetPathBuilder) Build() *DeleteAssetPath {
	return b.s
}
