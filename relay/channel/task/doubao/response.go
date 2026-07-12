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
	Error struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`
	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
}
