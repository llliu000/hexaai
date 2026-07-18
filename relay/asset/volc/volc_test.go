package volc

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/QuantumNous/new-api/dto"
)

func TestVolcAssetFullFlow(t *testing.T) {
	accessID := os.Getenv("VOLC_TEST_ACCESS_ID")
	secretKey := os.Getenv("VOLC_TEST_SECRET_KEY")
	if accessID == "" || secretKey == "" {
		t.Skip("set VOLC_TEST_ACCESS_ID and VOLC_TEST_SECRET_KEY to run integration test")
	}
	baseUrl := os.Getenv("VOLC_TEST_BASE_URL")
	if baseUrl == "" {
		baseUrl = "https://openai.pixmax.cn/ai-api/volcengine/openapi"
	}
	projectName := os.Getenv("VOLC_TEST_PROJECT_NAME")
	if projectName == "" {
		projectName = "default"
	}

	adaptor := Adaptor{AccessID: accessID, SecretKey: secretKey, BaseUrl: baseUrl, ProjectName: projectName}
	suffix := time.Now().Format("20060102150405")

	groupName := fmt.Sprintf("new-api-test-group-%s", "suffix")
	groupDescription := "new-api volc asset integration test"
	groupType := "AIGC"

	createGroupResp, err := adaptor.CreateAssetGroup(&dto.CreateAssetGroupRequest{
		BaseAssetRequest: dto.BaseAssetRequest{
			ProjectName: optionalPtr(projectName),
		},
		Description: ptr(groupDescription),
		Name:        groupName,
		GroupType:   groupType,
	})
	if err != nil {
		t.Fatalf("CreateAssetGroup failed: %v", err)
	}
	if createGroupResp == nil || createGroupResp.Id == "" {
		t.Fatalf("CreateAssetGroup returned empty response: %#v", createGroupResp)
	}

	assetName := fmt.Sprintf("new-api-test-asset-%s", suffix)
	assetURL := "https://b0.bdstatic.com/ugc/img/2024-12-28/16a7486e93dc56ef5aba9e0aa6144e01.png"
	assetType := "Image"
	createAssetResp, err := adaptor.CreateAssets(&dto.CreateAssetRequest{
		BaseAssetRequest: dto.BaseAssetRequest{
			ProjectName: optionalPtr(projectName),
		},
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

	getAssetResp, err := adaptor.GetAsset(&dto.GetAssetRequest{
		BaseAssetRequest: dto.BaseAssetRequest{
			ProjectName: optionalPtr(projectName),
		},
		Id: createAssetResp.Id,
	})
	if err != nil {
		t.Fatalf("GetAsset failed: %v", err)
	}
	assertAsset(t, getAssetResp, createAssetResp.Id, createGroupResp.Id, assetName, assetType)

	listAssetsResp, err := adaptor.ListAssets(&dto.ListAssetsRequest{
		BaseAssetRequest: dto.BaseAssetRequest{
			ProjectName: optionalPtr(projectName),
		},
		Filter: &dto.ListAssetsFilter{
			GroupIds:  []string{createGroupResp.Id},
			GroupType: ptr(groupType),
			Name:      ptr(assetName),
		},
		PageNumber: 1,
		PageSize:   10,
	})
	if err != nil {
		t.Fatalf("ListAssets failed: %v", err)
	}
	assertAssetListed(t, listAssetsResp, createAssetResp.Id, createGroupResp.Id, assetName, assetType)
}

func assertAssetGroupListed(t *testing.T, resp *dto.ListAssetGroupsResponse, groupID string, groupName string) {
	t.Helper()
	if resp == nil {
		t.Fatal("ListAssetGroups returned nil response")
	}
	for _, item := range resp.Items {
		if item.Id == groupID {
			if item.Name != "" && item.Name != groupName {
				t.Fatalf("listed asset group name mismatch: got %q, want %q", item.Name, groupName)
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

func optionalPtr(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
