// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateExportTaskRequest struct {
	CreateExportTaskBody *CreateExportTaskBody `json:"createExportTaskBody,omitempty"`
}

func (s CreateExportTaskRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateExportTaskRequest) GoString() string {
	return s.String()
}

func (s CreateExportTaskRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateExportTaskRequest) SetCreateExportTaskBody(v *CreateExportTaskBody) *CreateExportTaskRequest {
	s.CreateExportTaskBody = v
	return s
}

type CreateExportTaskRequestBuilder struct {
	s *CreateExportTaskRequest
}

func NewCreateExportTaskRequestBuilder() *CreateExportTaskRequestBuilder {
	s := &CreateExportTaskRequest{}
	b := &CreateExportTaskRequestBuilder{s: s}
	return b
}

func (b *CreateExportTaskRequestBuilder) CreateExportTaskBody(v *CreateExportTaskBody) *CreateExportTaskRequestBuilder {
	b.s.CreateExportTaskBody = v
	return b
}

func (b *CreateExportTaskRequestBuilder) Build() *CreateExportTaskRequest {
	return b.s
}
