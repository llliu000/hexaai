package anyfast

type createResponse struct {
	Id string `json:"Id"`
}

type listAssetGroupResponse struct {
	TotalCount int64            `json:"TotalCount"`
	PageNumber int              `json:"PageNumber"`
	PageSize   int              `json:"PageSize"`
	Items      []assetGroupItem `json:"Items"`
}

type assetGroupItem struct {
	Id          string `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	GroupType   string `json:"GroupType"`
	ProjectName string `json:"ProjectName"`
}

type listAssetsResponse struct {
	Items      []assetItem `json:"Items"`
	PageNumber int         `json:"PageNumber"`
	PageSize   int         `json:"PageSize"`
	TotalCount int64       `json:"TotalCount"`
}

type assetItem struct {
	Id          string `json:"Id"`
	Name        string `json:"Name"`
	URL         string `json:"URL"`
	AssetType   string `json:"AssetType"`
	GroupId     string `json:"GroupId"`
	Status      string `json:"Status"`
	ProjectName string `json:"ProjectName"`
	CreateTime  string `json:"CreateTime"`
	UpdateTime  string `json:"UpdateTime"`
}
