// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PublishRequest struct {
	PublishPath *PublishPath `json:"publishPath,omitempty"`

	PublishQuery *PublishQuery `json:"publishQuery,omitempty"`
}

func (s PublishRequest) String() string {
	return utils.Beautify(s)
}

func (s PublishRequest) GoString() string {
	return s.String()
}

func (s PublishRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PublishRequest) SetPublishPath(v *PublishPath) *PublishRequest {
	s.PublishPath = v
	return s
}

func (s *PublishRequest) SetPublishQuery(v *PublishQuery) *PublishRequest {
	s.PublishQuery = v
	return s
}

type PublishRequestBuilder struct {
	s *PublishRequest
}

func NewPublishRequestBuilder() *PublishRequestBuilder {
	s := &PublishRequest{}
	b := &PublishRequestBuilder{s: s}
	return b
}

func (b *PublishRequestBuilder) PublishPath(v *PublishPath) *PublishRequestBuilder {
	b.s.PublishPath = v
	return b
}

func (b *PublishRequestBuilder) PublishQuery(v *PublishQuery) *PublishRequestBuilder {
	b.s.PublishQuery = v
	return b
}

func (b *PublishRequestBuilder) Build() *PublishRequest {
	return b.s
}
