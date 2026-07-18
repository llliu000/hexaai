package volc

import (
	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/universal"
)

type Adaptor struct {
	BaseUrl     string `json:"base_url"`
	AccessID    string `json:"access_id"`
	SecretKey   string `json:"secret_key"`
	ProjectName string `json:"project_name"`
	Region      string `json:"region"`
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

func (a *Adaptor) doCall(action string, request, response any) error {
	config := volcengine.NewConfig().WithAkSk(a.AccessID, a.SecretKey).
		WithEndpoint(a.BaseUrl).WithRegion(a.Region)
	sess, err := session.NewSession(config)
	if err != nil {
		return err
	}
	marshal, err := common.Marshal(request)
	if err != nil {
		return err
	}
	params := map[string]any{}
	if err = common.Unmarshal(marshal, &params); nil != err {
		return err
	}
	resp, err := universal.New(sess).DoCall(
		universal.RequestUniversal{
			ServiceName: "ark",
			Action:      action,
			Version:     "2024-01-01",
			HttpMethod:  universal.POST,
			ContentType: universal.ApplicationJSON,
		},
		&params,
	)
	if err != nil {
		return err
	}
	result, ok := (*resp)["Result"]
	if !ok {
		return nil
	}
	respData, err := common.Marshal(result)
	if err != nil {
		return err
	}
	err = common.Unmarshal(respData, &response)
	return err
}
