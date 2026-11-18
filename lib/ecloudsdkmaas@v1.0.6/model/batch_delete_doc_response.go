// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchDeleteDocResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *BatchDeleteDocResponseBody `json:"body,omitempty"`
}

func (s BatchDeleteDocResponse) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteDocResponse) GoString() string {
	return s.String()
}

func (s BatchDeleteDocResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteDocResponse) SetRequestId(v string) *BatchDeleteDocResponse {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteDocResponse) SetErrorMessage(v string) *BatchDeleteDocResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BatchDeleteDocResponse) SetErrorCode(v string) *BatchDeleteDocResponse {
	s.ErrorCode = &v
	return s
}

func (s *BatchDeleteDocResponse) SetState(v string) *BatchDeleteDocResponse {
	s.State = &v
	return s
}

func (s *BatchDeleteDocResponse) SetBody(v *BatchDeleteDocResponseBody) *BatchDeleteDocResponse {
	s.Body = v
	return s
}

type BatchDeleteDocResponseBuilder struct {
	s *BatchDeleteDocResponse
}

func NewBatchDeleteDocResponseBuilder() *BatchDeleteDocResponseBuilder {
	s := &BatchDeleteDocResponse{}
	b := &BatchDeleteDocResponseBuilder{s: s}
	return b
}

func (b *BatchDeleteDocResponseBuilder) RequestId(v string) *BatchDeleteDocResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchDeleteDocResponseBuilder) ErrorMessage(v string) *BatchDeleteDocResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BatchDeleteDocResponseBuilder) ErrorCode(v string) *BatchDeleteDocResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BatchDeleteDocResponseBuilder) State(v string) *BatchDeleteDocResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BatchDeleteDocResponseBuilder) Body(v *BatchDeleteDocResponseBody) *BatchDeleteDocResponseBuilder {
	b.s.Body = v
	return b
}

func (b *BatchDeleteDocResponseBuilder) Build() *BatchDeleteDocResponse {
	return b.s
}
