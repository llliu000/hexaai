// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DetailResponseBodyShelfStatusEnum string

// List of ShelfStatus
const (
	DetailResponseBodyShelfStatusEnumOn  DetailResponseBodyShelfStatusEnum = "ON"
	DetailResponseBodyShelfStatusEnumOff DetailResponseBodyShelfStatusEnum = "OFF"
)

type DetailResponseBodyAgentModeEnum string

// List of AgentMode
const (
	DetailResponseBodyAgentModeEnumLlm      DetailResponseBodyAgentModeEnum = "LLM"
	DetailResponseBodyAgentModeEnumWorkflow DetailResponseBodyAgentModeEnum = "WORKFLOW"
	DetailResponseBodyAgentModeEnumChat     DetailResponseBodyAgentModeEnum = "CHAT"
)

type DetailResponseBody struct {

	// 智能体更新时间
	UpdatedTime *string `json:"updatedTime,omitempty"`
	// 智能体id
	AgentId *string `json:"agentId,omitempty"`
	// 上下架状态： ON上下、OFF下架
	ShelfStatus *DetailResponseBodyShelfStatusEnum `json:"shelfStatus,omitempty"`
	// 推理服务ID
	InferenceId *string `json:"inferenceId,omitempty"`
	// 智能体图标
	Icon *string `json:"icon,omitempty"`
	// 推荐问
	Questions *string `json:"questions,omitempty"`
	// 智能体模式
	AgentMode *DetailResponseBodyAgentModeEnum `json:"agentMode,omitempty"`
	// 推理模型名称
	ModelName *string `json:"modelName,omitempty"`
	// 智能体系统人设
	SystemPrompt *string `json:"systemPrompt,omitempty"`
	// 追问
	CloselyQuestions *string `json:"closelyQuestions,omitempty"`
	// 智能体名称
	Name *string `json:"name,omitempty"`
	// 智能体创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
	// 智能体分类
	Category *string `json:"category,omitempty"`
	// 是否开启推荐问
	SuggestQuestion *int32 `json:"suggestQuestion,omitempty"`
	// 智能体简介
	Introduction *string `json:"introduction,omitempty"`
	// 开场白
	Prologue *string `json:"prologue,omitempty"`
	// 工作流ID
	WorkflowId *string `json:"workflowId,omitempty"`
	// 发布状态：0已保存、1已发布
	Status *int32 `json:"status,omitempty"`
}

func (s DetailResponseBody) String() string {
	return utils.Beautify(s)
}

func (s DetailResponseBody) GoString() string {
	return s.String()
}

func (s DetailResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DetailResponseBody) SetUpdatedTime(v string) *DetailResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *DetailResponseBody) SetAgentId(v string) *DetailResponseBody {
	s.AgentId = &v
	return s
}

func (s *DetailResponseBody) SetShelfStatus(v DetailResponseBodyShelfStatusEnum) *DetailResponseBody {
	s.ShelfStatus = &v
	return s
}

func (s *DetailResponseBody) SetInferenceId(v string) *DetailResponseBody {
	s.InferenceId = &v
	return s
}

func (s *DetailResponseBody) SetIcon(v string) *DetailResponseBody {
	s.Icon = &v
	return s
}

func (s *DetailResponseBody) SetQuestions(v string) *DetailResponseBody {
	s.Questions = &v
	return s
}

func (s *DetailResponseBody) SetAgentMode(v DetailResponseBodyAgentModeEnum) *DetailResponseBody {
	s.AgentMode = &v
	return s
}

func (s *DetailResponseBody) SetModelName(v string) *DetailResponseBody {
	s.ModelName = &v
	return s
}

func (s *DetailResponseBody) SetSystemPrompt(v string) *DetailResponseBody {
	s.SystemPrompt = &v
	return s
}

func (s *DetailResponseBody) SetCloselyQuestions(v string) *DetailResponseBody {
	s.CloselyQuestions = &v
	return s
}

func (s *DetailResponseBody) SetName(v string) *DetailResponseBody {
	s.Name = &v
	return s
}

