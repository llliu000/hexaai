// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryAiccCreditDeductionRequest struct {
	QueryAiccCreditDeductionBody *QueryAiccCreditDeductionBody `json:"queryAiccCreditDeductionBody,omitempty"`
}

func (s QueryAiccCreditDeductionRequest) String() string {
	return utils.Beautify(s)
}

func (s QueryAiccCreditDeductionRequest) GoString() string {
	return s.String()
}

func (s QueryAiccCreditDeductionRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryAiccCreditDeductionRequest) SetQueryAiccCreditDeductionBody(v *QueryAiccCreditDeductionBody) *QueryAiccCreditDeductionRequest {
	s.QueryAiccCreditDeductionBody = v
	return s
}

type QueryAiccCreditDeductionRequestBuilder struct {
	s *QueryAiccCreditDeductionRequest
}

func NewQueryAiccCreditDeductionRequestBuilder() *QueryAiccCreditDeductionRequestBuilder {
	s := &QueryAiccCreditDeductionRequest{}
	b := &QueryAiccCreditDeductionRequestBuilder{s: s}
	return b
}

func (b *QueryAiccCreditDeductionRequestBuilder) QueryAiccCreditDeductionBody(v *QueryAiccCreditDeductionBody) *QueryAiccCreditDeductionRequestBuilder {
	b.s.QueryAiccCreditDeductionBody = v
	return b
}

func (b *QueryAiccCreditDeductionRequestBuilder) Build() *QueryAiccCreditDeductionRequest {
	return b.s
}
