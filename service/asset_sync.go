package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/model"
	assetrelay "github.com/QuantumNous/new-api/relay/asset"
	"gorm.io/gorm"
)

var assetSyncQueue = make(chan *assetSyncTask, 1024)
var assetStatusSyncQueue = make(chan int, 1024)
var assetChannelSyncQueue = make(chan int, 20)

type assetSyncTask struct {
	AssetId string
	Update  bool
}

func init() {
	go func() {
		for task := range assetSyncQueue {
			log.Printf("新增豆包视频资源开始上游通过:local asset_id=%s\n", task.AssetId)
			if err := syncAssetToUpstreams(task); err != nil {
				common.SysError(fmt.Sprintf("asset sync failed, asset_id=%s: %s", task.AssetId, err.Error()))
			}
		}
	}()
	go func() {
		for abId := range assetStatusSyncQueue {
			log.Printf("新增上游资源后开始同步上游状态到本地:asset_channel_id=%d\n", abId)
			if err := syncAssetStatusToLocal(abId); err != nil {
				common.SysError(fmt.Sprintf("asset url sync failed, ab_id=%d: %s", abId, err.Error()))
			}
		}
	}()
	go func() {
		for chId := range assetChannelSyncQueue {
			log.Printf("新增channel后开始同步豆包视频资源:channel_id=%d\n", chId)
			if err := syncChannelAsset(chId); err != nil {
				common.SysError(fmt.Sprintf("asset sync failed, ac_id=%d: %s", chId, err.Error()))
			}
		}
	}()
}

func enqueueAssetSync(assetId string, update bool) {
	task := assetSyncTask{AssetId: assetId, Update: update}
	select {
	case assetSyncQueue <- &task:
	default:
		common.SysError(fmt.Sprintf("asset sync queue is full, skip asset_id=%s", task.AssetId))
	}
}

func EnqueueChannelAssetSync(chId int) {
	select {
	case assetChannelSyncQueue <- chId:
	default:
		common.SysError(fmt.Sprintf("channel asset sync queue is full, skip ch_id=%d", chId))
	}
}

func syncAssetToUpstreams(task *assetSyncTask) error {
	// 获取本地资源信息
	var localAsset model.Asset
	localAsset, err := model.GetAssetById(task.AssetId)
	if nil != err {
		return err
	}
	// 如果是修改操作就删除映射关系后再新增
	if task.Update {
		if err = model.DeleteAssertChannelByAssetId(task.AssetId); nil != err {
			return err
		}
	}

	// 查询所有可用豆包视频渠道
	chs, err := model.ListChannelsByOpenAIOrganization()
	if err != nil {
		return err
	}
	var errs []string
	for i := range chs {
		if err = createAssetOnChannel(&localAsset, &chs[i]); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}
	return nil
}

func syncAssetStatusToLocal(acId int) error {
	ab, err := model.GetAssetChannelById(acId)
	if err != nil {
		return err
	}
	channel, err := model.GetChannelById(ab.ChannelId, true)
	if err != nil {
		return err
	}
	adaptor := assetrelay.GetAdaptor(channel)
	if adaptor == nil {
		return nil
	}
	request := dto.GetAssetRequest{
		BaseAssetRequest: dto.BaseAssetRequest{
			ProjectName: &ab.UpstreamGroupId,
		},
		Id: ab.UpstreamAssertId,
	}
	asset, err := adaptor.GetAsset(&request)
	if err != nil {
		return err
	}
	if asset.Status == "Processing" {
		go func() {
			<-time.After(time.Second * 3)
			assetStatusSyncQueue <- acId
		}()
		return nil
	}
	err = model.UpdateAssetById(ab.AssetId, asset.Status)
	// TODO 同步映射状态
	return err
}

func syncChannelAsset(chId int) error {
	channel, err := model.GetChannelById(chId, true)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return err
	}
	if channel.Type != constant.ChannelTypeDoubaoVideo {
		return nil
	}
	adaptor := assetrelay.GetAdaptor(channel)
	if adaptor == nil {
		return nil
	}
	upstreamAssetGroupId, err := getUpstreamAssetGroupId(adaptor, chId)
	if err != nil {
		return err
	}

	const pageSize = 100
	pageNumber := 1
	for {
		assets, err := model.ListAssetsByPage(pageNumber, pageSize)
		if err != nil {
			return err
		}
		if len(assets) == 0 {
			return nil
		}
		for i := range assets {
			_ = createUpstreamAsset(adaptor, &assets[i], upstreamAssetGroupId, channel.Id)
		}
		if len(assets) < pageSize {
			return nil
		}
		pageNumber++
	}
}

func createAssetOnChannel(localAsset *model.Asset, channel *model.Channel) error {
	adaptor := assetrelay.GetAdaptor(channel)
	if adaptor == nil {
		return nil
	}
	// 上游资源分组ID处理
	upstreamGroupId, err := getUpstreamAssetGroupId(adaptor, channel.Id)
	if err != nil {
		return err
	}
	err = createUpstreamAsset(adaptor, localAsset, upstreamGroupId, channel.Id)
	return err
}

func createUpstreamAsset(adaptor assetrelay.Adaptor, localAsset *model.Asset, upstreamAssetGroupId string, channelId int) error {
	// 判断资源是否免审
	if localAsset.ReviewSkip() && !adaptor.ReviewSkip() {
		return nil
	}
	request := dto.CreateAssetRequest{
		BaseAssetRequest: dto.BaseAssetRequest{
			ProjectName: &localAsset.ProjectName,
		},
		URL:       localAsset.URL,
		GroupId:   upstreamAssetGroupId,
		Name:      localAsset.Name,
		AssetType: localAsset.AssetType,
	}
	if localAsset.ReviewSkip() && adaptor.ReviewSkip() {
		data := []byte(*localAsset.Moderation)
		var moderation dto.AssetModeration
		if err := json.Unmarshal(data, &moderation); err == nil {
			request.Moderation = &moderation
		}
	}
	response, err := adaptor.CreateAssets(&request)
	if err != nil {
		return err
	}
	ac := model.AssertChannel{
		ChannelId:        channelId,
		AssetId:          localAsset.Id,
		UpstreamGroupId:  upstreamAssetGroupId,
		UpstreamAssertId: response.Id,
		CreateTime:       time.Now(),
		UpdateTime:       time.Now(),
	}
	if err = ac.Create(); nil != err {
		return err
	}
	assetStatusSyncQueue <- ac.Id
	return nil
}

func getUpstreamAssetGroupId(adaptor assetrelay.Adaptor, channelId int) (string, error) {
	channelUpstreamAssetGroup, err := model.GetChannelUpstreamAssetGroup(channelId)
	if err != nil {
		return "", err
	}
	if channelUpstreamAssetGroup.UpstreamGroupId != "" {
		return channelUpstreamAssetGroup.UpstreamGroupId, nil
	}
	group, err := adaptor.CreateAssetGroup(&dto.CreateAssetGroupRequest{
		Name:      assetrelay.UpstreamAssetGroupName,
		GroupType: assetrelay.GroupType,
	})
	if err != nil {
		return "", err
	}
	channelUpstreamAssetGroup = model.ChannelUpstreamAssetGroup{
		UpdateTime:      time.Now(),
		CreateTime:      time.Now(),
		ChannelId:       channelId,
		UpstreamGroupId: group.Id,
	}
	err = channelUpstreamAssetGroup.Create()
	return group.Id, err
}
