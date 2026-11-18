// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryBillModelsResponseBody struct {

	// 模型名称
	ModelName *string `json:"modelName,omitempty"`
	// 实例ID
	InstanceId *string `json:"instanceId,omitempty"`
	// 失效时间
	ExpireTime *string `json:"expireTime,omitempty"`
	// 生效时间
	EffectTime *string `json:"effectTime,omitempty"`
	// 模型类别
	ModelCategory *string `json:"modelCategory,omitempty"`
	// 服务商
	ProviderName *string `json:"providerName,omitempty"`
}

func (s QueryBillModelsResponseBody) String() string {
	return utils.Beautify(s)
}

func (s QueryBillModelsResponseBody) GoString() string {
	return s.String()
}

func (s QueryBillModelsResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryBillModelsResponseBody) SetModelName(v string) *QueryBillModelsResponseBody {
	s.ModelName = &v
	return s
}

func (s *QueryBillModelsResponseBody) SetInstanceId(v string) *QueryBillModelsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QueryBillModelsResponseBody) SetExpireTime(v string) *QueryBillModelsResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *QueryBillModelsResponseBody) SetEffectTime(v string) *QueryBillModelsResponseBody {
	s.EffectTime = &v
	return s
}

func (s *QueryBillModelsResponseBody) SetModelCategory(v string) *QueryBillModelsResponseBody {
	s.ModelCategory = &v
	return s
}

func (s *QueryBillModelsResponseBody) SetProviderName(v string) *QueryBillModelsResponseBody {
	s.ProviderName = &v
	return s
}

type QueryBillModelsResponseBodyBuilder struct {
	s *QueryBillModelsResponseBody
}

func NewQueryBillModelsResponseBodyBuilder() *QueryBillModelsResponseBodyBuilder {
	s := &QueryBillModelsResponseBody{}
	b := &QueryBillModelsResponseBodyBuilder{s: s}
	return b
}

func (b *QueryBillModelsResponseBodyBuilder) ModelName(v string) *QueryBillModelsResponseBodyBuilder {
	b.s.ModelName = &v
	return b
}

func (b *QueryBillModelsResponseBodyBuilder) InstanceId(v string) *QueryBillModelsResponseBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *QueryBillModelsResponseBodyBuilder) ExpireTime(v string) *QueryBillModelsResponseBodyBuilder {
	b.s.ExpireTime = &v
	return b
}

func (b *QueryBillModelsResponseBodyBuilder) EffectTime(v string) *QueryBillModelsResponseBodyBuilder {
	b.s.EffectTime = &v
	return b
}

func (b *QueryBillModelsResponseBodyBuilder) ModelCategory(v string) *QueryBillModelsResponseBodyBuilder {
	b.s.ModelCategory = &v
	return b
}

func (b *QueryBillModelsResponseBodyBuilder) ProviderName(v string) *QueryBillModelsResponseBodyBuilder {
	b.s.ProviderName = &v
	return b
}

func (b *QueryBillModelsResponseBodyBuilder) Build() *QueryBillModelsResponseBody {
	return b.s
}
