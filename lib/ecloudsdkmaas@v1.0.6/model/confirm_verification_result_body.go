// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConfirmVerificationResultBody struct {
	position.Body
	// 真人认证Token，由创建H5认证会话接口返回
	BytedToken *string `json:"bytedToken,omitempty"`
}

func (s ConfirmVerificationResultBody) String() string {
	return utils.Beautify(s)
}

func (s ConfirmVerificationResultBody) GoString() string {
	return s.String()
}

func (s ConfirmVerificationResultBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConfirmVerificationResultBody) SetBytedToken(v string) *ConfirmVerificationResultBody {
	s.BytedToken = &v
	return s
}

type ConfirmVerificationResultBodyBuilder struct {
	s *ConfirmVerificationResultBody
}

func NewConfirmVerificationResultBodyBuilder() *ConfirmVerificationResultBodyBuilder {
	s := &ConfirmVerificationResultBody{}
	b := &ConfirmVerificationResultBodyBuilder{s: s}
	return b
}

func (b *ConfirmVerificationResultBodyBuilder) BytedToken(v string) *ConfirmVerificationResultBodyBuilder {
	b.s.BytedToken = &v
	return b
}

func (b *ConfirmVerificationResultBodyBuilder) Build() *ConfirmVerificationResultBody {
	return b.s
}
