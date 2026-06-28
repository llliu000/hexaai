package anyfast

type createAssetGroupRequest struct {
	Model string `json:"model"`
	Name  string `json:"Name,omitempty"`
}

type listAssetGroupRequest struct {
	Model  string `json:"model"`
	Filter struct {
		Name      *string  `json:"Name"`
		GroupIds  []string `json:"GroupIds"`
		GroupType *string  `json:"GroupType"`
	} `json:"Filter"`
	PageNumber int `json:"PageNumber"`
	PageSize   int `json:"PageSize"`
}

type createAssetRequest struct {
	Model     string `json:"model"`
	GroupId   string `json:"GroupId,omitempty"`
	URL       string `json:"URL,omitempty"`
	Name      string `json:"Name,omitempty"`
	AssetType string `json:"AssetType,omitempty"`
}

type listAssetsFilter struct {
	Name      *string  `json:"Name,omitempty"`
	GroupIds  []string `json:"GroupIds,omitempty"`
	GroupType *string  `json:"GroupType,omitempty"`
}

type listAssetsRequest struct {
	Model      string            `json:"model"`
	Filter     *listAssetsFilter `json:"Filter,omitempty"`
	PageNumber int               `json:"PageNumber,omitempty"`
	PageSize   int               `json:"PageSize,omitempty"`
}
