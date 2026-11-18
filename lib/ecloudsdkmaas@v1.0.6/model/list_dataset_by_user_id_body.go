// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListDatasetByUserIdBody struct {
	position.Body
	// 页码（从 1 开始）
	PageNo *int32 `json:"pageNo,omitempty"`
	// 每页行数
	PageSize *int32 `json:"pageSize,omitempty"`
}

func (s ListDatasetByUserIdBody) String() string {
	return utils.Beautify(s)
}

func (s ListDatasetByUserIdBody) GoString() string {
	return s.String()
}

func (s ListDatasetByUserIdBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListDatasetByUserIdBody) SetPageNo(v int32) *ListDatasetByUserIdBody {
	s.PageNo = &v
	return s
}

func (s *ListDatasetByUserIdBody) SetPageSize(v int32) *ListDatasetByUserIdBody {
	s.PageSize = &v
	return s
}

type ListDatasetByUserIdBodyBuilder struct {
	s *ListDatasetByUserIdBody
}

func NewListDatasetByUserIdBodyBuilder() *ListDatasetByUserIdBodyBuilder {
	s := &ListDatasetByUserIdBody{}
	b := &ListDatasetByUserIdBodyBuilder{s: s}
	return b
}

func (b *ListDatasetByUserIdBodyBuilder) PageNo(v int32) *ListDatasetByUserIdBodyBuilder {
	b.s.PageNo = &v
	return b
}

func (b *ListDatasetByUserIdBodyBuilder) PageSize(v int32) *ListDatasetByUserIdBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListDatasetByUserIdBodyBuilder) Build() *ListDatasetByUserIdBody {
	return b.s
}
