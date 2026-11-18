// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAssetGroupBody struct {
	position.Body
	// 素材组类型，可选值：AIGC（虚拟人像素材库）、LivenessFace（真人素材库）。本接口仅支持创建AIGC类型，真人素材库需通过真人认证流程创建
	GroupType *string `json:"groupType,omitempty"`
	// 素材组名称，最长64个字符
	GroupName *string `json:"groupName,omitempty"`
	// 素材组简介，最长300个字符，非必填
	Description *string `json:"description,omitempty"`
}

func (s CreateAssetGroupBody) String() string {
	return utils.Beautify(s)
}

func (s CreateAssetGroupBody) GoString() string {
	return s.String()
}

func (s CreateAssetGroupBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAssetGroupBody) SetGroupType(v string) *CreateAssetGroupBody {
	s.GroupType = &v
	return s
}

func (s *CreateAssetGroupBody) SetGroupName(v string) *CreateAssetGroupBody {
	s.GroupName = &v
	return s
}

func (s *CreateAssetGroupBody) SetDescription(v string) *CreateAssetGroupBody {
	s.Description = &v
	return s
}

type CreateAssetGroupBodyBuilder struct {
	s *CreateAssetGroupBody
}

func NewCreateAssetGroupBodyBuilder() *CreateAssetGroupBodyBuilder {
	s := &CreateAssetGroupBody{}
	b := &CreateAssetGroupBodyBuilder{s: s}
	return b
}

func (b *CreateAssetGroupBodyBuilder) GroupType(v string) *CreateAssetGroupBodyBuilder {
	b.s.GroupType = &v
	return b
}

func (b *CreateAssetGroupBodyBuilder) GroupName(v string) *CreateAssetGroupBodyBuilder {
	b.s.GroupName = &v
	return b
}

func (b *CreateAssetGroupBodyBuilder) Description(v string) *CreateAssetGroupBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateAssetGroupBodyBuilder) Build() *CreateAssetGroupBody {
	return b.s
}
