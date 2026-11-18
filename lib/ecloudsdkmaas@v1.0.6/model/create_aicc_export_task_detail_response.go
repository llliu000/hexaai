// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAICCExportTaskDetailResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 业务返回体
	Body *CreateAICCExportTaskDetailResponseBody `json:"body,omitempty"`
}

func (s CreateAICCExportTaskDetailResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateAICCExportTaskDetailResponse) GoString() string {
	return s.String()
}

func (s CreateAICCExportTaskDetailResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAICCExportTaskDetailResponse) SetRequestId(v string) *CreateAICCExportTaskDetailResponse {
	s.RequestId = &v
	return s
}

func (s *CreateAICCExportTaskDetailResponse) SetErrorMessage(v string) *CreateAICCExportTaskDetailResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAICCExportTaskDetailResponse) SetErrorCode(v string) *CreateAICCExportTaskDetailResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateAICCExportTaskDetailResponse) SetState(v string) *CreateAICCExportTaskDetailResponse {
	s.State = &v
	return s
}

func (s *CreateAICCExportTaskDetailResponse) SetBody(v *CreateAICCExportTaskDetailResponseBody) *CreateAICCExportTaskDetailResponse {
	s.Body = v
	return s
}

type CreateAICCExportTaskDetailResponseBuilder struct {
	s *CreateAICCExportTaskDetailResponse
}

func NewCreateAICCExportTaskDetailResponseBuilder() *CreateAICCExportTaskDetailResponseBuilder {
	s := &CreateAICCExportTaskDetailResponse{}
	b := &CreateAICCExportTaskDetailResponseBuilder{s: s}
	return b
}

func (b *CreateAICCExportTaskDetailResponseBuilder) RequestId(v string) *CreateAICCExportTaskDetailResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateAICCExportTaskDetailResponseBuilder) ErrorMessage(v string) *CreateAICCExportTaskDetailResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateAICCExportTaskDetailResponseBuilder) ErrorCode(v string) *CreateAICCExportTaskDetailResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateAICCExportTaskDetailResponseBuilder) State(v string) *CreateAICCExportTaskDetailResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateAICCExportTaskDetailResponseBuilder) Body(v *CreateAICCExportTaskDetailResponseBody) *CreateAICCExportTaskDetailResponseBuilder {
	b.s.Body = v
	return b
}

func (b *CreateAICCExportTaskDetailResponseBuilder) Build() *CreateAICCExportTaskDetailResponse {
	return b.s
}
