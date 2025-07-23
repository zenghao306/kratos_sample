package model

import (
	"gorm.io/plugin/soft_delete"
)

type User struct {
	UserID            uint64                 `gorm:"column:user_id;primaryKey;autoIncrement;comment:主键自增id"`
	Username          string                 `gorm:"column:username;type:varchar(60);not null;default:'';comment:用户名"`
	Password          string                 `gorm:"column:password;type:char(40);not null;default:'';comment:登录密码；md5(psd)"`
	Email             string                 `gorm:"column:email;type:varchar(100);not null;default:'';comment:邮箱地址"`
	Avatar            string                 `gorm:"column:avatar;type:varchar(1000);not null;default:'';comment:用户头像"`
	LastLoginIP       string                 `gorm:"column:last_login_ip;type:varchar(16);not null;default:'';comment:最后登录ip"`
	LastLoginTime     int64                  `gorm:"column:last_login_time;type:int unsigned;not null;default:0;comment:最后登录时间"`
	Status            int8                   `gorm:"column:status;type:tinyint;not null;default:1;comment:用户状态 0：未激活；1：激活；2：禁用；3：已删除"`
	FailedLoginCount  int16                  `gorm:"column:failed_login_count;type:smallint;not null;default:0;comment:连续登陆失败次数"`
	FailedLoginTime   int64                  `gorm:"column:failed_login_time;type:int;not null;default:0;comment:封锁账号起始时间"`
	ActiveCode        string                 `gorm:"column:active_code;type:char(32);not null;default:'';comment:激活码"`
	CodeTime          int64                  `gorm:"column:code_time;type:int unsigned;not null;default:0;comment:激活码创建时间"`
	ActiveTime        int64                  `gorm:"column:active_time;type:int unsigned;not null;default:0;comment:激活时间"`
	Remark            string                 `gorm:"column:remark;type:varchar(1000);not null;default:''"`
	CustomerID        int64                  `gorm:"column:customer_id;type:bigint;not null;default:0;comment:用户所属客户，无客户为0"`
	CompanyID         int64                  `gorm:"column:company_id;type:bigint;not null;default:0;comment:公司ID"`
	CategoryID        int64                  `gorm:"column:category_id;type:bigint;not null;default:0;comment:分组id"`
	GroupID           int64                  `gorm:"column:group_id;type:bigint;not null;default:0;comment:账户所属用户组"`
	AccessWorkspaceID int64                  `gorm:"column:access_workspace_id;type:bigint;not null;default:0;comment:可使用workspace"`
	ManageWorkspaceID int64                  `gorm:"column:manage_workspace_id;type:bigint;not null;default:0;comment:可管理workspace"`
	Origin            int8                   `gorm:"column:origin;type:tinyint;not null;default:0;comment:账号注册来源：0:开放注册;1:系统管理员邀请;2:客户邀请;3:UserAdmin邀请"`
	ValidityStatus    int8                   `gorm:"column:validity_status;type:tinyint(1);not null;default:1;comment:有效状态 1:永久有效；0：暂时"`
	ValidityDate      string                 `gorm:"column:validity_date;type:char(10);not null;comment:有效日期"`
	CreateBy          uint64                 `gorm:"column:create_by;type:bigint unsigned;not null;default:0;comment:创建人"`
	CreateTime        int64                  `gorm:"column:create_time;type:int unsigned;not null;default:0"`
	UpdateTime        int64                  `gorm:"column:update_time;type:int unsigned;not null;default:0"`
	DeleteTime        *soft_delete.DeletedAt `gorm:"softDelete:milli;default:null;index:idx_subfileid_deletetime"`
}

// TableName 设置表名
func (User) TableName() string {
	return "ds_user"
}
