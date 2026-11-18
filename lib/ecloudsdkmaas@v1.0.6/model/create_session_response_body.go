// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateSessionResponseBody struct {

	// H5活体检测链接（用户复制到手机端完成认证）
	H5Link *string `json:"h5Link,omitempty"`
	// Token有效期（秒）
	ExpiresIn *int32 `json:"expiresIn,omitempty"`
	// 认证Token（用于查询认证状态）
	BytedToken *string `json:"bytedToken,omitempty"`
}

func (s CreateSessionResponseBody) String() string {
	return utils.Beautify(s)
}

func (s CreateSessionResponseBody) GoString() string {
	return s.String()
}

func (s CreateSessionResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSessionResponseBody) SetH5Link(v string) *CreateSessionResponseBody {
	s.H5Link = &v
	return s
}

func (s *CreateSessionResponseBody) SetExpiresIn(v int32) *CreateSessionResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *CreateSessionResponseBody) SetBytedToken(v string) *CreateSessionResponseBody {
	s.BytedToken = &v
	return s
}

type CreateSessionResponseBodyBuilder struct {
	s *CreateSessionResponseBody
}

func NewCreateSessionResponseBodyBuilder() *CreateSessionResponseBodyBuilder {
	s := &CreateSessionResponseBody{}
	b := &CreateSessionResponseBodyBuilder{s: s}
	return b
}

func (b *CreateSessionResponseBodyBuilder) H5Link(v string) *CreateSessionResponseBodyBuilder {
	b.s.H5Link = &v
	return b
}

func (b *CreateSessionResponseBodyBuilder) ExpiresIn(v int32) *CreateSessionResponseBodyBuilder {
	b.s.ExpiresIn = &v
	return b
}

func (b *CreateSessionResponseBodyBuilder) BytedToken(v string) *CreateSessionResponseBodyBuilder {
	b.s.BytedToken = &v
	return b
}

func (b *CreateSessionResponseBodyBuilder) Build() *CreateSessionResponseBody {
	return b.s
}
