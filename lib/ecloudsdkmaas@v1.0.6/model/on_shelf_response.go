// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OnShelfResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *OnShelfResponseBody `json:"body,omitempty"`
}

func (s OnShelfResponse) String() string {
	return utils.Beautify(s)
}

func (s OnShelfResponse) GoString() string {
	return s.String()
}

func (s OnShelfResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OnShelfResponse) SetRequestId(v string) *OnShelfResponse {
	s.RequestId = &v
	return s
}

func (s *OnShelfResponse) SetErrorMessage(v string) *OnShelfResponse {
	s.ErrorMessage = &v
	return s
}

func (s *OnShelfResponse) SetErrorCode(v string) *OnShelfResponse {
	s.ErrorCode = &v
	return s
}

func (s *OnShelfResponse) SetState(v string) *OnShelfResponse {
	s.State = &v
	return s
}

func (s *OnShelfResponse) SetBody(v *OnShelfResponseBody) *OnShelfResponse {
	s.Body = v
	return s
}

type OnShelfResponseBuilder struct {
	s *OnShelfResponse
}

func NewOnShelfResponseBuilder() *OnShelfResponseBuilder {
	s := &OnShelfResponse{}
	b := &OnShelfResponseBuilder{s: s}
	return b
}

func (b *OnShelfResponseBuilder) RequestId(v string) *OnShelfResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *OnShelfResponseBuilder) ErrorMessage(v string) *OnShelfResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *OnShelfResponseBuilder) ErrorCode(v string) *OnShelfResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *OnShelfResponseBuilder) State(v string) *OnShelfResponseBuilder {
	b.s.State = &v
	return b
}

func (b *OnShelfResponseBuilder) Body(v *OnShelfResponseBody) *OnShelfResponseBuilder {
	b.s.Body = v
	return b
}

func (b *OnShelfResponseBuilder) Build() *OnShelfResponse {
	return b.s
}
