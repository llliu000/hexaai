// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateDatasetRequestRetrievalConfigSearchMethodEnum string

// List of SearchMethod
const (
	CreateDatasetRequestRetrievalConfigSearchMethodEnumSemantic             CreateDatasetRequestRetrievalConfigSearchMethodEnum = "SEMANTIC"
	CreateDatasetRequestRetrievalConfigSearchMethodEnumFullText             CreateDatasetRequestRetrievalConfigSearchMethodEnum = "FULL_TEXT"
	CreateDatasetRequestRetrievalConfigSearchMethodEnumHybridRerankingModel CreateDatasetRequestRetrievalConfigSearchMethodEnum = "HYBRID_RERANKING_MODEL"
	CreateDatasetRequestRetrievalConfigSearchMethodEnumHybridWeightedScore  CreateDatasetRequestRetrievalConfigSearchMethodEnum = "HYBRID_WEIGHTED_SCORE"
)

type CreateDatasetRequestRetrievalConfig struct {

	// Score阈值 ，用于设置文本片段筛选的相似度阈值，只召回超过设置分数的文本片段，数值越高说明对于文本与问题要求的相似度越高，预期被召回的文本数量也越少,默认0.5，只有当score_threshold_enabled=true时，该配置才生效
	ScoreThreshold *float32 `json:"score_threshold,omitempty"`
	// 默认都开启，是否开启排序模型
	RerankingEnabled *bool `json:"reranking_enabled,omitempty"`
	// 是否设置Score阈值
	ScoreThresholdEnabled *bool `json:"score_threshold_enabled,omitempty"`
	// 召回数量 ,用于筛选与用户问题相似度最高的文本片段。系统同时会根据选用模型上下文窗口大小动态调整片段数量，数值越高，预期被召回的文本分段数量越多。
	TopK *int32 `json:"top_k,omitempty"`
	// 权重设置，search_method=HYBRID_WEIGHTED_SCORE时才需要设置
	Weights *CreateDatasetRequestWeights `json:"weights,omitempty"`
	// 检索方式： HYBRID_WEIGHTED_SCORE: 混合检索; HYBRID_RERANKING_MODEL: 混合检索 SEMANTIC：向量检索; FULL_TEXT：全文检索 HYBRID_WEIGHTED_SCORE和HYBRID_RERANKING_MODEL的区别就是是否要设置权重即weights的值
	SearchMethod *CreateDatasetRequestRetrievalConfigSearchMethodEnum `json:"search_method,omitempty"`
}

func (s CreateDatasetRequestRetrievalConfig) String() string {
	return utils.Beautify(s)
}

func (s CreateDatasetRequestRetrievalConfig) GoString() string {
	return s.String()
}

func (s CreateDatasetRequestRetrievalConfig) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateDatasetRequestRetrievalConfig) SetScoreThreshold(v float32) *CreateDatasetRequestRetrievalConfig {
	s.ScoreThreshold = &v
	return s
}

func (s *CreateDatasetRequestRetrievalConfig) SetRerankingEnabled(v bool) *CreateDatasetRequestRetrievalConfig {
	s.RerankingEnabled = &v
	return s
}

func (s *CreateDatasetRequestRetrievalConfig) SetScoreThresholdEnabled(v bool) *CreateDatasetRequestRetrievalConfig {
	s.ScoreThresholdEnabled = &v
	return s
}

func (s *CreateDatasetRequestRetrievalConfig) SetTopK(v int32) *CreateDatasetRequestRetrievalConfig {
	s.TopK = &v
	return s
}

func (s *CreateDatasetRequestRetrievalConfig) SetWeights(v *CreateDatasetRequestWeights) *CreateDatasetRequestRetrievalConfig {
	s.Weights = v
	return s
}

func (s *CreateDatasetRequestRetrievalConfig) SetSearchMethod(v CreateDatasetRequestRetrievalConfigSearchMethodEnum) *CreateDatasetRequestRetrievalConfig {
	s.SearchMethod = &v
	return s
}

type CreateDatasetRequestRetrievalConfigBuilder struct {
	s *CreateDatasetRequestRetrievalConfig
}

func NewCreateDatasetRequestRetrievalConfigBuilder() *CreateDatasetRequestRetrievalConfigBuilder {
	s := &CreateDatasetRequestRetrievalConfig{}
	b := &CreateDatasetRequestRetrievalConfigBuilder{s: s}
	return b
}

func (b *CreateDatasetRequestRetrievalConfigBuilder) ScoreThreshold(v float32) *CreateDatasetRequestRetrievalConfigBuilder {
	b.s.ScoreThreshold = &v
	return b
}

func (b *CreateDatasetRequestRetrievalConfigBuilder) RerankingEnabled(v bool) *CreateDatasetRequestRetrievalConfigBuilder {
	b.s.RerankingEnabled = &v
	return b
}

func (b *CreateDatasetRequestRetrievalConfigBuilder) ScoreThresholdEnabled(v bool) *CreateDatasetRequestRetrievalConfigBuilder {
	b.s.ScoreThresholdEnabled = &v
	return b
}

func (b *CreateDatasetRequestRetrievalConfigBuilder) TopK(v int32) *CreateDatasetRequestRetrievalConfigBuilder {
	b.s.TopK = &v
	return b
}

func (b *CreateDatasetRequestRetrievalConfigBuilder) Weights(v *CreateDatasetRequestWeights) *CreateDatasetRequestRetrievalConfigBuilder {
	b.s.Weights = v
	return b
}

func (b *CreateDatasetRequestRetrievalConfigBuilder) SearchMethod(v CreateDatasetRequestRetrievalConfigSearchMethodEnum) *CreateDatasetRequestRetrievalConfigBuilder {
	b.s.SearchMethod = &v
	return b
}

func (b *CreateDatasetRequestRetrievalConfigBuilder) Build() *CreateDatasetRequestRetrievalConfig {
	return b.s
}
