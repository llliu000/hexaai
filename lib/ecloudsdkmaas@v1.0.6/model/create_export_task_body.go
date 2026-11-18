// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateExportTaskBody struct {
	position.Body
	// 模型名称，当前默认 AICC-Doubao-Seedance-2.0；未来扩展 aicc2/aicc3 等
	ModelName *string `json:"modelName,omitempty"`
	// 移动云页面上apiKey名称，不是具体的apikey的值
	ApiKey *string `json:"apiKey,omitempty"`
	// RAM用户名（可选，不为空时仅导出该 RAM 账号的扣减明细）
	RamName *string `json:"ramName,omitempty"`
	// 扣减开始时间（包含），格式 yyyy-MM-dd HH:mm:ss
	BeginTime *string `json:"beginTime,omitempty"`
	// 扣减结束时间（不包含），格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"endTime,omitempty"`
}

func (s CreateExportTaskBody) String() string {
	return utils.Beautify(s)
}

func (s CreateExportTaskBody) GoString() string {
	return s.String()
}

func (s CreateExportTaskBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateExportTaskBody) SetModelName(v string) *CreateExportTaskBody {
	s.ModelName = &v
	return s
}

func (s *CreateExportTaskBody) SetApiKey(v string) *CreateExportTaskBody {
	s.ApiKey = &v
	return s
}

func (s *CreateExportTaskBody) SetRamName(v string) *CreateExportTaskBody {
	s.RamName = &v
	return s
}

func (s *CreateExportTaskBody) SetBeginTime(v string) *CreateExportTaskBody {
	s.BeginTime = &v
	return s
}

func (s *CreateExportTaskBody) SetEndTime(v string) *CreateExportTaskBody {
	s.EndTime = &v
	return s
}

type CreateExportTaskBodyBuilder struct {
	s *CreateExportTaskBody
}

func NewCreateExportTaskBodyBuilder() *CreateExportTaskBodyBuilder {
	s := &CreateExportTaskBody{}
	b := &CreateExportTaskBodyBuilder{s: s}
	return b
}

func (b *CreateExportTaskBodyBuilder) ModelName(v string) *CreateExportTaskBodyBuilder {
	b.s.ModelName = &v
	return b
}

func (b *CreateExportTaskBodyBuilder) ApiKey(v string) *CreateExportTaskBodyBuilder {
	b.s.ApiKey = &v
	return b
}

func (b *CreateExportTaskBodyBuilder) RamName(v string) *CreateExportTaskBodyBuilder {
	b.s.RamName = &v
	return b
}

func (b *CreateExportTaskBodyBuilder) BeginTime(v string) *CreateExportTaskBodyBuilder {
	b.s.BeginTime = &v
	return b
}

func (b *CreateExportTaskBodyBuilder) EndTime(v string) *CreateExportTaskBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *CreateExportTaskBodyBuilder) Build() *CreateExportTaskBody {
	return b.s
}
