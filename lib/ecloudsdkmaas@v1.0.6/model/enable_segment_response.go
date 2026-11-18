// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EnableSegmentResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 返回体，true 表示执行成功，false表示执行失败，失败原因可以参考errorMessage或者联系客服人员
	Body *bool `json:"body,omitempty"`
}

func (s EnableSegmentResponse) String() string {
	return utils.Beautify(s)
}

func (s EnableSegmentResponse) GoString() string {
	return s.String()
}

func (s EnableSegmentResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EnableSegmentResponse) SetRequestId(v string) *EnableSegmentResponse {
	s.RequestId = &v
	return s
}

func (s *EnableSegmentResponse) SetErrorMessage(v string) *EnableSegmentResponse {
	s.ErrorMessage = &v
	return s
}

func (s *EnableSegmentResponse) SetErrorCode(v string) *EnableSegmentResponse {
	s.ErrorCode = &v
	return s
}

func (s *EnableSegmentResponse) SetState(v string) *EnableSegmentResponse {
	s.State = &v
	return s
}

func (s *EnableSegmentResponse) SetBody(v bool) *EnableSegmentResponse {
	s.Body = &v
	return s
}

type EnableSegmentResponseBuilder struct {
	s *EnableSegmentResponse
}

func NewEnableSegmentResponseBuilder() *EnableSegmentResponseBuilder {
	s := &EnableSegmentResponse{}
	b := &EnableSegmentResponseBuilder{s: s}
	return b
}

func (b *EnableSegmentResponseBuilder) RequestId(v string) *EnableSegmentResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *EnableSegmentResponseBuilder) ErrorMessage(v string) *EnableSegmentResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *EnableSegmentResponseBuilder) ErrorCode(v string) *EnableSegmentResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *EnableSegmentResponseBuilder) State(v string) *EnableSegmentResponseBuilder {
	b.s.State = &v
	return b
}

func (b *EnableSegmentResponseBuilder) Body(v bool) *EnableSegmentResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *EnableSegmentResponseBuilder) Build() *EnableSegmentResponse {
	return b.s
}
