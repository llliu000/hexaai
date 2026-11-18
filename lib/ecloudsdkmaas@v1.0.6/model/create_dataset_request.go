// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateDatasetRequest struct {
	CreateDatasetQuery *CreateDatasetQuery `json:"createDatasetQuery,omitempty"`

	CreateDatasetBody *CreateDatasetBody `json:"createDatasetBody,omitempty"`
}

func (s CreateDatasetRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s CreateDatasetRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateDatasetRequest) SetCreateDatasetQuery(v *CreateDatasetQuery) *CreateDatasetRequest {
	s.CreateDatasetQuery = v
	return s
}

func (s *CreateDatasetRequest) SetCreateDatasetBody(v *CreateDatasetBody) *CreateDatasetRequest {
	s.CreateDatasetBody = v
	return s
}

type CreateDatasetRequestBuilder struct {
	s *CreateDatasetRequest
}

func NewCreateDatasetRequestBuilder() *CreateDatasetRequestBuilder {
	s := &CreateDatasetRequest{}
	b := &CreateDatasetRequestBuilder{s: s}
	return b
}

func (b *CreateDatasetRequestBuilder) CreateDatasetQuery(v *CreateDatasetQuery) *CreateDatasetRequestBuilder {
	b.s.CreateDatasetQuery = v
	return b
}

func (b *CreateDatasetRequestBuilder) CreateDatasetBody(v *CreateDatasetBody) *CreateDatasetRequestBuilder {
	b.s.CreateDatasetBody = v
	return b
}

func (b *CreateDatasetRequestBuilder) Build() *CreateDatasetRequest {
	return b.s
}
