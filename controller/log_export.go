package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// logOtherData mirrors the JSON structure stored in the Log.Other field.
type logOtherData struct {
	OriginQuota       int     `json:"origin_quota"`
	Discount          float64 `json:"discount"`
	ModelRatio        float64 `json:"model_ratio"`
	GroupRatio        float64 `json:"group_ratio"`
	CompletionRatio   float64 `json:"completion_ratio"`
	CacheTokens       int     `json:"cache_tokens"`
	CacheRatio        float64 `json:"cache_ratio"`
	ModelPrice        float64 `json:"model_price"`
	UserGroupRatio    float64 `json:"user_group_ratio"`
	IsModelMapped     bool    `json:"is_model_mapped"`
	UpstreamModelName string  `json:"upstream_model_name"`

	// Claude-specific
	CacheCreationTokens   int     `json:"cache_creation_tokens"`
	CacheCreationRatio    float64 `json:"cache_creation_ratio"`
	CacheCreationTokens5m int     `json:"cache_creation_tokens_5m"`
	CacheCreationRatio5m  float64 `json:"cache_creation_ratio_5m"`
	CacheCreationTokens1h int     `json:"cache_creation_tokens_1h"`
	CacheCreationRatio1h  float64 `json:"cache_creation_ratio_1h"`

	// Audio-specific
	AudioInput           int     `json:"audio_input"`
	AudioOutput          int     `json:"audio_output"`
	TextInput            int     `json:"text_input"`
	TextOutput           int     `json:"text_output"`
	AudioRatio           float64 `json:"audio_ratio"`
	AudioCompletionRatio float64 `json:"audio_completion_ratio"`

	// Image-specific
	ImageOutputTokens int     `json:"image_output_tokens"`
	ImageRatio        float64 `json:"image_ratio"`

	// Task
	TaskID string `json:"task_id"`

	// Billing
	BillingSource string `json:"billing_source"`

	// First response time (ms)
	Frt float64 `json:"frt"`
}

func parseLogOther(otherStr string) logOtherData {
	var data logOtherData
	if otherStr == "" {
		return data
	}
	_ = common.Unmarshal([]byte(otherStr), &data)
	return data
}

// logTypeName returns a human-readable Chinese name for the log type.
func logTypeName(logType int) string {
	switch logType {
	case model.LogTypeTopup:
		return "充值"
	case model.LogTypeConsume:
		return "消费"
	case model.LogTypeManage:
		return "管理"
	case model.LogTypeSystem:
		return "系统"
	case model.LogTypeError:
		return "错误"
	case model.LogTypeRefund:
		return "退款"
	default:
		return "未知"
	}
}

// billingTypeName determines the billing method from the log's Other data.
func billingTypeName(other logOtherData) string {
	if other.ModelPrice > 0 {
		return "按次"
	}
	if other.ModelRatio > 0 || other.CompletionRatio > 0 {
		return "按Token"
	}
	return ""
}

// quotaToUSD converts internal quota units to USD amount.
func quotaToUSD(quota int) float64 {
	return float64(quota) / common.QuotaPerUnit
}

// ratioToPrice converts a model_ratio value to $/1M tokens.
// model_ratio unit: 1 = $0.002/1K tokens, so $/1M = model_ratio * 2.0
func ratioToPrice(modelRatio float64) float64 {
	return modelRatio * 2.0
}

// Excel column headers for log export (23 columns, A-W).
var exportHeaders = []interface{}{
	"时间",                 // A
	"请求ID",               // B
	"任务ID",               // C
	"分组名称",               // D
	"模型名称",               // E
	"类型",                 // F
	"原始Quota",            // G
	"折扣",                 // H
	"Quota",              // I
	"计价方式",               // J
	"输入价格($/1M tokens)",  // K
	"输出价格($/1M tokens)",  // L
	"缓存读价格($/1M tokens)", // M
	"缓存写价格($/1M tokens)", // N
	"输入Token数",           // O
	"输出Token数",           // P
	"缓存命中Token数",         // Q
	"缓存写Token数",          // R
	"请求时间 (秒)",           // S
	"预扣金额 ($)",           // T
	"退款金额 ($)",           // U
	"实际扣费金额 ($)",         // V
	"执行状态",               // W
}

