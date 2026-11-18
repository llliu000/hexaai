// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListDocsResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 返回体
	Body *[]ListDocsResponseBody `json:"body,omitempty"`
}

func (s ListDocsResponse) String() string {
	return utils.Beautify(s)
}

func (s ListDocsResponse) GoString() string {
	return s.String()
}

func (s ListDocsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListDocsResponse) SetRequestId(v string) *ListDocsResponse {
	s.RequestId = &v
	return s
}

func (s *ListDocsResponse) SetErrorMessage(v string) *ListDocsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListDocsResponse) SetErrorCode(v string) *ListDocsResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListDocsResponse) SetState(v string) *ListDocsResponse {
	s.State = &v
	return s
}

func (s *ListDocsResponse) SetBody(v []ListDocsResponseBody) *ListDocsResponse {
	s.Body = &v
	return s
}

type ListDocsResponseBuilder struct {
	s *ListDocsResponse
}

func NewListDocsResponseBuilder() *ListDocsResponseBuilder {
	s := &ListDocsResponse{}
	b := &ListDocsResponseBuilder{s: s}
	return b
}

func (b *ListDocsResponseBuilder) RequestId(v string) *ListDocsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListDocsResponseBuilder) ErrorMessage(v string) *ListDocsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListDocsResponseBuilder) ErrorCode(v string) *ListDocsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListDocsResponseBuilder) State(v string) *ListDocsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListDocsResponseBuilder) Body(v []ListDocsResponseBody) *ListDocsResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *ListDocsResponseBuilder) Build() *ListDocsResponse {
	return b.s
}
