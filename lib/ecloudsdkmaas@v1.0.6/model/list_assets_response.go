// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListAssetsResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *ListAssetsResponseBody `json:"body,omitempty"`
}

func (s ListAssetsResponse) String() string {
	return utils.Beautify(s)
}

func (s ListAssetsResponse) GoString() string {
	return s.String()
}

func (s ListAssetsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListAssetsResponse) SetRequestId(v string) *ListAssetsResponse {
	s.RequestId = &v
	return s
}

func (s *ListAssetsResponse) SetErrorMessage(v string) *ListAssetsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListAssetsResponse) SetErrorCode(v string) *ListAssetsResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListAssetsResponse) SetState(v string) *ListAssetsResponse {
	s.State = &v
	return s
}

func (s *ListAssetsResponse) SetBody(v *ListAssetsResponseBody) *ListAssetsResponse {
	s.Body = v
	return s
}

type ListAssetsResponseBuilder struct {
	s *ListAssetsResponse
}

func NewListAssetsResponseBuilder() *ListAssetsResponseBuilder {
	s := &ListAssetsResponse{}
	b := &ListAssetsResponseBuilder{s: s}
	return b
}

func (b *ListAssetsResponseBuilder) RequestId(v string) *ListAssetsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListAssetsResponseBuilder) ErrorMessage(v string) *ListAssetsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListAssetsResponseBuilder) ErrorCode(v string) *ListAssetsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListAssetsResponseBuilder) State(v string) *ListAssetsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListAssetsResponseBuilder) Body(v *ListAssetsResponseBody) *ListAssetsResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListAssetsResponseBuilder) Build() *ListAssetsResponse {
	return b.s
}
