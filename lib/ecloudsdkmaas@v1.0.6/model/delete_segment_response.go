// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteSegmentResponse struct {

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

func (s DeleteSegmentResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteSegmentResponse) GoString() string {
	return s.String()
}

func (s DeleteSegmentResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteSegmentResponse) SetRequestId(v string) *DeleteSegmentResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteSegmentResponse) SetErrorMessage(v string) *DeleteSegmentResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteSegmentResponse) SetErrorCode(v string) *DeleteSegmentResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSegmentResponse) SetState(v string) *DeleteSegmentResponse {
	s.State = &v
	return s
}

func (s *DeleteSegmentResponse) SetBody(v bool) *DeleteSegmentResponse {
	s.Body = &v
	return s
}

type DeleteSegmentResponseBuilder struct {
	s *DeleteSegmentResponse
}

func NewDeleteSegmentResponseBuilder() *DeleteSegmentResponseBuilder {
	s := &DeleteSegmentResponse{}
	b := &DeleteSegmentResponseBuilder{s: s}
	return b
}

func (b *DeleteSegmentResponseBuilder) RequestId(v string) *DeleteSegmentResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteSegmentResponseBuilder) ErrorMessage(v string) *DeleteSegmentResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteSegmentResponseBuilder) ErrorCode(v string) *DeleteSegmentResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteSegmentResponseBuilder) State(v string) *DeleteSegmentResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteSegmentResponseBuilder) Body(v bool) *DeleteSegmentResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteSegmentResponseBuilder) Build() *DeleteSegmentResponse {
	return b.s
}
