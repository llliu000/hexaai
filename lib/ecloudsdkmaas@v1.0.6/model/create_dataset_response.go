// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateDatasetResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *CreateDatasetResponseBody `json:"body,omitempty"`
}

func (s CreateDatasetResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateDatasetResponse) GoString() string {
	return s.String()
}

func (s CreateDatasetResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateDatasetResponse) SetRequestId(v string) *CreateDatasetResponse {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetResponse) SetErrorMessage(v string) *CreateDatasetResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDatasetResponse) SetErrorCode(v string) *CreateDatasetResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateDatasetResponse) SetState(v string) *CreateDatasetResponse {
	s.State = &v
	return s
}

func (s *CreateDatasetResponse) SetBody(v *CreateDatasetResponseBody) *CreateDatasetResponse {
	s.Body = v
	return s
}

type CreateDatasetResponseBuilder struct {
	s *CreateDatasetResponse
}

func NewCreateDatasetResponseBuilder() *CreateDatasetResponseBuilder {
	s := &CreateDatasetResponse{}
	b := &CreateDatasetResponseBuilder{s: s}
	return b
}

func (b *CreateDatasetResponseBuilder) RequestId(v string) *CreateDatasetResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateDatasetResponseBuilder) ErrorMessage(v string) *CreateDatasetResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateDatasetResponseBuilder) ErrorCode(v string) *CreateDatasetResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateDatasetResponseBuilder) State(v string) *CreateDatasetResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateDatasetResponseBuilder) Body(v *CreateDatasetResponseBody) *CreateDatasetResponseBuilder {
	b.s.Body = v
	return b
}

func (b *CreateDatasetResponseBuilder) Build() *CreateDatasetResponse {
	return b.s
}
