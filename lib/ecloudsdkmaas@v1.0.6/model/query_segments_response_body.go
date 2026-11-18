// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QuerySegmentsResponseBody struct {

	// 当前页数据行
	DataRows *[]QuerySegmentsResponseDataRows `json:"dataRows,omitempty"`
	// 总条数
	TotalSize *int64 `json:"totalSize,omitempty"`
}

func (s QuerySegmentsResponseBody) String() string {
	return utils.Beautify(s)
}

func (s QuerySegmentsResponseBody) GoString() string {
	return s.String()
}

func (s QuerySegmentsResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QuerySegmentsResponseBody) SetDataRows(v []QuerySegmentsResponseDataRows) *QuerySegmentsResponseBody {
	s.DataRows = &v
	return s
}

func (s *QuerySegmentsResponseBody) SetTotalSize(v int64) *QuerySegmentsResponseBody {
	s.TotalSize = &v
	return s
}

type QuerySegmentsResponseBodyBuilder struct {
	s *QuerySegmentsResponseBody
}

func NewQuerySegmentsResponseBodyBuilder() *QuerySegmentsResponseBodyBuilder {
	s := &QuerySegmentsResponseBody{}
	b := &QuerySegmentsResponseBodyBuilder{s: s}
	return b
}

func (b *QuerySegmentsResponseBodyBuilder) DataRows(v []QuerySegmentsResponseDataRows) *QuerySegmentsResponseBodyBuilder {
	b.s.DataRows = &v
	return b
}

func (b *QuerySegmentsResponseBodyBuilder) TotalSize(v int64) *QuerySegmentsResponseBodyBuilder {
	b.s.TotalSize = &v
	return b
}

func (b *QuerySegmentsResponseBodyBuilder) Build() *QuerySegmentsResponseBody {
	return b.s
}
