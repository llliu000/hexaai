// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateAssetGroupBody struct {
	position.Body
	// 资产组名称，最长64个字符，不传则不更新
	GroupName *string `json:"groupName,omitempty"`
	// 资产组简介，最长300个字符，不传则不更新
	Description *string `json:"description,omitempty"`
}

func (s UpdateAssetGroupBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateAssetGroupBody) GoString() string {
	return s.String()
}

func (s UpdateAssetGroupBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateAssetGroupBody) SetGroupName(v string) *UpdateAssetGroupBody {
	s.GroupName = &v
	return s
}

func (s *UpdateAssetGroupBody) SetDescription(v string) *UpdateAssetGroupBody {
	s.Description = &v
	return s
}

type UpdateAssetGroupBodyBuilder struct {
	s *UpdateAssetGroupBody
}

func NewUpdateAssetGroupBodyBuilder() *UpdateAssetGroupBodyBuilder {
	s := &UpdateAssetGroupBody{}
	b := &UpdateAssetGroupBodyBuilder{s: s}
	return b
}

func (b *UpdateAssetGroupBodyBuilder) GroupName(v string) *UpdateAssetGroupBodyBuilder {
	b.s.GroupName = &v
	return b
}

func (b *UpdateAssetGroupBodyBuilder) Description(v string) *UpdateAssetGroupBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *UpdateAssetGroupBodyBuilder) Build() *UpdateAssetGroupBody {
	return b.s
}
