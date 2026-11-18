// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryUsageStatsResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *QueryUsageStatsResponseBody `json:"body,omitempty"`
}

func (s QueryUsageStatsResponse) String() string {
	return utils.Beautify(s)
}

func (s QueryUsageStatsResponse) GoString() string {
	return s.String()
}

func (s QueryUsageStatsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryUsageStatsResponse) SetRequestId(v string) *QueryUsageStatsResponse {
	s.RequestId = &v
	return s
}

func (s *QueryUsageStatsResponse) SetErrorMessage(v string) *QueryUsageStatsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *QueryUsageStatsResponse) SetErrorCode(v string) *QueryUsageStatsResponse {
	s.ErrorCode = &v
	return s
}

func (s *QueryUsageStatsResponse) SetState(v string) *QueryUsageStatsResponse {
	s.State = &v
	return s
}

func (s *QueryUsageStatsResponse) SetBody(v *QueryUsageStatsResponseBody) *QueryUsageStatsResponse {
	s.Body = v
	return s
}

type QueryUsageStatsResponseBuilder struct {
	s *QueryUsageStatsResponse
}

func NewQueryUsageStatsResponseBuilder() *QueryUsageStatsResponseBuilder {
	s := &QueryUsageStatsResponse{}
	b := &QueryUsageStatsResponseBuilder{s: s}
	return b
}

func (b *QueryUsageStatsResponseBuilder) RequestId(v string) *QueryUsageStatsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *QueryUsageStatsResponseBuilder) ErrorMessage(v string) *QueryUsageStatsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *QueryUsageStatsResponseBuilder) ErrorCode(v string) *QueryUsageStatsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *QueryUsageStatsResponseBuilder) State(v string) *QueryUsageStatsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *QueryUsageStatsResponseBuilder) Body(v *QueryUsageStatsResponseBody) *QueryUsageStatsResponseBuilder {
	b.s.Body = v
	return b
}

func (b *QueryUsageStatsResponseBuilder) Build() *QueryUsageStatsResponse {
	return b.s
}
