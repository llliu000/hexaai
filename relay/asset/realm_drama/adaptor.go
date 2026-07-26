package realm_drama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
)

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
	var result dto.CreateAssetGroupResponse
	req.ProjectName = &a.ProjectName
	err := a.doCall("CreateAssetGroup", req, &result)
	return &result, err
}

func (a *Adaptor) ListAssetGroups(req *dto.ListAssetGroupsRequest) (*dto.ListAssetGroupsResponse, error) {
	var result dto.ListAssetGroupsResponse
	req.ProjectName = &a.ProjectName
	err := a.doCall("ListAssetGroups", req, &result)
	return &result, err
}

func (a *Adaptor) CreateAssets(req *dto.CreateAssetRequest) (*dto.CreateAssetResponse, error) {
	var result dto.CreateAssetResponse
	req.ProjectName = &a.ProjectName
	err := a.doCall("CreateAsset", req, &result)
	return &result, err
}

func (a *Adaptor) GetAsset(req *dto.GetAssetRequest) (*dto.GetAssetResponse, error) {
	var result dto.GetAssetResponse
	req.ProjectName = &a.ProjectName
	err := a.doCall("GetAsset", req, &result)
	return &result, err
}

func (a *Adaptor) ListAssets(req *dto.ListAssetsRequest) (*dto.ListAssetsResponse, error) {
	var result dto.ListAssetsResponse
	req.ProjectName = &a.ProjectName
	err := a.doCall("ListAssets", req, &result)
	return &result, err
}

func (a *Adaptor) CreateVisualValidateSession(req *dto.CreateVisualValidateSessionRequest) (*dto.CreateVisualValidateSessionResponse, error) {
	var result dto.CreateVisualValidateSessionResponse
	req.ProjectName = &a.ProjectName
	err := a.doCall("CreateVisualValidateSession", req, &result)
	return &result, err
}

func (a *Adaptor) GetVisualValidateResult(req *dto.GetVisualValidateResultRequest) (*dto.GetVisualValidateResultResponse, error) {
	var result dto.GetVisualValidateResultResponse
	req.ProjectName = &a.ProjectName
	err := a.doCall("GetVisualValidateResult", req, &result)
	return &result, err
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
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("one asset failed: status %d: %s", resp.StatusCode, strings.TrimSpace(string(respBody)))
	}
	if len(respBody) == 0 || response == nil {
		return nil
	}
	if err = common.Unmarshal(respBody, response); err != nil {
		return fmt.Errorf("asset decode response failed: %w", err)
	}
	return nil
}
