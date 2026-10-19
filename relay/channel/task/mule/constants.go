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
		"480p":  1,
		"720p":  0.100 / 0.05,
		"1080p": 0.2 / 0.05,
	},
	"berry-1.0-pro": {
		"1080p": 1,
		"2k":    0.2 / 0.18,
		"4k":    0.23 / 0.18,
	},
	"berry-1.0-turbe": {
		"480p":  1,
		"720p":  0.14 / 0.068,
		"1080p": 0.28 / 0.068,
	},
	"berry-1.0-pro-turbe": {
		"1080p": 1,
		"2k":    0.28 / 0.26,
		"4k":    0.31 / 0.26,
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
