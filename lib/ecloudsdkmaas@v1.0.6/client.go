// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package ecloudsdkmaas

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/config"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/param"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/region"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/request"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
	"gitlab.ecloud.com/ecloud/ecloudsdkmaas/model"
)

type Client struct {
	apiClient   *ecloudsdkcore.APIClient
	config      *config.Config
	httpRequest *request.HttpRequest
	allRegions  []region.Region
}

func NewClient(config *config.Config) *Client {
	httpRequest := request.DefaultHttpRequest()
	httpRequest.Product = product
	httpRequest.Version = version
	httpRequest.SdkVersion = sdkVersion
	ecloudsdkcore.InitConfig(config)
	apiClient := ecloudsdkcore.DefaultApiClient(config, httpRequest)
	client := &Client{
		apiClient:   apiClient,
		config:      config,
		httpRequest: httpRequest,
	}
	client.allRegions = client.initRegions()
	client.setEndpoint(config, httpRequest)
	return client
}

const (
	product    string = "MaaS"
	version           = "v1"
	sdkVersion        = "1.0.6"
)

func (c *Client) initRegions() []region.Region {
	var regions []region.Region
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-jiangsu-1"),
		PoolId:   utils.String("CIDC-RP-25"),
		Endpoint: utils.String("https://api-wuxi-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-guangdong-1"),
		PoolId:   utils.String("CIDC-RP-26"),
		Endpoint: utils.String("https://api-dongguan-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-sichuan-1"),
		PoolId:   utils.String("CIDC-RP-27"),
		Endpoint: utils.String("https://api-yaan-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-henan-1"),
		PoolId:   utils.String("CIDC-RP-28"),
		Endpoint: utils.String("https://api-zhengzhou-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-beijing-1"),
		PoolId:   utils.String("CIDC-RP-29"),
		Endpoint: utils.String("https://api-beijing-2.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-hunan-1"),
		PoolId:   utils.String("CIDC-RP-30"),
		Endpoint: utils.String("https://api-zhuzhou-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-shandong-1"),
		PoolId:   utils.String("CIDC-RP-31"),
		Endpoint: utils.String("https://api-jinan-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-shaanxi-1"),
		PoolId:   utils.String("CIDC-RP-32"),
		Endpoint: utils.String("https://api-xian-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-shanghai-1"),
		PoolId:   utils.String("CIDC-RP-33"),
		Endpoint: utils.String("https://api-shanghai-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-chongqing-1"),
		PoolId:   utils.String("CIDC-RP-34"),
		Endpoint: utils.String("https://api-chongqing-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-zhejiang-1"),
		PoolId:   utils.String("CIDC-RP-35"),
		Endpoint: utils.String("https://api-ningbo-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-tianjin-1"),
		PoolId:   utils.String("CIDC-RP-36"),
		Endpoint: utils.String("https://api-tianjin-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-jilin-1"),
		PoolId:   utils.String("CIDC-RP-37"),
		Endpoint: utils.String("https://api-jilin-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-hubei-1"),
		PoolId:   utils.String("CIDC-RP-38"),
		Endpoint: utils.String("https://api-hubei-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-jiangxi-1"),
		PoolId:   utils.String("CIDC-RP-39"),
		Endpoint: utils.String("https://api-jiangxi-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-gansu-1"),
		PoolId:   utils.String("CIDC-RP-40"),
		Endpoint: utils.String("https://api-gansu-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-shangxi-1"),
		PoolId:   utils.String("CIDC-RP-41"),
		Endpoint: utils.String("https://api-shanxi-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-liaoning-1"),
		PoolId:   utils.String("CIDC-RP-42"),
		Endpoint: utils.String("https://api-liaoning-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-yunnan-1"),
		PoolId:   utils.String("CIDC-RP-43"),
		Endpoint: utils.String("https://api-yunnan-2.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-hebei-1"),
		PoolId:   utils.String("CIDC-RP-44"),
		Endpoint: utils.String("https://api-hebei-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-fujian-1"),
		PoolId:   utils.String("CIDC-RP-45"),
		Endpoint: utils.String("https://api-fujian-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-guangxi-1"),
		PoolId:   utils.String("CIDC-RP-46"),
		Endpoint: utils.String("https://api-guangxi-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-anhui-1"),
		PoolId:   utils.String("CIDC-RP-47"),
		Endpoint: utils.String("https://api-anhui-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-neimenggu-1"),
		PoolId:   utils.String("CIDC-RP-48"),
		Endpoint: utils.String("https://api-huhehaote-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-guzhou-1"),
		PoolId:   utils.String("CIDC-RP-49"),
		Endpoint: utils.String("https://api-guiyang-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-hainan-1"),
		PoolId:   utils.String("CIDC-RP-53"),
		Endpoint: utils.String("https://api-hainan-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-xinjiang-1"),
		PoolId:   utils.String("CIDC-RP-54"),
		Endpoint: utils.String("https://api-xinjiang-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-heilongjiang-1"),
		PoolId:   utils.String("CIDC-RP-55"),
		Endpoint: utils.String("https://api-heilongjiang-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-ningxia-1"),
		PoolId:   utils.String("CIDC-RP-60"),
		Endpoint: utils.String("https://api-ningxia-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-qinghai-1"),
		PoolId:   utils.String("CIDC-RP-61"),
		Endpoint: utils.String("https://api-qinghai-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-xizang-1"),
		PoolId:   utils.String("CIDC-RP-62"),
		Endpoint: utils.String("https://api-xizang-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-hubei-2"),
		PoolId:   utils.String("CIDC-RP-64"),
		Endpoint: utils.String("https://api-wuhan-1.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-xizang-2"),
		PoolId:   utils.String("CIDC-RP-68"),
		Endpoint: utils.String("https://api-xizang-2.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-qinghai-2"),
		PoolId:   utils.String("CIDC-RP-69"),
		Endpoint: utils.String("https://api-qinghai-2.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String("cn-ningxia-2"),
		PoolId:   utils.String("CIDC-RP-71"),
		Endpoint: utils.String("https://api-ningxia-2.cmecloud.cn:8443"),
	})
	regions = append(regions, region.Region{
		RegionId: utils.String(""),
		PoolId:   utils.String("CIDC-CORE-00"),
		Endpoint: utils.String("https://ecloud.10086.cn"),
	})
	return regions
}

