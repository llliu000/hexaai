// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryDocsPageByTypePath struct {
	position.Path

	DatasetId *string `json:"datasetId,omitempty"`
}

func (s QueryDocsPageByTypePath) String() string {
	return utils.Beautify(s)
}

func (s QueryDocsPageByTypePath) GoString() string {
	return s.String()
}

func (s QueryDocsPageByTypePath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryDocsPageByTypePath) SetDatasetId(v string) *QueryDocsPageByTypePath {
	s.DatasetId = &v
	return s
}

type QueryDocsPageByTypePathBuilder struct {
	s *QueryDocsPageByTypePath
}

func NewQueryDocsPageByTypePathBuilder() *QueryDocsPageByTypePathBuilder {
	s := &QueryDocsPageByTypePath{}
	b := &QueryDocsPageByTypePathBuilder{s: s}
	return b
}

func (b *QueryDocsPageByTypePathBuilder) DatasetId(v string) *QueryDocsPageByTypePathBuilder {
	b.s.DatasetId = &v
	return b
}

func (b *QueryDocsPageByTypePathBuilder) Build() *QueryDocsPageByTypePath {
	return b.s
}
