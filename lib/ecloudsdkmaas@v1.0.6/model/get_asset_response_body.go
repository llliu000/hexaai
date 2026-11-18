// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAssetResponseBodyAssetTypeEnum string

// List of AssetType
const (
	GetAssetResponseBodyAssetTypeEnumImage GetAssetResponseBodyAssetTypeEnum = "Image"
	GetAssetResponseBodyAssetTypeEnumVideo GetAssetResponseBodyAssetTypeEnum = "Video"
	GetAssetResponseBodyAssetTypeEnumAudio GetAssetResponseBodyAssetTypeEnum = "Audio"
)

type GetAssetResponseBodyStatusEnum string

// List of Status
const (
	GetAssetResponseBodyStatusEnumProcessing GetAssetResponseBodyStatusEnum = "PROCESSING"
	GetAssetResponseBodyStatusEnumActive     GetAssetResponseBodyStatusEnum = "ACTIVE"
	GetAssetResponseBodyStatusEnumFailed     GetAssetResponseBodyStatusEnum = "FAILED"
)

type GetAssetResponseBody struct {

	// 更新时间，格式：yyyy-MM-dd HH:mm:ss
	UpdatedTime *string `json:"updatedTime,omitempty"`
	// 火山素材ID
	AssetId *string `json:"assetId,omitempty"`
	// 火山资产组ID
	GroupId *string `json:"groupId,omitempty"`
	// 错误信息，素材状态为FAILED时返回失败原因
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 素材名称
	AssetName *string `json:"assetName,omitempty"`
	// 创建时间，格式：yyyy-MM-dd HH:mm:ss
	CreatedTime *string `json:"createdTime,omitempty"`
	// 素材预签名URL，动态生成，有效期12小时
	AssetUrl *string `json:"assetUrl,omitempty"`
	// AssetTypeEnum
	AssetType *GetAssetResponseBodyAssetTypeEnum `json:"assetType,omitempty"`
	// AssetStatusEnum
	Status *GetAssetResponseBodyStatusEnum `json:"status,omitempty"`
}

func (s GetAssetResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetAssetResponseBody) GoString() string {
	return s.String()
}

func (s GetAssetResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAssetResponseBody) SetUpdatedTime(v string) *GetAssetResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *GetAssetResponseBody) SetAssetId(v string) *GetAssetResponseBody {
	s.AssetId = &v
	return s
}

func (s *GetAssetResponseBody) SetGroupId(v string) *GetAssetResponseBody {
	s.GroupId = &v
	return s
}

func (s *GetAssetResponseBody) SetErrorMessage(v string) *GetAssetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAssetResponseBody) SetAssetName(v string) *GetAssetResponseBody {
	s.AssetName = &v
	return s
}

func (s *GetAssetResponseBody) SetCreatedTime(v string) *GetAssetResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetAssetResponseBody) SetAssetUrl(v string) *GetAssetResponseBody {
	s.AssetUrl = &v
	return s
}

func (s *GetAssetResponseBody) SetAssetType(v GetAssetResponseBodyAssetTypeEnum) *GetAssetResponseBody {
	s.AssetType = &v
	return s
}

func (s *GetAssetResponseBody) SetStatus(v GetAssetResponseBodyStatusEnum) *GetAssetResponseBody {
	s.Status = &v
	return s
}

type GetAssetResponseBodyBuilder struct {
	s *GetAssetResponseBody
}

func NewGetAssetResponseBodyBuilder() *GetAssetResponseBodyBuilder {
	s := &GetAssetResponseBody{}
	b := &GetAssetResponseBodyBuilder{s: s}
	return b
}

func (b *GetAssetResponseBodyBuilder) UpdatedTime(v string) *GetAssetResponseBodyBuilder {
	b.s.UpdatedTime = &v
	return b
}

func (b *GetAssetResponseBodyBuilder) AssetId(v string) *GetAssetResponseBodyBuilder {
	b.s.AssetId = &v
	return b
}

func (b *GetAssetResponseBodyBuilder) GroupId(v string) *GetAssetResponseBodyBuilder {
	b.s.GroupId = &v
	return b
}

func (b *GetAssetResponseBodyBuilder) ErrorMessage(v string) *GetAssetResponseBodyBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetAssetResponseBodyBuilder) AssetName(v string) *GetAssetResponseBodyBuilder {
	b.s.AssetName = &v
	return b
}

func (b *GetAssetResponseBodyBuilder) CreatedTime(v string) *GetAssetResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetAssetResponseBodyBuilder) AssetUrl(v string) *GetAssetResponseBodyBuilder {
	b.s.AssetUrl = &v
	return b
}

func (b *GetAssetResponseBodyBuilder) AssetType(v GetAssetResponseBodyAssetTypeEnum) *GetAssetResponseBodyBuilder {
	b.s.AssetType = &v
	return b
}

func (b *GetAssetResponseBodyBuilder) Status(v GetAssetResponseBodyStatusEnum) *GetAssetResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *GetAssetResponseBodyBuilder) Build() *GetAssetResponseBody {
	return b.s
}
