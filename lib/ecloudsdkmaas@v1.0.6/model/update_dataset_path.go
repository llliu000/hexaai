// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateDatasetPath struct {
	position.Path
	// 知识库Id（URL路径参数）
	DatasetId *string `json:"datasetId,omitempty"`
}

func (s UpdateDatasetPath) String() string {
	return utils.Beautify(s)
}

func (s UpdateDatasetPath) GoString() string {
	return s.String()
}

func (s UpdateDatasetPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateDatasetPath) SetDatasetId(v string) *UpdateDatasetPath {
	s.DatasetId = &v
	return s
}

type UpdateDatasetPathBuilder struct {
	s *UpdateDatasetPath
}

func NewUpdateDatasetPathBuilder() *UpdateDatasetPathBuilder {
	s := &UpdateDatasetPath{}
	b := &UpdateDatasetPathBuilder{s: s}
	return b
}

func (b *UpdateDatasetPathBuilder) DatasetId(v string) *UpdateDatasetPathBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *UpdateDatasetPathBuilder) Build() *UpdateDatasetPath {
	return b.s
}
