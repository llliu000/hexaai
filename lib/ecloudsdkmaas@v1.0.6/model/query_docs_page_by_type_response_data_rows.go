// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryDocsPageByTypeResponseDataRowsStatusEnum string

// List of Status
const (
	QueryDocsPageByTypeResponseDataRowsStatusEnumWaiting          QueryDocsPageByTypeResponseDataRowsStatusEnum = "WAITING"
	QueryDocsPageByTypeResponseDataRowsStatusEnumQueuing          QueryDocsPageByTypeResponseDataRowsStatusEnum = "QUEUING"
	QueryDocsPageByTypeResponseDataRowsStatusEnumParsing          QueryDocsPageByTypeResponseDataRowsStatusEnum = "PARSING"
	QueryDocsPageByTypeResponseDataRowsStatusEnumCleaning         QueryDocsPageByTypeResponseDataRowsStatusEnum = "CLEANING"
	QueryDocsPageByTypeResponseDataRowsStatusEnumSplitting        QueryDocsPageByTypeResponseDataRowsStatusEnum = "SPLITTING"
	QueryDocsPageByTypeResponseDataRowsStatusEnumIndexing         QueryDocsPageByTypeResponseDataRowsStatusEnum = "INDEXING"
	QueryDocsPageByTypeResponseDataRowsStatusEnumCompleted        QueryDocsPageByTypeResponseDataRowsStatusEnum = "COMPLETED"
	QueryDocsPageByTypeResponseDataRowsStatusEnumExceeded         QueryDocsPageByTypeResponseDataRowsStatusEnum = "EXCEEDED"
	QueryDocsPageByTypeResponseDataRowsStatusEnumFilenamenotmatch QueryDocsPageByTypeResponseDataRowsStatusEnum = "FILENAMENOTMATCH"
	QueryDocsPageByTypeResponseDataRowsStatusEnumEmpty            QueryDocsPageByTypeResponseDataRowsStatusEnum = "EMPTY"
	QueryDocsPageByTypeResponseDataRowsStatusEnumError            QueryDocsPageByTypeResponseDataRowsStatusEnum = "ERROR"
	QueryDocsPageByTypeResponseDataRowsStatusEnumDeleting         QueryDocsPageByTypeResponseDataRowsStatusEnum = "DELETING"
	QueryDocsPageByTypeResponseDataRowsStatusEnumDeleted          QueryDocsPageByTypeResponseDataRowsStatusEnum = "DELETED"
)

type QueryDocsPageByTypeResponseDataRows struct {

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
	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`
	// 文档状态（INDEXING、ERROR、COMPLETED）
	Status *QueryDocsPageByTypeResponseDataRowsStatusEnum `json:"status,omitempty"`
}

func (s QueryDocsPageByTypeResponseDataRows) String() string {
	return utils.Beautify(s)
}

func (s QueryDocsPageByTypeResponseDataRows) GoString() string {
	return s.String()
}

func (s QueryDocsPageByTypeResponseDataRows) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryDocsPageByTypeResponseDataRows) SetCreatedAt(v int64) *QueryDocsPageByTypeResponseDataRows {
	s.CreatedAt = &v
	return s
}

func (s *QueryDocsPageByTypeResponseDataRows) SetDocName(v string) *QueryDocsPageByTypeResponseDataRows {
	s.DocName = &v
	return s
}

func (s *QueryDocsPageByTypeResponseDataRows) SetExtension(v string) *QueryDocsPageByTypeResponseDataRows {
	s.Extension = &v
	return s
}

func (s *QueryDocsPageByTypeResponseDataRows) SetSize(v int64) *QueryDocsPageByTypeResponseDataRows {
	s.Size = &v
	return s
}

func (s *QueryDocsPageByTypeResponseDataRows) SetDocId(v string) *QueryDocsPageByTypeResponseDataRows {
	s.DocId = &v
	return s
}

func (s *QueryDocsPageByTypeResponseDataRows) SetEnabled(v bool) *QueryDocsPageByTypeResponseDataRows {
	s.Enabled = &v
	return s
}

func (s *QueryDocsPageByTypeResponseDataRows) SetStatus(v QueryDocsPageByTypeResponseDataRowsStatusEnum) *QueryDocsPageByTypeResponseDataRows {
	s.Status = &v
	return s
}

type QueryDocsPageByTypeResponseDataRowsBuilder struct {
	s *QueryDocsPageByTypeResponseDataRows
}

func NewQueryDocsPageByTypeResponseDataRowsBuilder() *QueryDocsPageByTypeResponseDataRowsBuilder {
	s := &QueryDocsPageByTypeResponseDataRows{}
	b := &QueryDocsPageByTypeResponseDataRowsBuilder{s: s}
	return b
}

func (b *QueryDocsPageByTypeResponseDataRowsBuilder) CreatedAt(v int64) *QueryDocsPageByTypeResponseDataRowsBuilder {
	b.s.CreatedAt = &v
	return b
}

func (b *QueryDocsPageByTypeResponseDataRowsBuilder) DocName(v string) *QueryDocsPageByTypeResponseDataRowsBuilder {
	b.s.DocName = &v
	return b
}

func (b *QueryDocsPageByTypeResponseDataRowsBuilder) Extension(v string) *QueryDocsPageByTypeResponseDataRowsBuilder {
	b.s.Extension = &v
	return b
}

func (b *QueryDocsPageByTypeResponseDataRowsBuilder) Size(v int64) *QueryDocsPageByTypeResponseDataRowsBuilder {
	b.s.Size = &v
	return b
}

func (b *QueryDocsPageByTypeResponseDataRowsBuilder) DocId(v string) *QueryDocsPageByTypeResponseDataRowsBuilder {
	b.s.DocId = &v
	return b
}

func (b *QueryDocsPageByTypeResponseDataRowsBuilder) Enabled(v bool) *QueryDocsPageByTypeResponseDataRowsBuilder {
	b.s.Enabled = &v
	return b
}

func (b *QueryDocsPageByTypeResponseDataRowsBuilder) Status(v QueryDocsPageByTypeResponseDataRowsStatusEnum) *QueryDocsPageByTypeResponseDataRowsBuilder {
	b.s.Status = &v
	return b
}

func (b *QueryDocsPageByTypeResponseDataRowsBuilder) Build() *QueryDocsPageByTypeResponseDataRows {
	return b.s
}
