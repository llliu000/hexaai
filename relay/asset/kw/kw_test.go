package kw

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/QuantumNous/new-api/dto"
)

func TestKwAssetFullFlow(t *testing.T) {
	apiKey := os.Getenv("KW_TEST_API_KEY")
	if apiKey == "" {
		t.Skip("set KW_TEST_API_KEY to run integration test")
	}
	baseURL := os.Getenv("KW_TEST_BASE_URL")
	if baseURL == "" {
		baseURL = "https://modeltop.ai"
	}
	model := os.Getenv("KW_TEST_MODEL")
	if model == "" {
		model = "doubao-seedance-2-0-fast"
	}

	adaptor := Adaptor{ApiKey: apiKey, BaseURL: baseURL, Model: model}
	suffix := time.Now().Format("20060102150405")
	groupName := fmt.Sprintf("new-api-test-group-%s", suffix)
	groupDescription := "new-api kw asset integration test"
	groupType := "AIGC"

	createGroupResp, err := adaptor.CreateAssetGroup(&dto.CreateAssetGroupRequest{
		Name:        groupName,
		Description: ptr(groupDescription),
		GroupType:   groupType,
	})
	if err != nil {
		t.Fatalf("CreateAssetGroup failed: %v", err)
	}
	if createGroupResp == nil || createGroupResp.Id == "" {
		t.Fatalf("CreateAssetGroup returned empty response: %#v", createGroupResp)
	}

	listGroupsResp, err := adaptor.ListAssetGroups(&dto.ListAssetGroupsRequest{
		Filter: &dto.ListAssetGroupsFilter{
			Name:      ptr(groupName),
			GroupType: ptr(groupType),
			GroupIds:  []string{createGroupResp.Id},
		},
		PageNumber: 1,
		PageSize:   10,
		SortBy:     "CreateTime",
		SortOrder:  "Desc",
	})
	if err != nil {
		t.Fatalf("ListAssetGroups failed: %v", err)
	}
	assertAssetGroupListed(t, listGroupsResp, createGroupResp.Id, groupName, groupType)

	assetName := fmt.Sprintf("new-api-test-asset-%s", suffix)
	assetURL := os.Getenv("KW_TEST_ASSET_URL")
	if assetURL == "" {
		assetURL = "https://b0.bdstatic.com/ugc/img/2024-12-28/16a7486e93dc56ef5aba9e0aa6144e01.png"
	}
	assetType := "Image"
	createAssetResp, err := adaptor.CreateAssets(&dto.CreateAssetRequest{
		GroupId:   createGroupResp.Id,
		URL:       assetURL,
		Name:      assetName,
		AssetType: assetType,
	})
	if err != nil {
		t.Fatalf("CreateAssets failed: %v", err)
	}
	if createAssetResp == nil || createAssetResp.Id == "" {
		t.Fatalf("CreateAssets returned empty response: %#v", createAssetResp)
	}

	getAssetResp, err := adaptor.GetAsset(&dto.GetAssetRequest{Id: createAssetResp.Id})
	if err != nil {
		t.Fatalf("GetAsset failed: %v", err)
	}
	assertAsset(t, getAssetResp, createAssetResp.Id, createGroupResp.Id, assetName, assetType)

	listAssetsResp, err := adaptor.ListAssets(&dto.ListAssetsRequest{
		Filter: &dto.ListAssetsFilter{
			GroupIds:  []string{createGroupResp.Id},
			GroupType: ptr(groupType),
			Name:      ptr(assetName),
		},
		PageNumber: 1,
		PageSize:   10,
		SortBy:     "CreateTime",
		SortOrder:  "Desc",
	})
	if err != nil {
		t.Fatalf("ListAssets failed: %v", err)
	}
	assertAssetListed(t, listAssetsResp, createAssetResp.Id, createGroupResp.Id, assetName, assetType)
}

func assertAssetGroupListed(t *testing.T, resp *dto.ListAssetGroupsResponse, groupID string, groupName string, groupType string) {
	t.Helper()
	if resp == nil {
		t.Fatal("ListAssetGroups returned nil response")
	}
	for _, item := range resp.Items {
		if item.Id == groupID {
			if item.Name != "" && item.Name != groupName {
				t.Fatalf("listed asset group name mismatch: got %q, want %q", item.Name, groupName)
			}
			if item.GroupType != "" && !strings.EqualFold(item.GroupType, groupType) {
				t.Fatalf("listed asset group type mismatch: got %q, want %q", item.GroupType, groupType)
			}
			return
		}
	}
	t.Fatalf("asset group %q not found in list response: %#v", groupID, resp)
}

func assertAssetListed(t *testing.T, resp *dto.ListAssetsResponse, assetID string, groupID string, assetName string, assetType string) {
	t.Helper()
	if resp == nil {
		t.Fatal("ListAssets returned nil response")
	}
	for _, item := range resp.Items {
		if item.Id == assetID {
			assertAsset(t, &item, assetID, groupID, assetName, assetType)
			return
		}
	}
	t.Fatalf("asset %q not found in list response: %#v", assetID, resp)
}

func assertAsset(t *testing.T, resp *dto.GetAssetResponse, assetID string, groupID string, assetName string, assetType string) {
	t.Helper()
	if resp == nil {
		t.Fatal("asset response is nil")
	}
	if resp.Id != assetID {
		t.Fatalf("asset id mismatch: got %q, want %q", resp.Id, assetID)
	}
	if resp.GroupId != "" && resp.GroupId != groupID {
		t.Fatalf("asset group id mismatch: got %q, want %q", resp.GroupId, groupID)
	}
	if resp.Name != "" && resp.Name != assetName {
		t.Fatalf("asset name mismatch: got %q, want %q", resp.Name, assetName)
	}
	if resp.AssetType != "" && !strings.EqualFold(resp.AssetType, assetType) {
		t.Fatalf("asset type mismatch: got %q, want %q", resp.AssetType, assetType)
	}
}

func ptr(value string) *string {
	return &value
}
