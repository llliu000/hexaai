// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateDatasetRequestWeights struct {

	// 语义占比，只有当search_method=HYBRID_WEIGHTED_SCORE时才需要设置，并且语义占比+关键词占比=1
	VectorWeight *float32 `json:"vector_weight,omitempty"`
	// 关键词占比，只有当search_method=HYBRID_WEIGHTED_SCORE时才需要设置，并且语义占比+关键词占比=1
	KeywordWeight *float32 `json:"keyword_weight,omitempty"`
}

func (s CreateDatasetRequestWeights) String() string {
	return utils.Beautify(s)
}

func (s CreateDatasetRequestWeights) GoString() string {
	return s.String()
}

func (s CreateDatasetRequestWeights) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateDatasetRequestWeights) SetVectorWeight(v float32) *CreateDatasetRequestWeights {
	s.VectorWeight = &v
	return s
}

func (s *CreateDatasetRequestWeights) SetKeywordWeight(v float32) *CreateDatasetRequestWeights {
	s.KeywordWeight = &v
	return s
}

type CreateDatasetRequestWeightsBuilder struct {
	s *CreateDatasetRequestWeights
}

func NewCreateDatasetRequestWeightsBuilder() *CreateDatasetRequestWeightsBuilder {
	s := &CreateDatasetRequestWeights{}
	b := &CreateDatasetRequestWeightsBuilder{s: s}
	return b
}

func (b *CreateDatasetRequestWeightsBuilder) VectorWeight(v float32) *CreateDatasetRequestWeightsBuilder {
	b.s.VectorWeight = &v
	return b
}

func (b *CreateDatasetRequestWeightsBuilder) KeywordWeight(v float32) *CreateDatasetRequestWeightsBuilder {
	b.s.KeywordWeight = &v
	return b
}

func (b *CreateDatasetRequestWeightsBuilder) Build() *CreateDatasetRequestWeights {
	return b.s
}
