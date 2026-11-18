// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EnableDocumentResponse struct {

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

func (s EnableDocumentResponse) String() string {
	return utils.Beautify(s)
}

func (s EnableDocumentResponse) GoString() string {
	return s.String()
}

func (s EnableDocumentResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EnableDocumentResponse) SetRequestId(v string) *EnableDocumentResponse {
	s.RequestId = &v
	return s
}

func (s *EnableDocumentResponse) SetErrorMessage(v string) *EnableDocumentResponse {
	s.ErrorMessage = &v
	return s
}

func (s *EnableDocumentResponse) SetErrorCode(v string) *EnableDocumentResponse {
	s.ErrorCode = &v
	return s
}

func (s *EnableDocumentResponse) SetState(v string) *EnableDocumentResponse {
	s.State = &v
	return s
}

func (s *EnableDocumentResponse) SetBody(v bool) *EnableDocumentResponse {
	s.Body = &v
	return s
}

type EnableDocumentResponseBuilder struct {
	s *EnableDocumentResponse
}

func NewEnableDocumentResponseBuilder() *EnableDocumentResponseBuilder {
	s := &EnableDocumentResponse{}
	b := &EnableDocumentResponseBuilder{s: s}
	return b
}

func (b *EnableDocumentResponseBuilder) RequestId(v string) *EnableDocumentResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *EnableDocumentResponseBuilder) ErrorMessage(v string) *EnableDocumentResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *EnableDocumentResponseBuilder) ErrorCode(v string) *EnableDocumentResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *EnableDocumentResponseBuilder) State(v string) *EnableDocumentResponseBuilder {
	b.s.State = &v
	return b
}

func (b *EnableDocumentResponseBuilder) Body(v bool) *EnableDocumentResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *EnableDocumentResponseBuilder) Build() *EnableDocumentResponse {
	return b.s
}
