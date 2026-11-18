// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryUsageStatsResponseData struct {

	// 图片生成总张数
	ImageCount *int64 `json:"imageCount,omitempty"`
	// 付费调用中非思考的输出 tokens（仅仅适用于需要按照思考计费的模型）
	PaidNoThinkTokens *int64 `json:"paidNoThinkTokens,omitempty"`
	// 失败调用次数（HTTP 4xx/5xx 或服务端内部错误）
	FailedCalls *int64 `json:"failedCalls,omitempty"`
	// 模型域类型，指示模型所属的能力域。TEXT=文本、IMAGE=图像、SPEECH=语音 等
	DomainType *string `json:"domainType,omitempty"`
	// 思考过程消耗的总 tokens（仅仅适用于需要按照思考计费的模型）
	ThinkTokens *int64 `json:"thinkTokens,omitempty"`
	// 成功调用次数（HTTP 2xx 或无服务端错误）
	SuccessfulCalls *int64 `json:"successfulCalls,omitempty"`
	// 推理名称（即模型名称）
	InferenceName *string `json:"inferenceName,omitempty"`
	// 非思考的输出总 tokens（仅仅适用于需要按照思考计费的模型）
	NoThinkTokens *int64 `json:"noThinkTokens,omitempty"`
	// 隐式缓存命中 tokens 总量（系统自动识别的重复内容命中缓存，客户无感知）
	ImplicitlyCacheTokens *int64 `json:"implicitlyCacheTokens,omitempty"`
	// 总调用次数（成功 + 失败）
	TotalCalls *int64 `json:"totalCalls,omitempty"`
	// 语音合成处理字符总数（万字符），计算公式: 总字符数 ÷ 10000
	Characters *float64 `json:"characters,omitempty"`
	// 显式缓存创建 tokens 总量（客户通过 API 参数主动写入缓存的 tokens）
	CacheCreationInputTokens *int64 `json:"cacheCreationInputTokens,omitempty"`
	// 是否付费订阅。null/0=免费；1=付费（需订阅或购买资源包）
	PayType *int32 `json:"payType,omitempty"`
	// 语音处理总时长（秒），适用于语音识别/合成等语音模型
	AudioSeconds *int64 `json:"audioSeconds,omitempty"`
	// 输入总 tokens（用户提问内容消耗的 tokens）
	PromptTokens *int64 `json:"promptTokens,omitempty"`
	// 付费调用中的输出 tokens
	PaidCompletionTokens *int64 `json:"paidCompletionTokens,omitempty"`
	// 模型供应商名称（如 移动云、阿里云 等）
	ProviderName *string `json:"providerName,omitempty"`
	// 输出总 tokens（模型回复内容消耗的 tokens，不含思考部分）
	CompletionTokens *int64 `json:"completionTokens,omitempty"`
	// 视频生成总时长（秒）
	VideoSeconds *int64 `json:"videoSeconds,omitempty"`
	// 付费调用中的输入 tokens
	PaidPromptTokens *int64 `json:"paidPromptTokens,omitempty"`
	// 调用消耗的总 tokens（输入 + 输出 + 思考），含免费和付费调用
	TotalTokens *int64 `json:"totalTokens,omitempty"`
	// 付费调用消耗的总 tokens（仅含付费订阅/资源包对应的调用）
	PaidTotalTokens *int64 `json:"paidTotalTokens,omitempty"`
	// 显式缓存命中 tokens 总量（客户通过 API 参数主动指定的缓存内容被命中）
	ExplicitCacheTokens *int64 `json:"explicitCacheTokens,omitempty"`
	// 付费调用中的输出 tokens（与 paidCompletionTokens 同义）
	PaidCompletions *int64 `json:"paidCompletions,omitempty"`
	// 付费调用中思考过程消耗的tokens（仅仅适用于需要按照思考计费的模型）
	PaidThinkTokens *int64 `json:"paidThinkTokens,omitempty"`
}

func (s QueryUsageStatsResponseData) String() string {
	return utils.Beautify(s)
}

func (s QueryUsageStatsResponseData) GoString() string {
	return s.String()
}

