package cme

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/QuantumNous/new-api/dto"
)

// 接口测试需要配置 CME_TEST_ACCESS_KEY、CME_TEST_SECRET_KEY。
// CreateAssets 还需要 CME_TEST_GROUP_ID、CME_TEST_ASSET_URL；
// GetAsset 需要 CME_TEST_ASSET_ID；CreateVisualValidateSession 需要
// CME_TEST_CALLBACK_URL；GetVisualValidateResult 需要已完成人脸核验的
// CME_TEST_BYTED_TOKEN。未配置必需参数时，对应用例会跳过。
// 使用 go test -v ./relay/asset/cme 查看每个方法的返回结果和错误信息。
// 创建类用例会在真实服务上创建数据；各用例独立，不依赖其他用例的执行顺序。

func TestCMEAdaptorReviewSkip(t *testing.T) {
	adaptor := &Adaptor{}
	result := adaptor.ReviewSkip()
	cmeTestLogResult(t, "ReviewSkip", result, nil)
	if result {
		t.Fatal("ReviewSkip should return false")
	}
}

func TestCMEAdaptorCreateAssetGroup(t *testing.T) {
	adaptor := cmeTestAdaptor(t)
	description := "new-api cme asset integration test"
	req := &dto.CreateAssetGroupRequest{
		Name:        cmeTestEnv("CME_TEST_GROUP_NAME", cmeTestName("group")),
		Description: &description,
		GroupType:   "AIGC",
	}

	result, err := adaptor.CreateAssetGroup(req)
	cmeTestLogResult(t, "CreateAssetGroup", result, err)
	if result == nil || result.Id == "" {
		t.Fatalf("CreateAssetGroup returned an empty group ID: %#v", result)
	}
}

func TestCMEAdaptorCreateAssets(t *testing.T) {
	adaptor := cmeTestAdaptor(t)
	req := &dto.CreateAssetRequest{
		GroupId:   "group-20261118151821-b9wkd",
		URL:       "https://images.pexels.com/photos/13039769/pexels-photo-13039769.jpeg",
		Name:      cmeTestName("asset"),
		AssetType: "Image",
	}

	result, err := adaptor.CreateAssets(req)
	cmeTestLogResult(t, "CreateAssets", result, err)
	if result == nil || result.Id == "" {
		t.Fatalf("CreateAssets returned an empty asset ID: %#v", result)
	}
}

func TestCMEAdaptorGetAsset(t *testing.T) {
	adaptor := cmeTestAdaptor(t)
	req := &dto.GetAssetRequest{Id: "asset-20261118152039-hlxqw"}

	result, err := adaptor.GetAsset(req)
	cmeTestLogResult(t, "GetAsset", result, err)
	if result == nil {
		t.Fatal("GetAsset returned a nil response")
	}
	if result.Id != req.Id {
		t.Errorf("asset ID: got %q, want %q", result.Id, req.Id)
	}
	if result.GroupId == "" {
		t.Error("GetAsset returned an empty group ID")
	}
	cmeTestAssertAssetStatus(t, result.Status)
}

func TestCMEAdaptorListAssets(t *testing.T) {
	adaptor := cmeTestAdaptor(t)
	s := "AIGC"
	filter := &dto.ListAssetsFilter{
		GroupType: &s,
		Name:      cmeTestOptionalEnv("CME_TEST_ASSET_NAME"),
	}
	if groupID := os.Getenv("CME_TEST_GROUP_ID"); groupID != "" {
		filter.GroupIds = []string{groupID}
	}
	if status := os.Getenv("CME_TEST_ASSET_STATUS"); status != "" {
		filter.Statuses = []string{status}
	}
	req := &dto.ListAssetsRequest{Filter: filter, PageNumber: 1, PageSize: 10}

	result, err := adaptor.ListAssets(req)
	cmeTestLogResult(t, "ListAssets", result, err)
	if result == nil {
		t.Fatal("ListAssets returned a nil response")
	}
	cmeTestAssertPagination(t, result.TotalCount, result.PageNumber, result.PageSize,
		len(result.Items), req.PageNumber, req.PageSize)
	for i, item := range result.Items {
		if item.Id == "" {
			t.Errorf("Items[%d] has an empty asset ID", i)
		}
		if len(filter.GroupIds) > 0 && item.GroupId != filter.GroupIds[0] {
			t.Errorf("Items[%d].GroupId: got %q, want %q", i, item.GroupId, filter.GroupIds[0])
		}
		cmeTestAssertAssetStatus(t, item.Status)
	}
}

func TestCMEAdaptorListAssetGroups(t *testing.T) {
	s := "AIGC"
	adaptor := cmeTestAdaptor(t)
	filter := &dto.ListAssetGroupsFilter{
		GroupType: &s,
		Name:      cmeTestOptionalEnv("CME_TEST_GROUP_NAME"),
	}
	if groupID := os.Getenv("CME_TEST_GROUP_ID"); groupID != "" {
		filter.GroupIds = []string{groupID}
	}
	req := &dto.ListAssetGroupsRequest{Filter: filter, PageNumber: 1, PageSize: 10}

	result, err := adaptor.ListAssetGroups(req)
	cmeTestLogResult(t, "ListAssetGroups", result, err)
	if result == nil {
		t.Fatal("ListAssetGroups returned a nil response")
	}
	cmeTestAssertPagination(t, result.TotalCount, result.PageNumber, result.PageSize,
		len(result.Items), req.PageNumber, req.PageSize)
	for i, item := range result.Items {
		if item.Id == "" {
			t.Errorf("Items[%d] has an empty group ID", i)
		}
		if len(filter.GroupIds) > 0 && item.Id != filter.GroupIds[0] {
			t.Errorf("Items[%d].Id: got %q, want %q", i, item.Id, filter.GroupIds[0])
		}
		if filter.GroupType != nil && item.GroupType != *filter.GroupType {
			t.Errorf("Items[%d].GroupType: got %q, want %q", i, item.GroupType, *filter.GroupType)
		}
	}
}

