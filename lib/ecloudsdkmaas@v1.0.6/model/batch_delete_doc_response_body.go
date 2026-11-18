// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchDeleteDocResponseBody struct {

	// 操作结果，true 表示执行成功，false表示执行失败，失败原因可以参考errorMessage或者联系客服人员
	Result *bool `json:"result,omitempty"`
}

func (s BatchDeleteDocResponseBody) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteDocResponseBody) GoString() string {
	return s.String()
}

func (s BatchDeleteDocResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteDocResponseBody) SetResult(v bool) *BatchDeleteDocResponseBody {
	s.Result = &v
	return s
}

type BatchDeleteDocResponseBodyBuilder struct {
	s *BatchDeleteDocResponseBody
}

func NewBatchDeleteDocResponseBodyBuilder() *BatchDeleteDocResponseBodyBuilder {
	s := &BatchDeleteDocResponseBody{}
	b := &BatchDeleteDocResponseBodyBuilder{s: s}
	return b
}

func (b *BatchDeleteDocResponseBodyBuilder) Result(v bool) *BatchDeleteDocResponseBodyBuilder {
	b.s.Result = &v
	return b
}

func (b *BatchDeleteDocResponseBodyBuilder) Build() *BatchDeleteDocResponseBody {
	return b.s
}
