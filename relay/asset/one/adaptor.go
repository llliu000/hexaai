package one

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
	ApiKey     string `json:"api_key"`
	BaseURL    string `json:"base_url"`
	Moderation struct {
		Strategy string `json:"strategy"`
	} `json:"moderation"`
}

func (a *Adaptor) ReviewSkip() bool {
	return a.Moderation.Strategy == "Skip"
}

func (a *Adaptor) CreateAssetGroup(req *dto.CreateAssetGroupRequest) (*dto.CreateAssetGroupResponse, error) {
	return &dto.CreateAssetGroupResponse{Id: "group-20250713223312-jnqvA"}, nil
}

func (a *Adaptor) ListAssetGroups(req *dto.ListAssetGroupsRequest) (*dto.ListAssetGroupsResponse, error) {
	return &dto.ListAssetGroupsResponse{}, nil
}

func (a *Adaptor) CreateAssets(req *dto.CreateAssetRequest) (*dto.CreateAssetResponse, error) {
	request := createAssetRequest{
		URL:       req.URL,
		Name:      req.Name,
		AssetType: req.AssetType,
	}
	marshal, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	reqUrl := fmt.Sprintf("%s/v1/sd/assets", a.BaseURL)
	var response Response[createResponse]
	reader := bytes.NewReader(marshal)
	if err = a.doRequest(http.MethodPost, reqUrl, reader, &response); err != nil {
		return nil, err
	}
	return &dto.CreateAssetResponse{Id: response.Data.Id}, nil
}

func (a *Adaptor) GetAsset(req *dto.GetAssetRequest) (*dto.GetAssetResponse, error) {
	reqUrl := fmt.Sprintf("%s/v1/sd/assets/%s", a.BaseURL, req.Id)
	var resp Response[assetResponse]
	if err := a.doRequest(http.MethodGet, reqUrl, nil, &resp); err != nil {
		return nil, err
	}
	ar := dto.GetAssetResponse{
		Id:          resp.Data.Id,
		GroupId:     resp.Data.GroupId,
		URL:         resp.Data.URL,
		Name:        resp.Data.Name,
		AssetType:   resp.Data.AssetType,
		ProjectName: resp.Data.ProjectName,
		Status:      resp.Data.Status,
		CreateTime:  resp.Data.CreateTime,
		UpdateTime:  resp.Data.UpdateTime,
	}
	if resp.Data.StatusCode != 0 || resp.Data.StatusMsg != "" {
		marshal, _ := json.Marshal(map[string]string{
			"Code":    fmt.Sprintf("%d", resp.Data.StatusCode),
			"Message": resp.Data.StatusMsg,
		})
		raw := json.RawMessage(marshal)
		ar.Error = &raw
	}
	return &ar, nil
}

func (a *Adaptor) ListAssets(req *dto.ListAssetsRequest) (*dto.ListAssetsResponse, error) {
	return &dto.ListAssetsResponse{}, nil
}

func (a *Adaptor) doRequest(method, url string, body io.Reader, out any) error {
	httpReq, err := http.NewRequest(method, url, body)
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
	if len(respBody) == 0 || out == nil {
		return nil
	}
	if err = common.Unmarshal(respBody, out); err != nil {
		return fmt.Errorf("one asset decode response failed: %w", err)
	}
	return nil
}
