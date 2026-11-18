// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListAssetGroupsResponseDataGroupTypeEnum string

// List of GroupType
const (
	ListAssetGroupsResponseDataGroupTypeEnumAigc         ListAssetGroupsResponseDataGroupTypeEnum = "AIGC"
	ListAssetGroupsResponseDataGroupTypeEnumLivenessface ListAssetGroupsResponseDataGroupTypeEnum = "LivenessFace"
)

type ListAssetGroupsResponseData struct {

	// 更新时间，格式：yyyy-MM-dd HH:mm:ss
	UpdatedTime *string `json:"updatedTime,omitempty"`
	// AssetGroupTypeEnum
	GroupType *ListAssetGroupsResponseDataGroupTypeEnum `json:"groupType,omitempty"`
	// 资产组名称
	GroupName *string `json:"groupName,omitempty"`
	// 火山资产组ID
	GroupId *string `json:"groupId,omitempty"`
	// 资产组简介
	Description *string `json:"description,omitempty"`
	// 创建时间，格式：yyyy-MM-dd HH:mm:ss
	CreatedTime *string `json:"createdTime,omitempty"`
}

func (s ListAssetGroupsResponseData) String() string {
	return utils.Beautify(s)
}

func (s ListAssetGroupsResponseData) GoString() string {
	return s.String()
}

func (s ListAssetGroupsResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListAssetGroupsResponseData) SetUpdatedTime(v string) *ListAssetGroupsResponseData {
	s.UpdatedTime = &v
	return s
}

func (s *ListAssetGroupsResponseData) SetGroupType(v ListAssetGroupsResponseDataGroupTypeEnum) *ListAssetGroupsResponseData {
	s.GroupType = &v
	return s
}

func (s *ListAssetGroupsResponseData) SetGroupName(v string) *ListAssetGroupsResponseData {
	s.GroupName = &v
	return s
}

func (s *ListAssetGroupsResponseData) SetGroupId(v string) *ListAssetGroupsResponseData {
	s.GroupId = &v
	return s
}

func (s *ListAssetGroupsResponseData) SetDescription(v string) *ListAssetGroupsResponseData {
	s.Description = &v
	return s
}

func (s *ListAssetGroupsResponseData) SetCreatedTime(v string) *ListAssetGroupsResponseData {
	s.CreatedTime = &v
	return s
}

type ListAssetGroupsResponseDataBuilder struct {
	s *ListAssetGroupsResponseData
}

func NewListAssetGroupsResponseDataBuilder() *ListAssetGroupsResponseDataBuilder {
	s := &ListAssetGroupsResponseData{}
	b := &ListAssetGroupsResponseDataBuilder{s: s}
	return b
}

func (b *ListAssetGroupsResponseDataBuilder) UpdatedTime(v string) *ListAssetGroupsResponseDataBuilder {
	b.s.UpdatedTime = &v
	return b
}

func (b *ListAssetGroupsResponseDataBuilder) GroupType(v ListAssetGroupsResponseDataGroupTypeEnum) *ListAssetGroupsResponseDataBuilder {
	b.s.GroupType = &v
	return b
}

func (b *ListAssetGroupsResponseDataBuilder) GroupName(v string) *ListAssetGroupsResponseDataBuilder {
	b.s.GroupName = &v
	return b
}

func (b *ListAssetGroupsResponseDataBuilder) GroupId(v string) *ListAssetGroupsResponseDataBuilder {
	b.s.GroupId = &v
	return b
}

func (b *ListAssetGroupsResponseDataBuilder) Description(v string) *ListAssetGroupsResponseDataBuilder {
	b.s.Description = &v
	return b
}

func (b *ListAssetGroupsResponseDataBuilder) CreatedTime(v string) *ListAssetGroupsResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListAssetGroupsResponseDataBuilder) Build() *ListAssetGroupsResponseData {
	return b.s
}