func (s QueryUsageStatsResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryUsageStatsResponseData) SetImageCount(v int64) *QueryUsageStatsResponseData {
	s.ImageCount = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetPaidNoThinkTokens(v int64) *QueryUsageStatsResponseData {
	s.PaidNoThinkTokens = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetFailedCalls(v int64) *QueryUsageStatsResponseData {
	s.FailedCalls = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetDomainType(v string) *QueryUsageStatsResponseData {
	s.DomainType = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetThinkTokens(v int64) *QueryUsageStatsResponseData {
	s.ThinkTokens = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetSuccessfulCalls(v int64) *QueryUsageStatsResponseData {
	s.SuccessfulCalls = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetInferenceName(v string) *QueryUsageStatsResponseData {
	s.InferenceName = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetNoThinkTokens(v int64) *QueryUsageStatsResponseData {
	s.NoThinkTokens = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetImplicitlyCacheTokens(v int64) *QueryUsageStatsResponseData {
	s.ImplicitlyCacheTokens = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetTotalCalls(v int64) *QueryUsageStatsResponseData {
	s.TotalCalls = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetCharacters(v float64) *QueryUsageStatsResponseData {
	s.Characters = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetCacheCreationInputTokens(v int64) *QueryUsageStatsResponseData {
	s.CacheCreationInputTokens = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetPayType(v int32) *QueryUsageStatsResponseData {
	s.PayType = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetAudioSeconds(v int64) *QueryUsageStatsResponseData {
	s.AudioSeconds = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetPromptTokens(v int64) *QueryUsageStatsResponseData {
	s.PromptTokens = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetPaidCompletionTokens(v int64) *QueryUsageStatsResponseData {
	s.PaidCompletionTokens = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetProviderName(v string) *QueryUsageStatsResponseData {
	s.ProviderName = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetCompletionTokens(v int64) *QueryUsageStatsResponseData {
	s.CompletionTokens = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetVideoSeconds(v int64) *QueryUsageStatsResponseData {
	s.VideoSeconds = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetPaidPromptTokens(v int64) *QueryUsageStatsResponseData {
	s.PaidPromptTokens = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetTotalTokens(v int64) *QueryUsageStatsResponseData {
	s.TotalTokens = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetPaidTotalTokens(v int64) *QueryUsageStatsResponseData {
	s.PaidTotalTokens = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetExplicitCacheTokens(v int64) *QueryUsageStatsResponseData {
	s.ExplicitCacheTokens = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetPaidCompletions(v int64) *QueryUsageStatsResponseData {
	s.PaidCompletions = &v
	return s
}

func (s *QueryUsageStatsResponseData) SetPaidThinkTokens(v int64) *QueryUsageStatsResponseData {
	s.PaidThinkTokens = &v
	return s
}

type QueryUsageStatsResponseDataBuilder struct {
	s *QueryUsageStatsResponseData
}

func NewQueryUsageStatsResponseDataBuilder() *QueryUsageStatsResponseDataBuilder {
	s := &QueryUsageStatsResponseData{}
	b := &QueryUsageStatsResponseDataBuilder{s: s}
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) ImageCount(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.ImageCount = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) PaidNoThinkTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.PaidNoThinkTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) FailedCalls(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.FailedCalls = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) DomainType(v string) *QueryUsageStatsResponseDataBuilder {
	b.s.DomainType = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) ThinkTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.ThinkTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) SuccessfulCalls(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.SuccessfulCalls = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) InferenceName(v string) *QueryUsageStatsResponseDataBuilder {
	b.s.InferenceName = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) NoThinkTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.NoThinkTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) ImplicitlyCacheTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.ImplicitlyCacheTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) TotalCalls(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.TotalCalls = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) Characters(v float64) *QueryUsageStatsResponseDataBuilder {
	b.s.Characters = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) CacheCreationInputTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.CacheCreationInputTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) PayType(v int32) *QueryUsageStatsResponseDataBuilder {
	b.s.PayType = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) AudioSeconds(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.AudioSeconds = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) PromptTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.PromptTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) PaidCompletionTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.PaidCompletionTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) ProviderName(v string) *QueryUsageStatsResponseDataBuilder {
	b.s.ProviderName = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) CompletionTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.CompletionTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) VideoSeconds(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.VideoSeconds = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) PaidPromptTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.PaidPromptTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) TotalTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.TotalTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) PaidTotalTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.PaidTotalTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) ExplicitCacheTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.ExplicitCacheTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) PaidCompletions(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.PaidCompletions = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) PaidThinkTokens(v int64) *QueryUsageStatsResponseDataBuilder {
	b.s.PaidThinkTokens = &v
	return b
}

func (b *QueryUsageStatsResponseDataBuilder) Build() *QueryUsageStatsResponseData {
	return b.s
}
