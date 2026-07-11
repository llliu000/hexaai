package doubao

var ModelList = []string{}

var ChannelName = "doubao-video"

type videoPriceKey struct {
	Resolution string // 480p、720p、1080p、4k
	HasVideo   bool   // 输入是否包含视频
	HasAudio   bool   // 输出是否有声音
}

var videoPriceTable = map[string]map[videoPriceKey]float64{
	"doubao-seedance-2-0-mini-260615": {
		{Resolution: "480p", HasVideo: false}: 23,
		{Resolution: "480p", HasVideo: true}:  14,

		{Resolution: "720p", HasVideo: false}: 23,
		{Resolution: "720p", HasVideo: true}:  14,
	},
	"doubao-seedance-2-0-260128": {
		{Resolution: "480p", HasVideo: false}: 46,
		{Resolution: "480p", HasVideo: true}:  28,

		{Resolution: "720p", HasVideo: false}: 46,
		{Resolution: "720p", HasVideo: true}:  28,

		{Resolution: "1080p", HasVideo: false}: 51,
		{Resolution: "1080p", HasVideo: true}:  31,

		{Resolution: "4k", HasVideo: false}: 26,
		{Resolution: "4k", HasVideo: true}:  16,
	},
	"doubao-seedance-2-0-fast-260128": {
		{Resolution: "480p", HasVideo: false}: 37,
		{Resolution: "480p", HasVideo: true}:  22,

		{Resolution: "720p", HasVideo: false}: 37,
		{Resolution: "720p", HasVideo: true}:  22,
	},
	"doubao-seedance-1-5-pro-251215": {
		{HasAudio: false}: 8,
		{HasAudio: true}:  16,
	},
	"doubao-seedance-1-0-pro-fast-251015": {
		{HasAudio: false}: 4.2,
		{HasAudio: true}:  4.2,
	},
	"doubao-seedance-1-0-pro-250528": {
		{HasAudio: false}: 15,
		{HasAudio: true}:  15,
	},
}

func GetVideoInputRatio(model string, metadata map[string]interface{}) (float64, bool) {
	hasVideo, hasAudio, resolution := hasInfo(metadata)
	prices, ok := videoPriceTable[model]
	if !ok {
		return 0, false
	}
	// 模型实际价格
	key := videoPriceKey{
		Resolution: resolution,
		HasVideo:   hasVideo,
		HasAudio:   hasAudio,
	}
	price, ok := prices[key]
	if !ok {
		return 0, false
	}
	// 基准价格
	if key.Resolution != "" {
		key.HasVideo = true
		key.HasAudio = false
	} else {
		key.HasAudio = false
	}
	basePrice, ok := prices[key]
	if !ok || basePrice == 0 {
		return 0, false
	}
	return price / basePrice, true
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
