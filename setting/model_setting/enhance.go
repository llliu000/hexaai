package model_setting

import "github.com/QuantumNous/new-api/setting/config"

type EnhanceSetting struct {
	Channel      string `json:"channel"`
	BaseUrl      string `json:"base_url"`
	ApiKey       string `json:"api_key"`
	Enabled      bool   `json:"enabled"`
	ToolVersion  string `json:"tool_version"`
	Scene        string `json:"scene"`
	BitrateLevel string `json:"bitrate_level"`
}

var defaultEnhanceSetting = EnhanceSetting{
	BaseUrl:      "https://mediakit.cn-beijing.volces.com",
	Channel:      "volc",
	ApiKey:       "AKLTYjM3ZTlhZWQ3Y2NkNGI2Y2IyNWFiY2MxMzMzMTUyY2I",
	ToolVersion:  "standard",
	Scene:        "common",
	BitrateLevel: "medium",
}

var enhanceSetting = defaultEnhanceSetting

func init() {
	config.GlobalConfig.Register("enhance", &enhanceSetting)
}

func GetEnhanceSetting() *EnhanceSetting {
	return &enhanceSetting
}
