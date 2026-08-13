package model_setting

import "github.com/QuantumNous/new-api/setting/config"

type EnhanceSetting struct {
	Channel string `json:"channel"`
	BaseUrl string `json:"base_url"`
	ApiKey  string `json:"api_key"`
}

var defaultEnhanceSetting = EnhanceSetting{
	BaseUrl: "https://mediakit.cn-beijing.volces.com",
	Channel: "volc",
	ApiKey:  "",
}

var enhanceSetting = defaultEnhanceSetting

func init() {
	config.GlobalConfig.Register("enhance", &enhanceSetting)
}

func GetEnhanceSetting() *EnhanceSetting {
	return &enhanceSetting
}
