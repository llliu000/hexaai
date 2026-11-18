// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DisableDocumentResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 返回体
	Body *bool `json:"body,omitempty"`
}

func (s DisableDocumentResponse) String() string {
	return utils.Beautify(s)
}

func (s DisableDocumentResponse) GoString() string {
	return s.String()
}

func (s DisableDocumentResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DisableDocumentResponse) SetRequestId(v string) *DisableDocumentResponse {
	s.RequestId = &v
	return s
}

func (s *DisableDocumentResponse) SetErrorMessage(v string) *DisableDocumentResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DisableDocumentResponse) SetErrorCode(v string) *DisableDocumentResponse {
	s.ErrorCode = &v
	return s
}

func (s *DisableDocumentResponse) SetState(v string) *DisableDocumentResponse {
	s.State = &v
	return s
}

func (s *DisableDocumentResponse) SetBody(v bool) *DisableDocumentResponse {
	s.Body = &v
	return s
}

type DisableDocumentResponseBuilder struct {
	s *DisableDocumentResponse
}

func NewDisableDocumentResponseBuilder() *DisableDocumentResponseBuilder {
	s := &DisableDocumentResponse{}
	b := &DisableDocumentResponseBuilder{s: s}
	return b
}

func (b *DisableDocumentResponseBuilder) RequestId(v string) *DisableDocumentResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DisableDocumentResponseBuilder) ErrorMessage(v string) *DisableDocumentResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DisableDocumentResponseBuilder) ErrorCode(v string) *DisableDocumentResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DisableDocumentResponseBuilder) State(v string) *DisableDocumentResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DisableDocumentResponseBuilder) Body(v bool) *DisableDocumentResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DisableDocumentResponseBuilder) Build() *DisableDocumentResponse {
	return b.s
}
