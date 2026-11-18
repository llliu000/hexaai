// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListDatasetByUserIdResponseBody struct {

	// 当前页数据行
	DataRows *[]ListDatasetByUserIdResponseDataRows `json:"dataRows,omitempty"`
	// 总条数
	TotalSize *int64 `json:"totalSize,omitempty"`
}

func (s ListDatasetByUserIdResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListDatasetByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s ListDatasetByUserIdResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListDatasetByUserIdResponseBody) SetDataRows(v []ListDatasetByUserIdResponseDataRows) *ListDatasetByUserIdResponseBody {
	s.DataRows = &v
	return s
}

func (s *ListDatasetByUserIdResponseBody) SetTotalSize(v int64) *ListDatasetByUserIdResponseBody {
	s.TotalSize = &v
	return s
}

type ListDatasetByUserIdResponseBodyBuilder struct {
	s *ListDatasetByUserIdResponseBody
}

func NewListDatasetByUserIdResponseBodyBuilder() *ListDatasetByUserIdResponseBodyBuilder {
	s := &ListDatasetByUserIdResponseBody{}
	b := &ListDatasetByUserIdResponseBodyBuilder{s: s}
	return b
}

func (b *ListDatasetByUserIdResponseBodyBuilder) DataRows(v []ListDatasetByUserIdResponseDataRows) *ListDatasetByUserIdResponseBodyBuilder {
	b.s.DataRows = &v
	return b
}

func (b *ListDatasetByUserIdResponseBodyBuilder) TotalSize(v int64) *ListDatasetByUserIdResponseBodyBuilder {
	b.s.TotalSize = &v
	return b
}

func (b *ListDatasetByUserIdResponseBodyBuilder) Build() *ListDatasetByUserIdResponseBody {
	return b.s
}
