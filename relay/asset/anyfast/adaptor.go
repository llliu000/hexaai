package anyfast

import (
	"bytes"
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
	request := createAssetGroupRequest{
		Model: "volc-asset",
		Name:  req.Name,
	}
	var response createResponse
	if err := a.post("CreateAssetGroup", request, &response); err != nil {
		return nil, err
	}
	return &dto.CreateAssetGroupResponse{Id: response.Id}, nil
}

func (a *Adaptor) ListAssetGroups(req *dto.ListAssetGroupsRequest) (*dto.ListAssetGroupsResponse, error) {
	request := listAssetGroupRequest{
		Model:      "volc-asset",
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
	}
	if req.Filter != nil {
		request.Filter.GroupType = req.Filter.GroupType
		request.Filter.GroupIds = req.Filter.GroupIds
		request.Filter.Name = req.Filter.Name
	}

	var response listAssetGroupResponse
	if err := a.post("ListAssetGroups", request, &response); err != nil {
		return nil, err
	}
	items := make([]dto.GetAssetGroupResponse, 0, len(response.Items))
	for i := range response.Items {
		item := response.Items[i]
		items = append(items, dto.GetAssetGroupResponse{
			Id:        item.Id,
			Name:      item.Name,
			GroupType: item.GroupType,
		})
	}
	return &dto.ListAssetGroupsResponse{
		Items:      items,
		TotalCount: response.TotalCount,
		PageNumber: response.PageNumber,
		PageSize:   response.PageSize,
	}, nil
}

func (a *Adaptor) CreateAssets(req *dto.CreateAssetRequest) (*dto.CreateAssetResponse, error) {
	request := createAssetRequest{
		GroupId:   req.GroupId,
		URL:       req.URL,
		Name:      req.Name,
		AssetType: req.AssetType,
	}
	switch strings.ToLower(req.AssetType) {
	case "video":
		request.Model = "volc-asset-video"
	case "audio":
		request.Model = "volc-asset-audio"
	default:
		request.Model = "volc-asset"
	}
	var response createResponse
	if err := a.post("CreateAsset", request, &response); err != nil {
		return nil, err
	}
	return &dto.CreateAssetResponse{Id: response.Id}, nil
}

func (a *Adaptor) GetAsset(req *dto.GetAssetRequest) (*dto.GetAssetResponse, error) {
	var gt = "AIGC"
	request := dto.ListAssetsRequest{
		Filter: &dto.ListAssetsFilter{
			GroupIds:  []string{*req.ProjectName},
			GroupType: &gt,
		},
		PageNumber: 1,
		PageSize:   100,
		SortBy:     "CreateTime",
		SortOrder:  "Desc",
	}
	assets, err := a.ListAssets(&request)
	if err != nil {
		return nil, err
	}
	var asset dto.GetAssetResponse
	for i := range assets.Items {
		if assets.Items[i].Id == req.Id {
			asset.AssetType = assets.Items[i].AssetType
			asset.Status = assets.Items[i].Status
			asset.Name = assets.Items[i].Name
			asset.URL = assets.Items[i].URL
			asset.Id = assets.Items[i].Id
			break
		}
	}
	return &asset, nil
}

func (a *Adaptor) ListAssets(req *dto.ListAssetsRequest) (*dto.ListAssetsResponse, error) {
	var filter *listAssetsFilter
	if req.Filter != nil {
		filter = &listAssetsFilter{
			Name:      req.Filter.Name,
			GroupIds:  req.Filter.GroupIds,
			GroupType: req.Filter.GroupType,
		}
	}
	request := listAssetsRequest{
		Model:      "volc-asset",
		Filter:     filter,
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
	}
	var response listAssetsResponse
	if err := a.post("ListAssets", request, &response); err != nil {
		return nil, err
	}
	items := make([]dto.GetAssetResponse, 0, len(response.Items))
	for i := range response.Items {
		item := response.Items[i]
		items = append(items, dto.GetAssetResponse{
			Id:          item.Id,
			GroupId:     item.GroupId,
			URL:         item.URL,
			Name:        item.Name,
			AssetType:   item.AssetType,
			ProjectName: item.ProjectName,
			Status:      item.Status,
			CreateTime:  item.CreateTime,
			UpdateTime:  item.UpdateTime,
		})
	}
	return &dto.ListAssetsResponse{
		Items:      items,
		TotalCount: response.TotalCount,
		PageNumber: response.PageNumber,
		PageSize:   response.PageSize,
	}, nil
}

func (a *Adaptor) post(action string, payload any, out any) error {
	body, err := common.Marshal(payload)
	if err != nil {
		return err
	}
	reqUrl := fmt.Sprintf("%s/volc/asset/%s", a.BaseURL, action)
	httpReq, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewReader(body))
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
		return fmt.Errorf("anyfast asset %s failed: status %d: %s", action, resp.StatusCode, strings.TrimSpace(string(respBody)))
	}
	if len(respBody) == 0 || out == nil {
		return nil
	}
	if err = common.Unmarshal(respBody, out); err != nil {
		return fmt.Errorf("anyfast asset %s decode response failed: %w", action, err)
	}
	return nil
}
