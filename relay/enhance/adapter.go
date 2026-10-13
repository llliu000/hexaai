package enhance

import (
	"context"
	"errors"
	"fmt"

	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/relay/enhance/volc"
	"github.com/QuantumNous/new-api/setting/model_setting"
)

// Adapter defines the provider-independent video enhancement workflow.
type Adapter interface {
	Submit(ctx context.Context, request dto.SubmitRequest) (*dto.SubmitResult, error)
	// GetTask returns a failed Task together with a non-nil error when the
	// provider task reaches a failed state. Request errors return a nil Task.
	GetTask(ctx context.Context, taskID string) (*dto.Task, error)
}

type retryableError interface {
	error
	Retryable() bool
}

func IsRetryable(err error) bool {
	var target retryableError
	return errors.As(err, &target) && target.Retryable()
}

func GetAdapter() (Adapter, error) {
	setting := model_setting.GetEnhanceSetting()
	switch setting.Channel {
	case "volc":
		return volc.New(volc.Config{
			APIKey:      setting.ApiKey,
			BaseURL:     setting.BaseUrl,
			ToolVersion: setting.ToolVersion,
		})
	}
	return nil, fmt.Errorf("unsupported channel: %s", setting.Channel)
}
