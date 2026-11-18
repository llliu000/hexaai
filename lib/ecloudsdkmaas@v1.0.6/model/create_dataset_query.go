// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateDatasetQuery struct {
	position.Query
	// 项目ID,参考智算平台openApi文档，项目空间列表接口
	ProjectId *string `json:"projectId,omitempty"`
}

func (s CreateDatasetQuery) String() string {
	return utils.Beautify(s)
}

func (s CreateDatasetQuery) GoString() string {
	return s.String()
}

func (s CreateDatasetQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateDatasetQuery) SetProjectId(v string) *CreateDatasetQuery {
	s.ProjectId = &v
	return s
}

type CreateDatasetQueryBuilder struct {
	s *CreateDatasetQuery
}

func NewCreateDatasetQueryBuilder() *CreateDatasetQueryBuilder {
	s := &CreateDatasetQuery{}
	b := &CreateDatasetQueryBuilder{s: s}
	return b
}

func (b *CreateDatasetQueryBuilder) ProjectId(v string) *CreateDatasetQueryBuilder {
	b.s.ProjectId = &v
	return b
}

func (b *CreateDatasetQueryBuilder) Build() *CreateDatasetQuery {
	return b.s
}
