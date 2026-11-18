// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAssetRequest struct {
	CreateAssetBody *CreateAssetBody `json:"createAssetBody,omitempty"`
}

func (s CreateAssetRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateAssetRequest) GoString() string {
	return s.String()
}

func (s CreateAssetRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAssetRequest) SetCreateAssetBody(v *CreateAssetBody) *CreateAssetRequest {
	s.CreateAssetBody = v
	return s
}

type CreateAssetRequestBuilder struct {
	s *CreateAssetRequest
}

func NewCreateAssetRequestBuilder() *CreateAssetRequestBuilder {
	s := &CreateAssetRequest{}
	b := &CreateAssetRequestBuilder{s: s}
	return b
}

func (b *CreateAssetRequestBuilder) CreateAssetBody(v *CreateAssetBody) *CreateAssetRequestBuilder {
	b.s.CreateAssetBody = v
	return b
}

func (b *CreateAssetRequestBuilder) Build() *CreateAssetRequest {
	return b.s
}
