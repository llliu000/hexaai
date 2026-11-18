// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateDatasetResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *UpdateDatasetResponseBody `json:"body,omitempty"`
}

func (s UpdateDatasetResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateDatasetResponse) GoString() string {
	return s.String()
}

func (s UpdateDatasetResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateDatasetResponse) SetRequestId(v string) *UpdateDatasetResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasetResponse) SetErrorMessage(v string) *UpdateDatasetResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDatasetResponse) SetErrorCode(v string) *UpdateDatasetResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDatasetResponse) SetState(v string) *UpdateDatasetResponse {
	s.State = &v
	return s
}

func (s *UpdateDatasetResponse) SetBody(v *UpdateDatasetResponseBody) *UpdateDatasetResponse {
	s.Body = v
	return s
}

type UpdateDatasetResponseBuilder struct {
	s *UpdateDatasetResponse
}

func NewUpdateDatasetResponseBuilder() *UpdateDatasetResponseBuilder {
	s := &UpdateDatasetResponse{}
	b := &UpdateDatasetResponseBuilder{s: s}
	return b
}

func (b *UpdateDatasetResponseBuilder) RequestId(v string) *UpdateDatasetResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateDatasetResponseBuilder) ErrorMessage(v string) *UpdateDatasetResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateDatasetResponseBuilder) ErrorCode(v string) *UpdateDatasetResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateDatasetResponseBuilder) State(v string) *UpdateDatasetResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateDatasetResponseBuilder) Body(v *UpdateDatasetResponseBody) *UpdateDatasetResponseBuilder {
	b.s.Body = v
	return b
}

func (b *UpdateDatasetResponseBuilder) Build() *UpdateDatasetResponse {
	return b.s
}
