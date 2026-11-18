// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryDocsPageByTypeBody struct {
	position.Body
	// 文档解析状态
	IndexingStatusList []string `json:"indexingStatusList,omitempty"`
	// 文件名
	DocName *string `json:"docName,omitempty"`
	// 文件扩展名
	FileTypeList []string `json:"fileTypeList,omitempty"`
	// 页码（从 1 开始）
	PageNo *int32 `json:"pageNo,omitempty"`
	// 每页行数
	PageSize *int32 `json:"pageSize,omitempty"`
}

func (s QueryDocsPageByTypeBody) String() string {
	return utils.Beautify(s)
}

func (s QueryDocsPageByTypeBody) GoString() string {
	return s.String()
}

func (s QueryDocsPageByTypeBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryDocsPageByTypeBody) SetIndexingStatusList(v []string) *QueryDocsPageByTypeBody {
	s.IndexingStatusList = v
	return s
}

func (s *QueryDocsPageByTypeBody) SetDocName(v string) *QueryDocsPageByTypeBody {
	s.DocName = &v
	return s
}

func (s *QueryDocsPageByTypeBody) SetFileTypeList(v []string) *QueryDocsPageByTypeBody {
	s.FileTypeList = v
	return s
}

func (s *QueryDocsPageByTypeBody) SetPageNo(v int32) *QueryDocsPageByTypeBody {
	s.PageNo = &v
	return s
}

func (s *QueryDocsPageByTypeBody) SetPageSize(v int32) *QueryDocsPageByTypeBody {
	s.PageSize = &v
	return s
}

type QueryDocsPageByTypeBodyBuilder struct {
	s *QueryDocsPageByTypeBody
}

func NewQueryDocsPageByTypeBodyBuilder() *QueryDocsPageByTypeBodyBuilder {
	s := &QueryDocsPageByTypeBody{}
	b := &QueryDocsPageByTypeBodyBuilder{s: s}
	return b
}

func (b *QueryDocsPageByTypeBodyBuilder) IndexingStatusList(v []string) *QueryDocsPageByTypeBodyBuilder {
	b.s.IndexingStatusList = v
	return b
}

func (b *QueryDocsPageByTypeBodyBuilder) DocName(v string) *QueryDocsPageByTypeBodyBuilder {
	b.s.DocName = &v
	return b
}

func (b *QueryDocsPageByTypeBodyBuilder) FileTypeList(v []string) *QueryDocsPageByTypeBodyBuilder {
	b.s.FileTypeList = v
	return b
}

func (b *QueryDocsPageByTypeBodyBuilder) PageNo(v int32) *QueryDocsPageByTypeBodyBuilder {
	b.s.PageNo = &v
	return b
}

func (b *QueryDocsPageByTypeBodyBuilder) PageSize(v int32) *QueryDocsPageByTypeBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *QueryDocsPageByTypeBodyBuilder) Build() *QueryDocsPageByTypeBody {
	return b.s
}
