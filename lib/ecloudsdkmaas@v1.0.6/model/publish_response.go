// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PublishResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *PublishResponseBody `json:"body,omitempty"`
}

func (s PublishResponse) String() string {
	return utils.Beautify(s)
}

func (s PublishResponse) GoString() string {
	return s.String()
}

func (s PublishResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PublishResponse) SetRequestId(v string) *PublishResponse {
	s.RequestId = &v
	return s
}

func (s *PublishResponse) SetErrorMessage(v string) *PublishResponse {
	s.ErrorMessage = &v
	return s
}

func (s *PublishResponse) SetErrorCode(v string) *PublishResponse {
	s.ErrorCode = &v
	return s
}

func (s *PublishResponse) SetState(v string) *PublishResponse {
	s.State = &v
	return s
}

func (s *PublishResponse) SetBody(v *PublishResponseBody) *PublishResponse {
	s.Body = v
	return s
}

type PublishResponseBuilder struct {
	s *PublishResponse
}

func NewPublishResponseBuilder() *PublishResponseBuilder {
	s := &PublishResponse{}
	b := &PublishResponseBuilder{s: s}
	return b
}

func (b *PublishResponseBuilder) RequestId(v string) *PublishResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *PublishResponseBuilder) ErrorMessage(v string) *PublishResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *PublishResponseBuilder) ErrorCode(v string) *PublishResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *PublishResponseBuilder) State(v string) *PublishResponseBuilder {
	b.s.State = &v
	return b
}

func (b *PublishResponseBuilder) Body(v *PublishResponseBody) *PublishResponseBuilder {
	b.s.Body = v
	return b
}

func (b *PublishResponseBuilder) Build() *PublishResponse {
	return b.s
}
