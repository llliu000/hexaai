// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListAssetsBody struct {
	position.Body
	// 素材组类型，可选值：AIGC（虚拟人像素材库）、LivenessFace（真人素材库）
	GroupType *string `json:"groupType,omitempty"`
	// 页码（从 1 开始）
	PageNo *int32 `json:"pageNo,omitempty"`
	// 素材组ID列表，按ID精确查询，不传则不过滤
	GroupIds []string `json:"groupIds,omitempty"`
	// 每页行数
	PageSize *int32 `json:"pageSize,omitempty"`
	// 素材名称模糊搜索关键字，最长64个字符，不传则不过滤
	AssetName *string `json:"assetName,omitempty"`
	// 素材状态列表，可选值：PROCESSING（处理中）、ACTIVE（可用）、FAILED（失败），不传则不过滤
	Statuses []string `json:"statuses,omitempty"`
}

func (s ListAssetsBody) String() string {
	return utils.Beautify(s)
}

func (s ListAssetsBody) GoString() string {
	return s.String()
}

func (s ListAssetsBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListAssetsBody) SetGroupType(v string) *ListAssetsBody {
	s.GroupType = &v
	return s
}

func (s *ListAssetsBody) SetPageNo(v int32) *ListAssetsBody {
	s.PageNo = &v
	return s
}

func (s *ListAssetsBody) SetGroupIds(v []string) *ListAssetsBody {
	s.GroupIds = v
	return s
}

func (s *ListAssetsBody) SetPageSize(v int32) *ListAssetsBody {
	s.PageSize = &v
	return s
}

func (s *ListAssetsBody) SetAssetName(v string) *ListAssetsBody {
	s.AssetName = &v
	return s
}

func (s *ListAssetsBody) SetStatuses(v []string) *ListAssetsBody {
	s.Statuses = v
	return s
}

type ListAssetsBodyBuilder struct {
	s *ListAssetsBody
}

func NewListAssetsBodyBuilder() *ListAssetsBodyBuilder {
	s := &ListAssetsBody{}
	b := &ListAssetsBodyBuilder{s: s}
	return b
}

func (b *ListAssetsBodyBuilder) GroupType(v string) *ListAssetsBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *ListAssetsBodyBuilder) PageNo(v int32) *ListAssetsBodyBuilder {
	b.s.PageNo = &v
	return b
}

func (b *ListAssetsBodyBuilder) GroupIds(v []string) *ListAssetsBodyBuilder {
	b.s.GroupIds = v
	return b
}

func (b *ListAssetsBodyBuilder) PageSize(v int32) *ListAssetsBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListAssetsBodyBuilder) AssetName(v string) *ListAssetsBodyBuilder {
	b.s.AssetName = &v
	return b
}

func (b *ListAssetsBodyBuilder) Statuses(v []string) *ListAssetsBodyBuilder {
	b.s.Statuses = v
	return b
}

func (b *ListAssetsBodyBuilder) Build() *ListAssetsBody {
	return b.s
}