// ExportLogsExcel exports filtered logs as an Excel file download.
// Supports up to 500,000 rows via streaming write + batched DB reads.
func ExportLogsExcel(c *gin.Context) {
	logType, _ := strconv.Atoi(c.Query("type"))
	startTimestamp, _ := strconv.ParseInt(c.Query("start_timestamp"), 10, 64)
	endTimestamp, _ := strconv.ParseInt(c.Query("end_timestamp"), 10, 64)
	username := c.Query("username")
	tokenName := c.Query("token_name")
	modelName := c.Query("model_name")
	channel, _ := strconv.Atoi(c.Query("channel"))
	group := c.Query("group")
	requestId := c.Query("request_id")

	f := excelize.NewFile()
	defer f.Close()
	sheetName := "Sheet1"

	// --- Pre-set column widths (must be done before StreamWriter) ---
	colWidths := []struct {
		col   string
		width float64
	}{
		{"A", 20}, {"B", 28}, {"C", 28}, {"D", 16}, {"E", 28},
		{"F", 8}, {"G", 12}, {"H", 10}, {"I", 22}, {"J", 22},
		{"K", 22}, {"L", 26}, {"M", 18}, {"N", 18}, {"O", 22},
		{"P", 24}, {"Q", 12}, {"R", 16}, {"S", 16}, {"T", 16},
		{"U", 16}, {"V", 16}, {"W", 12},
	}
	for _, cw := range colWidths {
		_ = f.SetColWidth(sheetName, cw.col, cw.col, cw.width)
	}

	// --- Header style ---
	headerStyleID, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#D9E1F2"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
	})

	// --- Create StreamWriter ---
	sw, err := f.NewStreamWriter(sheetName)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": "创建Excel流失败: " + err.Error()})
		return
	}

	// Write header row (row 1)
	_ = sw.SetRow("A1", exportHeaders, excelize.RowOpts{StyleID: headerStyleID, Height: 30})

	// --- Batched DB read + streaming write ---
	rowIdx := 2
	_, exportErr := model.ExportLogsBatch(
		logType, startTimestamp, endTimestamp,
		modelName, username, tokenName,
		channel, group, requestId,
		func(batch []*model.Log) error {
			for _, log := range batch {
				other := parseLogOther(log.Other)
				row := buildExportRow(log, other)
				cell, _ := excelize.CoordinatesToCellName(1, rowIdx)
				if err := sw.SetRow(cell, row); err != nil {
					return err
				}
				rowIdx++
			}
			return nil
		},
	)
	if exportErr != nil {
		c.JSON(500, gin.H{"success": false, "message": "导出日志失败: " + exportErr.Error()})
		return
	}

	// Flush stream writer
	if err := sw.Flush(); err != nil {
		c.JSON(500, gin.H{"success": false, "message": "写入Excel失败: " + err.Error()})
		return
	}

	// --- Stream response ---
	filename := fmt.Sprintf("logs_export_%s.xlsx", time.Now().Format("20060102_150405"))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")

	if err := f.Write(c.Writer); err != nil {
		// Headers already sent, can't send JSON error at this point.
		common.SysError("写入Excel到HTTP响应失败: " + err.Error())
	}
}

