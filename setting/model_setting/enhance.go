package model_setting

import (
	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/setting/config"
)

type EnhanceSetting struct {
	Channel      string `json:"channel"`
	BaseUrl      string `json:"base_url"`
	ApiKey       string `json:"api_key"`
	ToolVersion  string `json:"tool_version"`
	Scene        string `json:"scene"`
	BitrateLevel string `json:"bitrate_level"`
}

var defaultEnhanceSetting = EnhanceSetting{
	BaseUrl:      "https://mediakit.cn-beijing.volces.com",
	Channel:      "volc",
	ToolVersion:  "standard",
	Scene:        "aigc",
	BitrateLevel: "medium",
}

var enhanceSetting = defaultEnhanceSetting

func init() {
	config.GlobalConfig.Register("enhance", &enhanceSetting)
}

func GetEnhanceSetting() *EnhanceSetting {
	common.OptionMapRWMutex.RLock()
	defer common.OptionMapRWMutex.RUnlock()
	setting := enhanceSetting
	return &setting
}
