// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateSessionRequest struct {
	CreateSessionBody *CreateSessionBody `json:"createSessionBody,omitempty"`
}

func (s CreateSessionRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateSessionRequest) GoString() string {
	return s.String()
}

func (s CreateSessionRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSessionRequest) SetCreateSessionBody(v *CreateSessionBody) *CreateSessionRequest {
	s.CreateSessionBody = v
	return s
}

type CreateSessionRequestBuilder struct {
	s *CreateSessionRequest
}

func NewCreateSessionRequestBuilder() *CreateSessionRequestBuilder {
	s := &CreateSessionRequest{}
	b := &CreateSessionRequestBuilder{s: s}
	return b
}

func (b *CreateSessionRequestBuilder) CreateSessionBody(v *CreateSessionBody) *CreateSessionRequestBuilder {
	b.s.CreateSessionBody = v
	return b
}

func (b *CreateSessionRequestBuilder) Build() *CreateSessionRequest {
	return b.s
}
