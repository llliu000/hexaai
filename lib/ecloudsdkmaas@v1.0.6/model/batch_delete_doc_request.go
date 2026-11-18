// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchDeleteDocRequest struct {
	BatchDeleteDocPath *BatchDeleteDocPath `json:"batchDeleteDocPath,omitempty"`

	BatchDeleteDocBody *BatchDeleteDocBody `json:"batchDeleteDocBody,omitempty"`
}

func (s BatchDeleteDocRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteDocRequest) GoString() string {
	return s.String()
}

func (s BatchDeleteDocRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteDocRequest) SetBatchDeleteDocPath(v *BatchDeleteDocPath) *BatchDeleteDocRequest {
	s.BatchDeleteDocPath = v
	return s
}

func (s *BatchDeleteDocRequest) SetBatchDeleteDocBody(v *BatchDeleteDocBody) *BatchDeleteDocRequest {
	s.BatchDeleteDocBody = v
	return s
}

type BatchDeleteDocRequestBuilder struct {
	s *BatchDeleteDocRequest
}

func NewBatchDeleteDocRequestBuilder() *BatchDeleteDocRequestBuilder {
	s := &BatchDeleteDocRequest{}
	b := &BatchDeleteDocRequestBuilder{s: s}
	return b
}

func (b *BatchDeleteDocRequestBuilder) BatchDeleteDocPath(v *BatchDeleteDocPath) *BatchDeleteDocRequestBuilder {
	b.s.BatchDeleteDocPath = v
	return b
}

func (b *BatchDeleteDocRequestBuilder) BatchDeleteDocBody(v *BatchDeleteDocBody) *BatchDeleteDocRequestBuilder {
	b.s.BatchDeleteDocBody = v
	return b
}

func (b *BatchDeleteDocRequestBuilder) Build() *BatchDeleteDocRequest {
	return b.s
}
