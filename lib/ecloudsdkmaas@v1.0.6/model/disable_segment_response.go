// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DisableSegmentResponse struct {

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

func (s DisableSegmentResponse) String() string {
	return utils.Beautify(s)
}

func (s DisableSegmentResponse) GoString() string {
	return s.String()
}

func (s DisableSegmentResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DisableSegmentResponse) SetRequestId(v string) *DisableSegmentResponse {
	s.RequestId = &v
	return s
}

func (s *DisableSegmentResponse) SetErrorMessage(v string) *DisableSegmentResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DisableSegmentResponse) SetErrorCode(v string) *DisableSegmentResponse {
	s.ErrorCode = &v
	return s
}

func (s *DisableSegmentResponse) SetState(v string) *DisableSegmentResponse {
	s.State = &v
	return s
}

func (s *DisableSegmentResponse) SetBody(v bool) *DisableSegmentResponse {
	s.Body = &v
	return s
}

type DisableSegmentResponseBuilder struct {
	s *DisableSegmentResponse
}

func NewDisableSegmentResponseBuilder() *DisableSegmentResponseBuilder {
	s := &DisableSegmentResponse{}
	b := &DisableSegmentResponseBuilder{s: s}
	return b
}

func (b *DisableSegmentResponseBuilder) RequestId(v string) *DisableSegmentResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DisableSegmentResponseBuilder) ErrorMessage(v string) *DisableSegmentResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DisableSegmentResponseBuilder) ErrorCode(v string) *DisableSegmentResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DisableSegmentResponseBuilder) State(v string) *DisableSegmentResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DisableSegmentResponseBuilder) Body(v bool) *DisableSegmentResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DisableSegmentResponseBuilder) Build() *DisableSegmentResponse {
	return b.s
}
