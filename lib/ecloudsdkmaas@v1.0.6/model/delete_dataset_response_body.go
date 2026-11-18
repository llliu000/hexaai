// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteDatasetResponseBody struct {

	// 操作结果,true 表示删除成功，false表示不存在或者无权限访问
	Result *bool `json:"result,omitempty"`
}

func (s DeleteDatasetResponseBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteDatasetResponseBody) GoString() string {
	return s.String()
}

func (s DeleteDatasetResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteDatasetResponseBody) SetResult(v bool) *DeleteDatasetResponseBody {
	s.Result = &v
	return s
}

type DeleteDatasetResponseBodyBuilder struct {
	s *DeleteDatasetResponseBody
}

func NewDeleteDatasetResponseBodyBuilder() *DeleteDatasetResponseBodyBuilder {
	s := &DeleteDatasetResponseBody{}
	b := &DeleteDatasetResponseBodyBuilder{s: s}
	return b
}

func (b *DeleteDatasetResponseBodyBuilder) Result(v bool) *DeleteDatasetResponseBodyBuilder {
	b.s.Result = &v
	return b
}

func (b *DeleteDatasetResponseBodyBuilder) Build() *DeleteDatasetResponseBody {
	return b.s
}
