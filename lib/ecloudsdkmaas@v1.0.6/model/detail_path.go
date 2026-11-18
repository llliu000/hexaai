// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DetailPath struct {
	position.Path
	// 智能体Id（URL路径参数）
	AgentId *string `json:"agentId,omitempty"`
}

func (s DetailPath) String() string {
	return utils.Beautify(s)
}

func (s DetailPath) GoString() string {
	return s.String()
}

func (s DetailPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DetailPath) SetAgentId(v string) *DetailPath {
	s.AgentId = &v
	return s
}

type DetailPathBuilder struct {
	s *DetailPath
}

func NewDetailPathBuilder() *DetailPathBuilder {
	s := &DetailPath{}
	b := &DetailPathBuilder{s: s}
	return b
}

func (b *DetailPathBuilder) AgentId(v string) *DetailPathBuilder {
	b.s.AgentId = &v
	return b
}

func (b *DetailPathBuilder) Build() *DetailPath {
	return b.s
}
