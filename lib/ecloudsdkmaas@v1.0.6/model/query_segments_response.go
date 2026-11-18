// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QuerySegmentsResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *QuerySegmentsResponseBody `json:"body,omitempty"`
}

func (s QuerySegmentsResponse) String() string {
	return utils.Beautify(s)
}

func (s QuerySegmentsResponse) GoString() string {
	return s.String()
}

func (s QuerySegmentsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QuerySegmentsResponse) SetRequestId(v string) *QuerySegmentsResponse {
	s.RequestId = &v
	return s
}

func (s *QuerySegmentsResponse) SetErrorMessage(v string) *QuerySegmentsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *QuerySegmentsResponse) SetErrorCode(v string) *QuerySegmentsResponse {
	s.ErrorCode = &v
	return s
}

func (s *QuerySegmentsResponse) SetState(v string) *QuerySegmentsResponse {
	s.State = &v
	return s
}

func (s *QuerySegmentsResponse) SetBody(v *QuerySegmentsResponseBody) *QuerySegmentsResponse {
	s.Body = v
	return s
}

type QuerySegmentsResponseBuilder struct {
	s *QuerySegmentsResponse
}

func NewQuerySegmentsResponseBuilder() *QuerySegmentsResponseBuilder {
	s := &QuerySegmentsResponse{}
	b := &QuerySegmentsResponseBuilder{s: s}
	return b
}

func (b *QuerySegmentsResponseBuilder) RequestId(v string) *QuerySegmentsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *QuerySegmentsResponseBuilder) ErrorMessage(v string) *QuerySegmentsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *QuerySegmentsResponseBuilder) ErrorCode(v string) *QuerySegmentsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *QuerySegmentsResponseBuilder) State(v string) *QuerySegmentsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *QuerySegmentsResponseBuilder) Body(v *QuerySegmentsResponseBody) *QuerySegmentsResponseBuilder {
	b.s.Body = v
	return b
}

func (b *QuerySegmentsResponseBuilder) Build() *QuerySegmentsResponse {
	return b.s
}