// buildExportRow converts a single Log into a row of []interface{} for StreamWriter.
func buildExportRow(log *model.Log, other logOtherData) []interface{} {
	row := make([]interface{}, len(exportHeaders))
	// A: 时间 (精确到秒)
	row[0] = time.Unix(log.CreatedAt, 0).Format("2006-01-02 15:04:05")
	// B: 上游请求ID
	row[1] = log.RequestId
	// C: 任务ID (异步任务时 other 中的 task_id)
	if other.TaskID != "" {
		row[2] = other.TaskID
	}
	// D: 分组名称
	row[3] = log.Group
	// E: 模型名称
	row[4] = log.ModelName
	// F: 类型
	row[5] = logTypeName(log.Type)
	// G: 原始价格
	if other.OriginQuota == 0 {
		row[6] = log.Quota
	} else {
		row[6] = other.OriginQuota
	}
	// H: 折扣
	if other.OriginQuota == 0 {
		row[7] = 1
	} else {
		row[7] = other.Discount
	}
	// I: Quota
	row[8] = log.Quota
	// J: 计价方式
	row[9] = billingTypeName(other)
	// K: 单价——输入
	if other.ModelPrice > 0 {
		row[10] = fmt.Sprintf("%.6f/次", other.ModelPrice)
	} else if other.ModelRatio > 0 {
		row[10] = fmt.Sprintf("%.6f", ratioToPrice(other.ModelRatio))
	}
	// L: 单价——输出
	if other.ModelPrice <= 0 && other.ModelRatio > 0 {
		row[11] = fmt.Sprintf("%.6f", ratioToPrice(other.ModelRatio)*other.CompletionRatio)
	}
	// M: 单价——Cache
	if other.CacheRatio > 0 && other.ModelRatio > 0 {
		row[12] = fmt.Sprintf("%.6f", ratioToPrice(other.ModelRatio)*other.CacheRatio)
	}
	// N: 单价——Cache Write
	if other.CacheCreationRatio > 0 && other.ModelRatio > 0 {
		row[13] = fmt.Sprintf("%.6f", ratioToPrice(other.ModelRatio)*other.CacheCreationRatio)
	}
	// O: 输入 Token
	row[14] = log.PromptTokens
	// P: 输出 Token
	row[15] = log.CompletionTokens
	// Q: Cache 命中 Token
	if other.CacheTokens > 0 {
		row[16] = other.CacheTokens
	}
	// R: Cache Write Token
	cacheWriteTokens := other.CacheCreationTokens + other.CacheCreationTokens5m + other.CacheCreationTokens1h
	if cacheWriteTokens > 0 {
		row[17] = cacheWriteTokens
	}
	// S: 请求时间 (秒)
	row[18] = log.UseTime
	// T: 预扣金额 (6位小数) — 固定金额，不参与分组倍率计算
	// 按次计费: 单次费用 = model_price
	// 按Token计费: 100,000 tokens 的输入价格 = ratioToPrice(modelRatio) / 1M * 100000
	if other.ModelPrice > 0 {
		row[19] = fmt.Sprintf("%.6f", other.ModelPrice)
	} else if other.ModelRatio > 0 {
		preDeduct := ratioToPrice(other.ModelRatio) * 100000.0 / 1000000.0
		row[19] = fmt.Sprintf("%.6f", preDeduct)
	}
	// U: 退款金额 (6位小数)
	if log.Type == model.LogTypeRefund {
		refundAmt := quotaToUSD(log.Quota)
		if refundAmt < 0 {
			refundAmt = -refundAmt
		}
		row[20] = fmt.Sprintf("%.6f", refundAmt)
	}
	// V: 实际扣费金额 (6位小数) — 退款类型无实际扣费
	if log.Type != model.LogTypeRefund {
		actualAmount := quotaToUSD(log.Quota)
		if actualAmount < 0 {
			actualAmount = -actualAmount
		}
		row[21] = fmt.Sprintf("%.6f", actualAmount)
	}
	// W: 执行状态
	if log.Type == model.LogTypeError {
		row[22] = "失败"
	} else if log.Type == model.LogTypeConsume || log.Type == model.LogTypeRefund {
		row[22] = "成功"
	} else {
		row[22] = logTypeName(log.Type)
	}
	return row
}
