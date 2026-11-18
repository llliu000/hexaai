// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAssetResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *GetAssetResponseBody `json:"body,omitempty"`
}

func (s GetAssetResponse) String() string {
	return utils.Beautify(s)
}

func (s GetAssetResponse) GoString() string {
	return s.String()
}

func (s GetAssetResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAssetResponse) SetRequestId(v string) *GetAssetResponse {
	s.RequestId = &v
	return s
}

func (s *GetAssetResponse) SetErrorMessage(v string) *GetAssetResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetAssetResponse) SetErrorCode(v string) *GetAssetResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetAssetResponse) SetState(v string) *GetAssetResponse {
	s.State = &v
	return s
}

func (s *GetAssetResponse) SetBody(v *GetAssetResponseBody) *GetAssetResponse {
	s.Body = v
	return s
}

type GetAssetResponseBuilder struct {
	s *GetAssetResponse
}

func NewGetAssetResponseBuilder() *GetAssetResponseBuilder {
	s := &GetAssetResponse{}
	b := &GetAssetResponseBuilder{s: s}
	return b
}

func (b *GetAssetResponseBuilder) RequestId(v string) *GetAssetResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetAssetResponseBuilder) ErrorMessage(v string) *GetAssetResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetAssetResponseBuilder) ErrorCode(v string) *GetAssetResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetAssetResponseBuilder) State(v string) *GetAssetResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetAssetResponseBuilder) Body(v *GetAssetResponseBody) *GetAssetResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetAssetResponseBuilder) Build() *GetAssetResponse {
	return b.s
}
