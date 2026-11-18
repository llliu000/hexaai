// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryDocsPageByTypeResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *QueryDocsPageByTypeResponseBody `json:"body,omitempty"`
}

func (s QueryDocsPageByTypeResponse) String() string {
	return utils.Beautify(s)
}

func (s QueryDocsPageByTypeResponse) GoString() string {
	return s.String()
}

func (s QueryDocsPageByTypeResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryDocsPageByTypeResponse) SetRequestId(v string) *QueryDocsPageByTypeResponse {
	s.RequestId = &v
	return s
}

func (s *QueryDocsPageByTypeResponse) SetErrorMessage(v string) *QueryDocsPageByTypeResponse {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDocsPageByTypeResponse) SetErrorCode(v string) *QueryDocsPageByTypeResponse {
	s.ErrorCode = &v
	return s
}

func (s *QueryDocsPageByTypeResponse) SetState(v string) *QueryDocsPageByTypeResponse {
	s.State = &v
	return s
}

func (s *QueryDocsPageByTypeResponse) SetBody(v *QueryDocsPageByTypeResponseBody) *QueryDocsPageByTypeResponse {
	s.Body = v
	return s
}

type QueryDocsPageByTypeResponseBuilder struct {
	s *QueryDocsPageByTypeResponse
}

func NewQueryDocsPageByTypeResponseBuilder() *QueryDocsPageByTypeResponseBuilder {
	s := &QueryDocsPageByTypeResponse{}
	b := &QueryDocsPageByTypeResponseBuilder{s: s}
	return b
}

func (b *QueryDocsPageByTypeResponseBuilder) RequestId(v string) *QueryDocsPageByTypeResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *QueryDocsPageByTypeResponseBuilder) ErrorMessage(v string) *QueryDocsPageByTypeResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *QueryDocsPageByTypeResponseBuilder) ErrorCode(v string) *QueryDocsPageByTypeResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *QueryDocsPageByTypeResponseBuilder) State(v string) *QueryDocsPageByTypeResponseBuilder {
	b.s.State = &v
	return b
}

func (b *QueryDocsPageByTypeResponseBuilder) Body(v *QueryDocsPageByTypeResponseBody) *QueryDocsPageByTypeResponseBuilder {
	b.s.Body = v
	return b
}

func (b *QueryDocsPageByTypeResponseBuilder) Build() *QueryDocsPageByTypeResponse {
	return b.s
}
