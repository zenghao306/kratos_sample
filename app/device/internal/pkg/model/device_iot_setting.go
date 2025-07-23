package model

import (
	"encoding/json"
	"gorm.io/datatypes"
	"gorm.io/plugin/soft_delete"
)

// DeviceIotSetting 对应 ds_device_iot_setting 表
type DeviceIotSetting struct {
	SettingID    uint64                 `gorm:"primaryKey;column:setting_id;autoIncrement;comment:主键"`
	DeviceID     int64                  `gorm:"column:device_id;not null;comment:设备表自增id"`
	Settings     datatypes.JSON         `gorm:"column:settings;not null;comment:iot终端数据"`
	OriginalData datatypes.JSON         `gorm:"column:original_data;comment:iot终端原数据"`
	CreateTime   int64                  `gorm:"column:create_time;not null;default:0"`
	UpdateTime   int64                  `gorm:"column:update_time;not null;default:0"`
	DeleteTime   *soft_delete.DeletedAt `gorm:"softDelete:second;type:tinyint;default:null"`
}

// TableName 设置表名
func (DeviceIotSetting) TableName() string {
	return "ds_device_iot_setting"
}

type Setting struct {
	ClearCache       int32 `json:"clear_cache"`
	DisplayRatioType int32 `json:"display_ratio_type"`
}

func (d DeviceIotSetting) UnmarshalSettings() Setting {
	var t Setting
	json.Unmarshal(d.Settings, &t)
	return t
}
