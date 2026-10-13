package volc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/dto"
	"github.com/google/uuid"
)

const (
	DefaultBaseURL      = "https://mediakit.cn-beijing.volces.com"
	defaultToolVersion  = "standard"
	defaultScene        = "aigc"
	defaultBitrateLevel = "medium"
	maxRequestAttempts  = 3
)

type Config struct {
	APIKey      string
	BaseURL     string
	ToolVersion string
}

type Adapter struct {
	apiKey      string
	baseURL     string
	toolVersion string
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
	if config.BaseURL == "" {
		config.BaseURL = DefaultBaseURL
	}
	if config.ToolVersion == "" {
		config.ToolVersion = defaultToolVersion
	}
	return &Adapter{
		toolVersion: config.ToolVersion,
		baseURL:     config.BaseURL,
		apiKey:      config.APIKey,
	}, nil
}

func (a *Adapter) Submit(ctx context.Context, request dto.SubmitRequest) (*dto.SubmitResult, error) {
	payload := submitRequest{
		Resolution:   request.TargetResolution,
		BitrateLevel: defaultBitrateLevel,
		VideoURL:     request.VideoURL,
	}
	if request.ClientToken != "" {
		payload.ClientToken = request.ClientToken
	} else {
		payload.ClientToken = uuid.NewString()
	}
	var uri string
	switch a.toolVersion {
	case "standard", "professional":
		uri = "/api/v1/tools/enhance-video"
		payload.ToolVersion = a.toolVersion
		payload.Scene = defaultScene
	case "generative":
		uri = "/api/v1/tools/enhance-video-generative"
	case "fast":
		uri = "/api/v1/tools/enhance-video-fast"
	default:
		return nil, fmt.Errorf("unsupported volc video enhancement tool version %q", a.toolVersion)
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	reqUrl := a.baseURL + uri
	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost, reqUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	a.setHeaders(req)

	var response submitResponse
	if err = a.do(req, &response); err != nil {
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
	reqUrl := a.baseURL + "/api/v1/tasks/" + taskID
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	a.setHeaders(req)

	var response taskResponse
	if err = a.do(req, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, responseError(response.Error, response.RequestID)
	}

	status, err := normalizeStatus(response.Status)
	if err != nil {
		return nil, err
	}
	task := &dto.Task{
		VideoURL: response.Result.VideoURL,
		TaskID:   response.TaskID,
		Status:   status,
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
	var lastErr error
	for attempt := 0; attempt < maxRequestAttempts; attempt++ {
		attemptReq := req
		if attempt > 0 && req.Body != nil {
			if req.GetBody == nil {
				return fmt.Errorf("volc video enhancement request body cannot be replayed")
			}
			body, err := req.GetBody()
			if err != nil {
				return fmt.Errorf("rebuild volc video enhancement request body: %w", err)
			}
			attemptReq = req.Clone(req.Context())
			attemptReq.Body = body
		}

		resp, err := http.DefaultClient.Do(attemptReq)
		if err != nil {
			if req.Context().Err() != nil {
				return err
			}
			lastErr = err
		} else {
			body, readErr := io.ReadAll(resp.Body)
			_ = resp.Body.Close()
			if readErr != nil {
				if req.Context().Err() != nil {
					return readErr
				}
				lastErr = readErr
			} else if isRetryableStatus(resp.StatusCode) {
				lastErr = responseStatusError(resp.StatusCode, body)
			} else if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
				var errorResponse struct {
					RequestID string    `json:"request_id"`
					Error     *apiError `json:"error"`
				}
				if json.Unmarshal(body, &errorResponse) == nil && errorResponse.Error != nil {
					return responseError(errorResponse.Error, errorResponse.RequestID)
				}
				return fmt.Errorf("volc video enhancement api returned status %d: %s", resp.StatusCode, string(body))
			} else {
				return json.Unmarshal(body, target)
			}
		}

		if attempt+1 < maxRequestAttempts {
			if err := waitBeforeRetry(req.Context(), attempt); err != nil {
				return err
			}
			continue
		}
	}
	return &retryableAPIError{err: lastErr}
}

func isRetryableStatus(statusCode int) bool {
	return statusCode == http.StatusRequestTimeout ||
		statusCode == http.StatusTooEarly ||
		statusCode == http.StatusTooManyRequests ||
		statusCode >= http.StatusInternalServerError
}

func responseStatusError(statusCode int, body []byte) error {
	var errorResponse struct {
		RequestID string    `json:"request_id"`
		Error     *apiError `json:"error"`
	}
	if json.Unmarshal(body, &errorResponse) == nil && errorResponse.Error != nil {
		return responseError(errorResponse.Error, errorResponse.RequestID)
	}
	return fmt.Errorf("volc video enhancement api returned status %d: %s", statusCode, string(body))
}

func waitBeforeRetry(ctx context.Context, attempt int) error {
	delay := time.Duration(attempt+1) * 500 * time.Millisecond
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

type retryableAPIError struct {
	err error
}

func (e *retryableAPIError) Error() string {
	if e == nil || e.err == nil {
		return "volc video enhancement request failed after retries"
	}
	return fmt.Sprintf("volc video enhancement request failed after retries: %v", e.err)
}

func (e *retryableAPIError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

func (e *retryableAPIError) Retryable() bool {
	return true
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
