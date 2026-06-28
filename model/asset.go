package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm/clause"
)

// AssetGroup 资源分组
type AssetGroup struct {
	Id          string    `json:"id" gorm:"type:varchar(191);uniqueIndex"` // 分组ID，如 group-2026**********-*****
	UserId      int       `json:"user_id" gorm:"index"`
	Name        string    `json:"name" gorm:"type:varchar(191);not null;index"`      // 名称，上限为64个字符
	Description string    `json:"description" gorm:"type:varchar(500)"`              // 描述，上限为 300 字符
	GroupType   string    `json:"group_type" gorm:"type:varchar(50);not null;index"` // 分组类型，如 AIGC、LivenessFace
	ProjectName string    `json:"project_name" gorm:"type:varchar(191);index"`       // 项目名称，如 default
	CreateTime  time.Time `json:"create_time" gorm:"bigint"`
	UpdateTime  time.Time `json:"update_time" gorm:"bigint"`
}

func (AssetGroup) TableName() string {
	return "asset_group"
}

// Asset 资源
type Asset struct {
	Id          string    `json:"id" gorm:"type:varchar(191);uniqueIndex"` // 资源ID，如 Asset-2026**********-*****
	UserId      int       `json:"user_id" gorm:"index"`
	GroupId     string    `json:"group_id" gorm:"type:varchar(191);not null;index"`  // 分组ID
	URL         string    `json:"url" gorm:"type:text;not null"`                     // 资源地址
	Name        string    `json:"name" gorm:"type:varchar(191);not null;index"`      // 名称，上限为 64 个字符
	AssetType   string    `json:"asset_type" gorm:"type:varchar(50);not null;index"` // 资源类型，如 Image、Video、Audio
	ProjectName string    `json:"project_name" gorm:"type:varchar(191);index"`       // 项目名称，如 default
	Status      string    `json:"status" gorm:"type:varchar(50);not null;index"`     // 任务状态，如 Active、Processing、Failed
	Moderation  *string   `json:"moderation" gorm:"type:varchar(200)"`               // 审核配置 {"Strategy": "Skip"}
	CreateTime  time.Time `json:"create_time" gorm:"bigint"`
	UpdateTime  time.Time `json:"update_time" gorm:"bigint"`
}

func (*Asset) TableName() string {
	return "asset"
}

func (a *Asset) ReviewSkip() bool {
	if a.Moderation == nil {
		return false
	}
	m := make(map[string]any)
	_ = json.Unmarshal([]byte(*a.Moderation), &m)
	s, ok := m["Strategy"]
	if !ok {
		return false
	}
	strategy, ok := s.(string)
	if !ok {
		return false
	}
	return strategy == "Skip"
}

func ListAssetsByPage(pageNumber int, pageSize int) (assets []Asset, err error) {
	if pageNumber < 1 {
		pageNumber = 1
	}
	if pageSize < 1 {
		pageSize = 100
	}
	err = DB.Order("create_time ASC").
		Order("id ASC").
		Limit(pageSize).
		Offset((pageNumber - 1) * pageSize).
		Find(&assets).Error
	return
}

func GetAssetById(id string) (asset Asset, err error) {
	err = DB.Where("id = ?", id).First(&asset).Error
	return
}

func UpdateAssetById(id, status string) (err error) {
	return DB.Model(&Asset{}).Where("id = ?", id).Update("status", status).Error
}

type AssertChannel struct {
	Id               int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ChannelId        int       `json:"channel_id" gorm:"type:bigint;not null;uniqueIndex:idx_asset_channel_channel_asset"`           // 上游渠道ID
	AssetId          string    `json:"asset_id" gorm:"type:varchar(191);not null;index;uniqueIndex:idx_asset_channel_channel_asset"` // 本地资源ID
	UpstreamGroupId  string    `json:"upstream_group_id" gorm:"type:varchar(191);not null;index"`                                    // 上游资源分组ID
	UpstreamAssertId string    `json:"upstream_assert_id" gorm:"type:varchar(191);not null;index"`                                   // 上游资源ID
	UpstreamStatus   string    `json:"upstream_status" gorm:"type:varchar(50);not null;index"`                                       // 任务状态，如 Active、Processing、Failed
	Reason           string    `json:"reason" gorm:"type:varchar(500);not null"`                                                     // 失败原因
	CreateTime       time.Time `json:"create_time" gorm:"bigint"`
	UpdateTime       time.Time `json:"update_time" gorm:"bigint"`
}

func (*AssertChannel) TableName() string {
	return "asset_channel"
}

func (a *AssertChannel) Create() error {
	return DB.Create(a).Error
}
func GetAssetChannelByChannelIdAndAssetId(channelId int, assetId string) (ac AssertChannel, err error) {
	err = DB.Where("channel_id = ? AND asset_id = ?", channelId, assetId).First(&ac).Error
	return
}

func GetAssetChannelById(id int) (ab AssertChannel, err error) {
	err = DB.Model(&AssertChannel{}).Where("id=?", id).Find(&ab).Error
	return
}

func DeleteAssertChannelByAssetId(assetId string) error {
	return DB.Where("asset_id = ?", assetId).Delete(&AssertChannel{}).Error
}

type ChannelUpstreamAssetGroup struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ChannelId       int       `json:"channel_id" gorm:"type:bigint;not null;uniqueIndex"`
	UpstreamGroupId string    `json:"upstream_group_id" gorm:"type:varchar(191);not null;index"`
	CreateTime      time.Time `json:"create_time" gorm:"bigint"`
	UpdateTime      time.Time `json:"update_time" gorm:"bigint"`
}

func (*ChannelUpstreamAssetGroup) TableName() string {
	return "channel_upstream_asset_group"
}

func (a *ChannelUpstreamAssetGroup) Create() error {
	return DB.Clauses(clause.OnConflict{DoNothing: true}).Create(a).Error
}

func GetChannelUpstreamAssetGroup(channelId int) (ab ChannelUpstreamAssetGroup, err error) {
	err = DB.Where("channel_id = ?", channelId).Find(&ab).Error
	return
}
