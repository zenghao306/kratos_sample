package model

import (
	"gorm.io/datatypes"
	"gorm.io/plugin/soft_delete"
)

// DsDevice 设备列表
type DsDevice struct {
	//gorm.Model
	ID                    uint64                 `gorm:"primaryKey;autoIncrement;column:id;type:bigint unsigned;not null"`
	Name                  string                 `gorm:"column:name;type:varchar(256);not null;default:'';comment:设备名"`
	GroupID               int64                  `gorm:"column:group_id;type:bigint;not null;default:0;comment:所属组"`
	DeviceID              string                 `gorm:"column:device_id;type:char(40);not null;default:'';comment:机器的设备id"`
	Dnum                  int64                  `gorm:"column:dnum;type:bigint;not null;default:0;comment:设备编号（废弃 使用serial）"`
	Serial                string                 `gorm:"column:serial;type:char(60);not null;default:'';comment:序列号(TVID)"`
	ProjectID             uint                   `gorm:"column:projectid;type:int unsigned;not null;default:0;comment:Project ID"`
	IotID                 string                 `gorm:"column:iot_id;type:varchar(256);not null;default:'';comment:iot_id(加密)"`
	CertID                string                 `gorm:"column:cert_id;type:char(64);not null;default:'';comment:证书id"`
	Mac                   string                 `gorm:"column:mac;type:char(32);not null;default:'';comment:mac地址"`
	Sign                  *string                `gorm:"column:sign;type:char(64);comment:Mac的md5值"`
	ClientType            string                 `gorm:"column:client_type;type:varchar(100);not null;default:'';comment:Client Type"`
	SetShadow             uint8                  `gorm:"column:set_shadow;type:tinyint unsigned;not null;default:0;comment:设置影子"`
	Rotation              uint                   `gorm:"column:rotation;type:int unsigned;not null;default:0;comment:旋转角度 0~360"`
	SoftwareVersion       string                 `gorm:"column:software_version;type:char(30);not null;default:'';comment:系统版本"`
	AppVersion            string                 `gorm:"column:app_version;type:char(20);not null;default:'';comment:app版本"`
	Brand                 string                 `gorm:"column:brand;type:varchar(255);not null;default:'';comment:品牌"`
	Source                int8                   `gorm:"column:source;type:tinyint;not null;default:0;comment:数据来源: 0 moka主板Apk; 1 moka合作方Apk; 2 GoogleApk"`
	ModelID               uint64                 `gorm:"column:model_id;type:bigint unsigned;not null;default:0;comment:机型id"`
	ModelName             string                 `gorm:"column:model_name;type:varchar(100);not null;default:'';comment:机型名称"`
	HaveOss               int8                   `gorm:"column:have_oss;type:tinyint;not null;default:1;comment:是否拥有Oss分配空间 1:是 0:否（Model表冗余字段）"`
	FrontendStyle         int8                   `gorm:"column:frontend_style;type:tinyint;not null;default:0;comment:模型表-前端样式"`
	Status                uint8                  `gorm:"column:status;type:tinyint unsigned;not null;default:2;comment:设备状态 0 : 关机 1 : 开机 2 : 未激活"`
	ActiveBy              uint64                 `gorm:"column:active_by;type:bigint unsigned;not null;default:0;comment:激活人"`
	ActiveTime            int                    `gorm:"column:active_time;type:int;not null;default:0;comment:激活时间"`
	ConfirmBy             uint64                 `gorm:"column:confirm_by;type:bigint unsigned;not null;default:0;comment:激活审核人"`
	UserGroupID           int64                  `gorm:"column:user_group_id;type:bigint;not null;default:0;comment:所属用户组"`
	CreateTime            uint                   `gorm:"column:create_time;type:int unsigned;not null;default:0"`
	UpdateTime            uint                   `gorm:"column:update_time;type:int unsigned;not null;default:0"`
	DeleteTime            *soft_delete.DeletedAt `gorm:"softDelete:second;type:tinyint;default:null"`
	ScreenCaptureInterval int                    `gorm:"column:screen_capture_interval;type:int;not null;default:0;comment:捕捉屏幕间隔"`
	AutoScreenCapture     int16                  `gorm:"column:auto_screen_capture;type:smallint;not null;default:0;comment:自动刷新截图：0关闭；1开启"`
	PlayerFingerprint     string                 `gorm:"column:player_fingerprint;type:varchar(128);default:'';comment:播放器指纹"`
	PlayerComponents      datatypes.JSON         `gorm:"column:player_components;type:json;comment:播放器信息"`
}

func (DsDevice) TableName() string {
	return "ds_device"
}
