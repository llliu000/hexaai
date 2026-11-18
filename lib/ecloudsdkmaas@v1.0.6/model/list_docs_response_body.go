// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListDocsResponseBodyStatusEnum string

// List of Status
const (
	ListDocsResponseBodyStatusEnumWaiting   ListDocsResponseBodyStatusEnum = "WAITING"
	ListDocsResponseBodyStatusEnumQueuing   ListDocsResponseBodyStatusEnum = "QUEUING"
	ListDocsResponseBodyStatusEnumParsing   ListDocsResponseBodyStatusEnum = "PARSING"
	ListDocsResponseBodyStatusEnumCleaning  ListDocsResponseBodyStatusEnum = "CLEANING"
	ListDocsResponseBodyStatusEnumSplitting ListDocsResponseBodyStatusEnum = "SPLITTING"
	ListDocsResponseBodyStatusEnumIndexing  ListDocsResponseBodyStatusEnum = "INDEXING"
	ListDocsResponseBodyStatusEnumCompleted ListDocsResponseBodyStatusEnum = "COMPLETED"
	ListDocsResponseBodyStatusEnumExceeded  ListDocsResponseBodyStatusEnum = "EXCEEDED"
	ListDocsResponseBodyStatusEnumEmpty     ListDocsResponseBodyStatusEnum = "EMPTY"
	ListDocsResponseBodyStatusEnumError     ListDocsResponseBodyStatusEnum = "ERROR"
	ListDocsResponseBodyStatusEnumDeleting  ListDocsResponseBodyStatusEnum = "DELETING"
	ListDocsResponseBodyStatusEnumDeleted   ListDocsResponseBodyStatusEnum = "DELETED"
)

type ListDocsResponseBody struct {

	// 创建时间（UNIX 毫秒时间戳）
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// 文档名称
	DocName *string `json:"docName,omitempty"`
	// 文档扩展名
	Extension *string `json:"extension,omitempty"`
	// 文档大小（字节数）
	Size *int64 `json:"size,omitempty"`
	// 文档 Id
	DocId *string `json:"docId,omitempty"`
	// 文档状态（INDEXING、ERROR、COMPLETED）
	Status *ListDocsResponseBodyStatusEnum `json:"status,omitempty"`
}

func (s ListDocsResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListDocsResponseBody) GoString() string {
	return s.String()
}

func (s ListDocsResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListDocsResponseBody) SetCreatedAt(v int64) *ListDocsResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *ListDocsResponseBody) SetDocName(v string) *ListDocsResponseBody {
	s.DocName = &v
	return s
}

func (s *ListDocsResponseBody) SetExtension(v string) *ListDocsResponseBody {
	s.Extension = &v
	return s
}

func (s *ListDocsResponseBody) SetSize(v int64) *ListDocsResponseBody {
	s.Size = &v
	return s
}

func (s *ListDocsResponseBody) SetDocId(v string) *ListDocsResponseBody {
	s.DocId = &v
	return s
}

func (s *ListDocsResponseBody) SetStatus(v ListDocsResponseBodyStatusEnum) *ListDocsResponseBody {
	s.Status = &v
	return s
}

type ListDocsResponseBodyBuilder struct {
	s *ListDocsResponseBody
}

func NewListDocsResponseBodyBuilder() *ListDocsResponseBodyBuilder {
	s := &ListDocsResponseBody{}
	b := &ListDocsResponseBodyBuilder{s: s}
	return b
}

func (b *ListDocsResponseBodyBuilder) CreatedAt(v int64) *ListDocsResponseBodyBuilder {
	b.s.CreatedAt = &v
	return b
}

func (b *ListDocsResponseBodyBuilder) DocName(v string) *ListDocsResponseBodyBuilder {
	b.s.DocName = &v
	return b
}

func (b *ListDocsResponseBodyBuilder) Extension(v string) *ListDocsResponseBodyBuilder {
	b.s.Extension = &v
	return b
}

func (b *ListDocsResponseBodyBuilder) Size(v int64) *ListDocsResponseBodyBuilder {
	b.s.Size = &v
	return b
}

func (b *ListDocsResponseBodyBuilder) DocId(v string) *ListDocsResponseBodyBuilder {
	b.s.DocId = &v
	return b
}

func (b *ListDocsResponseBodyBuilder) Status(v ListDocsResponseBodyStatusEnum) *ListDocsResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *ListDocsResponseBodyBuilder) Build() *ListDocsResponseBody {
	return b.s
}
