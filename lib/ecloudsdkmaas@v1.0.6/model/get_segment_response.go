// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetSegmentResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *GetSegmentResponseBody `json:"body,omitempty"`
}

func (s GetSegmentResponse) String() string {
	return utils.Beautify(s)
}

func (s GetSegmentResponse) GoString() string {
	return s.String()
}

func (s GetSegmentResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSegmentResponse) SetRequestId(v string) *GetSegmentResponse {
	s.RequestId = &v
	return s
}

func (s *GetSegmentResponse) SetErrorMessage(v string) *GetSegmentResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetSegmentResponse) SetErrorCode(v string) *GetSegmentResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetSegmentResponse) SetState(v string) *GetSegmentResponse {
	s.State = &v
	return s
}

func (s *GetSegmentResponse) SetBody(v *GetSegmentResponseBody) *GetSegmentResponse {
	s.Body = v
	return s
}

type GetSegmentResponseBuilder struct {
	s *GetSegmentResponse
}

func NewGetSegmentResponseBuilder() *GetSegmentResponseBuilder {
	s := &GetSegmentResponse{}
	b := &GetSegmentResponseBuilder{s: s}
	return b
}

func (b *GetSegmentResponseBuilder) RequestId(v string) *GetSegmentResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetSegmentResponseBuilder) ErrorMessage(v string) *GetSegmentResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetSegmentResponseBuilder) ErrorCode(v string) *GetSegmentResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetSegmentResponseBuilder) State(v string) *GetSegmentResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetSegmentResponseBuilder) Body(v *GetSegmentResponseBody) *GetSegmentResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetSegmentResponseBuilder) Build() *GetSegmentResponse {
	return b.s
}
