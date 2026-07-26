package token_mart

type Response[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}
type BaseResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type createResponse struct {
	BaseResp
	Id string `json:"Id"`
}

type assetResponse struct {
	BaseResp
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
