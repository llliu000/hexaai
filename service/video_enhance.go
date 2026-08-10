package service

import (
	"context"
	"fmt"

	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/relay/enhance"
	"github.com/QuantumNous/new-api/relay/enhance/volc"
)

const videoEnhanceTargetResolution = "1080p"

func submitEnhanceTask(ctx context.Context, channelType int, apiKey, proxy, videoURL string) (string, error) {
	adapter, err := newVideoEnhanceAdapter(channelType, apiKey, proxy)
	if err != nil {
		return "", err
	}
	response, err := adapter.Submit(ctx, enhance.SubmitRequest{
		VideoURL:         videoURL,
		TargetResolution: videoEnhanceTargetResolution,
	})
	if err != nil {
		return "", err
	}
	return response.TaskID, nil
}

func getEnhanceTask(ctx context.Context, channelType int, apiKey, proxy, taskID string) (*enhance.Task, error) {
	adapter, err := newVideoEnhanceAdapter(channelType, apiKey, proxy)
	if err != nil {
		return nil, err
	}
	return adapter.GetTask(ctx, taskID)
}

func newVideoEnhanceAdapter(channelType int, apiKey, proxy string) (enhance.Adapter, error) {
	client, err := GetHttpClientWithProxy(proxy)
	if err != nil {
		return nil, fmt.Errorf("create video enhancement http client: %w", err)
	}
	switch channelType {
	case constant.ChannelTypeVolcEngine, constant.ChannelTypeDoubaoVideo:
		return volc.New(volc.Config{
			APIKey: apiKey,
			Client: client,
		})
	default:
		return nil, fmt.Errorf("video enhancement is not supported for channel type %d", channelType)
	}
}
