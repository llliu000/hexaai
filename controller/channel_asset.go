package controller

import (
	"strconv"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/service"
	"github.com/gin-gonic/gin"
)

func SyncAssetsToChannel(c *gin.Context) {
	channelID, err := strconv.Atoi(c.Param("id"))
	if err != nil || channelID <= 0 {
		common.ApiErrorMsg(c, "Invalid channel ID")
		return
	}
	task, err := service.SyncAssetsToChannel(c.Request.Context(), channelID)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, task.ToResponse())
}

func GetChannelAssetSyncProgress(c *gin.Context) {
	channelID, err := strconv.Atoi(c.Param("id"))
	if err != nil || channelID <= 0 {
		common.ApiErrorMsg(c, "Invalid channel ID")
		return
	}
	progress, err := service.GetChannelAssetSyncProgress(c.Request.Context(), channelID, c.Query("task_id"))
	if err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, progress)
}