func (c *Client) setEndpoint(config *config.Config, httpRequest *request.HttpRequest) {
	if utils.IsUnSet(config.RegionId) && utils.IsUnSet(config.PoolId) {
		httpRequest.Endpoint = utils.DefaultEndpoint
	} else if utils.IsSet(config.RegionId) {
		c.findAndSetEndpointByRegionId(config.RegionId, httpRequest)
	} else {
		c.findAndSetEndpointByPoolId(config.PoolId, httpRequest)
	}
}

func (c *Client) findAndSetEndpointByRegionId(regionId *string, httpRequest *request.HttpRequest) {
	for _, r := range c.allRegions {
		if utils.StringValue(r.RegionId) == utils.StringValue(regionId) {
			endpoint := r.Endpoint
			if utils.IsUnSet(endpoint) {
				httpRequest.Endpoint = utils.DefaultEndpoint
			} else {
				httpRequest.Endpoint = *endpoint
			}
			break
		}
	}
}

func (c *Client) findAndSetEndpointByPoolId(poolId *string, httpRequest *request.HttpRequest) {
	for _, r := range c.allRegions {
		if utils.StringValue(r.PoolId) == utils.StringValue(poolId) {
			endpoint := r.Endpoint
			if utils.IsUnSet(endpoint) {
				httpRequest.Endpoint = utils.DefaultEndpoint
			} else {
				httpRequest.Endpoint = *endpoint
			}
			break
		}
	}
}

// CreateAsset 创建素材
func (c *Client) CreateAsset(request *model.CreateAssetRequest) (*model.CreateAssetResponse, error) {
	return c.CreateAssetWithConfig(request, nil)
}

