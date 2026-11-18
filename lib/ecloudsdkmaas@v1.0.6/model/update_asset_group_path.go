// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateAssetGroupPath struct {
	position.Path
	// 素材组ID（火山资产组ID）
	GroupId *string `json:"groupId,omitempty"`
}

func (s UpdateAssetGroupPath) String() string {
	return utils.Beautify(s)
}

func (s UpdateAssetGroupPath) GoString() string {
	return s.String()
}

func (s UpdateAssetGroupPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateAssetGroupPath) SetGroupId(v string) *UpdateAssetGroupPath {
	s.GroupId = &v
	return s
}

type UpdateAssetGroupPathBuilder struct {
	s *UpdateAssetGroupPath
}

func NewUpdateAssetGroupPathBuilder() *UpdateAssetGroupPathBuilder {
	s := &UpdateAssetGroupPath{}
	b := &UpdateAssetGroupPathBuilder{s: s}
	return b
}

func (b *UpdateAssetGroupPathBuilder) GroupId(v string) *UpdateAssetGroupPathBuilder {
	b.s.GroupId = &v
	return b
}

func (b *UpdateAssetGroupPathBuilder) Build() *UpdateAssetGroupPath {
	return b.s
}
