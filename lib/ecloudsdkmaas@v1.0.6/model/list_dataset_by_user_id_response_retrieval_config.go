// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListDatasetByUserIdResponseRetrievalConfigSearchMethodEnum string

// List of SearchMethod
const (
	ListDatasetByUserIdResponseRetrievalConfigSearchMethodEnumSemantic             ListDatasetByUserIdResponseRetrievalConfigSearchMethodEnum = "SEMANTIC"
	ListDatasetByUserIdResponseRetrievalConfigSearchMethodEnumFullText             ListDatasetByUserIdResponseRetrievalConfigSearchMethodEnum = "FULL_TEXT"
	ListDatasetByUserIdResponseRetrievalConfigSearchMethodEnumHybridRerankingModel ListDatasetByUserIdResponseRetrievalConfigSearchMethodEnum = "HYBRID_RERANKING_MODEL"
	ListDatasetByUserIdResponseRetrievalConfigSearchMethodEnumHybridWeightedScore  ListDatasetByUserIdResponseRetrievalConfigSearchMethodEnum = "HYBRID_WEIGHTED_SCORE"
)

type ListDatasetByUserIdResponseRetrievalConfig struct {

	// Score阈值 ，用于设置文本片段筛选的相似度阈值，只召回超过设置分数的文本片段，数值越高说明对于文本与问题要求的相似度越高，预期被召回的文本数量也越少
	ScoreThreshold *float32 `json:"score_threshold,omitempty"`
	// 是否设置Score阈值
	ScoreThresholdEnabled *bool `json:"score_threshold_enabled,omitempty"`
	// 是否开启排序模型
	RerankingEnabled *bool `json:"reranking_enabled,omitempty"`
	// 召回数量 ,用于筛选与用户问题相似度最高的文本片段。系统同时会根据选用模型上下文窗口大小动态调整片段数量，数值越高，预期被召回的文本分段数量越多。
	TopK *int32 `json:"top_k,omitempty"`
	// 权重设置
	Weights *ListDatasetByUserIdResponseWeights `json:"weights,omitempty"`
	// 检索方式： HYBRID_WEIGHTED_SCORE: 混合检索; HYBRID_RERANKING_MODEL: 混合检索 SEMANTIC：向量检索; FULL_TEXT：全文检索 HYBRID_WEIGHTED_SCORE和HYBRID_RERANKING_MODEL的区别就是是否要设置权重即weights的值
	SearchMethod *ListDatasetByUserIdResponseRetrievalConfigSearchMethodEnum `json:"search_method,omitempty"`
}

func (s ListDatasetByUserIdResponseRetrievalConfig) String() string {
	return utils.Beautify(s)
}

func (s ListDatasetByUserIdResponseRetrievalConfig) GoString() string {
	return s.String()
}

func (s ListDatasetByUserIdResponseRetrievalConfig) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListDatasetByUserIdResponseRetrievalConfig) SetScoreThreshold(v float32) *ListDatasetByUserIdResponseRetrievalConfig {
	s.ScoreThreshold = &v
	return s
}

func (s *ListDatasetByUserIdResponseRetrievalConfig) SetScoreThresholdEnabled(v bool) *ListDatasetByUserIdResponseRetrievalConfig {
	s.ScoreThresholdEnabled = &v
	return s
}

func (s *ListDatasetByUserIdResponseRetrievalConfig) SetRerankingEnabled(v bool) *ListDatasetByUserIdResponseRetrievalConfig {
	s.RerankingEnabled = &v
	return s
}

func (s *ListDatasetByUserIdResponseRetrievalConfig) SetTopK(v int32) *ListDatasetByUserIdResponseRetrievalConfig {
	s.TopK = &v
	return s
}

func (s *ListDatasetByUserIdResponseRetrievalConfig) SetWeights(v *ListDatasetByUserIdResponseWeights) *ListDatasetByUserIdResponseRetrievalConfig {
	s.Weights = v
	return s
}

func (s *ListDatasetByUserIdResponseRetrievalConfig) SetSearchMethod(v ListDatasetByUserIdResponseRetrievalConfigSearchMethodEnum) *ListDatasetByUserIdResponseRetrievalConfig {
	s.SearchMethod = &v
	return s
}

type ListDatasetByUserIdResponseRetrievalConfigBuilder struct {
	s *ListDatasetByUserIdResponseRetrievalConfig
}

func NewListDatasetByUserIdResponseRetrievalConfigBuilder() *ListDatasetByUserIdResponseRetrievalConfigBuilder {
	s := &ListDatasetByUserIdResponseRetrievalConfig{}
	b := &ListDatasetByUserIdResponseRetrievalConfigBuilder{s: s}
	return b
}

func (b *ListDatasetByUserIdResponseRetrievalConfigBuilder) ScoreThreshold(v float32) *ListDatasetByUserIdResponseRetrievalConfigBuilder {
	b.s.ScoreThreshold = &v
	return b
}

func (b *ListDatasetByUserIdResponseRetrievalConfigBuilder) ScoreThresholdEnabled(v bool) *ListDatasetByUserIdResponseRetrievalConfigBuilder {
	b.s.ScoreThresholdEnabled = &v
	return b
}

func (b *ListDatasetByUserIdResponseRetrievalConfigBuilder) RerankingEnabled(v bool) *ListDatasetByUserIdResponseRetrievalConfigBuilder {
	b.s.RerankingEnabled = &v
	return b
}

func (b *ListDatasetByUserIdResponseRetrievalConfigBuilder) TopK(v int32) *ListDatasetByUserIdResponseRetrievalConfigBuilder {
	b.s.TopK = &v
	return b
}

func (b *ListDatasetByUserIdResponseRetrievalConfigBuilder) Weights(v *ListDatasetByUserIdResponseWeights) *ListDatasetByUserIdResponseRetrievalConfigBuilder {
	b.s.Weights = v
	return b
}

func (b *ListDatasetByUserIdResponseRetrievalConfigBuilder) SearchMethod(v ListDatasetByUserIdResponseRetrievalConfigSearchMethodEnum) *ListDatasetByUserIdResponseRetrievalConfigBuilder {
	b.s.SearchMethod = &v
	return b
}

func (b *ListDatasetByUserIdResponseRetrievalConfigBuilder) Build() *ListDatasetByUserIdResponseRetrievalConfig {
	return b.s
}
