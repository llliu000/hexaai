// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchDeleteDocBody struct {
	position.Body
	// 文档id列表
	DocIds []string `json:"docIds,omitempty"`
}

func (s BatchDeleteDocBody) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteDocBody) GoString() string {
	return s.String()
}

func (s BatchDeleteDocBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteDocBody) SetDocIds(v []string) *BatchDeleteDocBody {
	s.DocIds = v
	return s
}

type BatchDeleteDocBodyBuilder struct {
	s *BatchDeleteDocBody
}

func NewBatchDeleteDocBodyBuilder() *BatchDeleteDocBodyBuilder {
	s := &BatchDeleteDocBody{}
	b := &BatchDeleteDocBodyBuilder{s: s}
	return b
}

func (b *BatchDeleteDocBodyBuilder) DocIds(v []string) *BatchDeleteDocBodyBuilder {
	b.s.DocIds = v
	return b
}

func (b *BatchDeleteDocBodyBuilder) Build() *BatchDeleteDocBody {
	return b.s
}