func (s *DetailResponseBody) SetCreatedTime(v string) *DetailResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DetailResponseBody) SetCategory(v string) *DetailResponseBody {
	s.Category = &v
	return s
}

func (s *DetailResponseBody) SetSuggestQuestion(v int32) *DetailResponseBody {
	s.SuggestQuestion = &v
	return s
}

func (s *DetailResponseBody) SetIntroduction(v string) *DetailResponseBody {
	s.Introduction = &v
	return s
}

func (s *DetailResponseBody) SetPrologue(v string) *DetailResponseBody {
	s.Prologue = &v
	return s
}

func (s *DetailResponseBody) SetWorkflowId(v string) *DetailResponseBody {
	s.WorkflowId = &v
	return s
}

func (s *DetailResponseBody) SetStatus(v int32) *DetailResponseBody {
	s.Status = &v
	return s
}

type DetailResponseBodyBuilder struct {
	s *DetailResponseBody
}

func NewDetailResponseBodyBuilder() *DetailResponseBodyBuilder {
	s := &DetailResponseBody{}
	b := &DetailResponseBodyBuilder{s: s}
	return b
}

func (b *DetailResponseBodyBuilder) UpdatedTime(v string) *DetailResponseBodyBuilder {
	b.s.UpdatedTime = &v
	return b
}

func (b *DetailResponseBodyBuilder) AgentId(v string) *DetailResponseBodyBuilder {
	b.s.AgentId = &v
	return b
}

func (b *DetailResponseBodyBuilder) ShelfStatus(v DetailResponseBodyShelfStatusEnum) *DetailResponseBodyBuilder {
	b.s.ShelfStatus = &v
	return b
}

func (b *DetailResponseBodyBuilder) InferenceId(v string) *DetailResponseBodyBuilder {
	b.s.InferenceId = &v
	return b
}

func (b *DetailResponseBodyBuilder) Icon(v string) *DetailResponseBodyBuilder {
	b.s.Icon = &v
	return b
}

func (b *DetailResponseBodyBuilder) Questions(v string) *DetailResponseBodyBuilder {
	b.s.Questions = &v
	return b
}

func (b *DetailResponseBodyBuilder) AgentMode(v DetailResponseBodyAgentModeEnum) *DetailResponseBodyBuilder {
	b.s.AgentMode = &v
	return b
}

func (b *DetailResponseBodyBuilder) ModelName(v string) *DetailResponseBodyBuilder {
	b.s.ModelName = &v
	return b
}

func (b *DetailResponseBodyBuilder) SystemPrompt(v string) *DetailResponseBodyBuilder {
	b.s.SystemPrompt = &v
	return b
}

func (b *DetailResponseBodyBuilder) CloselyQuestions(v string) *DetailResponseBodyBuilder {
	b.s.CloselyQuestions = &v
	return b
}

func (b *DetailResponseBodyBuilder) Name(v string) *DetailResponseBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *DetailResponseBodyBuilder) CreatedTime(v string) *DetailResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *DetailResponseBodyBuilder) Category(v string) *DetailResponseBodyBuilder {
	b.s.Category = &v
	return b
}

func (b *DetailResponseBodyBuilder) SuggestQuestion(v int32) *DetailResponseBodyBuilder {
	b.s.SuggestQuestion = &v
	return b
}

func (b *DetailResponseBodyBuilder) Introduction(v string) *DetailResponseBodyBuilder {
	b.s.Introduction = &v
	return b
}

func (b *DetailResponseBodyBuilder) Prologue(v string) *DetailResponseBodyBuilder {
	b.s.Prologue = &v
	return b
}

func (b *DetailResponseBodyBuilder) WorkflowId(v string) *DetailResponseBodyBuilder {
	b.s.WorkflowId = &v
	return b
}

func (b *DetailResponseBodyBuilder) Status(v int32) *DetailResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *DetailResponseBodyBuilder) Build() *DetailResponseBody {
	return b.s
}
