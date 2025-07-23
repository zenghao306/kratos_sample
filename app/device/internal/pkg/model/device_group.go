package model

import (
	"gorm.io/plugin/soft_delete"
)

type DsDeviceGroup struct {
	GroupID     uint64                 `gorm:"column:group_id;primaryKey;autoIncrement;not null"`
	Name        string                 `gorm:"column:name;type:varchar(60);not null;default:''"`
	Description string                 `gorm:"column:description;type:varchar(500);not null;default:'';comment:描述"`
	Lock        int8                   `gorm:"column:lock;type:tinyint(4);not null;default:0;comment:系统创建：0:否, 1:是"`
	UserGroupID uint64                 `gorm:"column:user_group_id;type:bigint(20);not null;default:0;comment:所属用户组;index:fk_ds_device_group_ds_group1_idx"`
	CreateBy    uint64                 `gorm:"column:create_by;type:bigint(20);not null;default:0;comment:创建人;index:fk_ds_device_group_ds_user1_idx"`
	CreateTime  int64                  `gorm:"column:create_time;type:int(10) unsigned;not null;default:0"`
	UpdateTime  int64                  `gorm:"column:update_time;type:int(10) unsigned;not null;default:0"`
	DeleteTime  *soft_delete.DeletedAt `gorm:"softDelete:second;type:tinyint;default:null"`
}

// TableName 设置表名
func (DsDeviceGroup) TableName() string {
	return "ds_device_group"
}

type IdCount struct {
	ID    int32 `gorm:"primary_key;column:id"` //id
	Count int32 `gorm:"column:count"`
}
