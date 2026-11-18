// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OnShelfResponseBody struct {

	// 操作结果
	Result *bool `json:"result,omitempty"`
}

func (s OnShelfResponseBody) String() string {
	return utils.Beautify(s)
}

func (s OnShelfResponseBody) GoString() string {
	return s.String()
}

func (s OnShelfResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OnShelfResponseBody) SetResult(v bool) *OnShelfResponseBody {
	s.Result = &v
	return s
}

type OnShelfResponseBodyBuilder struct {
	s *OnShelfResponseBody
}

func NewOnShelfResponseBodyBuilder() *OnShelfResponseBodyBuilder {
	s := &OnShelfResponseBody{}
	b := &OnShelfResponseBodyBuilder{s: s}
	return b
}

func (b *OnShelfResponseBodyBuilder) Result(v bool) *OnShelfResponseBodyBuilder {
	b.s.Result = &v
	return b
}

func (b *OnShelfResponseBodyBuilder) Build() *OnShelfResponseBody {
	return b.s
}
