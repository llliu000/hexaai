// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAssetGroupResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *GetAssetGroupResponseBody `json:"body,omitempty"`
}

func (s GetAssetGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s GetAssetGroupResponse) GoString() string {
	return s.String()
}

func (s GetAssetGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAssetGroupResponse) SetRequestId(v string) *GetAssetGroupResponse {
	s.RequestId = &v
	return s
}

func (s *GetAssetGroupResponse) SetErrorMessage(v string) *GetAssetGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetAssetGroupResponse) SetErrorCode(v string) *GetAssetGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetAssetGroupResponse) SetState(v string) *GetAssetGroupResponse {
	s.State = &v
	return s
}

func (s *GetAssetGroupResponse) SetBody(v *GetAssetGroupResponseBody) *GetAssetGroupResponse {
	s.Body = v
	return s
}

type GetAssetGroupResponseBuilder struct {
	s *GetAssetGroupResponse
}

func NewGetAssetGroupResponseBuilder() *GetAssetGroupResponseBuilder {
	s := &GetAssetGroupResponse{}
	b := &GetAssetGroupResponseBuilder{s: s}
	return b
}

func (b *GetAssetGroupResponseBuilder) RequestId(v string) *GetAssetGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetAssetGroupResponseBuilder) ErrorMessage(v string) *GetAssetGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetAssetGroupResponseBuilder) ErrorCode(v string) *GetAssetGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetAssetGroupResponseBuilder) State(v string) *GetAssetGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetAssetGroupResponseBuilder) Body(v *GetAssetGroupResponseBody) *GetAssetGroupResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetAssetGroupResponseBuilder) Build() *GetAssetGroupResponse {
	return b.s
}
