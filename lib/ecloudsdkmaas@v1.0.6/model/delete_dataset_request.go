// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteDatasetRequest struct {
	DeleteDatasetPath *DeleteDatasetPath `json:"deleteDatasetPath,omitempty"`
}

func (s DeleteDatasetRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteDatasetRequest) GoString() string {
	return s.String()
}

func (s DeleteDatasetRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteDatasetRequest) SetDeleteDatasetPath(v *DeleteDatasetPath) *DeleteDatasetRequest {
	s.DeleteDatasetPath = v
	return s
}

type DeleteDatasetRequestBuilder struct {
	s *DeleteDatasetRequest
}

func NewDeleteDatasetRequestBuilder() *DeleteDatasetRequestBuilder {
	s := &DeleteDatasetRequest{}
	b := &DeleteDatasetRequestBuilder{s: s}
	return b
}

func (b *DeleteDatasetRequestBuilder) DeleteDatasetPath(v *DeleteDatasetPath) *DeleteDatasetRequestBuilder {
	b.s.DeleteDatasetPath = v
	return b
}

func (b *DeleteDatasetRequestBuilder) Build() *DeleteDatasetRequest {
	return b.s
}
