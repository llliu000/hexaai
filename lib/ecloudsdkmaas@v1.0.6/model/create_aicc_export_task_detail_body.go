// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAICCExportTaskDetailBody struct {
	position.Body
	// 移动云页面上apiKey名称，不是具体的apikey的值
	ApiKey *string `json:"apiKey,omitempty"`
	// RAM用户名（可选，非空时仅导出该RAM账号的明细）
	RamName *string `json:"ramName,omitempty"`
	// 查询开始时间（包含），格式：yyyy-MM-dd HH:mm:ss
	BeginTime *string `json:"beginTime,omitempty"`
	// 查询结束时间（不包含），格式：yyyy-MM-dd HH:mm:ss。与beginTime跨度不超过30天
	EndTime *string `json:"endTime,omitempty"`
}

func (s CreateAICCExportTaskDetailBody) String() string {
	return utils.Beautify(s)
}

func (s CreateAICCExportTaskDetailBody) GoString() string {
	return s.String()
}

func (s CreateAICCExportTaskDetailBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAICCExportTaskDetailBody) SetApiKey(v string) *CreateAICCExportTaskDetailBody {
	s.ApiKey = &v
	return s
}

func (s *CreateAICCExportTaskDetailBody) SetRamName(v string) *CreateAICCExportTaskDetailBody {
	s.RamName = &v
	return s
}

func (s *CreateAICCExportTaskDetailBody) SetBeginTime(v string) *CreateAICCExportTaskDetailBody {
	s.BeginTime = &v
	return s
}

func (s *CreateAICCExportTaskDetailBody) SetEndTime(v string) *CreateAICCExportTaskDetailBody {
	s.EndTime = &v
	return s
}

type CreateAICCExportTaskDetailBodyBuilder struct {
	s *CreateAICCExportTaskDetailBody
}

func NewCreateAICCExportTaskDetailBodyBuilder() *CreateAICCExportTaskDetailBodyBuilder {
	s := &CreateAICCExportTaskDetailBody{}
	b := &CreateAICCExportTaskDetailBodyBuilder{s: s}
	return b
}

func (b *CreateAICCExportTaskDetailBodyBuilder) ApiKey(v string) *CreateAICCExportTaskDetailBodyBuilder {
	b.s.ApiKey = &v
	return b
}

func (b *CreateAICCExportTaskDetailBodyBuilder) RamName(v string) *CreateAICCExportTaskDetailBodyBuilder {
	b.s.RamName = &v
	return b
}

func (b *CreateAICCExportTaskDetailBodyBuilder) BeginTime(v string) *CreateAICCExportTaskDetailBodyBuilder {
	b.s.BeginTime = &v
	return b
}

func (b *CreateAICCExportTaskDetailBodyBuilder) EndTime(v string) *CreateAICCExportTaskDetailBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *CreateAICCExportTaskDetailBodyBuilder) Build() *CreateAICCExportTaskDetailBody {
	return b.s
}
