// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryAiccCreditDeductionResponseBody struct {

	// 查询总条数
	Total *int64 `json:"total,omitempty"`
	// 分页查询列表
	Data *[]QueryAiccCreditDeductionResponseData `json:"data,omitempty"`
}

func (s QueryAiccCreditDeductionResponseBody) String() string {
	return utils.Beautify(s)
}

func (s QueryAiccCreditDeductionResponseBody) GoString() string {
	return s.String()
}

func (s QueryAiccCreditDeductionResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryAiccCreditDeductionResponseBody) SetTotal(v int64) *QueryAiccCreditDeductionResponseBody {
	s.Total = &v
	return s
}

func (s *QueryAiccCreditDeductionResponseBody) SetData(v []QueryAiccCreditDeductionResponseData) *QueryAiccCreditDeductionResponseBody {
	s.Data = &v
	return s
}

type QueryAiccCreditDeductionResponseBodyBuilder struct {
	s *QueryAiccCreditDeductionResponseBody
}

func NewQueryAiccCreditDeductionResponseBodyBuilder() *QueryAiccCreditDeductionResponseBodyBuilder {
	s := &QueryAiccCreditDeductionResponseBody{}
	b := &QueryAiccCreditDeductionResponseBodyBuilder{s: s}
	return b
}

func (b *QueryAiccCreditDeductionResponseBodyBuilder) Total(v int64) *QueryAiccCreditDeductionResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseBodyBuilder) Data(v []QueryAiccCreditDeductionResponseData) *QueryAiccCreditDeductionResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseBodyBuilder) Build() *QueryAiccCreditDeductionResponseBody {
	return b.s
}
