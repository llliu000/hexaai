// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteDatasetResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *DeleteDatasetResponseBody `json:"body,omitempty"`
}

func (s DeleteDatasetResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteDatasetResponse) GoString() string {
	return s.String()
}

func (s DeleteDatasetResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteDatasetResponse) SetRequestId(v string) *DeleteDatasetResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetResponse) SetErrorMessage(v string) *DeleteDatasetResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDatasetResponse) SetErrorCode(v string) *DeleteDatasetResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDatasetResponse) SetState(v string) *DeleteDatasetResponse {
	s.State = &v
	return s
}

func (s *DeleteDatasetResponse) SetBody(v *DeleteDatasetResponseBody) *DeleteDatasetResponse {
	s.Body = v
	return s
}

type DeleteDatasetResponseBuilder struct {
	s *DeleteDatasetResponse
}

func NewDeleteDatasetResponseBuilder() *DeleteDatasetResponseBuilder {
	s := &DeleteDatasetResponse{}
	b := &DeleteDatasetResponseBuilder{s: s}
	return b
}

func (b *DeleteDatasetResponseBuilder) RequestId(v string) *DeleteDatasetResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteDatasetResponseBuilder) ErrorMessage(v string) *DeleteDatasetResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteDatasetResponseBuilder) ErrorCode(v string) *DeleteDatasetResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteDatasetResponseBuilder) State(v string) *DeleteDatasetResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteDatasetResponseBuilder) Body(v *DeleteDatasetResponseBody) *DeleteDatasetResponseBuilder {
	b.s.Body = v
	return b
}

func (b *DeleteDatasetResponseBuilder) Build() *DeleteDatasetResponse {
	return b.s
}
