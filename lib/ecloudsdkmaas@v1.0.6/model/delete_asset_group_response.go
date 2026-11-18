// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteAssetGroupResponse struct {

	// 请求id
	RequestId *string `json:"requestId,omitempty"`
	// 异常消息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 异常码
	ErrorCode *string `json:"errorCode,omitempty"`
	// 状态是否正常，OK代表返回正常，ERROR代表返回错误
	State *string `json:"state,omitempty"`
	// 返回体
	Body *bool `json:"body,omitempty"`
}

func (s DeleteAssetGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteAssetGroupResponse) GoString() string {
	return s.String()
}

func (s DeleteAssetGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteAssetGroupResponse) SetRequestId(v string) *DeleteAssetGroupResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteAssetGroupResponse) SetErrorMessage(v string) *DeleteAssetGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteAssetGroupResponse) SetErrorCode(v string) *DeleteAssetGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAssetGroupResponse) SetState(v string) *DeleteAssetGroupResponse {
	s.State = &v
	return s
}

func (s *DeleteAssetGroupResponse) SetBody(v bool) *DeleteAssetGroupResponse {
	s.Body = &v
	return s
}

type DeleteAssetGroupResponseBuilder struct {
	s *DeleteAssetGroupResponse
}

func NewDeleteAssetGroupResponseBuilder() *DeleteAssetGroupResponseBuilder {
	s := &DeleteAssetGroupResponse{}
	b := &DeleteAssetGroupResponseBuilder{s: s}
	return b
}

func (b *DeleteAssetGroupResponseBuilder) RequestId(v string) *DeleteAssetGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteAssetGroupResponseBuilder) ErrorMessage(v string) *DeleteAssetGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteAssetGroupResponseBuilder) ErrorCode(v string) *DeleteAssetGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteAssetGroupResponseBuilder) State(v string) *DeleteAssetGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteAssetGroupResponseBuilder) Body(v bool) *DeleteAssetGroupResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteAssetGroupResponseBuilder) Build() *DeleteAssetGroupResponse {
	return b.s
}
