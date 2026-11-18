// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DetailRequest struct {
	DetailPath *DetailPath `json:"detailPath,omitempty"`
}

func (s DetailRequest) String() string {
	return utils.Beautify(s)
}

func (s DetailRequest) GoString() string {
	return s.String()
}

func (s DetailRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DetailRequest) SetDetailPath(v *DetailPath) *DetailRequest {
	s.DetailPath = v
	return s
}

type DetailRequestBuilder struct {
	s *DetailRequest
}

func NewDetailRequestBuilder() *DetailRequestBuilder {
	s := &DetailRequest{}
	b := &DetailRequestBuilder{s: s}
	return b
}

func (b *DetailRequestBuilder) DetailPath(v *DetailPath) *DetailRequestBuilder {
	b.s.DetailPath = v
	return b
}

func (b *DetailRequestBuilder) Build() *DetailRequest {
	return b.s
}
