// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryBillModelsResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 返回体
	Body *[]QueryBillModelsResponseBody `json:"body,omitempty"`
}

func (s QueryBillModelsResponse) String() string {
	return utils.Beautify(s)
}

func (s QueryBillModelsResponse) GoString() string {
	return s.String()
}

func (s QueryBillModelsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryBillModelsResponse) SetRequestId(v string) *QueryBillModelsResponse {
	s.RequestId = &v
	return s
}

func (s *QueryBillModelsResponse) SetErrorMessage(v string) *QueryBillModelsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *QueryBillModelsResponse) SetErrorCode(v string) *QueryBillModelsResponse {
	s.ErrorCode = &v
	return s
}

func (s *QueryBillModelsResponse) SetState(v string) *QueryBillModelsResponse {
	s.State = &v
	return s
}

func (s *QueryBillModelsResponse) SetBody(v []QueryBillModelsResponseBody) *QueryBillModelsResponse {
	s.Body = &v
	return s
}

type QueryBillModelsResponseBuilder struct {
	s *QueryBillModelsResponse
}

func NewQueryBillModelsResponseBuilder() *QueryBillModelsResponseBuilder {
	s := &QueryBillModelsResponse{}
	b := &QueryBillModelsResponseBuilder{s: s}
	return b
}

func (b *QueryBillModelsResponseBuilder) RequestId(v string) *QueryBillModelsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *QueryBillModelsResponseBuilder) ErrorMessage(v string) *QueryBillModelsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *QueryBillModelsResponseBuilder) ErrorCode(v string) *QueryBillModelsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *QueryBillModelsResponseBuilder) State(v string) *QueryBillModelsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *QueryBillModelsResponseBuilder) Body(v []QueryBillModelsResponseBody) *QueryBillModelsResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *QueryBillModelsResponseBuilder) Build() *QueryBillModelsResponse {
	return b.s
}
