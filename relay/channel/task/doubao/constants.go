package doubao

import (
	"strings"

	"github.com/QuantumNous/new-api/model"
)

var ModelList = []string{}
var AssetPrefix = "asset://"
var ChannelName = "doubao-video"

type videoPriceConfig struct {
	BasePrice float64
	Prices    map[videoPriceKey]float64
}

type videoPriceKey struct {
	Resolution string // 480p、720p、1080p、4k
	HasVideo   bool   // 输入是否包含视频
	HasAudio   bool   // 输出是否有声音
}

var videoPriceTable = map[string]*videoPriceConfig{
	"doubao-seedance-2-0-mini": {
		Prices: map[videoPriceKey]float64{
			{Resolution: "480p", HasVideo: false}: 23,
			{Resolution: "480p", HasVideo: true}:  14,

			{Resolution: "720p", HasVideo: false}: 23,
			{Resolution: "720p", HasVideo: true}:  14,
		},
		BasePrice: 14,
	},
	"doubao-seedance-2-0": {
		Prices: map[videoPriceKey]float64{
			{Resolution: "480p", HasVideo: false}: 46,
			{Resolution: "480p", HasVideo: true}:  28,

			{Resolution: "720p", HasVideo: false}: 46,
			{Resolution: "720p", HasVideo: true}:  28,

			{Resolution: "1080p", HasVideo: false}: 51,
			{Resolution: "1080p", HasVideo: true}:  31,

			{Resolution: "4k", HasVideo: false}: 26,
			{Resolution: "4k", HasVideo: true}:  16,
		},
		BasePrice: 16,
	},
	"doubao-seedance-2-0-fast": {
		Prices: map[videoPriceKey]float64{
			{Resolution: "480p", HasVideo: false}: 37,
			{Resolution: "480p", HasVideo: true}:  22,

			{Resolution: "720p", HasVideo: false}: 37,
			{Resolution: "720p", HasVideo: true}:  22,
		},
		BasePrice: 22,
	},
	"dreamina-seedance-2-0-mini": {
		Prices: map[videoPriceKey]float64{
			{Resolution: "480p", HasVideo: false}: 3.5,
			{Resolution: "480p", HasVideo: true}:  2.1,

			{Resolution: "720p", HasVideo: false}: 3.5,
			{Resolution: "720p", HasVideo: true}:  2.1,
		},
		BasePrice: 2.1,
	},
	"dreamina-seedance-2-0": {
		Prices: map[videoPriceKey]float64{
			{Resolution: "480p", HasVideo: false}: 7.0,
			{Resolution: "480p", HasVideo: true}:  4.3,

			{Resolution: "720p", HasVideo: false}: 7.0,
			{Resolution: "720p", HasVideo: true}:  4.3,

			{Resolution: "1080p", HasVideo: false}: 7.7,
			{Resolution: "1080p", HasVideo: true}:  4.7,

			{Resolution: "4k", HasVideo: false}: 4.0,
			{Resolution: "4k", HasVideo: true}:  2.4,
		},
		BasePrice: 2.4,
	},
	"dreamina-seedance-2-0-fast": {
		Prices: map[videoPriceKey]float64{
			{Resolution: "480p", HasVideo: false}: 5.6,
			{Resolution: "480p", HasVideo: true}:  3.3,

			{Resolution: "720p", HasVideo: false}: 5.6,
			{Resolution: "720p", HasVideo: true}:  3.3,
		},
		BasePrice: 3.3,
	},
	"doubao-seedance-1-5-pro": {
		Prices: map[videoPriceKey]float64{
			{HasAudio: false}: 8,
			{HasAudio: true}:  16,
		},
		BasePrice: 8,
	},
	"doubao-seedance-1-0-pro-fast": {
		Prices: map[videoPriceKey]float64{
			{HasAudio: false}: 4.2,
			{HasAudio: true}:  4.2,
		},
		BasePrice: 4.2,
	},
	"doubao-seedance-1-0-pro": {
		Prices: map[videoPriceKey]float64{
			{HasAudio: false}: 15,
			{HasAudio: true}:  15,
		},
		BasePrice: 15,
	},
}

func GetVideoInputRatio(model string, metadata map[string]interface{}) (float64, bool) {
	hasVideo, hasAudio, resolution := hasInfo(metadata)
	cfg, ok := videoPriceTable[model]
	if !ok {
		return 0, false
	}
	// 模型实际价格
	price, ok := cfg.Prices[videoPriceKey{
		Resolution: resolution,
		HasVideo:   hasVideo,
		HasAudio:   hasAudio,
	}]
	if !ok {
		return 0, false
	}
	return price / cfg.BasePrice, true
}

func hasInfo(metadata map[string]interface{}) (hasVideo, hasAudio bool, resolution string) {
	if resolution, _ = metadata["resolution"].(string); resolution == "" {
		resolution = "720p"
	}
	content, ok := metadata["content"].([]interface{})
	if !ok {
		return
	}
	for _, v := range content {
		item, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if item["type"] == "video_url" || item["video_url"] != nil {
			hasVideo = true
		}
		if item["type"] == "audio_url" || item["audio_url"] != nil {
			hasAudio = true
		}
	}
	return
}

func resolveAssetURL(channelId int, raw string) string {
	if !strings.HasPrefix(raw, AssetPrefix) {
		return raw
	}
	assetId := strings.TrimPrefix(raw, AssetPrefix)
	ac, err := model.GetAssetChannelByChannelIdAndAssetId(channelId, assetId)
	if nil != err {
		return raw
	}
	return AssetPrefix + ac.UpstreamAssertId
}
