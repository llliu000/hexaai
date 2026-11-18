// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteDocRequest struct {
	DeleteDocPath *DeleteDocPath `json:"deleteDocPath,omitempty"`
}

func (s DeleteDocRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteDocRequest) GoString() string {
	return s.String()
}

func (s DeleteDocRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteDocRequest) SetDeleteDocPath(v *DeleteDocPath) *DeleteDocRequest {
	s.DeleteDocPath = v
	return s
}

type DeleteDocRequestBuilder struct {
	s *DeleteDocRequest
}

func NewDeleteDocRequestBuilder() *DeleteDocRequestBuilder {
	s := &DeleteDocRequest{}
	b := &DeleteDocRequestBuilder{s: s}
	return b
}

func (b *DeleteDocRequestBuilder) DeleteDocPath(v *DeleteDocPath) *DeleteDocRequestBuilder {
	b.s.DeleteDocPath = v
	return b
}

func (b *DeleteDocRequestBuilder) Build() *DeleteDocRequest {
	return b.s
}
