// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryAICCExportTaskDetailResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 业务返回体
	Body *QueryAICCExportTaskDetailResponseBody `json:"body,omitempty"`
}

func (s QueryAICCExportTaskDetailResponse) String() string {
	return utils.Beautify(s)
}

func (s QueryAICCExportTaskDetailResponse) GoString() string {
	return s.String()
}

func (s QueryAICCExportTaskDetailResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryAICCExportTaskDetailResponse) SetRequestId(v string) *QueryAICCExportTaskDetailResponse {
	s.RequestId = &v
	return s
}

func (s *QueryAICCExportTaskDetailResponse) SetErrorMessage(v string) *QueryAICCExportTaskDetailResponse {
	s.ErrorMessage = &v
	return s
}

func (s *QueryAICCExportTaskDetailResponse) SetErrorCode(v string) *QueryAICCExportTaskDetailResponse {
	s.ErrorCode = &v
	return s
}

func (s *QueryAICCExportTaskDetailResponse) SetState(v string) *QueryAICCExportTaskDetailResponse {
	s.State = &v
	return s
}

func (s *QueryAICCExportTaskDetailResponse) SetBody(v *QueryAICCExportTaskDetailResponseBody) *QueryAICCExportTaskDetailResponse {
	s.Body = v
	return s
}

type QueryAICCExportTaskDetailResponseBuilder struct {
	s *QueryAICCExportTaskDetailResponse
}

func NewQueryAICCExportTaskDetailResponseBuilder() *QueryAICCExportTaskDetailResponseBuilder {
	s := &QueryAICCExportTaskDetailResponse{}
	b := &QueryAICCExportTaskDetailResponseBuilder{s: s}
	return b
}

func (b *QueryAICCExportTaskDetailResponseBuilder) RequestId(v string) *QueryAICCExportTaskDetailResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *QueryAICCExportTaskDetailResponseBuilder) ErrorMessage(v string) *QueryAICCExportTaskDetailResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *QueryAICCExportTaskDetailResponseBuilder) ErrorCode(v string) *QueryAICCExportTaskDetailResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *QueryAICCExportTaskDetailResponseBuilder) State(v string) *QueryAICCExportTaskDetailResponseBuilder {
	b.s.State = &v
	return b
}

func (b *QueryAICCExportTaskDetailResponseBuilder) Body(v *QueryAICCExportTaskDetailResponseBody) *QueryAICCExportTaskDetailResponseBuilder {
	b.s.Body = v
	return b
}

func (b *QueryAICCExportTaskDetailResponseBuilder) Build() *QueryAICCExportTaskDetailResponse {
	return b.s
}
