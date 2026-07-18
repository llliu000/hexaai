package asset

import (
	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/relay/asset/anyfast"
	"github.com/QuantumNous/new-api/relay/asset/kw"
	"github.com/QuantumNous/new-api/relay/asset/volc"
)

var (
	UpstreamAssetGroupName = "6666-xxxfkafjlasd-123"
	GroupType              = "AIGC"
)

type Adaptor interface {
	CreateAssetGroup(req *dto.CreateAssetGroupRequest) (*dto.CreateAssetGroupResponse, error)
	ListAssetGroups(req *dto.ListAssetGroupsRequest) (*dto.ListAssetGroupsResponse, error)
	CreateAssets(req *dto.CreateAssetRequest) (*dto.CreateAssetResponse, error)
	GetAsset(req *dto.GetAssetRequest) (*dto.GetAssetResponse, error)
	ListAssets(req *dto.ListAssetsRequest) (*dto.ListAssetsResponse, error)
	ReviewSkip() bool
}

func GetAdaptor(ch *model.Channel) Adaptor {
	if ch == nil || ch.OpenAIOrganization == nil {
		return nil
	}
	var baseUrl string
	if ch.BaseURL != nil {
		baseUrl = *ch.BaseURL
	}
	var a Adaptor
	switch *ch.OpenAIOrganization {
	case "kwjm":
		a = &kw.Adaptor{ApiKey: ch.Key, BaseURL: baseUrl}
	case "anyfast":
		a = &anyfast.Adaptor{ApiKey: ch.Key, BaseURL: baseUrl}
	case "volc":
		a = &volc.Adaptor{}
	}
	if remark := ch.Remark; remark != nil {
		_ = common.Unmarshal([]byte(*remark), a)
	}
	return a
}
