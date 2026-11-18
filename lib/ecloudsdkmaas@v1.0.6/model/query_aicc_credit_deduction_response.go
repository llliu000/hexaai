// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryAiccCreditDeductionResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *QueryAiccCreditDeductionResponseBody `json:"body,omitempty"`
}

func (s QueryAiccCreditDeductionResponse) String() string {
	return utils.Beautify(s)
}

func (s QueryAiccCreditDeductionResponse) GoString() string {
	return s.String()
}

func (s QueryAiccCreditDeductionResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryAiccCreditDeductionResponse) SetRequestId(v string) *QueryAiccCreditDeductionResponse {
	s.RequestId = &v
	return s
}

func (s *QueryAiccCreditDeductionResponse) SetErrorMessage(v string) *QueryAiccCreditDeductionResponse {
	s.ErrorMessage = &v
	return s
}

func (s *QueryAiccCreditDeductionResponse) SetErrorCode(v string) *QueryAiccCreditDeductionResponse {
	s.ErrorCode = &v
	return s
}

func (s *QueryAiccCreditDeductionResponse) SetState(v string) *QueryAiccCreditDeductionResponse {
	s.State = &v
	return s
}

func (s *QueryAiccCreditDeductionResponse) SetBody(v *QueryAiccCreditDeductionResponseBody) *QueryAiccCreditDeductionResponse {
	s.Body = v
	return s
}

type QueryAiccCreditDeductionResponseBuilder struct {
	s *QueryAiccCreditDeductionResponse
}

func NewQueryAiccCreditDeductionResponseBuilder() *QueryAiccCreditDeductionResponseBuilder {
	s := &QueryAiccCreditDeductionResponse{}
	b := &QueryAiccCreditDeductionResponseBuilder{s: s}
	return b
}

func (b *QueryAiccCreditDeductionResponseBuilder) RequestId(v string) *QueryAiccCreditDeductionResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseBuilder) ErrorMessage(v string) *QueryAiccCreditDeductionResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseBuilder) ErrorCode(v string) *QueryAiccCreditDeductionResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseBuilder) State(v string) *QueryAiccCreditDeductionResponseBuilder {
	b.s.State = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseBuilder) Body(v *QueryAiccCreditDeductionResponseBody) *QueryAiccCreditDeductionResponseBuilder {
	b.s.Body = v
	return b
}

func (b *QueryAiccCreditDeductionResponseBuilder) Build() *QueryAiccCreditDeductionResponse {
	return b.s
}
