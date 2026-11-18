// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteDocResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *DeleteDocResponseBody `json:"body,omitempty"`
}

func (s DeleteDocResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteDocResponse) GoString() string {
	return s.String()
}

func (s DeleteDocResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteDocResponse) SetRequestId(v string) *DeleteDocResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteDocResponse) SetErrorMessage(v string) *DeleteDocResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDocResponse) SetErrorCode(v string) *DeleteDocResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDocResponse) SetState(v string) *DeleteDocResponse {
	s.State = &v
	return s
}

func (s *DeleteDocResponse) SetBody(v *DeleteDocResponseBody) *DeleteDocResponse {
	s.Body = v
	return s
}

type DeleteDocResponseBuilder struct {
	s *DeleteDocResponse
}

func NewDeleteDocResponseBuilder() *DeleteDocResponseBuilder {
	s := &DeleteDocResponse{}
	b := &DeleteDocResponseBuilder{s: s}
	return b
}

func (b *DeleteDocResponseBuilder) RequestId(v string) *DeleteDocResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteDocResponseBuilder) ErrorMessage(v string) *DeleteDocResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteDocResponseBuilder) ErrorCode(v string) *DeleteDocResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteDocResponseBuilder) State(v string) *DeleteDocResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteDocResponseBuilder) Body(v *DeleteDocResponseBody) *DeleteDocResponseBuilder {
	b.s.Body = v
	return b
}

func (b *DeleteDocResponseBuilder) Build() *DeleteDocResponse {
	return b.s
}