func TestCMEAdaptorCreateVisualValidateSession(t *testing.T) {
	adaptor := cmeTestAdaptor(t)
	req := &dto.CreateVisualValidateSessionRequest{
		CallbackURL: cmeTestRequiredEnv(t, "CME_TEST_CALLBACK_URL"),
	}

	result, err := adaptor.CreateVisualValidateSession(req)
	cmeTestLogResult(t, "CreateVisualValidateSession", result, err)
	if result == nil {
		t.Fatal("CreateVisualValidateSession returned a nil response")
	}
	if result.BytedToken == "" {
		t.Error("CreateVisualValidateSession returned an empty BytedToken")
	}
	if result.H5Link == "" {
		t.Error("CreateVisualValidateSession returned an empty H5Link")
	}
	if result.CallbackURL != req.CallbackURL {
		t.Errorf("CallbackURL: got %q, want %q", result.CallbackURL, req.CallbackURL)
	}
}

func TestCMEAdaptorGetVisualValidateResult(t *testing.T) {
	adaptor := cmeTestAdaptor(t)
	req := &dto.GetVisualValidateResultRequest{
		BytedToken: cmeTestRequiredEnv(t, "CME_TEST_BYTED_TOKEN"),
	}

	result, err := adaptor.GetVisualValidateResult(req)
	cmeTestLogResult(t, "GetVisualValidateResult", result, err)
	if result == nil || result.GroupId == "" {
		t.Fatalf("GetVisualValidateResult returned an empty group ID: %#v", result)
	}
}

func TestCMEAdaptorClient(t *testing.T) {
	// client 仅构造 SDK 客户端，不发送请求，也不需要真实密钥。
	adaptor := &Adaptor{AccessKey: "test-access-key", SecretKey: "test-secret-key"}
	client := adaptor.client()
	if client == nil {
		t.Fatal("client returned nil")
	}
	t.Logf("client result: initialized %T", client)
}

func TestCMECapitalize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: ""},
		{name: "active", input: "ACTIVE", want: "Active"},
		{name: "processing", input: "PROCESSING", want: "Processing"},
		{name: "failed", input: "FAILED", want: "Failed"},
		{name: "lowercase", input: "active", want: "Active"},
		{name: "mixed_case", input: "pRoCeSsInG", want: "Processing"},
		{name: "capitalized", input: "Active", want: "Active"},
		{name: "single_character", input: "a", want: "A"},
		{name: "leading_digit", input: "1ACTIVE", want: "1active"},
		{name: "leading_space", input: " ACTIVE", want: " active"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := capitalize(tt.input)
			t.Logf("capitalize(%q) result: %q, want: %q", tt.input, result, tt.want)
			if result != tt.want {
				t.Errorf("capitalize(%q): got %q, want %q", tt.input, result, tt.want)
			}
		})
	}
}

func cmeTestAdaptor(t *testing.T) *Adaptor {
	t.Helper()
	return &Adaptor{
		AccessKey: "f58ecc51311847c9950ada687aa4d458",
		SecretKey: "d8ac2dd2d550405b8f344ef2cf042bbd",
	}
}

func cmeTestRequiredEnv(t *testing.T, name string) string {
	t.Helper()
	value := os.Getenv(name)
	if value == "" {
		t.Skipf("set %s to run this integration test", name)
	}
	return value
}

func cmeTestEnv(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}

func cmeTestOptionalEnv(name string) *string {
	if value := os.Getenv(name); value != "" {
		return &value
	}
	return nil
}

func cmeTestName(kind string) string {
	return fmt.Sprintf("new-api-test-%s-%d", kind, time.Now().UnixNano())
}

func cmeTestLogResult(t *testing.T, method string, result any, err error) {
	t.Helper()
	// 先输出结果和错误，再做断言，保证失败时也能看到接口返回的信息。
	data, marshalErr := json.MarshalIndent(result, "", "  ")
	if marshalErr != nil {
		t.Logf("%s result: %#v\nerror: %v", method, result, err)
		t.Fatalf("marshal %s result: %v", method, marshalErr)
	}
	t.Logf("%s result:\n%s\nerror: %v", method, data, err)
	if err != nil {
		t.Fatalf("%s failed: %v", method, err)
	}
}

func cmeTestAssertAssetStatus(t *testing.T, status string) {
	t.Helper()
	switch status {
	case "Processing", "Active", "Failed":
	default:
		t.Errorf("unexpected asset status: %q", status)
	}
}

func cmeTestAssertPagination(t *testing.T, total int64, pageNumber, pageSize, itemCount, wantPageNumber, wantPageSize int) {
	t.Helper()
	if pageNumber != wantPageNumber {
		t.Errorf("PageNumber: got %d, want %d", pageNumber, wantPageNumber)
	}
	if pageSize != wantPageSize {
		t.Errorf("PageSize: got %d, want %d", pageSize, wantPageSize)
	}
	if total < int64(itemCount) {
		t.Errorf("TotalCount %d is smaller than item count %d", total, itemCount)
	}
	if itemCount > pageSize {
		t.Errorf("item count %d exceeds PageSize %d", itemCount, pageSize)
	}
}
