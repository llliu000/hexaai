// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAssetResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 素材id
	Body *string `json:"body,omitempty"`
}

func (s CreateAssetResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateAssetResponse) GoString() string {
	return s.String()
}

func (s CreateAssetResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAssetResponse) SetRequestId(v string) *CreateAssetResponse {
	s.RequestId = &v
	return s
}

func (s *CreateAssetResponse) SetErrorMessage(v string) *CreateAssetResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAssetResponse) SetErrorCode(v string) *CreateAssetResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateAssetResponse) SetState(v string) *CreateAssetResponse {
	s.State = &v
	return s
}

func (s *CreateAssetResponse) SetBody(v string) *CreateAssetResponse {
	s.Body = &v
	return s
}

type CreateAssetResponseBuilder struct {
	s *CreateAssetResponse
}

func NewCreateAssetResponseBuilder() *CreateAssetResponseBuilder {
	s := &CreateAssetResponse{}
	b := &CreateAssetResponseBuilder{s: s}
	return b
}

func (b *CreateAssetResponseBuilder) RequestId(v string) *CreateAssetResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateAssetResponseBuilder) ErrorMessage(v string) *CreateAssetResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateAssetResponseBuilder) ErrorCode(v string) *CreateAssetResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateAssetResponseBuilder) State(v string) *CreateAssetResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateAssetResponseBuilder) Body(v string) *CreateAssetResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateAssetResponseBuilder) Build() *CreateAssetResponse {
	return b.s
}
