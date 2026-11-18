// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateAssetBody struct {
	position.Body
	// 素材名称，最长64个字符，不传则不更新
	AssetName *string `json:"assetName,omitempty"`
}

func (s UpdateAssetBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateAssetBody) GoString() string {
	return s.String()
}

func (s UpdateAssetBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateAssetBody) SetAssetName(v string) *UpdateAssetBody {
	s.AssetName = &v
	return s
}

type UpdateAssetBodyBuilder struct {
	s *UpdateAssetBody
}

func NewUpdateAssetBodyBuilder() *UpdateAssetBodyBuilder {
	s := &UpdateAssetBody{}
	b := &UpdateAssetBodyBuilder{s: s}
	return b
}

func (b *UpdateAssetBodyBuilder) AssetName(v string) *UpdateAssetBodyBuilder {
	b.s.AssetName = &v
	return b
}

func (b *UpdateAssetBodyBuilder) Build() *UpdateAssetBody {
	return b.s
}
