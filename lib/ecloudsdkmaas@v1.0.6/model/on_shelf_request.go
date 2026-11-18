// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OnShelfRequest struct {
	OnShelfPath *OnShelfPath `json:"onShelfPath,omitempty"`
}

func (s OnShelfRequest) String() string {
	return utils.Beautify(s)
}

func (s OnShelfRequest) GoString() string {
	return s.String()
}

func (s OnShelfRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OnShelfRequest) SetOnShelfPath(v *OnShelfPath) *OnShelfRequest {
	s.OnShelfPath = v
	return s
}

type OnShelfRequestBuilder struct {
	s *OnShelfRequest
}

func NewOnShelfRequestBuilder() *OnShelfRequestBuilder {
	s := &OnShelfRequest{}
	b := &OnShelfRequestBuilder{s: s}
	return b
}

func (b *OnShelfRequestBuilder) OnShelfPath(v *OnShelfPath) *OnShelfRequestBuilder {
	b.s.OnShelfPath = v
	return b
}

func (b *OnShelfRequestBuilder) Build() *OnShelfRequest {
	return b.s
}
