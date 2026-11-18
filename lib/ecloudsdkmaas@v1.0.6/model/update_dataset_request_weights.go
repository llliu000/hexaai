// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateDatasetRequestWeights struct {

	// 语义占比，只有当search_method=HYBRID_WEIGHTED_SCORE时才需要设置，并且语义占比+关键词占比=1
	VectorWeight *float32 `json:"vector_weight,omitempty"`
	// 关键词占比,只有当search_method=HYBRID_WEIGHTED_SCORE时才需要设置，并且语义占比+关键词占比=1
	KeywordWeight *float32 `json:"keyword_weight,omitempty"`
}

func (s UpdateDatasetRequestWeights) String() string {
	return utils.Beautify(s)
}

func (s UpdateDatasetRequestWeights) GoString() string {
	return s.String()
}

func (s UpdateDatasetRequestWeights) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateDatasetRequestWeights) SetVectorWeight(v float32) *UpdateDatasetRequestWeights {
	s.VectorWeight = &v
	return s
}

func (s *UpdateDatasetRequestWeights) SetKeywordWeight(v float32) *UpdateDatasetRequestWeights {
	s.KeywordWeight = &v
	return s
}

type UpdateDatasetRequestWeightsBuilder struct {
	s *UpdateDatasetRequestWeights
}

func NewUpdateDatasetRequestWeightsBuilder() *UpdateDatasetRequestWeightsBuilder {
	s := &UpdateDatasetRequestWeights{}
	b := &UpdateDatasetRequestWeightsBuilder{s: s}
	return b
}

func (b *UpdateDatasetRequestWeightsBuilder) VectorWeight(v float32) *UpdateDatasetRequestWeightsBuilder {
	b.s.VectorWeight = &v
	return b
}

func (b *UpdateDatasetRequestWeightsBuilder) KeywordWeight(v float32) *UpdateDatasetRequestWeightsBuilder {
	b.s.KeywordWeight = &v
	return b
}

func (b *UpdateDatasetRequestWeightsBuilder) Build() *UpdateDatasetRequestWeights {
	return b.s
}
