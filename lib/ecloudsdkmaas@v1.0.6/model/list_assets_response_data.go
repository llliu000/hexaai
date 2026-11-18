// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListAssetsResponseDataAssetTypeEnum string

// List of AssetType
const (
	ListAssetsResponseDataAssetTypeEnumImage ListAssetsResponseDataAssetTypeEnum = "Image"
	ListAssetsResponseDataAssetTypeEnumVideo ListAssetsResponseDataAssetTypeEnum = "Video"
	ListAssetsResponseDataAssetTypeEnumAudio ListAssetsResponseDataAssetTypeEnum = "Audio"
)

type ListAssetsResponseDataStatusEnum string

// List of Status
const (
	ListAssetsResponseDataStatusEnumProcessing ListAssetsResponseDataStatusEnum = "PROCESSING"
	ListAssetsResponseDataStatusEnumActive     ListAssetsResponseDataStatusEnum = "ACTIVE"
	ListAssetsResponseDataStatusEnumFailed     ListAssetsResponseDataStatusEnum = "FAILED"
)

type ListAssetsResponseData struct {

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
	AssetType *ListAssetsResponseDataAssetTypeEnum `json:"assetType,omitempty"`
	// AssetStatusEnum
	Status *ListAssetsResponseDataStatusEnum `json:"status,omitempty"`
}

func (s ListAssetsResponseData) String() string {
	return utils.Beautify(s)
}

func (s ListAssetsResponseData) GoString() string {
	return s.String()
}

func (s ListAssetsResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListAssetsResponseData) SetUpdatedTime(v string) *ListAssetsResponseData {
	s.UpdatedTime = &v
	return s
}

func (s *ListAssetsResponseData) SetAssetId(v string) *ListAssetsResponseData {
	s.AssetId = &v
	return s
}

func (s *ListAssetsResponseData) SetGroupId(v string) *ListAssetsResponseData {
	s.GroupId = &v
	return s
}

func (s *ListAssetsResponseData) SetErrorMessage(v string) *ListAssetsResponseData {
	s.ErrorMessage = &v
	return s
}

func (s *ListAssetsResponseData) SetAssetName(v string) *ListAssetsResponseData {
	s.AssetName = &v
	return s
}

func (s *ListAssetsResponseData) SetCreatedTime(v string) *ListAssetsResponseData {
	s.CreatedTime = &v
	return s
}

func (s *ListAssetsResponseData) SetAssetUrl(v string) *ListAssetsResponseData {
	s.AssetUrl = &v
	return s
}

func (s *ListAssetsResponseData) SetAssetType(v ListAssetsResponseDataAssetTypeEnum) *ListAssetsResponseData {
	s.AssetType = &v
	return s
}

func (s *ListAssetsResponseData) SetStatus(v ListAssetsResponseDataStatusEnum) *ListAssetsResponseData {
	s.Status = &v
	return s
}

type ListAssetsResponseDataBuilder struct {
	s *ListAssetsResponseData
}

func NewListAssetsResponseDataBuilder() *ListAssetsResponseDataBuilder {
	s := &ListAssetsResponseData{}
	b := &ListAssetsResponseDataBuilder{s: s}
	return b
}

func (b *ListAssetsResponseDataBuilder) UpdatedTime(v string) *ListAssetsResponseDataBuilder {
	b.s.UpdatedTime = &v
	return b
}

func (b *ListAssetsResponseDataBuilder) AssetId(v string) *ListAssetsResponseDataBuilder {
	b.s.AssetId = &v
	return b
}

func (b *ListAssetsResponseDataBuilder) GroupId(v string) *ListAssetsResponseDataBuilder {
	b.s.GroupId = &v
	return b
}

func (b *ListAssetsResponseDataBuilder) ErrorMessage(v string) *ListAssetsResponseDataBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListAssetsResponseDataBuilder) AssetName(v string) *ListAssetsResponseDataBuilder {
	b.s.AssetName = &v
	return b
}

func (b *ListAssetsResponseDataBuilder) CreatedTime(v string) *ListAssetsResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListAssetsResponseDataBuilder) AssetUrl(v string) *ListAssetsResponseDataBuilder {
	b.s.AssetUrl = &v
	return b
}

func (b *ListAssetsResponseDataBuilder) AssetType(v ListAssetsResponseDataAssetTypeEnum) *ListAssetsResponseDataBuilder {
	b.s.AssetType = &v
	return b
}

func (b *ListAssetsResponseDataBuilder) Status(v ListAssetsResponseDataStatusEnum) *ListAssetsResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *ListAssetsResponseDataBuilder) Build() *ListAssetsResponseData {
	return b.s
}
