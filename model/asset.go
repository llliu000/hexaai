package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm/clause"
)

const (
	AssetRegionDomestic = "cn-beijing"
	AssetRegionOverseas = "ap-southeast-1"

	AssetStatusPending    = "Pending"
	AssetStatusProcessing = "Processing"
	AssetStatusActive     = "Active"
	AssetStatusFailed     = "Failed"
)

// AssetGroup 资源分组
type AssetGroup struct {
	Id          string    `json:"id" gorm:"type:varchar(191);uniqueIndex"` // 分组ID，如 group-2026**********-*****
	Region      string    `json:"region" gorm:"type:varchar(32);not null;default:cn-beijing;index"`
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
	Id          string          `json:"id" gorm:"type:varchar(191);uniqueIndex"` // 资源ID，如 Asset-2026**********-*****
	Region      string          `json:"region" gorm:"type:varchar(32);not null;default:cn-beijing;index"`
	UserId      int             `json:"user_id" gorm:"index"`
	GroupId     string          `json:"group_id" gorm:"type:varchar(191);not null"`        // 分组ID
	URL         string          `json:"url" gorm:"type:text;not null"`                     // 资源地址
	Name        string          `json:"name" gorm:"type:varchar(191);not null;index"`      // 名称，上限为 64 个字符
	AssetType   string          `json:"asset_type" gorm:"type:varchar(50);not null;index"` // 资源类型，如 Image、Video、Audio
	ProjectName string          `json:"project_name" gorm:"type:varchar(191);index"`       // 项目名称，如 default
	Status      string          `json:"status" gorm:"type:varchar(50);not null;index"`     // 任务状态，如 Active、Processing、Failed
	Moderation  *string         `json:"moderation" gorm:"type:varchar(200)"`               // 审核配置 {"Strategy": "Skip"}
	Error       *datatypes.JSON `json:"error" gorm:"type:jsonb"`                           // 错误信息
	CreateTime  time.Time       `json:"create_time" gorm:"bigint"`
	UpdateTime  time.Time       `json:"update_time" gorm:"bigint"`
}

func (*Asset) TableName() string {
	return "asset"
}

func NormalizeAssetRegion(region string) string {
	if region == "" {
		return AssetRegionDomestic
	}
	return region
}

type AssertChannel struct {
	Id               int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Region           string    `json:"region" gorm:"type:varchar(32);not null;default:cn-beijing;index"`
	UserId           int       `json:"user_id" gorm:"not null;default:0;index"`
	ChannelId        int       `json:"channel_id" gorm:"type:bigint;not null;uniqueIndex:idx_asset_channel_channel_asset"`                                                          // 上游渠道ID
	AssetId          string    `json:"asset_id" gorm:"type:varchar(1000);not null;index;uniqueIndex:idx_asset_channel_channel_asset"`                                               // 本地资源ID
	Manual           bool      `json:"manual" gorm:"default:false;index:idx_asset_channel_manual_status_poll,priority:1"`                                                           //
	UpstreamGroupId  string    `json:"upstream_group_id" gorm:"type:varchar(191);not null;index"`                                                                                   // 上游资源分组ID
	UpstreamAssertId string    `json:"upstream_assert_id" gorm:"type:varchar(191);not null;index"`                                                                                  // 上游资源ID
	UpstreamStatus   string    `json:"upstream_status" gorm:"type:varchar(50);not null;index"`                                                                                      // 任务状态，如 Active、Processing、Failed
	Reason           string    `json:"reason" gorm:"type:varchar(500);not null"`                                                                                                    // 失败原因
	RetryCount       int       `json:"retry_count" gorm:"not null;default:0"`                                                                                                       // 同步请求连续失败次数，成功后清零。
	NextPollAt       int64     `json:"next_poll_at" gorm:"not null;default:0;index:idx_asset_channel_status_poll,priority:2;index:idx_asset_channel_manual_status_poll,priority:3"` // 下次同步的 Unix 秒，0 表示未调度。
	CreateTime       time.Time `json:"create_time" gorm:"bigint"`
	UpdateTime       time.Time `json:"update_time" gorm:"bigint"`

	Asset Asset `gorm:"foreignKey:AssetId;references:Id"`
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

type ChannelUpstreamAssetGroup struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Region          string    `json:"region" gorm:"type:varchar(32);not null;default:cn-beijing;index"`
	ChannelId       int       `json:"channel_id" gorm:"not null;index;uniqueIndex:uk_asset_group_channel"`
	GroupId         string    `json:"group_id" gorm:"type:varchar(191);not null;default:'';uniqueIndex:uk_asset_group_channel"`
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
