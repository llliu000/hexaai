package controller

import (
	"errors"
	"net/http"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const assetAPIVersion = "2024-01-01"

type AssetController struct{}

func (a *AssetController) Action(c *gin.Context) {
	action := c.Query("Action")
	version := c.Query("Version")
	if action == "" || version != assetAPIVersion {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	var err error
	var result dto.AssetBaseResult
	result.ResponseMetadata.RequestId = c.GetString(common.RequestIdKey)
	region := model.NormalizeAssetRegion(c.Query("Region"))
	result.ResponseMetadata.Version = version
	result.ResponseMetadata.Region = region
	result.ResponseMetadata.Action = action
	result.ResponseMetadata.Service = "ark"

	switch action {
	case "CreateAssetGroup":
		result.Result, err = a.CreateAssetGroup(c)
	case "CreateAsset":
		result.Result, err = a.CreateAsset(c)
	case "ListAssetGroups":
		result.Result, err = a.ListAssetGroups(c)
	case "ListAssets":
		result.Result, err = a.ListAssets(c)
	case "GetAsset":
		result.Result, err = a.GetAsset(c)
	case "GetAssetGroup":
		result.Result, err = a.GetAssetGroup(c)
	case "UpdateAssetGroup":
		result.Result, err = a.UpdateAssetGroup(c)
	case "UpdateAsset":
		result.Result, err = a.UpdateAsset(c)
	case "DeleteAsset":
		result.Result, err = a.DeleteAsset(c)
	case "DeleteAssetGroup":
		result.Result, err = a.DeleteAssetGroup(c)
	case "CreateVisualValidateSession":
		result.Result, err = a.CreateVisualValidateSession(c)
	case "GetVisualValidateResult":
		result.Result, err = a.GetVisualValidateResult(c)
	default:
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if err != nil {
		if errs, ok := errors.AsType[validator.ValidationErrors](err); ok {
			err = errors.New(service.TranslateValidationErrors(errs))
		}
		result.Result = map[string]any{"Code": "unknown", "Error": err.Error()}
	}
	c.JSON(http.StatusOK, result)
}

func (a *AssetController) CreateAssetGroup(c *gin.Context) (*dto.CreateAssetGroupResponse, error) {
	var req dto.CreateAssetGroupRequest
	if err := c.BindJSON(&req); nil != err {
		return nil, err
	}
	return service.CreateAssetGroup(c, &req)
}

func (a *AssetController) CreateAsset(c *gin.Context) (*dto.CreateAssetResponse, error) {
	var req dto.CreateAssetRequest
	if err := c.BindJSON(&req); nil != err {
		return nil, err
	}
	return service.CreateAsset(c, &req)
}

func (a *AssetController) ListAssetGroups(c *gin.Context) (*dto.ListAssetGroupsResponse, error) {
	var req dto.ListAssetGroupsRequest
	if err := c.BindJSON(&req); nil != err {
		return nil, err
	}
	return service.ListAssetGroups(c, &req)
}

func (a *AssetController) ListAssets(c *gin.Context) (*dto.ListAssetsResponse, error) {
	var req dto.ListAssetsRequest
	if err := c.BindJSON(&req); nil != err {
		return nil, err
	}
	return service.ListAssets(c, &req)
}

func (a *AssetController) GetAsset(c *gin.Context) (*dto.GetAssetResponse, error) {
	var req dto.GetAssetRequest
	if err := c.BindJSON(&req); nil != err {
		return nil, err
	}
	return service.GetAsset(c, &req)
}

func (a *AssetController) GetAssetGroup(c *gin.Context) (*dto.GetAssetGroupResponse, error) {
	var req dto.GetAssetGroupRequest
	if err := c.BindJSON(&req); nil != err {
		return nil, err
	}
	return service.GetAssetGroup(c, &req)
}

func (a *AssetController) UpdateAssetGroup(c *gin.Context) (*dto.UpdateAssetGroupResponse, error) {
	var req dto.UpdateAssetGroupRequest
	if err := c.BindJSON(&req); nil != err {
		return nil, err
	}
	return service.UpdateAssetGroup(c, &req)
}

func (a *AssetController) UpdateAsset(c *gin.Context) (*dto.UpdateAssetResponse, error) {
	var req dto.UpdateAssetRequest
	if err := c.BindJSON(&req); nil != err {
		return nil, err
	}
	return service.UpdateAsset(c, &req)
}

func (a *AssetController) DeleteAsset(c *gin.Context) (*dto.DeleteAssetResponse, error) {
	var req dto.DeleteAssetRequest
	if err := c.BindJSON(&req); nil != err {
		return nil, err
	}
	return service.DeleteAsset(c, &req)
}

func (a *AssetController) DeleteAssetGroup(c *gin.Context) (*dto.DeleteAssetGroupResponse, error) {
	var req dto.DeleteAssetGroupRequest
	if err := c.BindJSON(&req); nil != err {
		return nil, err
	}
	return service.DeleteAssetGroup(c, &req)
}

func (a *AssetController) CreateVisualValidateSession(c *gin.Context) (*dto.CreateVisualValidateSessionResponse, error) {
	var req dto.CreateVisualValidateSessionRequest
	if err := c.BindJSON(&req); nil != err {
		return nil, err
	}
	return service.CreateVisualValidateSession(c, &req)
}

func (a *AssetController) GetVisualValidateResult(c *gin.Context) (*dto.GetVisualValidateResultResponse, error) {
	var req dto.GetVisualValidateResultRequest
	if err := c.BindJSON(&req); nil != err {
		return nil, err
	}
	return service.GetVisualValidateResult(c, &req)
}
