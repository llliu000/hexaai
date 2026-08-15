package volc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/QuantumNous/new-api/dto"
	"github.com/google/uuid"
)

const (
	DefaultBaseURL      = "https://mediakit.cn-beijing.volces.com"
	defaultToolVersion  = "standard"
	defaultScene        = "aigc"
	defaultBitrateLevel = "medium"
)

type Config struct {
	APIKey       string
	BaseURL      string
	ToolVersion  string
	Scene        string
	BitrateLevel string
}

type Adapter struct {
	apiKey       string
	baseURL      string
	toolVersion  string
	scene        string
	bitrateLevel string
}

type submitRequest struct {
	VideoURL     string `json:"video_url"`
	ToolVersion  string `json:"tool_version,omitempty"`
	Scene        string `json:"scene,omitempty"`
	Resolution   string `json:"resolution,omitempty"`
	BitrateLevel string `json:"bitrate_level,omitempty"`
	ClientToken  string `json:"client_token,omitempty"`
}

type submitResponse struct {
	Success   bool      `json:"success"`
	TaskID    string    `json:"task_id"`
	RequestID string    `json:"request_id"`
	Error     *apiError `json:"error,omitempty"`
}

type taskResponse struct {
	Success   bool       `json:"success"`
	TaskID    string     `json:"task_id"`
	Status    string     `json:"status"`
	Result    taskResult `json:"result"`
	Error     *apiError  `json:"error,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
}

type taskResult struct {
	VideoURL string `json:"video_url"`
}

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Param   string `json:"param,omitempty"`
	Type    string `json:"type,omitempty"`
}

func New(config Config) (*Adapter, error) {
	apiKey := strings.TrimSpace(config.APIKey)
	if apiKey == "" {
		return nil, fmt.Errorf("volc video enhancement api key is required")
	}

	baseURL := strings.TrimRight(strings.TrimSpace(config.BaseURL), "/")
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	parsedURL, err := url.Parse(baseURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return nil, fmt.Errorf("invalid volc video enhancement base url: %q", baseURL)
	}

	toolVersion := strings.TrimSpace(config.ToolVersion)
	if toolVersion == "" {
		toolVersion = defaultToolVersion
	}
	scene := strings.TrimSpace(config.Scene)
	if scene == "" {
		scene = defaultScene
	}
	bitrateLevel := strings.TrimSpace(config.BitrateLevel)
	if bitrateLevel == "" {
		bitrateLevel = defaultBitrateLevel
	}
	return &Adapter{
		apiKey:       apiKey,
		baseURL:      baseURL,
		toolVersion:  toolVersion,
		scene:        scene,
		bitrateLevel: bitrateLevel,
	}, nil
}

func (a *Adapter) Submit(ctx context.Context, request dto.SubmitRequest) (*dto.SubmitResult, error) {
	if strings.TrimSpace(request.VideoURL) == "" {
		return nil, fmt.Errorf("video url is required")
	}

	body, err := json.Marshal(submitRequest{
		VideoURL:     request.VideoURL,
		ToolVersion:  a.toolVersion,
		Scene:        a.scene,
		Resolution:   request.TargetResolution,
		BitrateLevel: a.bitrateLevel,
		ClientToken:  uuid.NewString(),
	})
	if err != nil {
		return nil, fmt.Errorf("marshal volc video enhancement request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.baseURL+"/api/v1/tools/enhance-video", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create volc video enhancement request: %w", err)
	}
	a.setHeaders(req)

	var response submitResponse
	if err := a.do(req, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, responseError(response.Error, response.RequestID)
	}
	if response.TaskID == "" {
		return nil, fmt.Errorf("volc video enhancement returned an empty task_id")
	}
	return &dto.SubmitResult{TaskID: response.TaskID}, nil
}

func (a *Adapter) GetTask(ctx context.Context, taskID string) (*dto.Task, error) {
	taskID = strings.TrimSpace(taskID)
	if taskID == "" {
		return nil, fmt.Errorf("task id is required")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.baseURL+"/api/v1/tasks/"+url.PathEscape(taskID), nil)
	if err != nil {
		return nil, fmt.Errorf("create volc video enhancement task request: %w", err)
	}
	a.setHeaders(req)

	var response taskResponse
	if err := a.do(req, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, responseError(response.Error, response.RequestID)
	}
	if response.TaskID == "" {
		return nil, fmt.Errorf("volc video enhancement returned an empty task_id")
	}

	status, err := normalizeStatus(response.Status)
	if err != nil {
		return nil, err
	}
	task := &dto.Task{
		TaskID:   response.TaskID,
		Status:   status,
		VideoURL: response.Result.VideoURL,
	}
	if status == dto.TaskStatusFailed {
		return task, responseError(response.Error, response.RequestID)
	}
	return task, nil
}

func normalizeStatus(status string) (dto.TaskStatus, error) {
	switch strings.ToLower(strings.TrimSpace(status)) {
	case "queued", "running", "processing":
		return dto.TaskStatusProcessing, nil
	case "completed":
		return dto.TaskStatusCompleted, nil
	case "failed":
		return dto.TaskStatusFailed, nil
	default:
		return "", fmt.Errorf("unknown volc video enhancement task status %q", status)
	}
}

func (a *Adapter) setHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.apiKey)
	if req.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
}

func (a *Adapter) do(req *http.Request, target any) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request volc video enhancement api: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read volc video enhancement response: %w", err)
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		var errorResponse struct {
			RequestID string    `json:"request_id"`
			Error     *apiError `json:"error"`
		}
		if json.Unmarshal(body, &errorResponse) == nil && errorResponse.Error != nil {
			return responseError(errorResponse.Error, errorResponse.RequestID)
		}
		return fmt.Errorf("volc video enhancement api returned status %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	if err = json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("decode volc video enhancement response: %w", err)
	}
	return nil
}

func responseError(err *apiError, requestID string) error {
	if err == nil {
		if requestID == "" {
			return fmt.Errorf("volc video enhancement request failed")
		}
		return fmt.Errorf("volc video enhancement request failed (request_id: %s)", requestID)
	}
	if requestID == "" {
		return err
	}
	return fmt.Errorf("%w (request_id: %s)", err, requestID)
}

func (e *apiError) Error() string {
	if e == nil {
		return ""
	}
	if e.Code == "" {
		return e.Message
	}
	if e.Message == "" {
		return e.Code
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
