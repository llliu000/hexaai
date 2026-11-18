// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PublishResponseBody struct {

	// 操作结果
	Result *bool `json:"result,omitempty"`
}

func (s PublishResponseBody) String() string {
	return utils.Beautify(s)
}

func (s PublishResponseBody) GoString() string {
	return s.String()
}

func (s PublishResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PublishResponseBody) SetResult(v bool) *PublishResponseBody {
	s.Result = &v
	return s
}

type PublishResponseBodyBuilder struct {
	s *PublishResponseBody
}

func NewPublishResponseBodyBuilder() *PublishResponseBodyBuilder {
	s := &PublishResponseBody{}
	b := &PublishResponseBodyBuilder{s: s}
	return b
}

func (b *PublishResponseBodyBuilder) Result(v bool) *PublishResponseBodyBuilder {
	b.s.Result = &v
	return b
}

func (b *PublishResponseBodyBuilder) Build() *PublishResponseBody {
	return b.s
}
