// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchDeleteDocPath struct {
	position.Path

	DatasetId *string `json:"datasetId,omitempty"`
}

func (s BatchDeleteDocPath) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteDocPath) GoString() string {
	return s.String()
}

func (s BatchDeleteDocPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteDocPath) SetDatasetId(v string) *BatchDeleteDocPath {
	s.DatasetId = &v
	return s
}

type BatchDeleteDocPathBuilder struct {
	s *BatchDeleteDocPath
}

func NewBatchDeleteDocPathBuilder() *BatchDeleteDocPathBuilder {
	s := &BatchDeleteDocPath{}
	b := &BatchDeleteDocPathBuilder{s: s}
	return b
}

func (b *BatchDeleteDocPathBuilder) DatasetId(v string) *BatchDeleteDocPathBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *BatchDeleteDocPathBuilder) Build() *BatchDeleteDocPath {
	return b.s
}
