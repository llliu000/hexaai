package doubao

import "strings"

var ModelList = []string{}

var ChannelName = "doubao-video"

func priceRatio(req *requestPayload) float64 {
	var (
		hasAudio = false
		hasVideo = false
	)
	for i := range len(req.Content) {
		if audio := req.Content[i].AudioURL; audio != nil {
			hasAudio = audio.URL != ""
		}
		if video := req.Content[i].VideoURL; video != nil {
			hasVideo = video.URL != ""
		}
	}
	switch req.Model {
	case "doubao-seedance-1-0-pro-250528":
	case "doubao-seedance-1-0-pro-fast-251015":
	case "doubao-seedance-1-5-pro-251215":
		// 有声视频 0.0160 元/千tokens
		// 无声视频 0.0080 元/千tokens
		return inlineIf(hasAudio, 16/8, 1.0)
	case "doubao-seedance-2-0-260128", "seedance":
		// 1080p
		// 不含视频：51.00
		// 包含视频：31.00
		if strings.ToLower(req.Resolution) == "1080p" {
			return inlineIf(hasVideo, 31/28.00, 51/28.00)
		}
		// 480p，720p
		// 不含视频：46.00
		// 包含视频：28.00
		return inlineIf(hasVideo, 1.0, 46/28.00)
	case "doubao-seedance-2-0-fast-260128", "seedance-fast":
		return inlineIf(hasVideo, 1.0, 37/22.00)
	case "doubao-seedance-2-0-sea", "seedance-2-sea":
		if strings.ToLower(req.Resolution) == "1080p" {
			return inlineIf(hasVideo, 4.7/4.3, 7.7/4.3)
		}
		return inlineIf(hasVideo, 1.0, 7/4.3)
	case "doubao-seedance-2-0-fast-sea", "seedance-2-fast-sea":
		return inlineIf(hasVideo, 1.0, 5.6/3.3)
	}
	return 1.0
}

func inlineIf[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
