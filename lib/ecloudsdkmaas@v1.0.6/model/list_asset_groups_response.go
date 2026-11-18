// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListAssetGroupsResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *ListAssetGroupsResponseBody `json:"body,omitempty"`
}

func (s ListAssetGroupsResponse) String() string {
	return utils.Beautify(s)
}

func (s ListAssetGroupsResponse) GoString() string {
	return s.String()
}

func (s ListAssetGroupsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListAssetGroupsResponse) SetRequestId(v string) *ListAssetGroupsResponse {
	s.RequestId = &v
	return s
}

func (s *ListAssetGroupsResponse) SetErrorMessage(v string) *ListAssetGroupsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListAssetGroupsResponse) SetErrorCode(v string) *ListAssetGroupsResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListAssetGroupsResponse) SetState(v string) *ListAssetGroupsResponse {
	s.State = &v
	return s
}

func (s *ListAssetGroupsResponse) SetBody(v *ListAssetGroupsResponseBody) *ListAssetGroupsResponse {
	s.Body = v
	return s
}

type ListAssetGroupsResponseBuilder struct {
	s *ListAssetGroupsResponse
}

func NewListAssetGroupsResponseBuilder() *ListAssetGroupsResponseBuilder {
	s := &ListAssetGroupsResponse{}
	b := &ListAssetGroupsResponseBuilder{s: s}
	return b
}

func (b *ListAssetGroupsResponseBuilder) RequestId(v string) *ListAssetGroupsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListAssetGroupsResponseBuilder) ErrorMessage(v string) *ListAssetGroupsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListAssetGroupsResponseBuilder) ErrorCode(v string) *ListAssetGroupsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListAssetGroupsResponseBuilder) State(v string) *ListAssetGroupsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListAssetGroupsResponseBuilder) Body(v *ListAssetGroupsResponseBody) *ListAssetGroupsResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListAssetGroupsResponseBuilder) Build() *ListAssetGroupsResponse {
	return b.s
}
