package kw

type kwCreateAssetGroupRequest struct {
	Model       string  `json:"model"`
	Name        string  `json:"Name,omitempty"`
	GroupType   string  `json:"GroupType,omitempty"`
	Description *string `json:"Description,omitempty"`
}

type kwListAssetGroupRequest struct {
	Model  string `json:"model"`
	Filter struct {
		GroupType *string `json:"GroupType"`
		Name      *string `json:"name"`
	} `json:"Filter"`
	PageNumber int    `json:"PageNumber"`
	PageSize   int    `json:"PageSize"`
	SortBy     string `json:"SortBy"`
	SortOrder  string `json:"SortOrder"`
}

type kwCreateAssetRequest struct {
	Model      string        `json:"model"`
	GroupId    string        `json:"GroupId,omitempty"`
	URL        string        `json:"URL,omitempty"`
	Name       string        `json:"Name,omitempty"`
	AssetType  string        `json:"AssetType,omitempty"`
	Moderation *kwModeration `json:"Moderation,omitempty"`
}

type kwModeration struct {
	Strategy string `json:"Strategy"`
}

type kwGetAssetRequest struct {
	Model string `json:"model"`
	Id    string `json:"Id"`
}

type kwListAssetsFilter struct {
	GroupType *string  `json:"GroupType,omitempty"`
	GroupIds  []string `json:"GroupIds,omitempty"`
	Statuses  []string `json:"Statuses,omitempty"`
	Name      *string  `json:"Name,omitempty"`
}

type kwListAssetsRequest struct {
	Model      string              `json:"model"`
	Filter     *kwListAssetsFilter `json:"Filter,omitempty"`
	PageNumber int                 `json:"PageNumber"`
	PageSize   int                 `json:"PageSize"`
	SortBy     string              `json:"SortBy,omitempty"`
	SortOrder  string              `json:"SortOrder,omitempty"`
}
