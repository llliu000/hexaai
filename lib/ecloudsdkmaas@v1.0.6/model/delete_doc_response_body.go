// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteDocResponseBody struct {

	// 操作结果,true 表示删除成功，false表示不存在或者无权限访问
	Result *bool `json:"result,omitempty"`
}

func (s DeleteDocResponseBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteDocResponseBody) GoString() string {
	return s.String()
}

func (s DeleteDocResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteDocResponseBody) SetResult(v bool) *DeleteDocResponseBody {
	s.Result = &v
	return s
}

type DeleteDocResponseBodyBuilder struct {
	s *DeleteDocResponseBody
}

func NewDeleteDocResponseBodyBuilder() *DeleteDocResponseBodyBuilder {
	s := &DeleteDocResponseBody{}
	b := &DeleteDocResponseBodyBuilder{s: s}
	return b
}

func (b *DeleteDocResponseBodyBuilder) Result(v bool) *DeleteDocResponseBodyBuilder {
	b.s.Result = &v
	return b
}

func (b *DeleteDocResponseBodyBuilder) Build() *DeleteDocResponseBody {
	return b.s
}
