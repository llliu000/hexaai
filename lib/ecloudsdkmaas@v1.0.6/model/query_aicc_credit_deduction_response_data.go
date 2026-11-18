// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryAiccCreditDeductionResponseData struct {

	// 无视频输入tokens
	NoVideoInputTokens *int64 `json:"noVideoInputTokens,omitempty"`
	// 有视频输入tokens
	VideoInputTokens *int64 `json:"videoInputTokens,omitempty"`
	// 无视频输入1080p tokens
	NoVideoInput1080pTokens *int64 `json:"noVideoInput1080pTokens,omitempty"`
	// 扣减时间
	DeductTime *string `json:"deductTime,omitempty"`
	// 资费包扣减量（实际扣减资源点）
	CostAmount *float32 `json:"costAmount,omitempty"`
	// 输入tokens（任务实际消耗）
	InputTokens *int64 `json:"inputTokens,omitempty"`
	// 用户名（主账号取主账号用户名；RAM账号取RAM用户名，已删除则取ramId）
	UserName *string `json:"userName,omitempty"`
	// 输出tokens（任务实际消耗）
	OutputTokens *int64 `json:"outputTokens,omitempty"`
	// 总tokens（input + output）
	TotalTokens *int64 `json:"totalTokens,omitempty"`
	// 任务id
	TaskId *string `json:"taskId,omitempty"`
	// 有视频输入1080p tokens
	VideoInput1080pTokens *int64 `json:"videoInput1080pTokens,omitempty"`
}

func (s QueryAiccCreditDeductionResponseData) String() string {
	return utils.Beautify(s)
}

func (s QueryAiccCreditDeductionResponseData) GoString() string {
	return s.String()
}

func (s QueryAiccCreditDeductionResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryAiccCreditDeductionResponseData) SetNoVideoInputTokens(v int64) *QueryAiccCreditDeductionResponseData {
	s.NoVideoInputTokens = &v
	return s
}

func (s *QueryAiccCreditDeductionResponseData) SetVideoInputTokens(v int64) *QueryAiccCreditDeductionResponseData {
	s.VideoInputTokens = &v
	return s
}

func (s *QueryAiccCreditDeductionResponseData) SetNoVideoInput1080pTokens(v int64) *QueryAiccCreditDeductionResponseData {
	s.NoVideoInput1080pTokens = &v
	return s
}

func (s *QueryAiccCreditDeductionResponseData) SetDeductTime(v string) *QueryAiccCreditDeductionResponseData {
	s.DeductTime = &v
	return s
}

func (s *QueryAiccCreditDeductionResponseData) SetCostAmount(v float32) *QueryAiccCreditDeductionResponseData {
	s.CostAmount = &v
	return s
}

func (s *QueryAiccCreditDeductionResponseData) SetInputTokens(v int64) *QueryAiccCreditDeductionResponseData {
	s.InputTokens = &v
	return s
}

func (s *QueryAiccCreditDeductionResponseData) SetUserName(v string) *QueryAiccCreditDeductionResponseData {
	s.UserName = &v
	return s
}

func (s *QueryAiccCreditDeductionResponseData) SetOutputTokens(v int64) *QueryAiccCreditDeductionResponseData {
	s.OutputTokens = &v
	return s
}

func (s *QueryAiccCreditDeductionResponseData) SetTotalTokens(v int64) *QueryAiccCreditDeductionResponseData {
	s.TotalTokens = &v
	return s
}

func (s *QueryAiccCreditDeductionResponseData) SetTaskId(v string) *QueryAiccCreditDeductionResponseData {
	s.TaskId = &v
	return s
}

func (s *QueryAiccCreditDeductionResponseData) SetVideoInput1080pTokens(v int64) *QueryAiccCreditDeductionResponseData {
	s.VideoInput1080pTokens = &v
	return s
}

type QueryAiccCreditDeductionResponseDataBuilder struct {
	s *QueryAiccCreditDeductionResponseData
}

func NewQueryAiccCreditDeductionResponseDataBuilder() *QueryAiccCreditDeductionResponseDataBuilder {
	s := &QueryAiccCreditDeductionResponseData{}
	b := &QueryAiccCreditDeductionResponseDataBuilder{s: s}
	return b
}

func (b *QueryAiccCreditDeductionResponseDataBuilder) NoVideoInputTokens(v int64) *QueryAiccCreditDeductionResponseDataBuilder {
	b.s.NoVideoInputTokens = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseDataBuilder) VideoInputTokens(v int64) *QueryAiccCreditDeductionResponseDataBuilder {
	b.s.VideoInputTokens = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseDataBuilder) NoVideoInput1080pTokens(v int64) *QueryAiccCreditDeductionResponseDataBuilder {
	b.s.NoVideoInput1080pTokens = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseDataBuilder) DeductTime(v string) *QueryAiccCreditDeductionResponseDataBuilder {
	b.s.DeductTime = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseDataBuilder) CostAmount(v float32) *QueryAiccCreditDeductionResponseDataBuilder {
	b.s.CostAmount = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseDataBuilder) InputTokens(v int64) *QueryAiccCreditDeductionResponseDataBuilder {
	b.s.InputTokens = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseDataBuilder) UserName(v string) *QueryAiccCreditDeductionResponseDataBuilder {
	b.s.UserName = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseDataBuilder) OutputTokens(v int64) *QueryAiccCreditDeductionResponseDataBuilder {
	b.s.OutputTokens = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseDataBuilder) TotalTokens(v int64) *QueryAiccCreditDeductionResponseDataBuilder {
	b.s.TotalTokens = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseDataBuilder) TaskId(v string) *QueryAiccCreditDeductionResponseDataBuilder {
	b.s.TaskId = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseDataBuilder) VideoInput1080pTokens(v int64) *QueryAiccCreditDeductionResponseDataBuilder {
	b.s.VideoInput1080pTokens = &v
	return b
}

func (b *QueryAiccCreditDeductionResponseDataBuilder) Build() *QueryAiccCreditDeductionResponseData {
	return b.s
}
