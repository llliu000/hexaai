// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type QueryExportTaskResponseBody struct {

	// 更新时间
	UpdatedTime *string `json:"updatedTime,omitempty"`
	// 下载链接（SUCCESS 时有值，预签名 URL，24h 有效）
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	// 失败原因（FAILED 时有值）
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
	// 导出记录数（SUCCESS 时有值）
	TotalRows *int64 `json:"totalRows,omitempty"`
	// 导出任务ID
	TaskId *string `json:"taskId,omitempty"`
	// 任务状态：PENDING/RUNNING/SUCCESS/FAILED
	Status *string `json:"status,omitempty"`
}

func (s QueryExportTaskResponseBody) String() string {
	return utils.Beautify(s)
}

func (s QueryExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s QueryExportTaskResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *QueryExportTaskResponseBody) SetUpdatedTime(v string) *QueryExportTaskResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *QueryExportTaskResponseBody) SetDownloadUrl(v string) *QueryExportTaskResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *QueryExportTaskResponseBody) SetErrorMessage(v string) *QueryExportTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryExportTaskResponseBody) SetCreatedTime(v string) *QueryExportTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *QueryExportTaskResponseBody) SetTotalRows(v int64) *QueryExportTaskResponseBody {
	s.TotalRows = &v
	return s
}

func (s *QueryExportTaskResponseBody) SetTaskId(v string) *QueryExportTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *QueryExportTaskResponseBody) SetStatus(v string) *QueryExportTaskResponseBody {
	s.Status = &v
	return s
}

type QueryExportTaskResponseBodyBuilder struct {
	s *QueryExportTaskResponseBody
}

func NewQueryExportTaskResponseBodyBuilder() *QueryExportTaskResponseBodyBuilder {
	s := &QueryExportTaskResponseBody{}
	b := &QueryExportTaskResponseBodyBuilder{s: s}
	return b
}

func (b *QueryExportTaskResponseBodyBuilder) UpdatedTime(v string) *QueryExportTaskResponseBodyBuilder {
	b.s.UpdatedTime = &v
	return b
}

func (b *QueryExportTaskResponseBodyBuilder) DownloadUrl(v string) *QueryExportTaskResponseBodyBuilder {
	b.s.DownloadUrl = &v
	return b
}

func (b *QueryExportTaskResponseBodyBuilder) ErrorMessage(v string) *QueryExportTaskResponseBodyBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *QueryExportTaskResponseBodyBuilder) CreatedTime(v string) *QueryExportTaskResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *QueryExportTaskResponseBodyBuilder) TotalRows(v int64) *QueryExportTaskResponseBodyBuilder {
	b.s.TotalRows = &v
	return b
}

func (b *QueryExportTaskResponseBodyBuilder) TaskId(v string) *QueryExportTaskResponseBodyBuilder {
	b.s.TaskId = &v
	return b
}

func (b *QueryExportTaskResponseBodyBuilder) Status(v string) *QueryExportTaskResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *QueryExportTaskResponseBodyBuilder) Build() *QueryExportTaskResponseBody {
	return b.s
}
