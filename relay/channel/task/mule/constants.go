package mule

import (
	"time"
)

const (
	ChannelName = "mulerouter"
)

// muleResolutionPriceRatio stores each model's resolution price relative to
// its 1080p price. The model price configured in the channel is the 1080p
// price, so the ratio can be applied independently of the requested duration.
// Prices are from the MuleRouter Wan 3.0 Video model pages.
var muleResolutionPriceRatio = map[string]map[string]float64{
	"berry-1.0": {
		"480p":  0.050 / 0.200,
		"720p":  0.100 / 0.200,
		"1080p": 1,
	},
	"wan3.0-video": {
		"480p":  0.050 / 0.200,
		"720p":  0.100 / 0.200,
		"1080p": 1,
	},
	"wan3.0-video-prime": {
		"480p":  0.068 / 0.280,
		"720p":  0.140 / 0.280,
		"1080p": 1,
	},
}

type MuleRequest struct {
	Prompt          string   `json:"prompt,omitempty"`
	FirstFrame      string   `json:"first_frame,omitempty"`
	LastFrame       string   `json:"last_frame,omitempty"`
	ReferenceImages []string `json:"reference_images,omitempty"`
	ReferenceVideos []string `json:"reference_videos,omitempty"`
	ReferenceAudios []string `json:"reference_audios,omitempty"`
	File            string   `json:"file,omitempty"`
	Link            string   `json:"link,omitempty"`
	Resolution      string   `json:"resolution,omitempty"`
	Ratio           string   `json:"ratio,omitempty"`
	Duration        int      `json:"duration,omitempty"`
	Audio           bool     `json:"audio"`
	Seed            int      `json:"seed,omitempty"`
	PromptExtend    bool     `json:"prompt_extend"`

	durationProvided     bool
	resolutionProvided   bool
	ratioProvided        bool
	audioProvided        bool
	promptExtendProvided bool
	seedProvided         bool
}

type MuleResponse struct {
	Status    int    `json:"status,omitempty"`
	Title     string `json:"title,omitempty"`
	Detail    string `json:"detail,omitempty"`
	Instance  string `json:"instance,omitempty"`
	ErrorCode int    `json:"error_code,omitempty"`
	TaskInfo  struct {
		Id        string    `json:"id"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Error     struct {
			Code   int    `json:"code,omitempty"`
			Title  string `json:"title,omitempty"`
			Detail string `json:"detail,omitempty"`
		} `json:"error,omitempty"`
	} `json:"task_info"`
	Videos []string `json:"videos,omitempty"`
}
