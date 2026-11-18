// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAssetGroupResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *CreateAssetGroupResponseBody `json:"body,omitempty"`
}

func (s CreateAssetGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateAssetGroupResponse) GoString() string {
	return s.String()
}

func (s CreateAssetGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAssetGroupResponse) SetRequestId(v string) *CreateAssetGroupResponse {
	s.RequestId = &v
	return s
}

func (s *CreateAssetGroupResponse) SetErrorMessage(v string) *CreateAssetGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAssetGroupResponse) SetErrorCode(v string) *CreateAssetGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateAssetGroupResponse) SetState(v string) *CreateAssetGroupResponse {
	s.State = &v
	return s
}

func (s *CreateAssetGroupResponse) SetBody(v *CreateAssetGroupResponseBody) *CreateAssetGroupResponse {
	s.Body = v
	return s
}

type CreateAssetGroupResponseBuilder struct {
	s *CreateAssetGroupResponse
}

func NewCreateAssetGroupResponseBuilder() *CreateAssetGroupResponseBuilder {
	s := &CreateAssetGroupResponse{}
	b := &CreateAssetGroupResponseBuilder{s: s}
	return b
}

func (b *CreateAssetGroupResponseBuilder) RequestId(v string) *CreateAssetGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateAssetGroupResponseBuilder) ErrorMessage(v string) *CreateAssetGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateAssetGroupResponseBuilder) ErrorCode(v string) *CreateAssetGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateAssetGroupResponseBuilder) State(v string) *CreateAssetGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateAssetGroupResponseBuilder) Body(v *CreateAssetGroupResponseBody) *CreateAssetGroupResponseBuilder {
	b.s.Body = v
	return b
}

func (b *CreateAssetGroupResponseBuilder) Build() *CreateAssetGroupResponse {
	return b.s
}
