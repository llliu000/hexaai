// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListAssetGroupsBody struct {
	position.Body
	// 素材组类型过滤，可选值：AIGC（虚拟人像素材库）、LivenessFace（真人素材库）
	GroupType *string `json:"groupType,omitempty"`
	// 素材组名称模糊搜索关键字，最长64个字符，不传则不过滤
	GroupName *string `json:"groupName,omitempty"`
	// 页码（从 1 开始）
	PageNo *int32 `json:"pageNo,omitempty"`
	// 素材组ID列表，按ID精确查询
	GroupIds []string `json:"groupIds,omitempty"`
	// 每页行数
	PageSize *int32 `json:"pageSize,omitempty"`
}

func (s ListAssetGroupsBody) String() string {
	return utils.Beautify(s)
}

func (s ListAssetGroupsBody) GoString() string {
	return s.String()
}

func (s ListAssetGroupsBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListAssetGroupsBody) SetGroupType(v string) *ListAssetGroupsBody {
	s.GroupType = &v
	return s
}

func (s *ListAssetGroupsBody) SetGroupName(v string) *ListAssetGroupsBody {
	s.GroupName = &v
	return s
}

func (s *ListAssetGroupsBody) SetPageNo(v int32) *ListAssetGroupsBody {
	s.PageNo = &v
	return s
}

func (s *ListAssetGroupsBody) SetGroupIds(v []string) *ListAssetGroupsBody {
	s.GroupIds = v
	return s
}

func (s *ListAssetGroupsBody) SetPageSize(v int32) *ListAssetGroupsBody {
	s.PageSize = &v
	return s
}

type ListAssetGroupsBodyBuilder struct {
	s *ListAssetGroupsBody
}

func NewListAssetGroupsBodyBuilder() *ListAssetGroupsBodyBuilder {
	s := &ListAssetGroupsBody{}
	b := &ListAssetGroupsBodyBuilder{s: s}
	return b
}

func (b *ListAssetGroupsBodyBuilder) GroupType(v string) *ListAssetGroupsBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *ListAssetGroupsBodyBuilder) GroupName(v string) *ListAssetGroupsBodyBuilder {
	b.s.GroupName = &v
	return b
}

func (b *ListAssetGroupsBodyBuilder) PageNo(v int32) *ListAssetGroupsBodyBuilder {
	b.s.PageNo = &v
	return b
}

func (b *ListAssetGroupsBodyBuilder) GroupIds(v []string) *ListAssetGroupsBodyBuilder {
	b.s.GroupIds = v
	return b
}

func (b *ListAssetGroupsBodyBuilder) PageSize(v int32) *ListAssetGroupsBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListAssetGroupsBodyBuilder) Build() *ListAssetGroupsBody {
	return b.s
}
