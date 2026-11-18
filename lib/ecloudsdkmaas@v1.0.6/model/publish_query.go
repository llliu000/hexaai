// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PublishQuery struct {
	position.Query
	// 项目ID,参考智算平台openApi文档，项目空间列表接口
	ProjectId *string `json:"projectId,omitempty"`
}

func (s PublishQuery) String() string {
	return utils.Beautify(s)
}

func (s PublishQuery) GoString() string {
	return s.String()
}

func (s PublishQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PublishQuery) SetProjectId(v string) *PublishQuery {
	s.ProjectId = &v
	return s
}

type PublishQueryBuilder struct {
	s *PublishQuery
}

func NewPublishQueryBuilder() *PublishQueryBuilder {
	s := &PublishQuery{}
	b := &PublishQueryBuilder{s: s}
	return b
}

func (b *PublishQueryBuilder) ProjectId(v string) *PublishQueryBuilder {
	b.s.ProjectId = &v
	return b
}

func (b *PublishQueryBuilder) Build() *PublishQuery {
	return b.s
}
