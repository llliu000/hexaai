// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryUsageDetailBodyInvokeTypeEnum int32

// List of InvokeType
const (
	QueryUsageDetailBodyInvokeTypeEnum0  QueryUsageDetailBodyInvokeTypeEnum = 0
	QueryUsageDetailBodyInvokeTypeEnum1  QueryUsageDetailBodyInvokeTypeEnum = 1
	QueryUsageDetailBodyInvokeTypeEnum2  QueryUsageDetailBodyInvokeTypeEnum = 2
	QueryUsageDetailBodyInvokeTypeEnum3  QueryUsageDetailBodyInvokeTypeEnum = 3
	QueryUsageDetailBodyInvokeTypeEnum51 QueryUsageDetailBodyInvokeTypeEnum = 51
)

type QueryUsageDetailBody struct {
	position.Body
	// RAM 用户名称列表。不传或传空列表表示查询所有 RAM 用户（含主账号自身）的调用记录
	RamUserNames []string `json:"ramUserNames,omitempty"`
	// 调用类型筛选。null/0=全部、1=在线体验、2=API调用、3=智能体调用、51=批量推理
	InvokeType *QueryUsageDetailBodyInvokeTypeEnum `json:"invokeType,omitempty"`
	// 请求 ID 精确匹配。传入后仅返回该 requestId 对应的单条记录；不传或传空字符串则不按 requestId 过滤
	RequestId *string `json:"requestId,omitempty"`
	// 页码，从 1 开始。默认 1
	PageNo *int32 `json:"pageNo,omitempty"`
	// 实际请求的模型名称列表，如 DeepSeek-v4-flash、zhipu/glm-5.1，不区分大小写；不传或传空列表表示查询全部模型
	InferenceNames []string `json:"inferenceNames,omitempty"`
	// API Key 名称列表，用于按密钥筛选调用日志。不传或传空列表表示不限制密钥。
	ApiKeyNames []string `json:"apiKeyNames,omitempty"`
	// 每页返回的条数，范围 [1, 100]。默认 10
	PageSize *int32 `json:"pageSize,omitempty"`
	// 查询开始时间（包含），格式 yyyy-MM-dd HH:mm:ss。不能晚于结束时间，最多可查询 14 天的跨度
	BeginTime *string `json:"beginTime,omitempty"`
	// 查询结束时间（不包含），格式 yyyy-MM-dd HH:mm:ss。不能早于开始时间
	EndTime *string `json:"endTime,omitempty"`
	// 状态码分类筛选列表（多值 OR）：1=成功（200）、2=用户行为导致的异常（400-499）、3=服务不可用（其他状态码）
	StatusTypes []int32 `json:"statusTypes,omitempty"`
}

func (s QueryUsageDetailBody) String() string {
	return utils.Beautify(s)
}

func (s QueryUsageDetailBody) GoString() string {
	return s.String()
}

func (s QueryUsageDetailBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryUsageDetailBody) SetRamUserNames(v []string) *QueryUsageDetailBody {
	s.RamUserNames = v
	return s
}

func (s *QueryUsageDetailBody) SetInvokeType(v QueryUsageDetailBodyInvokeTypeEnum) *QueryUsageDetailBody {
	s.InvokeType = &v
	return s
}

func (s *QueryUsageDetailBody) SetRequestId(v string) *QueryUsageDetailBody {
	s.RequestId = &v
	return s
}

func (s *QueryUsageDetailBody) SetPageNo(v int32) *QueryUsageDetailBody {
	s.PageNo = &v
	return s
}

func (s *QueryUsageDetailBody) SetInferenceNames(v []string) *QueryUsageDetailBody {
	s.InferenceNames = v
	return s
}

func (s *QueryUsageDetailBody) SetApiKeyNames(v []string) *QueryUsageDetailBody {
	s.ApiKeyNames = v
	return s
}

func (s *QueryUsageDetailBody) SetPageSize(v int32) *QueryUsageDetailBody {
	s.PageSize = &v
	return s
}

func (s *QueryUsageDetailBody) SetBeginTime(v string) *QueryUsageDetailBody {
	s.BeginTime = &v
	return s
}

func (s *QueryUsageDetailBody) SetEndTime(v string) *QueryUsageDetailBody {
	s.EndTime = &v
	return s
}

func (s *QueryUsageDetailBody) SetStatusTypes(v []int32) *QueryUsageDetailBody {
	s.StatusTypes = v
	return s
}

type QueryUsageDetailBodyBuilder struct {
	s *QueryUsageDetailBody
}

func NewQueryUsageDetailBodyBuilder() *QueryUsageDetailBodyBuilder {
	s := &QueryUsageDetailBody{}
	b := &QueryUsageDetailBodyBuilder{s: s}
	return b
}

func (b *QueryUsageDetailBodyBuilder) RamUserNames(v []string) *QueryUsageDetailBodyBuilder {
	b.s.RamUserNames = v
	return b
}

func (b *QueryUsageDetailBodyBuilder) InvokeType(v QueryUsageDetailBodyInvokeTypeEnum) *QueryUsageDetailBodyBuilder {
	b.s.InvokeType = &v
	return b
}

func (b *QueryUsageDetailBodyBuilder) RequestId(v string) *QueryUsageDetailBodyBuilder {
	b.s.RequestId = &v
	return b
}

func (b *QueryUsageDetailBodyBuilder) PageNo(v int32) *QueryUsageDetailBodyBuilder {
	b.s.PageNo = &v
	return b
}

func (b *QueryUsageDetailBodyBuilder) InferenceNames(v []string) *QueryUsageDetailBodyBuilder {
	b.s.InferenceNames = v
	return b
}

func (b *QueryUsageDetailBodyBuilder) ApiKeyNames(v []string) *QueryUsageDetailBodyBuilder {
	b.s.ApiKeyNames = v
	return b
}

func (b *QueryUsageDetailBodyBuilder) PageSize(v int32) *QueryUsageDetailBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *QueryUsageDetailBodyBuilder) BeginTime(v string) *QueryUsageDetailBodyBuilder {
	b.s.BeginTime = &v
	return b
}

func (b *QueryUsageDetailBodyBuilder) EndTime(v string) *QueryUsageDetailBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *QueryUsageDetailBodyBuilder) StatusTypes(v []int32) *QueryUsageDetailBodyBuilder {
	b.s.StatusTypes = v
	return b
}

func (b *QueryUsageDetailBodyBuilder) Build() *QueryUsageDetailBody {
	return b.s
}
