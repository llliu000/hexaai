// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAssetBodyAssetTypeEnum string

// List of AssetType
const (
	CreateAssetBodyAssetTypeEnumImage CreateAssetBodyAssetTypeEnum = "Image"
	CreateAssetBodyAssetTypeEnumVideo CreateAssetBodyAssetTypeEnum = "Video"
	CreateAssetBodyAssetTypeEnumAudio CreateAssetBodyAssetTypeEnum = "Audio"
)

type CreateAssetBody struct {
	position.Body
	// 素材组ID（火山资产组ID）
	GroupId *string `json:"groupId,omitempty"`
	// 素材名称，最长64个字符
	AssetName *string `json:"assetName,omitempty"`
	// 素材URL，需为公网可访问的HTTP/HTTPS地址，火山侧将通过该URL下载素材
	AssetUrl *string `json:"assetUrl,omitempty"`
	// 素材类型，可选值：Image（图片）、Video（视频）、Audio（音频）
	AssetType *CreateAssetBodyAssetTypeEnum `json:"assetType,omitempty"`
}

func (s CreateAssetBody) String() string {
	return utils.Beautify(s)
}

func (s CreateAssetBody) GoString() string {
	return s.String()
}

func (s CreateAssetBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAssetBody) SetGroupId(v string) *CreateAssetBody {
	s.GroupId = &v
	return s
}

func (s *CreateAssetBody) SetAssetName(v string) *CreateAssetBody {
	s.AssetName = &v
	return s
}

func (s *CreateAssetBody) SetAssetUrl(v string) *CreateAssetBody {
	s.AssetUrl = &v
	return s
}

func (s *CreateAssetBody) SetAssetType(v CreateAssetBodyAssetTypeEnum) *CreateAssetBody {
	s.AssetType = &v
	return s
}

type CreateAssetBodyBuilder struct {
	s *CreateAssetBody
}

func NewCreateAssetBodyBuilder() *CreateAssetBodyBuilder {
	s := &CreateAssetBody{}
	b := &CreateAssetBodyBuilder{s: s}
	return b
}

func (b *CreateAssetBodyBuilder) GroupId(v string) *CreateAssetBodyBuilder {
	b.s.GroupId = &v
	return b
}

func (b *CreateAssetBodyBuilder) AssetName(v string) *CreateAssetBodyBuilder {
	b.s.AssetName = &v
	return b
}

func (b *CreateAssetBodyBuilder) AssetUrl(v string) *CreateAssetBodyBuilder {
	b.s.AssetUrl = &v
	return b
}

func (b *CreateAssetBodyBuilder) AssetType(v CreateAssetBodyAssetTypeEnum) *CreateAssetBodyBuilder {
	b.s.AssetType = &v
	return b
}

func (b *CreateAssetBodyBuilder) Build() *CreateAssetBody {
	return b.s
}
