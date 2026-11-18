// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAssetGroupPath struct {
	position.Path
	// 素材组ID（火山资产组ID）
	GroupId *string `json:"groupId,omitempty"`
}

func (s GetAssetGroupPath) String() string {
	return utils.Beautify(s)
}

func (s GetAssetGroupPath) GoString() string {
	return s.String()
}

func (s GetAssetGroupPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAssetGroupPath) SetGroupId(v string) *GetAssetGroupPath {
	s.GroupId = &v
	return s
}

type GetAssetGroupPathBuilder struct {
	s *GetAssetGroupPath
}

func NewGetAssetGroupPathBuilder() *GetAssetGroupPathBuilder {
	s := &GetAssetGroupPath{}
	b := &GetAssetGroupPathBuilder{s: s}
	return b
}

func (b *GetAssetGroupPathBuilder) GroupId(v string) *GetAssetGroupPathBuilder {
	b.s.GroupId = &v
	return b
}

func (b *GetAssetGroupPathBuilder) Build() *GetAssetGroupPath {
	return b.s
}
