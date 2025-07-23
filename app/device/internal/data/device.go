package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	device_v1 "kratos_sample/api/device/v1"
	"kratos_sample/app/device/internal/biz"
	"kratos_sample/app/device/internal/pkg/model"
)

type deviceRepo struct {
	data *Data
	log  *log.Helper
}

// NewDeviceRepo .
func NewDeviceRepo(data *Data, logger log.Logger) biz.DeviceRepo {
	return &deviceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *deviceRepo) ListDevice(ctx context.Context, params *device_v1.ListDeviceRequest) (*device_v1.ListDeviceReply, error) {
	var (
		sli    []*device_v1.DeviceInfo
		total  int64
		result = &device_v1.ListDeviceReply{}
	)

	query := r.data.db.Model(&model.DsDevice{})
	if params.Name != "" {
		query = query.Where("name LIKE ?", params.Name+"%")
	}
	if params.GroupId > 0 {
		query = query.Where("group_id = ?", params.GroupId)
	}
	if params.PlayerFingerprint != "" { //播放器指纹查找
		query = query.Where("player_fingerprint = ?", params.PlayerFingerprint)
	}

	if err := query.Count(&total).Error; err != nil { //获取总条数
		log.Errorf("[data.deviceRepo.ListDevice] 查询总条数失败,err:%s", err)
		return result, err
	}
	result.Total = total

	offset := (params.Page - 1) * params.PageSize
	if err := query.Limit(int(params.PageSize)).Offset(int(offset)).Find(&sli).Error; err != nil { //执行分页查询
		log.Errorf("[data.deviceRepo.ListDevice] 查询entity.DsDevice数据对象失败,err:%s", err)
		return result, err
	}
	result.Data = sli

	return result, nil
}

func (r *deviceRepo) DeviceGroupAllList(ctx context.Context, params *device_v1.IDRequest) (*device_v1.DeviceGroupListResponse, error) {
	var (
		sli    []*device_v1.DeviceGroupInfo
		result = &device_v1.DeviceGroupListResponse{}
	)

	//只查询group_id,name，lock这几个字段，过多数据库字段查询会拖慢速度
	if err := r.data.db.Model(&model.DsDeviceGroup{}).Select("group_id", "name", "lock").Where("user_group_id = ?", params.Id).Order("group_id desc").Find(&sli).Error; err != nil {
		log.Errorf("[data.deviceRepo.DeviceGroupAllList] 查询entity.DsDevice数据对象失败,err:%s", err)
		return result, err
	}

	//查询每个分组下的设备数量
	groupIds := make([]int32, 0)
	for _, g := range sli {
		groupIds = append(groupIds, g.GroupId)
	}
	//依次查询设备列表，统计每个group下设备数量
	if len(groupIds) > 0 {
		sliC := make([]model.IdCount, 0)
		if err := r.data.db.Model(&model.DsDevice{}).Select("group_id as id", "COUNT(id) AS count").
			Where("group_id in (?)", groupIds).Group("group_id").Find(&sliC).Error; err != nil {
		}
		for i := 0; i < len(sli); i++ {
			for _, g := range sliC {
				if sli[i].GroupId == g.ID {
					sli[i].Count = g.Count
				}
			}
		}
	}
	result.List = sli

	return result, nil
}
