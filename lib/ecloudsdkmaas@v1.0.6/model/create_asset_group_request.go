// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAssetGroupRequest struct {
	CreateAssetGroupBody *CreateAssetGroupBody `json:"createAssetGroupBody,omitempty"`
}

func (s CreateAssetGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateAssetGroupRequest) GoString() string {
	return s.String()
}

func (s CreateAssetGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAssetGroupRequest) SetCreateAssetGroupBody(v *CreateAssetGroupBody) *CreateAssetGroupRequest {
	s.CreateAssetGroupBody = v
	return s
}

type CreateAssetGroupRequestBuilder struct {
	s *CreateAssetGroupRequest
}

func NewCreateAssetGroupRequestBuilder() *CreateAssetGroupRequestBuilder {
	s := &CreateAssetGroupRequest{}
	b := &CreateAssetGroupRequestBuilder{s: s}
	return b
}

func (b *CreateAssetGroupRequestBuilder) CreateAssetGroupBody(v *CreateAssetGroupBody) *CreateAssetGroupRequestBuilder {
	b.s.CreateAssetGroupBody = v
	return b
}

func (b *CreateAssetGroupRequestBuilder) Build() *CreateAssetGroupRequest {
	return b.s
}
