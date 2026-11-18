// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateSessionBody struct {
	position.Body
	// 回调URL（H5表单页面地址），默认回调方式不传 callbackUrl 或留空，自定义回调方式必传
	CallbackUrl *string `json:"callbackUrl,omitempty"`
}

func (s CreateSessionBody) String() string {
	return utils.Beautify(s)
}

func (s CreateSessionBody) GoString() string {
	return s.String()
}

func (s CreateSessionBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSessionBody) SetCallbackUrl(v string) *CreateSessionBody {
	s.CallbackUrl = &v
	return s
}

type CreateSessionBodyBuilder struct {
	s *CreateSessionBody
}

func NewCreateSessionBodyBuilder() *CreateSessionBodyBuilder {
	s := &CreateSessionBody{}
	b := &CreateSessionBodyBuilder{s: s}
	return b
}

func (b *CreateSessionBodyBuilder) CallbackUrl(v string) *CreateSessionBodyBuilder {
	b.s.CallbackUrl = &v
	return b
}

func (b *CreateSessionBodyBuilder) Build() *CreateSessionBody {
	return b.s
}
