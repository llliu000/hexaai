// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteAssetResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 返回体
	Body *bool `json:"body,omitempty"`
}

func (s DeleteAssetResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteAssetResponse) GoString() string {
	return s.String()
}

func (s DeleteAssetResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteAssetResponse) SetRequestId(v string) *DeleteAssetResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteAssetResponse) SetErrorMessage(v string) *DeleteAssetResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteAssetResponse) SetErrorCode(v string) *DeleteAssetResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAssetResponse) SetState(v string) *DeleteAssetResponse {
	s.State = &v
	return s
}

func (s *DeleteAssetResponse) SetBody(v bool) *DeleteAssetResponse {
	s.Body = &v
	return s
}

type DeleteAssetResponseBuilder struct {
	s *DeleteAssetResponse
}

func NewDeleteAssetResponseBuilder() *DeleteAssetResponseBuilder {
	s := &DeleteAssetResponse{}
	b := &DeleteAssetResponseBuilder{s: s}
	return b
}

func (b *DeleteAssetResponseBuilder) RequestId(v string) *DeleteAssetResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteAssetResponseBuilder) ErrorMessage(v string) *DeleteAssetResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteAssetResponseBuilder) ErrorCode(v string) *DeleteAssetResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteAssetResponseBuilder) State(v string) *DeleteAssetResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteAssetResponseBuilder) Body(v bool) *DeleteAssetResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteAssetResponseBuilder) Build() *DeleteAssetResponse {
	return b.s
}
