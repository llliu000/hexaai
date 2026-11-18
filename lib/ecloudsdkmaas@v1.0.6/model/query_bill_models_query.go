// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryBillModelsQuery struct {
	position.Query
	// 模型名称（模糊匹配，可选）
	ModelName *string `json:"modelName,omitempty"`
}

func (s QueryBillModelsQuery) String() string {
	return utils.Beautify(s)
}

func (s QueryBillModelsQuery) GoString() string {
	return s.String()
}

func (s QueryBillModelsQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryBillModelsQuery) SetModelName(v string) *QueryBillModelsQuery {
	s.ModelName = &v
	return s
}

type QueryBillModelsQueryBuilder struct {
	s *QueryBillModelsQuery
}

func NewQueryBillModelsQueryBuilder() *QueryBillModelsQueryBuilder {
	s := &QueryBillModelsQuery{}
	b := &QueryBillModelsQueryBuilder{s: s}
	return b
}

func (b *QueryBillModelsQueryBuilder) ModelName(v string) *QueryBillModelsQueryBuilder {
	b.s.ModelName = &v
	return b
}

func (b *QueryBillModelsQueryBuilder) Build() *QueryBillModelsQuery {
	return b.s
}
