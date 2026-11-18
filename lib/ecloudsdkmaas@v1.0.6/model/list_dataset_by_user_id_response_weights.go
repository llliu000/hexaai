// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListDatasetByUserIdResponseWeights struct {

	// 语义占比
	VectorWeight *float32 `json:"vector_weight,omitempty"`
	// 关键词占比
	KeywordWeight *float32 `json:"keyword_weight,omitempty"`
}

func (s ListDatasetByUserIdResponseWeights) String() string {
	return utils.Beautify(s)
}

func (s ListDatasetByUserIdResponseWeights) GoString() string {
	return s.String()
}

func (s ListDatasetByUserIdResponseWeights) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListDatasetByUserIdResponseWeights) SetVectorWeight(v float32) *ListDatasetByUserIdResponseWeights {
	s.VectorWeight = &v
	return s
}

func (s *ListDatasetByUserIdResponseWeights) SetKeywordWeight(v float32) *ListDatasetByUserIdResponseWeights {
	s.KeywordWeight = &v
	return s
}

type ListDatasetByUserIdResponseWeightsBuilder struct {
	s *ListDatasetByUserIdResponseWeights
}

func NewListDatasetByUserIdResponseWeightsBuilder() *ListDatasetByUserIdResponseWeightsBuilder {
	s := &ListDatasetByUserIdResponseWeights{}
	b := &ListDatasetByUserIdResponseWeightsBuilder{s: s}
	return b
}

func (b *ListDatasetByUserIdResponseWeightsBuilder) VectorWeight(v float32) *ListDatasetByUserIdResponseWeightsBuilder {
	b.s.VectorWeight = &v
	return b
}

func (b *ListDatasetByUserIdResponseWeightsBuilder) KeywordWeight(v float32) *ListDatasetByUserIdResponseWeightsBuilder {
	b.s.KeywordWeight = &v
	return b
}

func (b *ListDatasetByUserIdResponseWeightsBuilder) Build() *ListDatasetByUserIdResponseWeights {
	return b.s
}
