// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateDatasetRequestRetrievalConfigSearchMethodEnum string

// List of SearchMethod
const (
	UpdateDatasetRequestRetrievalConfigSearchMethodEnumSemantic             UpdateDatasetRequestRetrievalConfigSearchMethodEnum = "SEMANTIC"
	UpdateDatasetRequestRetrievalConfigSearchMethodEnumFullText             UpdateDatasetRequestRetrievalConfigSearchMethodEnum = "FULL_TEXT"
	UpdateDatasetRequestRetrievalConfigSearchMethodEnumHybridRerankingModel UpdateDatasetRequestRetrievalConfigSearchMethodEnum = "HYBRID_RERANKING_MODEL"
	UpdateDatasetRequestRetrievalConfigSearchMethodEnumHybridWeightedScore  UpdateDatasetRequestRetrievalConfigSearchMethodEnum = "HYBRID_WEIGHTED_SCORE"
)

type UpdateDatasetRequestRetrievalConfig struct {

	// Score阈值 ，用于设置文本片段筛选的相似度阈值，只召回超过设置分数的文本片段，数值越高说明对于文本与问题要求的相似度越高，预期被召回的文本数量也越少,默认0.5，只有当score_threshold_enabled=true时，该配置才生效
	ScoreThreshold *float32 `json:"score_threshold,omitempty"`
	// 默认都开启，是否开启排序模型
	RerankingEnabled *bool `json:"reranking_enabled,omitempty"`
	// 是否设置Score阈值
	ScoreThresholdEnabled *bool `json:"score_threshold_enabled,omitempty"`
	// 召回数量 ,用于筛选与用户问题相似度最高的文本片段。系统同时会根据选用模型上下文窗口大小动态调整片段数量，数值越高，预期被召回的文本分段数量越多。
	TopK *int32 `json:"top_k,omitempty"`
	// 权重设置，search_method=HYBRID_WEIGHTED_SCORE时才需要设置
	Weights *UpdateDatasetRequestWeights `json:"weights,omitempty"`
	// 检索方式： HYBRID_WEIGHTED_SCORE: 混合检索(带权重);  HYBRID_RERANKING_MODEL: 混合检索  SEMANTIC：向量检索;  FULL_TEXT：全文检索  HYBRID_WEIGHTED_SCORE和HYBRID_RERANKING_MODEL的区别就是是否要设置权重即weights的值
	SearchMethod *UpdateDatasetRequestRetrievalConfigSearchMethodEnum `json:"search_method,omitempty"`
}

func (s UpdateDatasetRequestRetrievalConfig) String() string {
	return utils.Beautify(s)
}

func (s UpdateDatasetRequestRetrievalConfig) GoString() string {
	return s.String()
}

func (s UpdateDatasetRequestRetrievalConfig) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateDatasetRequestRetrievalConfig) SetScoreThreshold(v float32) *UpdateDatasetRequestRetrievalConfig {
	s.ScoreThreshold = &v
	return s
}

func (s *UpdateDatasetRequestRetrievalConfig) SetRerankingEnabled(v bool) *UpdateDatasetRequestRetrievalConfig {
	s.RerankingEnabled = &v
	return s
}

func (s *UpdateDatasetRequestRetrievalConfig) SetScoreThresholdEnabled(v bool) *UpdateDatasetRequestRetrievalConfig {
	s.ScoreThresholdEnabled = &v
	return s
}

func (s *UpdateDatasetRequestRetrievalConfig) SetTopK(v int32) *UpdateDatasetRequestRetrievalConfig {
	s.TopK = &v
	return s
}

func (s *UpdateDatasetRequestRetrievalConfig) SetWeights(v *UpdateDatasetRequestWeights) *UpdateDatasetRequestRetrievalConfig {
	s.Weights = v
	return s
}

func (s *UpdateDatasetRequestRetrievalConfig) SetSearchMethod(v UpdateDatasetRequestRetrievalConfigSearchMethodEnum) *UpdateDatasetRequestRetrievalConfig {
	s.SearchMethod = &v
	return s
}

type UpdateDatasetRequestRetrievalConfigBuilder struct {
	s *UpdateDatasetRequestRetrievalConfig
}

func NewUpdateDatasetRequestRetrievalConfigBuilder() *UpdateDatasetRequestRetrievalConfigBuilder {
	s := &UpdateDatasetRequestRetrievalConfig{}
	b := &UpdateDatasetRequestRetrievalConfigBuilder{s: s}
	return b
}

func (b *UpdateDatasetRequestRetrievalConfigBuilder) ScoreThreshold(v float32) *UpdateDatasetRequestRetrievalConfigBuilder {
	b.s.ScoreThreshold = &v
	return b
}

func (b *UpdateDatasetRequestRetrievalConfigBuilder) RerankingEnabled(v bool) *UpdateDatasetRequestRetrievalConfigBuilder {
	b.s.RerankingEnabled = &v
	return b
}

func (b *UpdateDatasetRequestRetrievalConfigBuilder) ScoreThresholdEnabled(v bool) *UpdateDatasetRequestRetrievalConfigBuilder {
	b.s.ScoreThresholdEnabled = &v
	return b
}

func (b *UpdateDatasetRequestRetrievalConfigBuilder) TopK(v int32) *UpdateDatasetRequestRetrievalConfigBuilder {
	b.s.TopK = &v
	return b
}

func (b *UpdateDatasetRequestRetrievalConfigBuilder) Weights(v *UpdateDatasetRequestWeights) *UpdateDatasetRequestRetrievalConfigBuilder {
	b.s.Weights = v
	return b
}

func (b *UpdateDatasetRequestRetrievalConfigBuilder) SearchMethod(v UpdateDatasetRequestRetrievalConfigSearchMethodEnum) *UpdateDatasetRequestRetrievalConfigBuilder {
	b.s.SearchMethod = &v
	return b
}

func (b *UpdateDatasetRequestRetrievalConfigBuilder) Build() *UpdateDatasetRequestRetrievalConfig {
	return b.s
}
