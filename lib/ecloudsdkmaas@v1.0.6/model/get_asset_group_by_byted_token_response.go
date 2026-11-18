// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAssetGroupByBytedTokenResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 返回体
	Body *string `json:"body,omitempty"`
}

func (s GetAssetGroupByBytedTokenResponse) String() string {
	return utils.Beautify(s)
}

func (s GetAssetGroupByBytedTokenResponse) GoString() string {
	return s.String()
}

func (s GetAssetGroupByBytedTokenResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAssetGroupByBytedTokenResponse) SetRequestId(v string) *GetAssetGroupByBytedTokenResponse {
	s.RequestId = &v
	return s
}

func (s *GetAssetGroupByBytedTokenResponse) SetErrorMessage(v string) *GetAssetGroupByBytedTokenResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetAssetGroupByBytedTokenResponse) SetErrorCode(v string) *GetAssetGroupByBytedTokenResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetAssetGroupByBytedTokenResponse) SetState(v string) *GetAssetGroupByBytedTokenResponse {
	s.State = &v
	return s
}

func (s *GetAssetGroupByBytedTokenResponse) SetBody(v string) *GetAssetGroupByBytedTokenResponse {
	s.Body = &v
	return s
}

type GetAssetGroupByBytedTokenResponseBuilder struct {
	s *GetAssetGroupByBytedTokenResponse
}

func NewGetAssetGroupByBytedTokenResponseBuilder() *GetAssetGroupByBytedTokenResponseBuilder {
	s := &GetAssetGroupByBytedTokenResponse{}
	b := &GetAssetGroupByBytedTokenResponseBuilder{s: s}
	return b
}

func (b *GetAssetGroupByBytedTokenResponseBuilder) RequestId(v string) *GetAssetGroupByBytedTokenResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetAssetGroupByBytedTokenResponseBuilder) ErrorMessage(v string) *GetAssetGroupByBytedTokenResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetAssetGroupByBytedTokenResponseBuilder) ErrorCode(v string) *GetAssetGroupByBytedTokenResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetAssetGroupByBytedTokenResponseBuilder) State(v string) *GetAssetGroupByBytedTokenResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetAssetGroupByBytedTokenResponseBuilder) Body(v string) *GetAssetGroupByBytedTokenResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *GetAssetGroupByBytedTokenResponseBuilder) Build() *GetAssetGroupByBytedTokenResponse {
	return b.s
}
