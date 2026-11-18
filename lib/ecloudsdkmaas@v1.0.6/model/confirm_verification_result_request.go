// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConfirmVerificationResultRequest struct {
	ConfirmVerificationResultBody *ConfirmVerificationResultBody `json:"confirmVerificationResultBody,omitempty"`
}

func (s ConfirmVerificationResultRequest) String() string {
	return utils.Beautify(s)
}

func (s ConfirmVerificationResultRequest) GoString() string {
	return s.String()
}

func (s ConfirmVerificationResultRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConfirmVerificationResultRequest) SetConfirmVerificationResultBody(v *ConfirmVerificationResultBody) *ConfirmVerificationResultRequest {
	s.ConfirmVerificationResultBody = v
	return s
}

type ConfirmVerificationResultRequestBuilder struct {
	s *ConfirmVerificationResultRequest
}

func NewConfirmVerificationResultRequestBuilder() *ConfirmVerificationResultRequestBuilder {
	s := &ConfirmVerificationResultRequest{}
	b := &ConfirmVerificationResultRequestBuilder{s: s}
	return b
}

func (b *ConfirmVerificationResultRequestBuilder) ConfirmVerificationResultBody(v *ConfirmVerificationResultBody) *ConfirmVerificationResultRequestBuilder {
	b.s.ConfirmVerificationResultBody = v
	return b
}

func (b *ConfirmVerificationResultRequestBuilder) Build() *ConfirmVerificationResultRequest {
	return b.s
}
