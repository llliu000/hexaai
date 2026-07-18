package ali

var ModelList = []string{}

var ChannelName = "ali"

// 以最低价格最为1倍倍率
var priceRatio = map[string]map[string]any{
	// 文生视频
	"wan2.7-t2v": {
		"720P":  0.6,
		"1080P": 1.0 / 0.6,
	},
	"wan2.6-t2v": {
		"720P":  1.0,
		"1080P": 1 / 0.6,
	},
	"wan2.5-t2v-preview": {
		"480P":  1.0,
		"720P":  0.6 / 0.3,
		"1080P": 1 / 0.3,
	},
	"wan2.2-t2v-plus": {
		"480P":  1.0,
		"1080P": 0.70 / 0.14,
	},

	// 图生视频-基于首帧
	"wan2.6-i2v-flash": {
		"720P": map[string]float64{
			"silent": 1.0,
			"audio":  0.3 / 0.15,
		},
		"1080P": map[string]float64{
			"silent": 0.25 / 0.15,
			"audio":  0.5 / 0.15,
		},
	},
	"wan2.7-i2v": {
		"720P":  1.0,
		"1080P": 1.0 / 0.6,
	},
	"wan2.6-i2v": {
		"720P":  1.0,
		"1080P": 1.0 / 0.6,
	},
	"wan2.5-i2v-preview": {
		"480P":  1.0,
		"720P":  0.6 / 0.3,
		"1080P": 1.0 / 0.3,
	},
	"wan2.2-i2v-flash": {
		"480P":  1.0,
		"720P":  0.2 / 0.1,
		"1080P": 0.48 / 0.1,
	},
	"wan2.2-i2v-plus": {
		"480P":  1.0,
		"1080P": 0.7 / 0.14,
	},

	// 图生视频-基于首尾帧
	"wan2.2-kf2v-flash": {
		"480P":  1.0,
		"720P":  0.2 / 0.1,
		"1080P": 0.7 / 0.1,
	},

	// 参考生视频
	"wan2.6-r2v-flash": {
		"720P": map[string]float64{
			"silent": 1.0,
			"audio":  0.3 / 0.15,
		},
		"1080P": map[string]float64{
			"silent": 0.25 / 0.15,
			"audio":  0.5 / 0.15,
		},
	},
	"wan2.7-r2v": {
		"720P":  1.0,
		"1080P": 1 / 0.6,
	},
	"wan2.6-r2v": {
		"720P":  1.0,
		"1080P": 1 / 0.6,
	},
}
