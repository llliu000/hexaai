package kw

type kwCreateResponse struct {
	IdLower string `json:"id"`
	IdUpper string `json:"Id"`
}

func (r kwCreateResponse) id() string {
	if r.IdLower != "" {
		return r.IdLower
	}
	return r.IdUpper
}

type kwListAssetGroupResponse struct {
	TotalCount int64              `json:"TotalCount"`
	PageNumber int                `json:"PageNumber"`
	PageSize   int                `json:"PageSize"`
	Items      []kwAssetGroupItem `json:"Items"`
}

type kwAssetGroupItem struct {
	Id         string `json:"Id"`
	Name       string `json:"Name"`
	GroupType  string `json:"GroupType"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type kwListAssetsResponse struct {
	Items      []kwAssetItem `json:"Items"`
	TotalCount int64         `json:"TotalCount"`
	PageNumber int           `json:"PageNumber"`
	PageSize   int           `json:"PageSize"`
}

type kwAssetItem struct {
	Id          string `json:"Id"`
	Name        string `json:"Name"`
	URL         string `json:"URL"`
	GroupId     string `json:"GroupId"`
	AssetType   string `json:"AssetType"`
	Status      string `json:"Status"`
	CreatedTime string `json:"createTime"`
	UpdatedTime string `json:"updateTime"`
}
