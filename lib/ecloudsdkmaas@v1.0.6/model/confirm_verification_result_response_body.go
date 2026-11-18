// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConfirmVerificationResultResponseBody struct {

	// 火山资产组ID
	GroupId *string `json:"groupId,omitempty"`
}

func (s ConfirmVerificationResultResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ConfirmVerificationResultResponseBody) GoString() string {
	return s.String()
}

func (s ConfirmVerificationResultResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConfirmVerificationResultResponseBody) SetGroupId(v string) *ConfirmVerificationResultResponseBody {
	s.GroupId = &v
	return s
}

type ConfirmVerificationResultResponseBodyBuilder struct {
	s *ConfirmVerificationResultResponseBody
}

func NewConfirmVerificationResultResponseBodyBuilder() *ConfirmVerificationResultResponseBodyBuilder {
	s := &ConfirmVerificationResultResponseBody{}
	b := &ConfirmVerificationResultResponseBodyBuilder{s: s}
	return b
}

func (b *ConfirmVerificationResultResponseBodyBuilder) GroupId(v string) *ConfirmVerificationResultResponseBodyBuilder {
	b.s.GroupId = &v
	return b
}

func (b *ConfirmVerificationResultResponseBodyBuilder) Build() *ConfirmVerificationResultResponseBody {
	return b.s
}
