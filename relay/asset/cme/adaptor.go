package cme

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/QuantumNous/new-api/dto"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/config"
	"gitlab.ecloud.com/ecloud/ecloudsdkmaas"
	"gitlab.ecloud.com/ecloud/ecloudsdkmaas/model"
)

type Adaptor struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

func (a *Adaptor) ReviewSkip() bool { return false }

func (a *Adaptor) CreateAssetGroup(req *dto.CreateAssetGroupRequest) (*dto.CreateAssetGroupResponse, error) {
	group, err := a.client().CreateAssetGroup(&model.CreateAssetGroupRequest{
		CreateAssetGroupBody: &model.CreateAssetGroupBody{
			Description: req.Description,
			GroupType:   &req.GroupType,
			GroupName:   &req.Name,
		},
	})
	if err != nil {
		return nil, err
	}
	if group == nil || group.State == nil {
		return nil, errors.New("create asset group failed")
	}
	if *group.State != "OK" {
		if message := group.ErrorMessage; message != nil {
			return nil, errors.New(*message)
		}
		return nil, errors.New("create asset group failed")
	}
	return &dto.CreateAssetGroupResponse{Id: *group.Body.GroupId}, nil
}

func (a *Adaptor) CreateAssets(req *dto.CreateAssetRequest) (*dto.CreateAssetResponse, error) {
	assetTypeEnum := model.CreateAssetBodyAssetTypeEnum(req.AssetType)
	asset, err := a.client().CreateAsset(&model.CreateAssetRequest{
		CreateAssetBody: &model.CreateAssetBody{
			AssetType: &assetTypeEnum,
			GroupId:   &req.GroupId,
			AssetName: &req.Name,
			AssetUrl:  &req.URL,
		},
	})
	if err != nil {
		return nil, err
	}
	if asset == nil || asset.State == nil {
		return nil, errors.New("create asset group failed")
	}
	if *asset.State != "OK" {
		if message := asset.ErrorMessage; message != nil {
			return nil, errors.New(*message)
		}
		return nil, errors.New("create asset group failed")
	}
	return &dto.CreateAssetResponse{Id: *asset.Body}, nil
}

func (a *Adaptor) GetAsset(req *dto.GetAssetRequest) (*dto.GetAssetResponse, error) {
	asset, err := a.client().GetAsset(&model.GetAssetRequest{
		GetAssetPath: &model.GetAssetPath{
			AssetId: &req.Id,
		},
	})
	if err != nil {
		return nil, err
	}
	if asset == nil || asset.State == nil {
		return nil, errors.New("create asset group failed")
	}
	if *asset.State != "OK" {
		if message := asset.ErrorMessage; message != nil {
			return nil, errors.New(*message)
		}
		return nil, errors.New("create asset group failed")
	}
	return &dto.GetAssetResponse{
		Status:     capitalize(string(*asset.Body.Status)),
		AssetType:  string(*asset.Body.AssetType),
		UpdateTime: *asset.Body.UpdatedTime,
		CreateTime: *asset.Body.CreatedTime,
		Name:       *asset.Body.AssetName,
		GroupId:    *asset.Body.GroupId,
		URL:        *asset.Body.AssetUrl,
		Id:         *asset.Body.AssetId,
	}, nil
}

func (a *Adaptor) ListAssets(req *dto.ListAssetsRequest) (*dto.ListAssetsResponse, error) {
	pageNo := int32(req.PageNumber)
	pageSize := int32(req.PageSize)
	assets, err := a.client().ListAssets(&model.ListAssetsRequest{
		ListAssetsBody: &model.ListAssetsBody{
			GroupType: req.Filter.GroupType,
			Statuses:  req.Filter.Statuses,
			GroupIds:  req.Filter.GroupIds,
			AssetName: req.Filter.Name,
			PageSize:  &pageSize,
			PageNo:    &pageNo,
		},
	})
	if err != nil {
		return nil, err
	}
	if assets == nil || assets.State == nil {
		return nil, errors.New("create asset group failed")
	}
	if *assets.State != "OK" {
		if message := assets.ErrorMessage; message != nil {
			return nil, errors.New(*message)
		}
		return nil, errors.New("create asset group failed")
	}
	var items []dto.GetAssetResponse
	for _, data := range *assets.Body.Data {
		item := dto.GetAssetResponse{
			Status:     capitalize(string(*data.Status)),
			AssetType:  string(*data.AssetType),
			CreateTime: *data.CreatedTime,
			UpdateTime: *data.UpdatedTime,
			Name:       *data.AssetName,
			Id:         *data.AssetId,
			GroupId:    *data.GroupId,
			URL:        *data.AssetUrl,
		}
		if message := data.ErrorMessage; message != nil {
			rawMessage := json.RawMessage(*message)
			item.Error = &rawMessage
		}
		items = append(items, item)
	}
	return &dto.ListAssetsResponse{
		TotalCount: *assets.Body.Total,
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
		Items:      items,
	}, nil
}