// CreateAssetWithConfig 创建素材
func (c *Client) CreateAssetWithConfig(request *model.CreateAssetRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateAssetResponse, error) {
	params := param.NewParamsBuilder().
		Action("createAsset").
		Uri("/exp/aicc/v2/asset/").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/asset").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.CreateAssetResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// ListAssets 查询素材列表
func (c *Client) ListAssets(request *model.ListAssetsRequest) (*model.ListAssetsResponse, error) {
	return c.ListAssetsWithConfig(request, nil)
}

// ListAssetsWithConfig 查询素材列表
func (c *Client) ListAssetsWithConfig(request *model.ListAssetsRequest, runtimeConfig *config.RuntimeConfig) (*model.ListAssetsResponse, error) {
	params := param.NewParamsBuilder().
		Action("listAssets").
		Uri("/exp/aicc/v2/asset/query").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/asset/query").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.ListAssetsResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// DeleteAssetGroup 删除素材组
func (c *Client) DeleteAssetGroup(request *model.DeleteAssetGroupRequest) (*model.DeleteAssetGroupResponse, error) {
	return c.DeleteAssetGroupWithConfig(request, nil)
}

// DeleteAssetGroupWithConfig 删除素材组
func (c *Client) DeleteAssetGroupWithConfig(request *model.DeleteAssetGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteAssetGroupResponse, error) {
	params := param.NewParamsBuilder().
		Action("deleteAssetGroup").
		Uri("/exp/aicc/v2/asset-group/{groupId}").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/asset-group/{groupId}").
		Protocol("http").
		ContentType("").
		Method("DELETE").
		Request(request).
		Build()
	returnValue := &model.DeleteAssetGroupResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// GetAssetGroup 查询素材组详情
func (c *Client) GetAssetGroup(request *model.GetAssetGroupRequest) (*model.GetAssetGroupResponse, error) {
	return c.GetAssetGroupWithConfig(request, nil)
}

// GetAssetGroupWithConfig 查询素材组详情
func (c *Client) GetAssetGroupWithConfig(request *model.GetAssetGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.GetAssetGroupResponse, error) {
	params := param.NewParamsBuilder().
		Action("getAssetGroup").
		Uri("/exp/aicc/v2/asset-group/{groupId}").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/asset-group/{groupId}").
		Protocol("http").
		ContentType("").
		Method("GET").
		Request(request).
		Build()
	returnValue := &model.GetAssetGroupResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// UpdateAsset 更新素材
func (c *Client) UpdateAsset(request *model.UpdateAssetRequest) (*model.UpdateAssetResponse, error) {
	return c.UpdateAssetWithConfig(request, nil)
}

// UpdateAssetWithConfig 更新素材
func (c *Client) UpdateAssetWithConfig(request *model.UpdateAssetRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateAssetResponse, error) {
	params := param.NewParamsBuilder().
		Action("updateAsset").
		Uri("/exp/aicc/v2/asset/{assetId}").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/asset/{assetId}").
		Protocol("http").
		ContentType("application/json").
		Method("PUT").
		Request(request).
		Build()
	returnValue := &model.UpdateAssetResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// CreateAssetGroup 创建素材组
func (c *Client) CreateAssetGroup(request *model.CreateAssetGroupRequest) (*model.CreateAssetGroupResponse, error) {
	return c.CreateAssetGroupWithConfig(request, nil)
}

// CreateAssetGroupWithConfig 创建素材组
func (c *Client) CreateAssetGroupWithConfig(request *model.CreateAssetGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateAssetGroupResponse, error) {
	params := param.NewParamsBuilder().
		Action("createAssetGroup").
		Uri("/exp/aicc/v2/asset-group/").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/asset-group").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.CreateAssetGroupResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// ListAssetGroups 查询素材组列表
func (c *Client) ListAssetGroups(request *model.ListAssetGroupsRequest) (*model.ListAssetGroupsResponse, error) {
	return c.ListAssetGroupsWithConfig(request, nil)
}

// ListAssetGroupsWithConfig 查询素材组列表
func (c *Client) ListAssetGroupsWithConfig(request *model.ListAssetGroupsRequest, runtimeConfig *config.RuntimeConfig) (*model.ListAssetGroupsResponse, error) {
	params := param.NewParamsBuilder().
		Action("listAssetGroups").
		Uri("/exp/aicc/v2/asset-group/query").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/asset-group/query").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.ListAssetGroupsResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// CreateOrder 按Token用量计费规格--订购
func (c *Client) CreateOrder(request *model.CreateOrderRequest) (*model.CreateOrderResponse, error) {
	return c.CreateOrderWithConfig(request, nil)
}

// CreateOrderWithConfig 按Token用量计费规格--订购
func (c *Client) CreateOrderWithConfig(request *model.CreateOrderRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateOrderResponse, error) {
	params := param.NewParamsBuilder().
		Action("createOrder").
		Uri("/console/studio/order").
		GatewayUri("/api/openapi-maas/console/studio/order").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.CreateOrderResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// GetSegment 获取分段详情
func (c *Client) GetSegment(request *model.GetSegmentRequest) (*model.GetSegmentResponse, error) {
	return c.GetSegmentWithConfig(request, nil)
}

// GetSegmentWithConfig 获取分段详情
func (c *Client) GetSegmentWithConfig(request *model.GetSegmentRequest, runtimeConfig *config.RuntimeConfig) (*model.GetSegmentResponse, error) {
	params := param.NewParamsBuilder().
		Action("getSegment").
		Uri("/works/datasets/{datasetId}/documents/{documentId}/segments/{segmentId}").
		GatewayUri("/api/openapi-maas/works/datasets/{datasetId}/documents/{documentId}/segments/{segmentId}").
		Protocol("http").
		ContentType("").
		Method("GET").
		Request(request).
		Build()
	returnValue := &model.GetSegmentResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// GetAsset 查询素材详情
func (c *Client) GetAsset(request *model.GetAssetRequest) (*model.GetAssetResponse, error) {
	return c.GetAssetWithConfig(request, nil)
}

// GetAssetWithConfig 查询素材详情
func (c *Client) GetAssetWithConfig(request *model.GetAssetRequest, runtimeConfig *config.RuntimeConfig) (*model.GetAssetResponse, error) {
	params := param.NewParamsBuilder().
		Action("getAsset").
		Uri("/exp/aicc/v2/asset/{assetId}").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/asset/{assetId}").
		Protocol("http").
		ContentType("").
		Method("GET").
		Request(request).
		Build()
	returnValue := &model.GetAssetResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// GetAssetGroupByBytedToken 通过bytedToken查询素材组
func (c *Client) GetAssetGroupByBytedToken(request *model.GetAssetGroupByBytedTokenRequest) (*model.GetAssetGroupByBytedTokenResponse, error) {
	return c.GetAssetGroupByBytedTokenWithConfig(request, nil)
}

// GetAssetGroupByBytedTokenWithConfig 通过bytedToken查询素材组
func (c *Client) GetAssetGroupByBytedTokenWithConfig(request *model.GetAssetGroupByBytedTokenRequest, runtimeConfig *config.RuntimeConfig) (*model.GetAssetGroupByBytedTokenResponse, error) {
	params := param.NewParamsBuilder().
		Action("getAssetGroupByBytedToken").
		Uri("/exp/aicc/v2/real-person-auth/asset-group/by-byted-token").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/real-person-auth/asset-group/by-byted-token").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.GetAssetGroupByBytedTokenResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// CreateSession 创建H5认证会话
func (c *Client) CreateSession(request *model.CreateSessionRequest) (*model.CreateSessionResponse, error) {
	return c.CreateSessionWithConfig(request, nil)
}

// CreateSessionWithConfig 创建H5认证会话
func (c *Client) CreateSessionWithConfig(request *model.CreateSessionRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateSessionResponse, error) {
	params := param.NewParamsBuilder().
		Action("createSession").
		Uri("/exp/aicc/v2/real-person-auth/sessions").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/real-person-auth/sessions").
		Protocol("http").
		ContentType("").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.CreateSessionResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// UpdateAssetGroup 更新素材组
func (c *Client) UpdateAssetGroup(request *model.UpdateAssetGroupRequest) (*model.UpdateAssetGroupResponse, error) {
	return c.UpdateAssetGroupWithConfig(request, nil)
}

// UpdateAssetGroupWithConfig 更新素材组
func (c *Client) UpdateAssetGroupWithConfig(request *model.UpdateAssetGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateAssetGroupResponse, error) {
	params := param.NewParamsBuilder().
		Action("updateAssetGroup").
		Uri("/exp/aicc/v2/asset-group/{groupId}").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/asset-group/{groupId}").
		Protocol("http").
		ContentType("application/json").
		Method("PUT").
		Request(request).
		Build()
	returnValue := &model.UpdateAssetGroupResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// DeleteAsset 删除素材
func (c *Client) DeleteAsset(request *model.DeleteAssetRequest) (*model.DeleteAssetResponse, error) {
	return c.DeleteAssetWithConfig(request, nil)
}

// DeleteAssetWithConfig 删除素材
func (c *Client) DeleteAssetWithConfig(request *model.DeleteAssetRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteAssetResponse, error) {
	params := param.NewParamsBuilder().
		Action("deleteAsset").
		Uri("/exp/aicc/v2/asset/{assetId}").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/asset/{assetId}").
		Protocol("http").
		ContentType("").
		Method("DELETE").
		Request(request).
		Build()
	returnValue := &model.DeleteAssetResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// ListDocs 列出文档
func (c *Client) ListDocs(request *model.ListDocsRequest) (*model.ListDocsResponse, error) {
	return c.ListDocsWithConfig(request, nil)
}

// ListDocsWithConfig 列出文档
func (c *Client) ListDocsWithConfig(request *model.ListDocsRequest, runtimeConfig *config.RuntimeConfig) (*model.ListDocsResponse, error) {
	params := param.NewParamsBuilder().
		Action("listDocs").
		Uri("/works/docs").
		GatewayUri("/api/openapi-maas/works/docs").
		Protocol("http").
		ContentType("").
		Method("GET").
		Request(request).
		Build()
	returnValue := &model.ListDocsResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// AutoRenewOrder Coding Plan自动续订
func (c *Client) AutoRenewOrder(request *model.AutoRenewOrderRequest) (*model.AutoRenewOrderResponse, error) {
	return c.AutoRenewOrderWithConfig(request, nil)
}

// AutoRenewOrderWithConfig Coding Plan自动续订
func (c *Client) AutoRenewOrderWithConfig(request *model.AutoRenewOrderRequest, runtimeConfig *config.RuntimeConfig) (*model.AutoRenewOrderResponse, error) {
	params := param.NewParamsBuilder().
		Action("autoRenewOrder").
		Uri("/console/studio/auto-renew").
		GatewayUri("/api/openapi-maas/console/studio/auto-renew").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.AutoRenewOrderResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// RenewOrder 手动续订Coding Plan，返回MOP-O订单号
func (c *Client) RenewOrder(request *model.RenewOrderRequest) (*model.RenewOrderResponse, error) {
	return c.RenewOrderWithConfig(request, nil)
}

// RenewOrderWithConfig 手动续订Coding Plan，返回MOP-O订单号
func (c *Client) RenewOrderWithConfig(request *model.RenewOrderRequest, runtimeConfig *config.RuntimeConfig) (*model.RenewOrderResponse, error) {
	params := param.NewParamsBuilder().
		Action("renewOrder").
		Uri("/console/studio/renew").
		GatewayUri("/api/openapi-maas/console/studio/renew").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.RenewOrderResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// CancelOrder 按Token用量计费规格--退订
func (c *Client) CancelOrder(request *model.CancelOrderRequest) (*model.CancelOrderResponse, error) {
	return c.CancelOrderWithConfig(request, nil)
}

// CancelOrderWithConfig 按Token用量计费规格--退订
func (c *Client) CancelOrderWithConfig(request *model.CancelOrderRequest, runtimeConfig *config.RuntimeConfig) (*model.CancelOrderResponse, error) {
	params := param.NewParamsBuilder().
		Action("cancelOrder").
		Uri("/console/studio/cancel").
		GatewayUri("/api/openapi-maas/console/studio/cancel").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.CancelOrderResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// DeleteDataset 删除知识库
func (c *Client) DeleteDataset(request *model.DeleteDatasetRequest) (*model.DeleteDatasetResponse, error) {
	return c.DeleteDatasetWithConfig(request, nil)
}

// DeleteDatasetWithConfig 删除知识库
func (c *Client) DeleteDatasetWithConfig(request *model.DeleteDatasetRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteDatasetResponse, error) {
	params := param.NewParamsBuilder().
		Action("deleteDataset").
		Uri("/works/datasets/{datasetId}").
		GatewayUri("/api/openapi-maas/works/datasets/{datasetId}").
		Protocol("http").
		ContentType("").
		Method("DELETE").
		Request(request).
		Build()
	returnValue := &model.DeleteDatasetResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// QueryUsageStats 模型用量统计
func (c *Client) QueryUsageStats(request *model.QueryUsageStatsRequest) (*model.QueryUsageStatsResponse, error) {
	return c.QueryUsageStatsWithConfig(request, nil)
}

// QueryUsageStatsWithConfig 模型用量统计
func (c *Client) QueryUsageStatsWithConfig(request *model.QueryUsageStatsRequest, runtimeConfig *config.RuntimeConfig) (*model.QueryUsageStatsResponse, error) {
	params := param.NewParamsBuilder().
		Action("queryUsageStats").
		Uri("/model/usages/statistics").
		GatewayUri("/api/openapi-maas/model/usages/statistics").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.QueryUsageStatsResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// UpdateDataset 修改知识库
func (c *Client) UpdateDataset(request *model.UpdateDatasetRequest) (*model.UpdateDatasetResponse, error) {
	return c.UpdateDatasetWithConfig(request, nil)
}

// UpdateDatasetWithConfig 修改知识库
func (c *Client) UpdateDatasetWithConfig(request *model.UpdateDatasetRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateDatasetResponse, error) {
	params := param.NewParamsBuilder().
		Action("updateDataset").
		Uri("/works/datasets/{datasetId}").
		GatewayUri("/api/openapi-maas/works/datasets/{datasetId}").
		Protocol("http").
		ContentType("application/json").
		Method("PUT").
		Request(request).
		Build()
	returnValue := &model.UpdateDatasetResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// CreateAICCExportTaskDetail 创建AICC任务级用量明细导出任务
func (c *Client) CreateAICCExportTaskDetail(request *model.CreateAICCExportTaskDetailRequest) (*model.CreateAICCExportTaskDetailResponse, error) {
	return c.CreateAICCExportTaskDetailWithConfig(request, nil)
}

// CreateAICCExportTaskDetailWithConfig 创建AICC任务级用量明细导出任务
func (c *Client) CreateAICCExportTaskDetailWithConfig(request *model.CreateAICCExportTaskDetailRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateAICCExportTaskDetailResponse, error) {
	params := param.NewParamsBuilder().
		Action("createAICCExportTaskDetail").
		Uri("/model/aicc/task-details/export-task").
		GatewayUri("/api/openapi-maas/model/aicc/task-details/export-task").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.CreateAICCExportTaskDetailResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// QueryUsageDetail 模型用量明细查询
func (c *Client) QueryUsageDetail(request *model.QueryUsageDetailRequest) (*model.QueryUsageDetailResponse, error) {
	return c.QueryUsageDetailWithConfig(request, nil)
}

// QueryUsageDetailWithConfig 模型用量明细查询
func (c *Client) QueryUsageDetailWithConfig(request *model.QueryUsageDetailRequest, runtimeConfig *config.RuntimeConfig) (*model.QueryUsageDetailResponse, error) {
	params := param.NewParamsBuilder().
		Action("queryUsageDetail").
		Uri("/model/usages/detail").
		GatewayUri("/api/openapi-maas/model/usages/detail").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.QueryUsageDetailResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// Page 分页查询智能体
func (c *Client) Page(request *model.PageRequest) (*model.PageResponse, error) {
	return c.PageWithConfig(request, nil)
}

// PageWithConfig 分页查询智能体
func (c *Client) PageWithConfig(request *model.PageRequest, runtimeConfig *config.RuntimeConfig) (*model.PageResponse, error) {
	params := param.NewParamsBuilder().
		Action("page").
		Uri("/works/agent/page").
		GatewayUri("/api/openapi-maas/works/agent/page").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.PageResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// QuerySegments 分段列表（分页查询）
func (c *Client) QuerySegments(request *model.QuerySegmentsRequest) (*model.QuerySegmentsResponse, error) {
	return c.QuerySegmentsWithConfig(request, nil)
}

// QuerySegmentsWithConfig 分段列表（分页查询）
func (c *Client) QuerySegmentsWithConfig(request *model.QuerySegmentsRequest, runtimeConfig *config.RuntimeConfig) (*model.QuerySegmentsResponse, error) {
	params := param.NewParamsBuilder().
		Action("querySegments").
		Uri("/works/datasets/{datasetId}/documents/{documentId}/segment-list").
		GatewayUri("/api/openapi-maas/works/datasets/{datasetId}/documents/{documentId}/segment-list").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.QuerySegmentsResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// DeleteDoc 删除文档
func (c *Client) DeleteDoc(request *model.DeleteDocRequest) (*model.DeleteDocResponse, error) {
	return c.DeleteDocWithConfig(request, nil)
}

// DeleteDocWithConfig 删除文档
func (c *Client) DeleteDocWithConfig(request *model.DeleteDocRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteDocResponse, error) {
	params := param.NewParamsBuilder().
		Action("deleteDoc").
		Uri("/works/docs/{docId}").
		GatewayUri("/api/openapi-maas/works/docs/{docId}").
		Protocol("http").
		ContentType("").
		Method("DELETE").
		Request(request).
		Build()
	returnValue := &model.DeleteDocResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// GetDocument 根据 docId 获取文档信息
func (c *Client) GetDocument(request *model.GetDocumentRequest) (*model.GetDocumentResponse, error) {
	return c.GetDocumentWithConfig(request, nil)
}

// GetDocumentWithConfig 根据 docId 获取文档信息
func (c *Client) GetDocumentWithConfig(request *model.GetDocumentRequest, runtimeConfig *config.RuntimeConfig) (*model.GetDocumentResponse, error) {
	params := param.NewParamsBuilder().
		Action("getDocument").
		Uri("/works/docs/{documentId}").
		GatewayUri("/api/openapi-maas/works/docs/{documentId}").
		Protocol("http").
		ContentType("").
		Method("GET").
		Request(request).
		Build()
	returnValue := &model.GetDocumentResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// BatchDeleteDoc 批量删除文档
func (c *Client) BatchDeleteDoc(request *model.BatchDeleteDocRequest) (*model.BatchDeleteDocResponse, error) {
	return c.BatchDeleteDocWithConfig(request, nil)
}

// BatchDeleteDocWithConfig 批量删除文档
func (c *Client) BatchDeleteDocWithConfig(request *model.BatchDeleteDocRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchDeleteDocResponse, error) {
	params := param.NewParamsBuilder().
		Action("batchDeleteDoc").
		Uri("/works/datasets/{datasetId}/documents/batch-delete").
		GatewayUri("/api/openapi-maas/works/datasets/{datasetId}/documents/batch-delete").
		Protocol("http").
		ContentType("application/json").
		Method("DELETE").
		Request(request).
		Build()
	returnValue := &model.BatchDeleteDocResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// DisableSegment 禁用分段
func (c *Client) DisableSegment(request *model.DisableSegmentRequest) (*model.DisableSegmentResponse, error) {
	return c.DisableSegmentWithConfig(request, nil)
}

// DisableSegmentWithConfig 禁用分段
func (c *Client) DisableSegmentWithConfig(request *model.DisableSegmentRequest, runtimeConfig *config.RuntimeConfig) (*model.DisableSegmentResponse, error) {
	params := param.NewParamsBuilder().
		Action("disableSegment").
		Uri("/works/datasets/{datasetId}/documents/{documentId}/segments/{segmentId}/disable").
		GatewayUri("/api/openapi-maas/works/datasets/{datasetId}/documents/{documentId}/segments/{segmentId}/disable").
		Protocol("http").
		ContentType("").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.DisableSegmentResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// EnableDocument 启用文档
func (c *Client) EnableDocument(request *model.EnableDocumentRequest) (*model.EnableDocumentResponse, error) {
	return c.EnableDocumentWithConfig(request, nil)
}

// EnableDocumentWithConfig 启用文档
func (c *Client) EnableDocumentWithConfig(request *model.EnableDocumentRequest, runtimeConfig *config.RuntimeConfig) (*model.EnableDocumentResponse, error) {
	params := param.NewParamsBuilder().
		Action("enableDocument").
		Uri("/works/datasets/{datasetId}/documents/{documentId}/enable").
		GatewayUri("/api/openapi-maas/works/datasets/{datasetId}/documents/{documentId}/enable").
		Protocol("http").
		ContentType("").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.EnableDocumentResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// Publish 发布智能体
func (c *Client) Publish(request *model.PublishRequest) (*model.PublishResponse, error) {
	return c.PublishWithConfig(request, nil)
}

// PublishWithConfig 发布智能体
func (c *Client) PublishWithConfig(request *model.PublishRequest, runtimeConfig *config.RuntimeConfig) (*model.PublishResponse, error) {
	params := param.NewParamsBuilder().
		Action("publish").
		Uri("/works/agent/{agentId}/publish").
		GatewayUri("/api/openapi-maas/works/agent/{agentId}/publish").
		Protocol("http").
		ContentType("").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.PublishResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// Detail 查询智能体详情
func (c *Client) Detail(request *model.DetailRequest) (*model.DetailResponse, error) {
	return c.DetailWithConfig(request, nil)
}

// DetailWithConfig 查询智能体详情
func (c *Client) DetailWithConfig(request *model.DetailRequest, runtimeConfig *config.RuntimeConfig) (*model.DetailResponse, error) {
	params := param.NewParamsBuilder().
		Action("detail").
		Uri("/works/agent/{agentId}/detail").
		GatewayUri("/api/openapi-maas/works/agent/{agentId}/detail").
		Protocol("http").
		ContentType("").
		Method("GET").
		Request(request).
		Build()
	returnValue := &model.DetailResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// QueryBillModels 查询用户订购的按量模型列表
func (c *Client) QueryBillModels(request *model.QueryBillModelsRequest) (*model.QueryBillModelsResponse, error) {
	return c.QueryBillModelsWithConfig(request, nil)
}

// QueryBillModelsWithConfig 查询用户订购的按量模型列表
func (c *Client) QueryBillModelsWithConfig(request *model.QueryBillModelsRequest, runtimeConfig *config.RuntimeConfig) (*model.QueryBillModelsResponse, error) {
	params := param.NewParamsBuilder().
		Action("queryBillModels").
		Uri("/console/studio/query-bill").
		GatewayUri("/api/openapi-maas/console/studio/query-bill").
		Protocol("http").
		ContentType("").
		Method("GET").
		Request(request).
		Build()
	returnValue := &model.QueryBillModelsResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// DeleteSegment 删除分段
func (c *Client) DeleteSegment(request *model.DeleteSegmentRequest) (*model.DeleteSegmentResponse, error) {
	return c.DeleteSegmentWithConfig(request, nil)
}

// DeleteSegmentWithConfig 删除分段
func (c *Client) DeleteSegmentWithConfig(request *model.DeleteSegmentRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteSegmentResponse, error) {
	params := param.NewParamsBuilder().
		Action("deleteSegment").
		Uri("/works/datasets/{datasetId}/documents/{documentId}/segments/{segmentId}").
		GatewayUri("/api/openapi-maas/works/datasets/{datasetId}/documents/{documentId}/segments/{segmentId}").
		Protocol("http").
		ContentType("").
		Method("DELETE").
		Request(request).
		Build()
	returnValue := &model.DeleteSegmentResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// EnableSegment 启用分段
func (c *Client) EnableSegment(request *model.EnableSegmentRequest) (*model.EnableSegmentResponse, error) {
	return c.EnableSegmentWithConfig(request, nil)
}

// EnableSegmentWithConfig 启用分段
func (c *Client) EnableSegmentWithConfig(request *model.EnableSegmentRequest, runtimeConfig *config.RuntimeConfig) (*model.EnableSegmentResponse, error) {
	params := param.NewParamsBuilder().
		Action("enableSegment").
		Uri("/works/datasets/{datasetId}/documents/{documentId}/segments/{segmentId}/enable").
		GatewayUri("/api/openapi-maas/works/datasets/{datasetId}/documents/{documentId}/segments/{segmentId}/enable").
		Protocol("http").
		ContentType("").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.EnableSegmentResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// CreateDataset 创建知识库
func (c *Client) CreateDataset(request *model.CreateDatasetRequest) (*model.CreateDatasetResponse, error) {
	return c.CreateDatasetWithConfig(request, nil)
}

// CreateDatasetWithConfig 创建知识库
func (c *Client) CreateDatasetWithConfig(request *model.CreateDatasetRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateDatasetResponse, error) {
	params := param.NewParamsBuilder().
		Action("createDataset").
		Uri("/works/datasets").
		GatewayUri("/api/openapi-maas/works/datasets").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.CreateDatasetResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// QueryAICCExportTaskDetail 查询AICC任务级用量明细导出任务状态
func (c *Client) QueryAICCExportTaskDetail(request *model.QueryAICCExportTaskDetailRequest) (*model.QueryAICCExportTaskDetailResponse, error) {
	return c.QueryAICCExportTaskDetailWithConfig(request, nil)
}

// QueryAICCExportTaskDetailWithConfig 查询AICC任务级用量明细导出任务状态
func (c *Client) QueryAICCExportTaskDetailWithConfig(request *model.QueryAICCExportTaskDetailRequest, runtimeConfig *config.RuntimeConfig) (*model.QueryAICCExportTaskDetailResponse, error) {
	params := param.NewParamsBuilder().
		Action("queryAICCExportTaskDetail").
		Uri("/model/aicc/task-details/export-task/{taskId}").
		GatewayUri("/api/openapi-maas/model/aicc/task-details/export-task/{taskId}").
		Protocol("http").
		ContentType("").
		Method("GET").
		Request(request).
		Build()
	returnValue := &model.QueryAICCExportTaskDetailResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// UpdateSegment 修改分段
func (c *Client) UpdateSegment(request *model.UpdateSegmentRequest) (*model.UpdateSegmentResponse, error) {
	return c.UpdateSegmentWithConfig(request, nil)
}

// UpdateSegmentWithConfig 修改分段
func (c *Client) UpdateSegmentWithConfig(request *model.UpdateSegmentRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateSegmentResponse, error) {
	params := param.NewParamsBuilder().
		Action("updateSegment").
		Uri("/works/datasets/{datasetId}/documents/{documentId}/segments/{segmentId}").
		GatewayUri("/api/openapi-maas/works/datasets/{datasetId}/documents/{documentId}/segments/{segmentId}").
		Protocol("http").
		ContentType("").
		Method("PUT").
		Request(request).
		Build()
	returnValue := &model.UpdateSegmentResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// DisableDocument 禁用文档
func (c *Client) DisableDocument(request *model.DisableDocumentRequest) (*model.DisableDocumentResponse, error) {
	return c.DisableDocumentWithConfig(request, nil)
}

// DisableDocumentWithConfig 禁用文档
func (c *Client) DisableDocumentWithConfig(request *model.DisableDocumentRequest, runtimeConfig *config.RuntimeConfig) (*model.DisableDocumentResponse, error) {
	params := param.NewParamsBuilder().
		Action("disableDocument").
		Uri("/works/datasets/{datasetId}/documents/{documentId}/disable").
		GatewayUri("/api/openapi-maas/works/datasets/{datasetId}/documents/{documentId}/disable").
		Protocol("http").
		ContentType("").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.DisableDocumentResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// QueryAiccCreditDeduction 查询资费包扣减明细
func (c *Client) QueryAiccCreditDeduction(request *model.QueryAiccCreditDeductionRequest) (*model.QueryAiccCreditDeductionResponse, error) {
	return c.QueryAiccCreditDeductionWithConfig(request, nil)
}

// QueryAiccCreditDeductionWithConfig 查询资费包扣减明细
func (c *Client) QueryAiccCreditDeductionWithConfig(request *model.QueryAiccCreditDeductionRequest, runtimeConfig *config.RuntimeConfig) (*model.QueryAiccCreditDeductionResponse, error) {
	params := param.NewParamsBuilder().
		Action("queryAiccCreditDeduction").
		Uri("/model/aicc/deduction").
		GatewayUri("/api/openapi-maas/model/aicc/deduction").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.QueryAiccCreditDeductionResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// QueryExportTask 查询资费包扣减明细导出任务状态
func (c *Client) QueryExportTask(request *model.QueryExportTaskRequest) (*model.QueryExportTaskResponse, error) {
	return c.QueryExportTaskWithConfig(request, nil)
}

// QueryExportTaskWithConfig 查询资费包扣减明细导出任务状态
func (c *Client) QueryExportTaskWithConfig(request *model.QueryExportTaskRequest, runtimeConfig *config.RuntimeConfig) (*model.QueryExportTaskResponse, error) {
	params := param.NewParamsBuilder().
		Action("queryExportTask").
		Uri("/model/aicc/deduction/export-task/{taskId}").
		GatewayUri("/api/openapi-maas/model/aicc/deduction/export-task/{taskId}").
		Protocol("http").
		ContentType("").
		Method("GET").
		Request(request).
		Build()
	returnValue := &model.QueryExportTaskResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// CreateExportTask 创建资费包扣减明细导出任务
func (c *Client) CreateExportTask(request *model.CreateExportTaskRequest) (*model.CreateExportTaskResponse, error) {
	return c.CreateExportTaskWithConfig(request, nil)
}

// CreateExportTaskWithConfig 创建资费包扣减明细导出任务
func (c *Client) CreateExportTaskWithConfig(request *model.CreateExportTaskRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateExportTaskResponse, error) {
	params := param.NewParamsBuilder().
		Action("createExportTask").
		Uri("/model/aicc/deduction/export-task").
		GatewayUri("/api/openapi-maas/model/aicc/deduction/export-task").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.CreateExportTaskResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// ListDatasetByUserId 根据用户查询某个项目空间下的知识库列表
func (c *Client) ListDatasetByUserId(request *model.ListDatasetByUserIdRequest) (*model.ListDatasetByUserIdResponse, error) {
	return c.ListDatasetByUserIdWithConfig(request, nil)
}

// ListDatasetByUserIdWithConfig 根据用户查询某个项目空间下的知识库列表
func (c *Client) ListDatasetByUserIdWithConfig(request *model.ListDatasetByUserIdRequest, runtimeConfig *config.RuntimeConfig) (*model.ListDatasetByUserIdResponse, error) {
	params := param.NewParamsBuilder().
		Action("listDatasetByUserId").
		Uri("/works/datasets/list/by-user-project").
		GatewayUri("/api/openapi-maas/works/datasets/list/by-user-project").
		Protocol("http").
		ContentType("application/json").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.ListDatasetByUserIdResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// OnShelf 智能体上架
func (c *Client) OnShelf(request *model.OnShelfRequest) (*model.OnShelfResponse, error) {
	return c.OnShelfWithConfig(request, nil)
}

// OnShelfWithConfig 智能体上架
func (c *Client) OnShelfWithConfig(request *model.OnShelfRequest, runtimeConfig *config.RuntimeConfig) (*model.OnShelfResponse, error) {
	params := param.NewParamsBuilder().
		Action("onShelf").
		Uri("/works/agentstore/{agentId}/on").
		GatewayUri("/api/openapi-maas/works/agentstore/{agentId}/on").
		Protocol("http").
		ContentType("").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.OnShelfResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// ConfirmVerificationResult 确认真人认证结果并获取素材组（自定义回调场景）
func (c *Client) ConfirmVerificationResult(request *model.ConfirmVerificationResultRequest) (*model.ConfirmVerificationResultResponse, error) {
	return c.ConfirmVerificationResultWithConfig(request, nil)
}

// ConfirmVerificationResultWithConfig 确认真人认证结果并获取素材组（自定义回调场景）
func (c *Client) ConfirmVerificationResultWithConfig(request *model.ConfirmVerificationResultRequest, runtimeConfig *config.RuntimeConfig) (*model.ConfirmVerificationResultResponse, error) {
	params := param.NewParamsBuilder().
		Action("confirmVerificationResult").
		Uri("/exp/aicc/v2/real-person-auth/verification-result").
		GatewayUri("/api/openapi-maas/exp/aicc/v2/real-person-auth/verification-result").
		Protocol("http").
		ContentType("").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.ConfirmVerificationResultResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}

// QueryDocsPageByType 根据文档状态和类型查询文档列表
func (c *Client) QueryDocsPageByType(request *model.QueryDocsPageByTypeRequest) (*model.QueryDocsPageByTypeResponse, error) {
	return c.QueryDocsPageByTypeWithConfig(request, nil)
}

// QueryDocsPageByTypeWithConfig 根据文档状态和类型查询文档列表
func (c *Client) QueryDocsPageByTypeWithConfig(request *model.QueryDocsPageByTypeRequest, runtimeConfig *config.RuntimeConfig) (*model.QueryDocsPageByTypeResponse, error) {
	params := param.NewParamsBuilder().
		Action("queryDocsPageByType").
		Uri("/works/datasets/{datasetId}/documents/query-by-extenxions").
		GatewayUri("/api/openapi-maas/works/datasets/{datasetId}/documents/query-by-extenxions").
		Protocol("http").
		ContentType("").
		Method("POST").
		Request(request).
		Build()
	returnValue := &model.QueryDocsPageByTypeResponse{}
	if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
		return nil, err
	} else {
		return returnValue, nil
	}
}
