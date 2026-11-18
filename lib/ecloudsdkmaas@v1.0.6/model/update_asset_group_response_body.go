// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateAssetGroupResponseBodyGroupTypeEnum string

// List of GroupType
const (
	UpdateAssetGroupResponseBodyGroupTypeEnumAigc         UpdateAssetGroupResponseBodyGroupTypeEnum = "AIGC"
	UpdateAssetGroupResponseBodyGroupTypeEnumLivenessface UpdateAssetGroupResponseBodyGroupTypeEnum = "LivenessFace"
)

type UpdateAssetGroupResponseBody struct {

	// 更新时间，格式：yyyy-MM-dd HH:mm:ss
	UpdatedTime *string `json:"updatedTime,omitempty"`
	// AssetGroupTypeEnum
	GroupType *UpdateAssetGroupResponseBodyGroupTypeEnum `json:"groupType,omitempty"`
	// 资产组名称
	GroupName *string `json:"groupName,omitempty"`
	// 火山资产组ID
	GroupId *string `json:"groupId,omitempty"`
	// 资产组简介
	Description *string `json:"description,omitempty"`
	// 创建时间，格式：yyyy-MM-dd HH:mm:ss
	CreatedTime *string `json:"createdTime,omitempty"`
}

func (s UpdateAssetGroupResponseBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateAssetGroupResponseBody) GoString() string {
	return s.String()
}

func (s UpdateAssetGroupResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateAssetGroupResponseBody) SetUpdatedTime(v string) *UpdateAssetGroupResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *UpdateAssetGroupResponseBody) SetGroupType(v UpdateAssetGroupResponseBodyGroupTypeEnum) *UpdateAssetGroupResponseBody {
	s.GroupType = &v
	return s
}

func (s *UpdateAssetGroupResponseBody) SetGroupName(v string) *UpdateAssetGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *UpdateAssetGroupResponseBody) SetGroupId(v string) *UpdateAssetGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *UpdateAssetGroupResponseBody) SetDescription(v string) *UpdateAssetGroupResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateAssetGroupResponseBody) SetCreatedTime(v string) *UpdateAssetGroupResponseBody {
	s.CreatedTime = &v
	return s
}

type UpdateAssetGroupResponseBodyBuilder struct {
	s *UpdateAssetGroupResponseBody
}

func NewUpdateAssetGroupResponseBodyBuilder() *UpdateAssetGroupResponseBodyBuilder {
	s := &UpdateAssetGroupResponseBody{}
	b := &UpdateAssetGroupResponseBodyBuilder{s: s}
	return b
}

func (b *UpdateAssetGroupResponseBodyBuilder) UpdatedTime(v string) *UpdateAssetGroupResponseBodyBuilder {
	b.s.UpdatedTime = &v
	return b
}

func (b *UpdateAssetGroupResponseBodyBuilder) GroupType(v UpdateAssetGroupResponseBodyGroupTypeEnum) *UpdateAssetGroupResponseBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *UpdateAssetGroupResponseBodyBuilder) GroupName(v string) *UpdateAssetGroupResponseBodyBuilder {
	b.s.GroupName = &v
	return b
}

func (b *UpdateAssetGroupResponseBodyBuilder) GroupId(v string) *UpdateAssetGroupResponseBodyBuilder {
	b.s.GroupId = &v
	return b
}

func (b *UpdateAssetGroupResponseBodyBuilder) Description(v string) *UpdateAssetGroupResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *UpdateAssetGroupResponseBodyBuilder) CreatedTime(v string) *UpdateAssetGroupResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *UpdateAssetGroupResponseBodyBuilder) Build() *UpdateAssetGroupResponseBody {
	return b.s
}
