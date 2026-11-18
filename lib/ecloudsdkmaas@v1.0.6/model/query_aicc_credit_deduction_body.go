// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryAiccCreditDeductionBody struct {
	position.Body
	// 模型名称，当前默认 AICC-Doubao-Seedance-2.0；未来扩展 aicc2/aicc3 等
	ModelName *string `json:"modelName,omitempty"`
	// 移动云页面上apiKey名称，不是具体的apikey的值
	ApiKey *string `json:"apiKey,omitempty"`
	// 页码（从 1 开始）
	PageNo *int32 `json:"pageNo,omitempty"`
	// ram用户名
	RamName *string `json:"ramName,omitempty"`
	// 每页行数
	PageSize *int32 `json:"pageSize,omitempty"`
	// 扣减开始时间（包含）
	BeginTime *string `json:"beginTime,omitempty"`
	// 扣减结束时间（不包含）
	EndTime *string `json:"endTime,omitempty"`
}

func (s QueryAiccCreditDeductionBody) String() string {
	return utils.Beautify(s)
}

func (s QueryAiccCreditDeductionBody) GoString() string {
	return s.String()
}

func (s QueryAiccCreditDeductionBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryAiccCreditDeductionBody) SetModelName(v string) *QueryAiccCreditDeductionBody {
	s.ModelName = &v
	return s
}

func (s *QueryAiccCreditDeductionBody) SetApiKey(v string) *QueryAiccCreditDeductionBody {
	s.ApiKey = &v
	return s
}

func (s *QueryAiccCreditDeductionBody) SetPageNo(v int32) *QueryAiccCreditDeductionBody {
	s.PageNo = &v
	return s
}

func (s *QueryAiccCreditDeductionBody) SetRamName(v string) *QueryAiccCreditDeductionBody {
	s.RamName = &v
	return s
}

func (s *QueryAiccCreditDeductionBody) SetPageSize(v int32) *QueryAiccCreditDeductionBody {
	s.PageSize = &v
	return s
}

func (s *QueryAiccCreditDeductionBody) SetBeginTime(v string) *QueryAiccCreditDeductionBody {
	s.BeginTime = &v
	return s
}

func (s *QueryAiccCreditDeductionBody) SetEndTime(v string) *QueryAiccCreditDeductionBody {
	s.EndTime = &v
	return s
}

type QueryAiccCreditDeductionBodyBuilder struct {
	s *QueryAiccCreditDeductionBody
}

func NewQueryAiccCreditDeductionBodyBuilder() *QueryAiccCreditDeductionBodyBuilder {
	s := &QueryAiccCreditDeductionBody{}
	b := &QueryAiccCreditDeductionBodyBuilder{s: s}
	return b
}

func (b *QueryAiccCreditDeductionBodyBuilder) ModelName(v string) *QueryAiccCreditDeductionBodyBuilder {
	b.s.ModelName = &v
	return b
}

func (b *QueryAiccCreditDeductionBodyBuilder) ApiKey(v string) *QueryAiccCreditDeductionBodyBuilder {
	b.s.ApiKey = &v
	return b
}

func (b *QueryAiccCreditDeductionBodyBuilder) PageNo(v int32) *QueryAiccCreditDeductionBodyBuilder {
	b.s.PageNo = &v
	return b
}

func (b *QueryAiccCreditDeductionBodyBuilder) RamName(v string) *QueryAiccCreditDeductionBodyBuilder {
	b.s.RamName = &v
	return b
}

func (b *QueryAiccCreditDeductionBodyBuilder) PageSize(v int32) *QueryAiccCreditDeductionBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *QueryAiccCreditDeductionBodyBuilder) BeginTime(v string) *QueryAiccCreditDeductionBodyBuilder {
	b.s.BeginTime = &v
	return b
}

func (b *QueryAiccCreditDeductionBodyBuilder) EndTime(v string) *QueryAiccCreditDeductionBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *QueryAiccCreditDeductionBodyBuilder) Build() *QueryAiccCreditDeductionBody {
	return b.s
}
