// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateAssetResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *UpdateAssetResponseBody `json:"body,omitempty"`
}

func (s UpdateAssetResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateAssetResponse) GoString() string {
	return s.String()
}

func (s UpdateAssetResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateAssetResponse) SetRequestId(v string) *UpdateAssetResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateAssetResponse) SetErrorMessage(v string) *UpdateAssetResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAssetResponse) SetErrorCode(v string) *UpdateAssetResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAssetResponse) SetState(v string) *UpdateAssetResponse {
	s.State = &v
	return s
}

func (s *UpdateAssetResponse) SetBody(v *UpdateAssetResponseBody) *UpdateAssetResponse {
	s.Body = v
	return s
}

type UpdateAssetResponseBuilder struct {
	s *UpdateAssetResponse
}

func NewUpdateAssetResponseBuilder() *UpdateAssetResponseBuilder {
	s := &UpdateAssetResponse{}
	b := &UpdateAssetResponseBuilder{s: s}
	return b
}

func (b *UpdateAssetResponseBuilder) RequestId(v string) *UpdateAssetResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateAssetResponseBuilder) ErrorMessage(v string) *UpdateAssetResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateAssetResponseBuilder) ErrorCode(v string) *UpdateAssetResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateAssetResponseBuilder) State(v string) *UpdateAssetResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateAssetResponseBuilder) Body(v *UpdateAssetResponseBody) *UpdateAssetResponseBuilder {
	b.s.Body = v
	return b
}

func (b *UpdateAssetResponseBuilder) Build() *UpdateAssetResponse {
	return b.s
}
