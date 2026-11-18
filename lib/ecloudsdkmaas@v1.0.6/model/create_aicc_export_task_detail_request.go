// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAICCExportTaskDetailRequest struct {
	CreateAICCExportTaskDetailBody *CreateAICCExportTaskDetailBody `json:"createAICCExportTaskDetailBody,omitempty"`
}

func (s CreateAICCExportTaskDetailRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateAICCExportTaskDetailRequest) GoString() string {
	return s.String()
}

func (s CreateAICCExportTaskDetailRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAICCExportTaskDetailRequest) SetCreateAICCExportTaskDetailBody(v *CreateAICCExportTaskDetailBody) *CreateAICCExportTaskDetailRequest {
	s.CreateAICCExportTaskDetailBody = v
	return s
}

type CreateAICCExportTaskDetailRequestBuilder struct {
	s *CreateAICCExportTaskDetailRequest
}

func NewCreateAICCExportTaskDetailRequestBuilder() *CreateAICCExportTaskDetailRequestBuilder {
	s := &CreateAICCExportTaskDetailRequest{}
	b := &CreateAICCExportTaskDetailRequestBuilder{s: s}
	return b
}

func (b *CreateAICCExportTaskDetailRequestBuilder) CreateAICCExportTaskDetailBody(v *CreateAICCExportTaskDetailBody) *CreateAICCExportTaskDetailRequestBuilder {
	b.s.CreateAICCExportTaskDetailBody = v
	return b
}

func (b *CreateAICCExportTaskDetailRequestBuilder) Build() *CreateAICCExportTaskDetailRequest {
	return b.s
}
