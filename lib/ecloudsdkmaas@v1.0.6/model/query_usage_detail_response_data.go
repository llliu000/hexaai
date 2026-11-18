// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryUsageDetailResponseData struct {

	// 图片生成张数。图像类模型有值；非图像模型为 null
	ImageCount *int32 `json:"imageCount,omitempty"`
	// 视频生成时长（秒）。视频类模型有值；非视频模型为 null
	VideoSeconds *int32 `json:"videoSeconds,omitempty"`
	// 思考过程 tokens（仅仅适用于需要按照思考计费的模型）
	ThinkTokens *int32 `json:"thinkTokens,omitempty"`
	// 推理名称（模型名称）
	InferenceName *string `json:"inferenceName,omitempty"`
	// API Key 名称
	ApiKeyName *string `json:"apiKeyName,omitempty"`
	// 非思考的输出 tokens（仅仅适用于需要按照思考计费的模型）
	NoThinkTokens *int32 `json:"noThinkTokens,omitempty"`
	// 隐式缓存命中 tokens（系统自动识别重复内容命中缓存，客户无感知）
	ImplicitlyCacheTokens *int32 `json:"implicitlyCacheTokens,omitempty"`
	// 本次调用消耗的总 tokens（输入 + 输出 + 思考）
	TotalTokens *int32 `json:"totalTokens,omitempty"`
	// RAM 用户名称（实际发起调用的子账号名称）
	RamUserName *string `json:"ramUserName,omitempty"`
	// 语音合成处理的字符数。TTS 语音合成模型有值；非语音合成模型为 null
	Characters *int32 `json:"characters,omitempty"`
	// 显式缓存创建 tokens（客户通过 API 参数新建缓存条目消耗的 tokens）
	CacheCreationInputTokens *int32 `json:"cacheCreationInputTokens,omitempty"`
	// 请求 ID（每次 API 调用的唯一标识）
	RequestId *string `json:"requestId,omitempty"`
	// 语音处理时长（秒）。语音类模型（ASR/TTS）有值；非语音模型为 null
	AudioSeconds *int32 `json:"audioSeconds,omitempty"`
	// 显式缓存命中 tokens（客户通过 API 参数主动指定的缓存内容被命中）
	ExplicitCacheTokens *int32 `json:"explicitCacheTokens,omitempty"`
	// 输入 tokens（用户提问及上下文消耗的 tokens）
	PromptTokens *int32 `json:"promptTokens,omitempty"`
	// 调用开始时间（服务端收到请求的时间），格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"startTime,omitempty"`
	// 首 token 延迟（毫秒），从收到请求到返回第一个 token 的耗时
	FirstTokenCostTime *int32 `json:"firstTokenCostTime,omitempty"`
	// 调用结果。true=成功；false=失败（HTTP 非 2xx 或服务端内部异常）
	CallResult *bool `json:"callResult,omitempty"`
	// 输出 tokens（模型回复文本消耗的 tokens，不含思考过程）
	CompletionTokens *int32 `json:"completionTokens,omitempty"`
	// HTTP 状态码。200=成功；400=请求参数错误；429=频率限制；500=服务端错误
	StatusCode *int32 `json:"statusCode,omitempty"`
}

func (s QueryUsageDetailResponseData) String() string {
	return utils.Beautify(s)
}

func (s QueryUsageDetailResponseData) GoString() string {
	return s.String()
}

func (s QueryUsageDetailResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryUsageDetailResponseData) SetImageCount(v int32) *QueryUsageDetailResponseData {
	s.ImageCount = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetVideoSeconds(v int32) *QueryUsageDetailResponseData {
	s.VideoSeconds = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetThinkTokens(v int32) *QueryUsageDetailResponseData {
	s.ThinkTokens = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetInferenceName(v string) *QueryUsageDetailResponseData {
	s.InferenceName = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetApiKeyName(v string) *QueryUsageDetailResponseData {
	s.ApiKeyName = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetNoThinkTokens(v int32) *QueryUsageDetailResponseData {
	s.NoThinkTokens = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetImplicitlyCacheTokens(v int32) *QueryUsageDetailResponseData {
	s.ImplicitlyCacheTokens = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetTotalTokens(v int32) *QueryUsageDetailResponseData {
	s.TotalTokens = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetRamUserName(v string) *QueryUsageDetailResponseData {
	s.RamUserName = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetCharacters(v int32) *QueryUsageDetailResponseData {
	s.Characters = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetCacheCreationInputTokens(v int32) *QueryUsageDetailResponseData {
	s.CacheCreationInputTokens = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetRequestId(v string) *QueryUsageDetailResponseData {
	s.RequestId = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetAudioSeconds(v int32) *QueryUsageDetailResponseData {
	s.AudioSeconds = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetExplicitCacheTokens(v int32) *QueryUsageDetailResponseData {
	s.ExplicitCacheTokens = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetPromptTokens(v int32) *QueryUsageDetailResponseData {
	s.PromptTokens = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetStartTime(v string) *QueryUsageDetailResponseData {
	s.StartTime = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetFirstTokenCostTime(v int32) *QueryUsageDetailResponseData {
	s.FirstTokenCostTime = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetCallResult(v bool) *QueryUsageDetailResponseData {
	s.CallResult = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetCompletionTokens(v int32) *QueryUsageDetailResponseData {
	s.CompletionTokens = &v
	return s
}

func (s *QueryUsageDetailResponseData) SetStatusCode(v int32) *QueryUsageDetailResponseData {
	s.StatusCode = &v
	return s
}

type QueryUsageDetailResponseDataBuilder struct {
	s *QueryUsageDetailResponseData
}

func NewQueryUsageDetailResponseDataBuilder() *QueryUsageDetailResponseDataBuilder {
	s := &QueryUsageDetailResponseData{}
	b := &QueryUsageDetailResponseDataBuilder{s: s}
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) ImageCount(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.ImageCount = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) VideoSeconds(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.VideoSeconds = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) ThinkTokens(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.ThinkTokens = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) InferenceName(v string) *QueryUsageDetailResponseDataBuilder {
	b.s.InferenceName = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) ApiKeyName(v string) *QueryUsageDetailResponseDataBuilder {
	b.s.ApiKeyName = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) NoThinkTokens(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.NoThinkTokens = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) ImplicitlyCacheTokens(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.ImplicitlyCacheTokens = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) TotalTokens(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.TotalTokens = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) RamUserName(v string) *QueryUsageDetailResponseDataBuilder {
	b.s.RamUserName = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) Characters(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.Characters = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) CacheCreationInputTokens(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.CacheCreationInputTokens = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) RequestId(v string) *QueryUsageDetailResponseDataBuilder {
	b.s.RequestId = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) AudioSeconds(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.AudioSeconds = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) ExplicitCacheTokens(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.ExplicitCacheTokens = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) PromptTokens(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.PromptTokens = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) StartTime(v string) *QueryUsageDetailResponseDataBuilder {
	b.s.StartTime = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) FirstTokenCostTime(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.FirstTokenCostTime = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) CallResult(v bool) *QueryUsageDetailResponseDataBuilder {
	b.s.CallResult = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) CompletionTokens(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.CompletionTokens = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) StatusCode(v int32) *QueryUsageDetailResponseDataBuilder {
	b.s.StatusCode = &v
	return b
}

func (b *QueryUsageDetailResponseDataBuilder) Build() *QueryUsageDetailResponseData {
	return b.s
}
