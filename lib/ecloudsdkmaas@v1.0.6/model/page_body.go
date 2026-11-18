// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PageBodyAgentModeEnum string

// List of AgentMode
const (
	PageBodyAgentModeEnumLlm        PageBodyAgentModeEnum = "LLM"
	PageBodyAgentModeEnumWorkflow   PageBodyAgentModeEnum = "WORKFLOW"
	PageBodyAgentModeEnumChat       PageBodyAgentModeEnum = "CHAT"
	PageBodyAgentModeEnumSupervisor PageBodyAgentModeEnum = "SUPERVISOR"
)

type PageBodyShelfStatusEnum string

// List of ShelfStatus
const (
	PageBodyShelfStatusEnumOn  PageBodyShelfStatusEnum = "ON"
	PageBodyShelfStatusEnumOff PageBodyShelfStatusEnum = "OFF"
)

type PageBody struct {
	position.Body
	// 智能体模式
	AgentMode *PageBodyAgentModeEnum `json:"agentMode,omitempty"`
	// 上下架状态： ON上下、OFF下架
	ShelfStatus *PageBodyShelfStatusEnum `json:"shelfStatus,omitempty"`
	// 页码（从 1 开始）
	PageNo *int32 `json:"pageNo,omitempty"`
	// 智能体名称
	Name *string `json:"name,omitempty"`
	// 每页行数
	PageSize *int32 `json:"pageSize,omitempty"`
	// 发布状态：0已保存、1已发布
	Status *int32 `json:"status,omitempty"`
}

func (s PageBody) String() string {
	return utils.Beautify(s)
}

func (s PageBody) GoString() string {
	return s.String()
}

func (s PageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PageBody) SetAgentMode(v PageBodyAgentModeEnum) *PageBody {
	s.AgentMode = &v
	return s
}

func (s *PageBody) SetShelfStatus(v PageBodyShelfStatusEnum) *PageBody {
	s.ShelfStatus = &v
	return s
}

func (s *PageBody) SetPageNo(v int32) *PageBody {
	s.PageNo = &v
	return s
}

func (s *PageBody) SetName(v string) *PageBody {
	s.Name = &v
	return s
}

func (s *PageBody) SetPageSize(v int32) *PageBody {
	s.PageSize = &v
	return s
}

func (s *PageBody) SetStatus(v int32) *PageBody {
	s.Status = &v
	return s
}

type PageBodyBuilder struct {
	s *PageBody
}

func NewPageBodyBuilder() *PageBodyBuilder {
	s := &PageBody{}
	b := &PageBodyBuilder{s: s}
	return b
}

func (b *PageBodyBuilder) AgentMode(v PageBodyAgentModeEnum) *PageBodyBuilder {
	b.s.AgentMode = &v
	return b
}

func (b *PageBodyBuilder) ShelfStatus(v PageBodyShelfStatusEnum) *PageBodyBuilder {
	b.s.ShelfStatus = &v
	return b
}

func (b *PageBodyBuilder) PageNo(v int32) *PageBodyBuilder {
	b.s.PageNo = &v
	return b
}

func (b *PageBodyBuilder) Name(v string) *PageBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *PageBodyBuilder) PageSize(v int32) *PageBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *PageBodyBuilder) Status(v int32) *PageBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *PageBodyBuilder) Build() *PageBody {
	return b.s
}
