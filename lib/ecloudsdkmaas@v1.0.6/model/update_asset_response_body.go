// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateAssetResponseBodyAssetTypeEnum string

// List of AssetType
const (
	UpdateAssetResponseBodyAssetTypeEnumImage UpdateAssetResponseBodyAssetTypeEnum = "Image"
	UpdateAssetResponseBodyAssetTypeEnumVideo UpdateAssetResponseBodyAssetTypeEnum = "Video"
	UpdateAssetResponseBodyAssetTypeEnumAudio UpdateAssetResponseBodyAssetTypeEnum = "Audio"
)

type UpdateAssetResponseBodyStatusEnum string

// List of Status
const (
	UpdateAssetResponseBodyStatusEnumProcessing UpdateAssetResponseBodyStatusEnum = "PROCESSING"
	UpdateAssetResponseBodyStatusEnumActive     UpdateAssetResponseBodyStatusEnum = "ACTIVE"
	UpdateAssetResponseBodyStatusEnumFailed     UpdateAssetResponseBodyStatusEnum = "FAILED"
)

type UpdateAssetResponseBody struct {

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
	AssetType *UpdateAssetResponseBodyAssetTypeEnum `json:"assetType,omitempty"`
	// AssetStatusEnum
	Status *UpdateAssetResponseBodyStatusEnum `json:"status,omitempty"`
}

func (s UpdateAssetResponseBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateAssetResponseBody) GoString() string {
	return s.String()
}

func (s UpdateAssetResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateAssetResponseBody) SetUpdatedTime(v string) *UpdateAssetResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *UpdateAssetResponseBody) SetAssetId(v string) *UpdateAssetResponseBody {
	s.AssetId = &v
	return s
}

func (s *UpdateAssetResponseBody) SetGroupId(v string) *UpdateAssetResponseBody {
	s.GroupId = &v
	return s
}

func (s *UpdateAssetResponseBody) SetErrorMessage(v string) *UpdateAssetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAssetResponseBody) SetAssetName(v string) *UpdateAssetResponseBody {
	s.AssetName = &v
	return s
}

func (s *UpdateAssetResponseBody) SetCreatedTime(v string) *UpdateAssetResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *UpdateAssetResponseBody) SetAssetUrl(v string) *UpdateAssetResponseBody {
	s.AssetUrl = &v
	return s
}

func (s *UpdateAssetResponseBody) SetAssetType(v UpdateAssetResponseBodyAssetTypeEnum) *UpdateAssetResponseBody {
	s.AssetType = &v
	return s
}

func (s *UpdateAssetResponseBody) SetStatus(v UpdateAssetResponseBodyStatusEnum) *UpdateAssetResponseBody {
	s.Status = &v
	return s
}

type UpdateAssetResponseBodyBuilder struct {
	s *UpdateAssetResponseBody
}

func NewUpdateAssetResponseBodyBuilder() *UpdateAssetResponseBodyBuilder {
	s := &UpdateAssetResponseBody{}
	b := &UpdateAssetResponseBodyBuilder{s: s}
	return b
}

func (b *UpdateAssetResponseBodyBuilder) UpdatedTime(v string) *UpdateAssetResponseBodyBuilder {
	b.s.UpdatedTime = &v
	return b
}

func (b *UpdateAssetResponseBodyBuilder) AssetId(v string) *UpdateAssetResponseBodyBuilder {
	b.s.AssetId = &v
	return b
}

func (b *UpdateAssetResponseBodyBuilder) GroupId(v string) *UpdateAssetResponseBodyBuilder {
	b.s.GroupId = &v
	return b
}

func (b *UpdateAssetResponseBodyBuilder) ErrorMessage(v string) *UpdateAssetResponseBodyBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateAssetResponseBodyBuilder) AssetName(v string) *UpdateAssetResponseBodyBuilder {
	b.s.AssetName = &v
	return b
}

func (b *UpdateAssetResponseBodyBuilder) CreatedTime(v string) *UpdateAssetResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *UpdateAssetResponseBodyBuilder) AssetUrl(v string) *UpdateAssetResponseBodyBuilder {
	b.s.AssetUrl = &v
	return b
}

func (b *UpdateAssetResponseBodyBuilder) AssetType(v UpdateAssetResponseBodyAssetTypeEnum) *UpdateAssetResponseBodyBuilder {
	b.s.AssetType = &v
	return b
}

func (b *UpdateAssetResponseBodyBuilder) Status(v UpdateAssetResponseBodyStatusEnum) *UpdateAssetResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *UpdateAssetResponseBodyBuilder) Build() *UpdateAssetResponseBody {
	return b.s
}
