// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryUsageDetailResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *QueryUsageDetailResponseBody `json:"body,omitempty"`
}

func (s QueryUsageDetailResponse) String() string {
	return utils.Beautify(s)
}

func (s QueryUsageDetailResponse) GoString() string {
	return s.String()
}

func (s QueryUsageDetailResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryUsageDetailResponse) SetRequestId(v string) *QueryUsageDetailResponse {
	s.RequestId = &v
	return s
}

func (s *QueryUsageDetailResponse) SetErrorMessage(v string) *QueryUsageDetailResponse {
	s.ErrorMessage = &v
	return s
}

func (s *QueryUsageDetailResponse) SetErrorCode(v string) *QueryUsageDetailResponse {
	s.ErrorCode = &v
	return s
}

func (s *QueryUsageDetailResponse) SetState(v string) *QueryUsageDetailResponse {
	s.State = &v
	return s
}

func (s *QueryUsageDetailResponse) SetBody(v *QueryUsageDetailResponseBody) *QueryUsageDetailResponse {
	s.Body = v
	return s
}

type QueryUsageDetailResponseBuilder struct {
	s *QueryUsageDetailResponse
}

func NewQueryUsageDetailResponseBuilder() *QueryUsageDetailResponseBuilder {
	s := &QueryUsageDetailResponse{}
	b := &QueryUsageDetailResponseBuilder{s: s}
	return b
}

func (b *QueryUsageDetailResponseBuilder) RequestId(v string) *QueryUsageDetailResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *QueryUsageDetailResponseBuilder) ErrorMessage(v string) *QueryUsageDetailResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *QueryUsageDetailResponseBuilder) ErrorCode(v string) *QueryUsageDetailResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *QueryUsageDetailResponseBuilder) State(v string) *QueryUsageDetailResponseBuilder {
	b.s.State = &v
	return b
}

func (b *QueryUsageDetailResponseBuilder) Body(v *QueryUsageDetailResponseBody) *QueryUsageDetailResponseBuilder {
	b.s.Body = v
	return b
}

func (b *QueryUsageDetailResponseBuilder) Build() *QueryUsageDetailResponse {
	return b.s
}
