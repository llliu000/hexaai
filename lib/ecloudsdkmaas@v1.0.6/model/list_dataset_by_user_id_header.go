// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListDatasetByUserIdHeader struct {
	position.Header
	// 项目空间id
	ProjectId *string `json:"project_id,omitempty"`
}

func (s ListDatasetByUserIdHeader) String() string {
	return utils.Beautify(s)
}

func (s ListDatasetByUserIdHeader) GoString() string {
	return s.String()
}

func (s ListDatasetByUserIdHeader) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListDatasetByUserIdHeader) SetProjectId(v string) *ListDatasetByUserIdHeader {
	s.ProjectId = &v
	return s
}

type ListDatasetByUserIdHeaderBuilder struct {
	s *ListDatasetByUserIdHeader
}

func NewListDatasetByUserIdHeaderBuilder() *ListDatasetByUserIdHeaderBuilder {
	s := &ListDatasetByUserIdHeader{}
	b := &ListDatasetByUserIdHeaderBuilder{s: s}
	return b
}

func (b *ListDatasetByUserIdHeaderBuilder) ProjectId(v string) *ListDatasetByUserIdHeaderBuilder {
	b.s.ProjectId = &v
	return b
}

func (b *ListDatasetByUserIdHeaderBuilder) Build() *ListDatasetByUserIdHeader {
	return b.s
}
