package doubao

type responsePayload struct {
	ID string `json:"id"` // task_id
}

type responseTask struct {
	ID      string `json:"id,omitempty"`
	Model   string `json:"model,omitempty"`
	Status  string `json:"status,omitempty"`
	Content struct {
		VideoURL string `json:"video_url,omitempty"`
	} `json:"content,omitempty"`
	Seed            int    `json:"seed,omitempty"`
	Resolution      string `json:"resolution,omitempty"`
	Duration        int    `json:"duration,omitempty"`
	Ratio           string `json:"ratio,omitempty"`
	FramesPerSecond int    `json:"framespersecond,omitempty"`
	ServiceTier     string `json:"service_tier,omitempty"`
	Tools           []struct {
		Type string `json:"type,omitempty"`
	} `json:"tools,omitempty"`
	Usage struct {
		CompletionTokens int `json:"completion_tokens,omitempty"`
		TotalTokens      int `json:"total_tokens,omitempty"`
		ToolUsage        struct {
			WebSearch int `json:"web_search,omitempty"`
		} `json:"tool_usage,omitempty"`
	} `json:"usage,omitempty"`
	Error *struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`
	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

type tokenMartResponse struct {
	Task struct {
		Id              string   `json:"id"`
		Status          string   `json:"status"`
		Model           string   `json:"model"`
		DurationSeconds int      `json:"duration_seconds"`
		Outputs         []string `json:"outputs"`
		Error           *string  `json:"error"`
		CreatedAt       any      `json:"created_at"`
		CompletedAt     any      `json:"completed_at"`
		Usage           struct {
			CompletionTokens int `json:"completion_tokens"`
			TotalTokens      int `json:"total_tokens"`
		} `json:"usage"`
		Metadata responseTask `json:"metadata"`
	} `json:"task"`
}

type viduResponse struct {
	Id      string `json:"id"`
	Model   string `json:"model"`
	Status  string `json:"status"`
	Content *struct {
		VideoUrl     string `json:"video_url"`
		LastFrameUrl string `json:"last_frame_url"`
	} `json:"content"`
	Usage *struct {
		PromptTokens     string `json:"prompt_tokens"`
		CompletionTokens string `json:"completion_tokens"`
		TotalTokens      string `json:"total_tokens"`
	} `json:"usage"`
	Error *struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`
	Tools []struct {
		Type string `json:"type,omitempty"`
	} `json:"tools,omitempty"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
	Seed                  string `json:"seed"`
	Resolution            string `json:"resolution"`
	Ratio                 string `json:"ratio"`
	Duration              string `json:"duration"`
	FramesPerSecond       string `json:"frames_per_second"`
	ServiceTier           string `json:"service_tier"`
	ExecutionExpiresAfter string `json:"execution_expires_after"`
	GenerateAudio         bool   `json:"generate_audio"`
	Draft                 bool   `json:"draft"`
	OmniReferenceTaskType string `json:"omni_reference_task_type"`
	OutputFormat          string `json:"output_format"`
}
