// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetDocumentResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *GetDocumentResponseBody `json:"body,omitempty"`
}

func (s GetDocumentResponse) String() string {
	return utils.Beautify(s)
}

func (s GetDocumentResponse) GoString() string {
	return s.String()
}

func (s GetDocumentResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetDocumentResponse) SetRequestId(v string) *GetDocumentResponse {
	s.RequestId = &v
	return s
}

func (s *GetDocumentResponse) SetErrorMessage(v string) *GetDocumentResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetDocumentResponse) SetErrorCode(v string) *GetDocumentResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetDocumentResponse) SetState(v string) *GetDocumentResponse {
	s.State = &v
	return s
}

func (s *GetDocumentResponse) SetBody(v *GetDocumentResponseBody) *GetDocumentResponse {
	s.Body = v
	return s
}

type GetDocumentResponseBuilder struct {
	s *GetDocumentResponse
}

func NewGetDocumentResponseBuilder() *GetDocumentResponseBuilder {
	s := &GetDocumentResponse{}
	b := &GetDocumentResponseBuilder{s: s}
	return b
}

func (b *GetDocumentResponseBuilder) RequestId(v string) *GetDocumentResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetDocumentResponseBuilder) ErrorMessage(v string) *GetDocumentResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetDocumentResponseBuilder) ErrorCode(v string) *GetDocumentResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetDocumentResponseBuilder) State(v string) *GetDocumentResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetDocumentResponseBuilder) Body(v *GetDocumentResponseBody) *GetDocumentResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetDocumentResponseBuilder) Build() *GetDocumentResponse {
	return b.s
}
