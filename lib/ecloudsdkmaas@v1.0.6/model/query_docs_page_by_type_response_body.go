// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryDocsPageByTypeResponseBody struct {

	// 当前页数据行
	DataRows *[]QueryDocsPageByTypeResponseDataRows `json:"dataRows,omitempty"`
	// 总条数
	TotalSize *int64 `json:"totalSize,omitempty"`
}

func (s QueryDocsPageByTypeResponseBody) String() string {
	return utils.Beautify(s)
}

func (s QueryDocsPageByTypeResponseBody) GoString() string {
	return s.String()
}

func (s QueryDocsPageByTypeResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryDocsPageByTypeResponseBody) SetDataRows(v []QueryDocsPageByTypeResponseDataRows) *QueryDocsPageByTypeResponseBody {
	s.DataRows = &v
	return s
}

func (s *QueryDocsPageByTypeResponseBody) SetTotalSize(v int64) *QueryDocsPageByTypeResponseBody {
	s.TotalSize = &v
	return s
}

type QueryDocsPageByTypeResponseBodyBuilder struct {
	s *QueryDocsPageByTypeResponseBody
}

func NewQueryDocsPageByTypeResponseBodyBuilder() *QueryDocsPageByTypeResponseBodyBuilder {
	s := &QueryDocsPageByTypeResponseBody{}
	b := &QueryDocsPageByTypeResponseBodyBuilder{s: s}
	return b
}

func (b *QueryDocsPageByTypeResponseBodyBuilder) DataRows(v []QueryDocsPageByTypeResponseDataRows) *QueryDocsPageByTypeResponseBodyBuilder {
	b.s.DataRows = &v
	return b
}

func (b *QueryDocsPageByTypeResponseBodyBuilder) TotalSize(v int64) *QueryDocsPageByTypeResponseBodyBuilder {
	b.s.TotalSize = &v
	return b
}

func (b *QueryDocsPageByTypeResponseBodyBuilder) Build() *QueryDocsPageByTypeResponseBody {
	return b.s
}
