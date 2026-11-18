// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryAICCExportTaskDetailResponseBody struct {

	// 最后更新时间
	UpdatedTime *string `json:"updatedTime,omitempty"`
	// CSV文件预签名下载链接（status=SUCCESS时有值，24h有效）
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	// 失败原因（status=FAILED时有值）
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
	// 导出记录数（status=SUCCESS时有值）
	TotalRows *int64 `json:"totalRows,omitempty"`
	// 导出任务ID
	TaskId *string `json:"taskId,omitempty"`
	// 任务状态：PENDING（待执行）/ RUNNING（执行中）/ SUCCESS（成功）/ FAILED（失败）
	Status *string `json:"status,omitempty"`
}

func (s QueryAICCExportTaskDetailResponseBody) String() string {
	return utils.Beautify(s)
}

func (s QueryAICCExportTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s QueryAICCExportTaskDetailResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryAICCExportTaskDetailResponseBody) SetUpdatedTime(v string) *QueryAICCExportTaskDetailResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *QueryAICCExportTaskDetailResponseBody) SetDownloadUrl(v string) *QueryAICCExportTaskDetailResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *QueryAICCExportTaskDetailResponseBody) SetErrorMessage(v string) *QueryAICCExportTaskDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryAICCExportTaskDetailResponseBody) SetCreatedTime(v string) *QueryAICCExportTaskDetailResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *QueryAICCExportTaskDetailResponseBody) SetTotalRows(v int64) *QueryAICCExportTaskDetailResponseBody {
	s.TotalRows = &v
	return s
}

func (s *QueryAICCExportTaskDetailResponseBody) SetTaskId(v string) *QueryAICCExportTaskDetailResponseBody {
	s.TaskId = &v
	return s
}

func (s *QueryAICCExportTaskDetailResponseBody) SetStatus(v string) *QueryAICCExportTaskDetailResponseBody {
	s.Status = &v
	return s
}

type QueryAICCExportTaskDetailResponseBodyBuilder struct {
	s *QueryAICCExportTaskDetailResponseBody
}

func NewQueryAICCExportTaskDetailResponseBodyBuilder() *QueryAICCExportTaskDetailResponseBodyBuilder {
	s := &QueryAICCExportTaskDetailResponseBody{}
	b := &QueryAICCExportTaskDetailResponseBodyBuilder{s: s}
	return b
}

func (b *QueryAICCExportTaskDetailResponseBodyBuilder) UpdatedTime(v string) *QueryAICCExportTaskDetailResponseBodyBuilder {
	b.s.UpdatedTime = &v
	return b
}

func (b *QueryAICCExportTaskDetailResponseBodyBuilder) DownloadUrl(v string) *QueryAICCExportTaskDetailResponseBodyBuilder {
	b.s.DownloadUrl = &v
	return b
}

func (b *QueryAICCExportTaskDetailResponseBodyBuilder) ErrorMessage(v string) *QueryAICCExportTaskDetailResponseBodyBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *QueryAICCExportTaskDetailResponseBodyBuilder) CreatedTime(v string) *QueryAICCExportTaskDetailResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *QueryAICCExportTaskDetailResponseBodyBuilder) TotalRows(v int64) *QueryAICCExportTaskDetailResponseBodyBuilder {
	b.s.TotalRows = &v
	return b
}

func (b *QueryAICCExportTaskDetailResponseBodyBuilder) TaskId(v string) *QueryAICCExportTaskDetailResponseBodyBuilder {
	b.s.TaskId = &v
	return b
}

func (b *QueryAICCExportTaskDetailResponseBodyBuilder) Status(v string) *QueryAICCExportTaskDetailResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *QueryAICCExportTaskDetailResponseBodyBuilder) Build() *QueryAICCExportTaskDetailResponseBody {
	return b.s
}
