// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PageQuery struct {
	position.Query
	// 项目ID,参考智算平台openApi文档，项目空间列表接口，不传或者非法值会报“拒绝访问”之类的提示
	ProjectId *string `json:"projectId,omitempty"`
}

func (s PageQuery) String() string {
	return utils.Beautify(s)
}

func (s PageQuery) GoString() string {
	return s.String()
}

func (s PageQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PageQuery) SetProjectId(v string) *PageQuery {
	s.ProjectId = &v
	return s
}

type PageQueryBuilder struct {
	s *PageQuery
}

func NewPageQueryBuilder() *PageQueryBuilder {
	s := &PageQuery{}
	b := &PageQueryBuilder{s: s}
	return b
}

func (b *PageQueryBuilder) ProjectId(v string) *PageQueryBuilder {
	b.s.ProjectId = &v
	return b
}

func (b *PageQueryBuilder) Build() *PageQuery {
	return b.s
}
