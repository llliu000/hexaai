// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAssetGroupByBytedTokenBody struct {
	position.Body
	// 真人认证Token，由创建H5认证会话接口返回
	BytedToken *string `json:"bytedToken,omitempty"`
}

func (s GetAssetGroupByBytedTokenBody) String() string {
	return utils.Beautify(s)
}

func (s GetAssetGroupByBytedTokenBody) GoString() string {
	return s.String()
}

func (s GetAssetGroupByBytedTokenBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAssetGroupByBytedTokenBody) SetBytedToken(v string) *GetAssetGroupByBytedTokenBody {
	s.BytedToken = &v
	return s
}

type GetAssetGroupByBytedTokenBodyBuilder struct {
	s *GetAssetGroupByBytedTokenBody
}

func NewGetAssetGroupByBytedTokenBodyBuilder() *GetAssetGroupByBytedTokenBodyBuilder {
	s := &GetAssetGroupByBytedTokenBody{}
	b := &GetAssetGroupByBytedTokenBodyBuilder{s: s}
	return b
}

func (b *GetAssetGroupByBytedTokenBodyBuilder) BytedToken(v string) *GetAssetGroupByBytedTokenBodyBuilder {
	b.s.BytedToken = &v
	return b
}

func (b *GetAssetGroupByBytedTokenBodyBuilder) Build() *GetAssetGroupByBytedTokenBody {
	return b.s
}
