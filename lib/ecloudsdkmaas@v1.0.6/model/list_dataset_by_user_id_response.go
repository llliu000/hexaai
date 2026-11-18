// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListDatasetByUserIdResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *ListDatasetByUserIdResponseBody `json:"body,omitempty"`
}

func (s ListDatasetByUserIdResponse) String() string {
	return utils.Beautify(s)
}

func (s ListDatasetByUserIdResponse) GoString() string {
	return s.String()
}

func (s ListDatasetByUserIdResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListDatasetByUserIdResponse) SetRequestId(v string) *ListDatasetByUserIdResponse {
	s.RequestId = &v
	return s
}

func (s *ListDatasetByUserIdResponse) SetErrorMessage(v string) *ListDatasetByUserIdResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListDatasetByUserIdResponse) SetErrorCode(v string) *ListDatasetByUserIdResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListDatasetByUserIdResponse) SetState(v string) *ListDatasetByUserIdResponse {
	s.State = &v
	return s
}

func (s *ListDatasetByUserIdResponse) SetBody(v *ListDatasetByUserIdResponseBody) *ListDatasetByUserIdResponse {
	s.Body = v
	return s
}

type ListDatasetByUserIdResponseBuilder struct {
	s *ListDatasetByUserIdResponse
}

func NewListDatasetByUserIdResponseBuilder() *ListDatasetByUserIdResponseBuilder {
	s := &ListDatasetByUserIdResponse{}
	b := &ListDatasetByUserIdResponseBuilder{s: s}
	return b
}

func (b *ListDatasetByUserIdResponseBuilder) RequestId(v string) *ListDatasetByUserIdResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListDatasetByUserIdResponseBuilder) ErrorMessage(v string) *ListDatasetByUserIdResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListDatasetByUserIdResponseBuilder) ErrorCode(v string) *ListDatasetByUserIdResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListDatasetByUserIdResponseBuilder) State(v string) *ListDatasetByUserIdResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListDatasetByUserIdResponseBuilder) Body(v *ListDatasetByUserIdResponseBody) *ListDatasetByUserIdResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListDatasetByUserIdResponseBuilder) Build() *ListDatasetByUserIdResponse {
	return b.s
}
