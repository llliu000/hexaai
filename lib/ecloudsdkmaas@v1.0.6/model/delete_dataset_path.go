// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteDatasetPath struct {
	position.Path
	// 知识库Id（URL路径参数）
	DatasetId *string `json:"datasetId,omitempty"`
}

func (s DeleteDatasetPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteDatasetPath) GoString() string {
	return s.String()
}

func (s DeleteDatasetPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteDatasetPath) SetDatasetId(v string) *DeleteDatasetPath {
	s.DatasetId = &v
	return s
}

type DeleteDatasetPathBuilder struct {
	s *DeleteDatasetPath
}

func NewDeleteDatasetPathBuilder() *DeleteDatasetPathBuilder {
	s := &DeleteDatasetPath{}
	b := &DeleteDatasetPathBuilder{s: s}
	return b
}

func (b *DeleteDatasetPathBuilder) DatasetId(v string) *DeleteDatasetPathBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *DeleteDatasetPathBuilder) Build() *DeleteDatasetPath {
	return b.s
}
