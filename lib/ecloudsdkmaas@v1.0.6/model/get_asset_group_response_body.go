// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAssetGroupResponseBodyGroupTypeEnum string

// List of GroupType
const (
	GetAssetGroupResponseBodyGroupTypeEnumAigc         GetAssetGroupResponseBodyGroupTypeEnum = "AIGC"
	GetAssetGroupResponseBodyGroupTypeEnumLivenessface GetAssetGroupResponseBodyGroupTypeEnum = "LivenessFace"
)

type GetAssetGroupResponseBody struct {

	// 更新时间，格式：yyyy-MM-dd HH:mm:ss
	UpdatedTime *string `json:"updatedTime,omitempty"`
	// AssetGroupTypeEnum
	GroupType *GetAssetGroupResponseBodyGroupTypeEnum `json:"groupType,omitempty"`
	// 资产组名称
	GroupName *string `json:"groupName,omitempty"`
	// 火山资产组ID
	GroupId *string `json:"groupId,omitempty"`
	// 资产组简介
	Description *string `json:"description,omitempty"`
	// 创建时间，格式：yyyy-MM-dd HH:mm:ss
	CreatedTime *string `json:"createdTime,omitempty"`
	// 主键ID（内部字段，OpenAPI不返回）
	Id *string `json:"id,omitempty"`
}

func (s GetAssetGroupResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetAssetGroupResponseBody) GoString() string {
	return s.String()
}

func (s GetAssetGroupResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAssetGroupResponseBody) SetUpdatedTime(v string) *GetAssetGroupResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *GetAssetGroupResponseBody) SetGroupType(v GetAssetGroupResponseBodyGroupTypeEnum) *GetAssetGroupResponseBody {
	s.GroupType = &v
	return s
}

func (s *GetAssetGroupResponseBody) SetGroupName(v string) *GetAssetGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *GetAssetGroupResponseBody) SetGroupId(v string) *GetAssetGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *GetAssetGroupResponseBody) SetDescription(v string) *GetAssetGroupResponseBody {
	s.Description = &v
	return s
}

func (s *GetAssetGroupResponseBody) SetCreatedTime(v string) *GetAssetGroupResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetAssetGroupResponseBody) SetId(v string) *GetAssetGroupResponseBody {
	s.Id = &v
	return s
}

type GetAssetGroupResponseBodyBuilder struct {
	s *GetAssetGroupResponseBody
}

func NewGetAssetGroupResponseBodyBuilder() *GetAssetGroupResponseBodyBuilder {
	s := &GetAssetGroupResponseBody{}
	b := &GetAssetGroupResponseBodyBuilder{s: s}
	return b
}

func (b *GetAssetGroupResponseBodyBuilder) UpdatedTime(v string) *GetAssetGroupResponseBodyBuilder {
	b.s.UpdatedTime = &v
	return b
}

func (b *GetAssetGroupResponseBodyBuilder) GroupType(v GetAssetGroupResponseBodyGroupTypeEnum) *GetAssetGroupResponseBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *GetAssetGroupResponseBodyBuilder) GroupName(v string) *GetAssetGroupResponseBodyBuilder {
	b.s.GroupName = &v
	return b
}

func (b *GetAssetGroupResponseBodyBuilder) GroupId(v string) *GetAssetGroupResponseBodyBuilder {
	b.s.GroupId = &v
	return b
}

func (b *GetAssetGroupResponseBodyBuilder) Description(v string) *GetAssetGroupResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *GetAssetGroupResponseBodyBuilder) CreatedTime(v string) *GetAssetGroupResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetAssetGroupResponseBodyBuilder) Id(v string) *GetAssetGroupResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetAssetGroupResponseBodyBuilder) Build() *GetAssetGroupResponseBody {
	return b.s
}
