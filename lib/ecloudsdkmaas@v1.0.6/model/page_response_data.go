// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PageResponseDataAgentModeEnum string

// List of AgentMode
const (
	PageResponseDataAgentModeEnumLlm        PageResponseDataAgentModeEnum = "LLM"
	PageResponseDataAgentModeEnumWorkflow   PageResponseDataAgentModeEnum = "WORKFLOW"
	PageResponseDataAgentModeEnumChat       PageResponseDataAgentModeEnum = "CHAT"
	PageResponseDataAgentModeEnumSupervisor PageResponseDataAgentModeEnum = "SUPERVISOR"
)

type PageResponseDataShelfStatusEnum string

// List of ShelfStatus
const (
	PageResponseDataShelfStatusEnumOn  PageResponseDataShelfStatusEnum = "ON"
	PageResponseDataShelfStatusEnumOff PageResponseDataShelfStatusEnum = "OFF"
)

type PageResponseData struct {

	// 智能体模式
	AgentMode *PageResponseDataAgentModeEnum `json:"agentMode,omitempty"`
	// 推理模型名称
	ModelName *string `json:"modelName,omitempty"`
	// 智能体更新时间
	UpdatedTime *string `json:"updatedTime,omitempty"`
	// 智能体ID
	AgentId *string `json:"agentId,omitempty"`
	// 智能体上下架状态（ON上下、OFF下架）
	ShelfStatus *PageResponseDataShelfStatusEnum `json:"shelfStatus,omitempty"`
	// 智能体名称
	Name *string `json:"name,omitempty"`
	// 智能体图标
	Icon *string `json:"icon,omitempty"`
	// 智能体创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
	// 智能体分类，openApi 创建的智能体一般默认OTHER，表示其它
	Category *string `json:"category,omitempty"`
	// 工作流ID
	WorkflowId *string `json:"workflowId,omitempty"`
	// 智能体简介
	Introduction *string `json:"introduction,omitempty"`
	// 智能体状态（0已保存、1已发布）
	Status *int32 `json:"status,omitempty"`
}

func (s PageResponseData) String() string {
	return utils.Beautify(s)
}

func (s PageResponseData) GoString() string {
	return s.String()
}

func (s PageResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PageResponseData) SetAgentMode(v PageResponseDataAgentModeEnum) *PageResponseData {
	s.AgentMode = &v
	return s
}

func (s *PageResponseData) SetModelName(v string) *PageResponseData {
	s.ModelName = &v
	return s
}

func (s *PageResponseData) SetUpdatedTime(v string) *PageResponseData {
	s.UpdatedTime = &v
	return s
}

func (s *PageResponseData) SetAgentId(v string) *PageResponseData {
	s.AgentId = &v
	return s
}

func (s *PageResponseData) SetShelfStatus(v PageResponseDataShelfStatusEnum) *PageResponseData {
	s.ShelfStatus = &v
	return s
}

func (s *PageResponseData) SetName(v string) *PageResponseData {
	s.Name = &v
	return s
}

func (s *PageResponseData) SetIcon(v string) *PageResponseData {
	s.Icon = &v
	return s
}

func (s *PageResponseData) SetCreatedTime(v string) *PageResponseData {
	s.CreatedTime = &v
	return s
}

func (s *PageResponseData) SetCategory(v string) *PageResponseData {
	s.Category = &v
	return s
}

func (s *PageResponseData) SetWorkflowId(v string) *PageResponseData {
	s.WorkflowId = &v
	return s
}

func (s *PageResponseData) SetIntroduction(v string) *PageResponseData {
	s.Introduction = &v
	return s
}

func (s *PageResponseData) SetStatus(v int32) *PageResponseData {
	s.Status = &v
	return s
}

type PageResponseDataBuilder struct {
	s *PageResponseData
}

func NewPageResponseDataBuilder() *PageResponseDataBuilder {
	s := &PageResponseData{}
	b := &PageResponseDataBuilder{s: s}
	return b
}

func (b *PageResponseDataBuilder) AgentMode(v PageResponseDataAgentModeEnum) *PageResponseDataBuilder {
	b.s.AgentMode = &v
	return b
}

func (b *PageResponseDataBuilder) ModelName(v string) *PageResponseDataBuilder {
	b.s.ModelName = &v
	return b
}

func (b *PageResponseDataBuilder) UpdatedTime(v string) *PageResponseDataBuilder {
	b.s.UpdatedTime = &v
	return b
}

func (b *PageResponseDataBuilder) AgentId(v string) *PageResponseDataBuilder {
	b.s.AgentId = &v
	return b
}

func (b *PageResponseDataBuilder) ShelfStatus(v PageResponseDataShelfStatusEnum) *PageResponseDataBuilder {
	b.s.ShelfStatus = &v
	return b
}

func (b *PageResponseDataBuilder) Name(v string) *PageResponseDataBuilder {
	b.s.Name = &v
	return b
}

func (b *PageResponseDataBuilder) Icon(v string) *PageResponseDataBuilder {
	b.s.Icon = &v
	return b
}

func (b *PageResponseDataBuilder) CreatedTime(v string) *PageResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *PageResponseDataBuilder) Category(v string) *PageResponseDataBuilder {
	b.s.Category = &v
	return b
}

func (b *PageResponseDataBuilder) WorkflowId(v string) *PageResponseDataBuilder {
	b.s.WorkflowId = &v
	return b
}

func (b *PageResponseDataBuilder) Introduction(v string) *PageResponseDataBuilder {
	b.s.Introduction = &v
	return b
}

func (b *PageResponseDataBuilder) Status(v int32) *PageResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *PageResponseDataBuilder) Build() *PageResponseData {
	return b.s
}
