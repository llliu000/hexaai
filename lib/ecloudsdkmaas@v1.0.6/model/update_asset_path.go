// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateAssetPath struct {
	position.Path
	// 素材ID（火山素材ID）
	AssetId *string `json:"assetId,omitempty"`
}

func (s UpdateAssetPath) String() string {
	return utils.Beautify(s)
}

func (s UpdateAssetPath) GoString() string {
	return s.String()
}

func (s UpdateAssetPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateAssetPath) SetAssetId(v string) *UpdateAssetPath {
	s.AssetId = &v
	return s
}

type UpdateAssetPathBuilder struct {
	s *UpdateAssetPath
}

func NewUpdateAssetPathBuilder() *UpdateAssetPathBuilder {
	s := &UpdateAssetPath{}
	b := &UpdateAssetPathBuilder{s: s}
	return b
}

func (b *UpdateAssetPathBuilder) AssetId(v string) *UpdateAssetPathBuilder {
	b.s.AssetId = &v
	return b
}

func (b *UpdateAssetPathBuilder) Build() *UpdateAssetPath {
	return b.s
}
