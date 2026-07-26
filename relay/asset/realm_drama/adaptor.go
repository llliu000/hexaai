package realm_drama

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
)

type AssetBaseResult[T any] struct {
	ResponseMetadata struct {
		RequestId string `json:"RequestId"`
		Action    string `json:"Action"`
		Version   string `json:"Version"`
		Service   string `json:"Service"`
		Region    string `json:"Region"`
		Error     struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error"`
	} `json:"ResponseMetadata"`
	Result T `json:"Result"`
}

type Adaptor struct {
	ApiKey      string `json:"api_key"`
	BaseURL     string `json:"base_url"`
	ProjectName string `json:"project_name"`
	Moderation  struct {
		Strategy string `json:"strategy"`
	} `json:"moderation"`
}

func (a *Adaptor) ReviewSkip() bool {
	return a.Moderation.Strategy == "Skip"
}

func (a *Adaptor) CreateAssetGroup(req *dto.CreateAssetGroupRequest) (*dto.CreateAssetGroupResponse, error) {
	req.ProjectName = &a.ProjectName
	var result AssetBaseResult[dto.CreateAssetGroupResponse]
	err := a.doCall("CreateAssetGroup", req, &result)
	return &result.Result, err
}

func (a *Adaptor) ListAssetGroups(req *dto.ListAssetGroupsRequest) (*dto.ListAssetGroupsResponse, error) {
	req.ProjectName = &a.ProjectName
	var result AssetBaseResult[dto.ListAssetGroupsResponse]
	err := a.doCall("ListAssetGroups", req, &result)
	return &result.Result, err
}

func (a *Adaptor) CreateAssets(req *dto.CreateAssetRequest) (*dto.CreateAssetResponse, error) {
	req.ProjectName = &a.ProjectName
	var result AssetBaseResult[dto.CreateAssetResponse]
	err := a.doCall("CreateAsset", req, &result)
	return &result.Result, err
}

func (a *Adaptor) GetAsset(req *dto.GetAssetRequest) (*dto.GetAssetResponse, error) {
	req.ProjectName = &a.ProjectName
	var result AssetBaseResult[dto.GetAssetResponse]
	err := a.doCall("GetAsset", req, &result)
	return &result.Result, err
}

func (a *Adaptor) ListAssets(req *dto.ListAssetsRequest) (*dto.ListAssetsResponse, error) {
	req.ProjectName = &a.ProjectName
	var result AssetBaseResult[dto.ListAssetsResponse]
	err := a.doCall("ListAssets", req, &result)
	return &result.Result, err
}

func (a *Adaptor) CreateVisualValidateSession(req *dto.CreateVisualValidateSessionRequest) (*dto.CreateVisualValidateSessionResponse, error) {
	req.ProjectName = &a.ProjectName
	var result AssetBaseResult[dto.CreateVisualValidateSessionResponse]
	err := a.doCall("CreateVisualValidateSession", req, &result)
	return &result.Result, err
}

func (a *Adaptor) GetVisualValidateResult(req *dto.GetVisualValidateResultRequest) (*dto.GetVisualValidateResultResponse, error) {
	req.ProjectName = &a.ProjectName
	var result AssetBaseResult[dto.GetVisualValidateResultResponse]
	err := a.doCall("GetVisualValidateResult", req, &result)
	if err != nil {
		return nil, err
	}
	if message := result.ResponseMetadata.Error.Message; message != "" {
		return nil, errors.New(message)
	}
	return &result.Result, nil
}

func (a *Adaptor) doCall(action string, request, response any) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)
	baseUrl := fmt.Sprintf("%s/?Action=%s&Version=2024-01-01", a.BaseURL, action)
	httpReq, err := http.NewRequest(http.MethodPost, baseUrl, reader)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Authorization", "Bearer "+a.ApiKey)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(respBody) == 0 || response == nil {
		return nil
	}
	if err = common.Unmarshal(respBody, response); err != nil {
		return fmt.Errorf("asset decode response failed: %w", err)
	}
	return nil
}
