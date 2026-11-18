// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteAssetGroupPath struct {
	position.Path
	// 素材组ID（火山资产组ID）
	GroupId *string `json:"groupId,omitempty"`
}

func (s DeleteAssetGroupPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteAssetGroupPath) GoString() string {
	return s.String()
}

func (s DeleteAssetGroupPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteAssetGroupPath) SetGroupId(v string) *DeleteAssetGroupPath {
	s.GroupId = &v
	return s
}

type DeleteAssetGroupPathBuilder struct {
	s *DeleteAssetGroupPath
}

func NewDeleteAssetGroupPathBuilder() *DeleteAssetGroupPathBuilder {
	s := &DeleteAssetGroupPath{}
	b := &DeleteAssetGroupPathBuilder{s: s}
	return b
}

func (b *DeleteAssetGroupPathBuilder) GroupId(v string) *DeleteAssetGroupPathBuilder {
	b.s.GroupId = &v
	return b
}

func (b *DeleteAssetGroupPathBuilder) Build() *DeleteAssetGroupPath {
	return b.s
}
