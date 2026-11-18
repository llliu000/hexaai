// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryUsageStatsBodyInvokeTypeEnum int32

// List of InvokeType
const (
	QueryUsageStatsBodyInvokeTypeEnum0  QueryUsageStatsBodyInvokeTypeEnum = 0
	QueryUsageStatsBodyInvokeTypeEnum1  QueryUsageStatsBodyInvokeTypeEnum = 1
	QueryUsageStatsBodyInvokeTypeEnum2  QueryUsageStatsBodyInvokeTypeEnum = 2
	QueryUsageStatsBodyInvokeTypeEnum3  QueryUsageStatsBodyInvokeTypeEnum = 3
	QueryUsageStatsBodyInvokeTypeEnum51 QueryUsageStatsBodyInvokeTypeEnum = 51
)

type QueryUsageStatsBody struct {
	position.Body
	// RAM 用户名称列表。不传或传空列表表示查询所有 RAM 用户（含主账号自身）的用量
	RamUserNames []string `json:"ramUserNames,omitempty"`
	// 调用类型筛选。0=全部（默认）、1=在线体验、2=API调用、3=智能体调用、51=批量推理
	InvokeType *QueryUsageStatsBodyInvokeTypeEnum `json:"invokeType,omitempty"`
	// 页码，从 1 开始。默认 1
	PageNo *int32 `json:"pageNo,omitempty"`
	// 实际请求的模型名称列表，如 DeepSeek-v4-flash、zhipu/glm-5.1，不区分大小写；不传或传空列表表示查询全部模型
	InferenceNames []string `json:"inferenceNames,omitempty"`
	// API Key 名称列表，用于按密钥筛选。不传或传空列表表示不限制密钥。
	ApiKeyNames []string `json:"apiKeyNames,omitempty"`
	// 每页返回的条数，范围 [1, 100]。默认 10
	PageSize *int32 `json:"pageSize,omitempty"`
	// 统计开始时间（包含），格式 yyyy-MM-dd HH:mm:ss，精确到秒。不能晚于结束时间，且时间跨度不超过14天
	BeginTime *string `json:"beginTime,omitempty"`
	// 统计结束时间（不包含），格式 yyyy-MM-dd HH:mm:ss，精确到秒。不能早于开始时间
	EndTime *string `json:"endTime,omitempty"`
}

func (s QueryUsageStatsBody) String() string {
	return utils.Beautify(s)
}

func (s QueryUsageStatsBody) GoString() string {
	return s.String()
}

func (s QueryUsageStatsBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryUsageStatsBody) SetRamUserNames(v []string) *QueryUsageStatsBody {
	s.RamUserNames = v
	return s
}

func (s *QueryUsageStatsBody) SetInvokeType(v QueryUsageStatsBodyInvokeTypeEnum) *QueryUsageStatsBody {
	s.InvokeType = &v
	return s
}

func (s *QueryUsageStatsBody) SetPageNo(v int32) *QueryUsageStatsBody {
	s.PageNo = &v
	return s
}

func (s *QueryUsageStatsBody) SetInferenceNames(v []string) *QueryUsageStatsBody {
	s.InferenceNames = v
	return s
}

func (s *QueryUsageStatsBody) SetApiKeyNames(v []string) *QueryUsageStatsBody {
	s.ApiKeyNames = v
	return s
}

func (s *QueryUsageStatsBody) SetPageSize(v int32) *QueryUsageStatsBody {
	s.PageSize = &v
	return s
}

func (s *QueryUsageStatsBody) SetBeginTime(v string) *QueryUsageStatsBody {
	s.BeginTime = &v
	return s
}

func (s *QueryUsageStatsBody) SetEndTime(v string) *QueryUsageStatsBody {
	s.EndTime = &v
	return s
}

type QueryUsageStatsBodyBuilder struct {
	s *QueryUsageStatsBody
}

func NewQueryUsageStatsBodyBuilder() *QueryUsageStatsBodyBuilder {
	s := &QueryUsageStatsBody{}
	b := &QueryUsageStatsBodyBuilder{s: s}
	return b
}

func (b *QueryUsageStatsBodyBuilder) RamUserNames(v []string) *QueryUsageStatsBodyBuilder {
	b.s.RamUserNames = v
	return b
}

func (b *QueryUsageStatsBodyBuilder) InvokeType(v QueryUsageStatsBodyInvokeTypeEnum) *QueryUsageStatsBodyBuilder {
	b.s.InvokeType = &v
	return b
}

func (b *QueryUsageStatsBodyBuilder) PageNo(v int32) *QueryUsageStatsBodyBuilder {
	b.s.PageNo = &v
	return b
}

func (b *QueryUsageStatsBodyBuilder) InferenceNames(v []string) *QueryUsageStatsBodyBuilder {
	b.s.InferenceNames = v
	return b
}

func (b *QueryUsageStatsBodyBuilder) ApiKeyNames(v []string) *QueryUsageStatsBodyBuilder {
	b.s.ApiKeyNames = v
	return b
}

func (b *QueryUsageStatsBodyBuilder) PageSize(v int32) *QueryUsageStatsBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *QueryUsageStatsBodyBuilder) BeginTime(v string) *QueryUsageStatsBodyBuilder {
	b.s.BeginTime = &v
	return b
}

func (b *QueryUsageStatsBodyBuilder) EndTime(v string) *QueryUsageStatsBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *QueryUsageStatsBodyBuilder) Build() *QueryUsageStatsBody {
	return b.s
}
