// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryBillModelsRequest struct {
	QueryBillModelsQuery *QueryBillModelsQuery `json:"queryBillModelsQuery,omitempty"`
}

func (s QueryBillModelsRequest) String() string {
	return utils.Beautify(s)
}

func (s QueryBillModelsRequest) GoString() string {
	return s.String()
}

func (s QueryBillModelsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryBillModelsRequest) SetQueryBillModelsQuery(v *QueryBillModelsQuery) *QueryBillModelsRequest {
	s.QueryBillModelsQuery = v
	return s
}

type QueryBillModelsRequestBuilder struct {
	s *QueryBillModelsRequest
}

func NewQueryBillModelsRequestBuilder() *QueryBillModelsRequestBuilder {
	s := &QueryBillModelsRequest{}
	b := &QueryBillModelsRequestBuilder{s: s}
	return b
}

func (b *QueryBillModelsRequestBuilder) QueryBillModelsQuery(v *QueryBillModelsQuery) *QueryBillModelsRequestBuilder {
	b.s.QueryBillModelsQuery = v
	return b
}

func (b *QueryBillModelsRequestBuilder) Build() *QueryBillModelsRequest {
	return b.s
}
