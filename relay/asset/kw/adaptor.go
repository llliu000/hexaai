package kw

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
	Model      string `json:"model"`
	Moderation struct {
		Strategy string `json:"strategy"`
	} `json:"moderation"`
}

func (a *Adaptor) ReviewSkip() bool {
	return a.Moderation.Strategy == "Skip"
}

func (a *Adaptor) CreateAssetGroup(req *dto.CreateAssetGroupRequest) (*dto.CreateAssetGroupResponse, error) {
	request := kwCreateAssetGroupRequest{
		Model:       a.Model,
		Name:        req.Name,
		Description: req.Description,
		GroupType:   req.GroupType,
	}
	var response kwCreateResponse
	if err := a.post("CreateAssetGroup", request, &response); err != nil {
		return nil, err
	}
	return &dto.CreateAssetGroupResponse{Id: response.id()}, nil
}

func (a *Adaptor) ListAssetGroups(req *dto.ListAssetGroupsRequest) (*dto.ListAssetGroupsResponse, error) {
	request := kwListAssetGroupRequest{
		Model:      a.Model,
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
		SortBy:     req.SortBy,
		SortOrder:  req.SortOrder,
	}
	if req.Filter != nil {
		request.Filter.GroupType = req.Filter.GroupType
		request.Filter.Name = req.Filter.Name
	}

	var response kwListAssetGroupResponse
	if err := a.post("ListAssetGroups", request, &response); err != nil {
		return nil, err
	}
	items := make([]dto.GetAssetGroupResponse, 0, len(response.Items))
	for i := range response.Items {
		item := response.Items[i]
		items = append(items, dto.GetAssetGroupResponse{
			Id:         item.Id,
			Name:       item.Name,
			GroupType:  item.GroupType,
			CreateTime: item.CreateTime,
			UpdateTime: item.UpdateTime,
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
	request := kwCreateAssetRequest{
		Model:     a.Model,
		GroupId:   req.GroupId,
		URL:       req.URL,
		Name:      req.Name,
		AssetType: req.AssetType,
	}
	if req.Moderation != nil {
		request.Moderation = &kwModeration{
			Strategy: req.Moderation.Strategy,
		}
	}
	var response kwCreateResponse
	if err := a.post("CreateAsset", request, &response); err != nil {
		return nil, err
	}
	return &dto.CreateAssetResponse{Id: response.id()}, nil
}

func (a *Adaptor) GetAsset(req *dto.GetAssetRequest) (*dto.GetAssetResponse, error) {
	request := kwGetAssetRequest{
		Model: a.Model,
		Id:    req.Id,
	}
	var response kwAssetItem
	if err := a.post("GetAsset", request, &response); err != nil {
		return nil, err
	}
	return &dto.GetAssetResponse{
		Id:        response.Id,
		GroupId:   response.GroupId,
		URL:       response.URL,
		Name:      response.Name,
		AssetType: response.AssetType,
		Status:    response.Status,
	}, nil
}

func (a *Adaptor) ListAssets(req *dto.ListAssetsRequest) (*dto.ListAssetsResponse, error) {
	var filter *kwListAssetsFilter
	if req.Filter != nil {
		filter = &kwListAssetsFilter{
			GroupType: req.Filter.GroupType,
			GroupIds:  req.Filter.GroupIds,
			Statuses:  req.Filter.Statuses,
			Name:      req.Filter.Name,
		}
	}
	request := kwListAssetsRequest{
		Model:      a.Model,
		Filter:     filter,
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
		SortBy:     req.SortBy,
		SortOrder:  req.SortOrder,
	}
	var response kwListAssetsResponse
	if err := a.post("ListAssets", request, &response); err != nil {
		return nil, err
	}
	items := make([]dto.GetAssetResponse, 0, len(response.Items))
	for i := range response.Items {
		item := response.Items[i]
		items = append(items, dto.GetAssetResponse{
			Id:         item.Id,
			GroupId:    item.GroupId,
			URL:        item.URL,
			Name:       item.Name,
			AssetType:  item.AssetType,
			Status:     item.Status,
			CreateTime: item.CreatedTime,
			UpdateTime: item.UpdatedTime,
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
	reqUrl := fmt.Sprintf("%s/v3/open/%s", a.BaseURL, action)
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
		return fmt.Errorf("kw asset %s failed: status %d: %s", action, resp.StatusCode, strings.TrimSpace(string(respBody)))
	}
	if len(respBody) == 0 || out == nil {
		return nil
	}
	if err = common.Unmarshal(respBody, out); err != nil {
		return fmt.Errorf("kw asset %s decode response failed: %w", action, err)
	}
	return nil
}
