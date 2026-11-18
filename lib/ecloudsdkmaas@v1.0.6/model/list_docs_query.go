// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListDocsQuery struct {
	position.Query
	// 关键词(模糊查询)，查询参数二选一即可
	DocName *string `json:"docName,omitempty"`
	// 文档 Id（精确匹配），查询参数二选一即可
	DocId *string `json:"docId,omitempty"`
}

func (s ListDocsQuery) String() string {
	return utils.Beautify(s)
}

func (s ListDocsQuery) GoString() string {
	return s.String()
}

func (s ListDocsQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListDocsQuery) SetDocName(v string) *ListDocsQuery {
	s.DocName = &v
	return s
}

func (s *ListDocsQuery) SetDocId(v string) *ListDocsQuery {
	s.DocId = &v
	return s
}

type ListDocsQueryBuilder struct {
	s *ListDocsQuery
}

func NewListDocsQueryBuilder() *ListDocsQueryBuilder {
	s := &ListDocsQuery{}
	b := &ListDocsQueryBuilder{s: s}
	return b
}

func (b *ListDocsQueryBuilder) DocName(v string) *ListDocsQueryBuilder {
	b.s.DocName = &v
	return b
}

func (b *ListDocsQueryBuilder) DocId(v string) *ListDocsQueryBuilder {
	b.s.DocId = &v
	return b
}

func (b *ListDocsQueryBuilder) Build() *ListDocsQuery {
	return b.s
}
