// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAssetGroupResponseBodyGroupTypeEnum string

// List of GroupType
const (
	CreateAssetGroupResponseBodyGroupTypeEnumAigc         CreateAssetGroupResponseBodyGroupTypeEnum = "AIGC"
	CreateAssetGroupResponseBodyGroupTypeEnumLivenessface CreateAssetGroupResponseBodyGroupTypeEnum = "LivenessFace"
)

type CreateAssetGroupResponseBody struct {

	// 更新时间，格式：yyyy-MM-dd HH:mm:ss
	UpdatedTime *string `json:"updatedTime,omitempty"`
	// AssetGroupTypeEnum
	GroupType *CreateAssetGroupResponseBodyGroupTypeEnum `json:"groupType,omitempty"`
	// 资产组名称
	GroupName *string `json:"groupName,omitempty"`
	// 火山资产组ID
	GroupId *string `json:"groupId,omitempty"`
	// 资产组简介
	Description *string `json:"description,omitempty"`
	// 创建时间，格式：yyyy-MM-dd HH:mm:ss
	CreatedTime *string `json:"createdTime,omitempty"`
}

func (s CreateAssetGroupResponseBody) String() string {
	return utils.Beautify(s)
}

func (s CreateAssetGroupResponseBody) GoString() string {
	return s.String()
}

func (s CreateAssetGroupResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAssetGroupResponseBody) SetUpdatedTime(v string) *CreateAssetGroupResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *CreateAssetGroupResponseBody) SetGroupType(v CreateAssetGroupResponseBodyGroupTypeEnum) *CreateAssetGroupResponseBody {
	s.GroupType = &v
	return s
}

func (s *CreateAssetGroupResponseBody) SetGroupName(v string) *CreateAssetGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *CreateAssetGroupResponseBody) SetGroupId(v string) *CreateAssetGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateAssetGroupResponseBody) SetDescription(v string) *CreateAssetGroupResponseBody {
	s.Description = &v
	return s
}

func (s *CreateAssetGroupResponseBody) SetCreatedTime(v string) *CreateAssetGroupResponseBody {
	s.CreatedTime = &v
	return s
}

type CreateAssetGroupResponseBodyBuilder struct {
	s *CreateAssetGroupResponseBody
}

func NewCreateAssetGroupResponseBodyBuilder() *CreateAssetGroupResponseBodyBuilder {
	s := &CreateAssetGroupResponseBody{}
	b := &CreateAssetGroupResponseBodyBuilder{s: s}
	return b
}

func (b *CreateAssetGroupResponseBodyBuilder) UpdatedTime(v string) *CreateAssetGroupResponseBodyBuilder {
	b.s.UpdatedTime = &v
	return b
}

func (b *CreateAssetGroupResponseBodyBuilder) GroupType(v CreateAssetGroupResponseBodyGroupTypeEnum) *CreateAssetGroupResponseBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *CreateAssetGroupResponseBodyBuilder) GroupName(v string) *CreateAssetGroupResponseBodyBuilder {
	b.s.GroupName = &v
	return b
}

func (b *CreateAssetGroupResponseBodyBuilder) GroupId(v string) *CreateAssetGroupResponseBodyBuilder {
	b.s.GroupId = &v
	return b
}

func (b *CreateAssetGroupResponseBodyBuilder) Description(v string) *CreateAssetGroupResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateAssetGroupResponseBodyBuilder) CreatedTime(v string) *CreateAssetGroupResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *CreateAssetGroupResponseBodyBuilder) Build() *CreateAssetGroupResponseBody {
	return b.s
}