func (a *Adaptor) ListAssetGroups(req *dto.ListAssetGroupsRequest) (*dto.ListAssetGroupsResponse, error) {
	pageNo := int32(req.PageNumber)
	pageSize := int32(req.PageSize)
	groups, err := a.client().ListAssetGroups(&model.ListAssetGroupsRequest{
		ListAssetGroupsBody: &model.ListAssetGroupsBody{
			GroupType: req.Filter.GroupType,
			GroupIds:  req.Filter.GroupIds,
			GroupName: req.Filter.Name,
			PageSize:  &pageSize,
			PageNo:    &pageNo,
		},
	})
	if err != nil {
		return nil, err
	}
	if groups == nil || groups.State == nil {
		return nil, errors.New("create asset group failed")
	}
	if *groups.State != "OK" {
		if message := groups.ErrorMessage; message != nil {
			return nil, errors.New(*message)
		}
		return nil, errors.New("create asset group failed")
	}
	var items []dto.GetAssetGroupResponse
	for _, data := range *groups.Body.Data {
		item := dto.GetAssetGroupResponse{
			GroupType:   string(*data.GroupType),
			Description: *data.Description,
			CreateTime:  *data.CreatedTime,
			UpdateTime:  *data.UpdatedTime,
			Name:        *data.GroupName,
			Id:          *data.GroupId,
		}
		items = append(items, item)
	}
	return &dto.ListAssetGroupsResponse{
		TotalCount: *groups.Body.Total,
		PageNumber: req.PageNumber,
		PageSize:   req.PageSize,
		Items:      items,
	}, nil
}

func (a *Adaptor) CreateVisualValidateSession(req *dto.CreateVisualValidateSessionRequest) (*dto.CreateVisualValidateSessionResponse, error) {
	session, err := a.client().CreateSession(&model.CreateSessionRequest{
		CreateSessionBody: &model.CreateSessionBody{
			CallbackUrl: &req.CallbackURL,
		},
	})
	if err != nil {
		return nil, err
	}
	if session == nil || session.State == nil {
		return nil, errors.New("create asset group failed")
	}
	if *session.State != "OK" {
		if message := session.ErrorMessage; message != nil {
			return nil, errors.New(*message)
		}
		return nil, errors.New("create asset group failed")
	}
	return &dto.CreateVisualValidateSessionResponse{
		BytedToken:  *session.Body.BytedToken,
		H5Link:      *session.Body.H5Link,
		CallbackURL: req.CallbackURL,
	}, nil
}

func (a *Adaptor) GetVisualValidateResult(req *dto.GetVisualValidateResultRequest) (*dto.GetVisualValidateResultResponse, error) {
	token, err := a.client().GetAssetGroupByBytedToken(&model.GetAssetGroupByBytedTokenRequest{
		GetAssetGroupByBytedTokenBody: &model.GetAssetGroupByBytedTokenBody{
			BytedToken: &req.BytedToken,
		},
	})
	if err != nil {
		return nil, err
	}
	if token == nil || token.State == nil {
		return nil, errors.New("create asset group failed")
	}
	if *token.State != "OK" {
		if message := token.ErrorMessage; message != nil {
			return nil, errors.New(*message)
		}
		return nil, errors.New("create asset group failed")
	}

	return &dto.GetVisualValidateResultResponse{
		GroupId: *token.Body,
	}, nil
}

func (a *Adaptor) client() *ecloudsdkmaas.Client {
	c := &config.Config{
		AccessKey: &a.AccessKey,
		SecretKey: &a.SecretKey,
	}
	return ecloudsdkmaas.NewClient(c)
}

func capitalize(s string) string {
	if s == "" {
		return s
	}

	s = strings.ToLower(s)
	return strings.ToUpper(s[:1]) + s[1:]
}
