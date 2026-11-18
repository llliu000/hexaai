// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateAssetGroupResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`

	Body *UpdateAssetGroupResponseBody `json:"body,omitempty"`
}

func (s UpdateAssetGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateAssetGroupResponse) GoString() string {
	return s.String()
}

func (s UpdateAssetGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateAssetGroupResponse) SetRequestId(v string) *UpdateAssetGroupResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateAssetGroupResponse) SetErrorMessage(v string) *UpdateAssetGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAssetGroupResponse) SetErrorCode(v string) *UpdateAssetGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAssetGroupResponse) SetState(v string) *UpdateAssetGroupResponse {
	s.State = &v
	return s
}

func (s *UpdateAssetGroupResponse) SetBody(v *UpdateAssetGroupResponseBody) *UpdateAssetGroupResponse {
	s.Body = v
	return s
}

type UpdateAssetGroupResponseBuilder struct {
	s *UpdateAssetGroupResponse
}

func NewUpdateAssetGroupResponseBuilder() *UpdateAssetGroupResponseBuilder {
	s := &UpdateAssetGroupResponse{}
	b := &UpdateAssetGroupResponseBuilder{s: s}
	return b
}

func (b *UpdateAssetGroupResponseBuilder) RequestId(v string) *UpdateAssetGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateAssetGroupResponseBuilder) ErrorMessage(v string) *UpdateAssetGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateAssetGroupResponseBuilder) ErrorCode(v string) *UpdateAssetGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateAssetGroupResponseBuilder) State(v string) *UpdateAssetGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateAssetGroupResponseBuilder) Body(v *UpdateAssetGroupResponseBody) *UpdateAssetGroupResponseBuilder {
	b.s.Body = v
	return b
}

func (b *UpdateAssetGroupResponseBuilder) Build() *UpdateAssetGroupResponse {
	return b.s
}
