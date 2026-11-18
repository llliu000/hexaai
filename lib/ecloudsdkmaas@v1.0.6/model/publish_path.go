// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PublishPath struct {
	position.Path
	// 智能体Id（URL路径参数）
	AgentId *string `json:"agentId,omitempty"`
}

func (s PublishPath) String() string {
	return utils.Beautify(s)
}

func (s PublishPath) GoString() string {
	return s.String()
}

func (s PublishPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PublishPath) SetAgentId(v string) *PublishPath {
	s.AgentId = &v
	return s
}

type PublishPathBuilder struct {
	s *PublishPath
}

func NewPublishPathBuilder() *PublishPathBuilder {
	s := &PublishPath{}
	b := &PublishPathBuilder{s: s}
	return b
}

func (b *PublishPathBuilder) AgentId(v string) *PublishPathBuilder {
	b.s.AgentId = &v
	return b
}

func (b *PublishPathBuilder) Build() *PublishPath {
	return b.s
}
