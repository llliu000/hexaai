// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OnShelfPath struct {
	position.Path
	// 智能体Id（URL路径参数）
	AgentId *string `json:"agentId,omitempty"`
}

func (s OnShelfPath) String() string {
	return utils.Beautify(s)
}

func (s OnShelfPath) GoString() string {
	return s.String()
}

func (s OnShelfPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OnShelfPath) SetAgentId(v string) *OnShelfPath {
	s.AgentId = &v
	return s
}

type OnShelfPathBuilder struct {
	s *OnShelfPath
}

func NewOnShelfPathBuilder() *OnShelfPathBuilder {
	s := &OnShelfPath{}
	b := &OnShelfPathBuilder{s: s}
	return b
}

func (b *OnShelfPathBuilder) AgentId(v string) *OnShelfPathBuilder {
	b.s.AgentId = &v
	return b
}

func (b *OnShelfPathBuilder) Build() *OnShelfPath {
	return b.s
}
