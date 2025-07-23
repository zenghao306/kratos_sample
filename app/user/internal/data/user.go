package data

import (
	"context"
	user_v1 "kratos_sample/api/user/v1"
	"kratos_sample/app/user/internal/pkg/model"

	"github.com/go-kratos/kratos/v2/log"
	//"kratos_sample/app/user/internal/pkg/model"
	//"fmt"
	"kratos_sample/app/user/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewDeviceRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) GetUserGroupId(ctx context.Context, params *user_v1.GetUserGroupRequest) (*user_v1.GroupIdResponse, error) {
	var (
		user model.User
	)

	if err := r.data.db.Where("user_id = ?", params.UserId).First(&user).Error; err != nil { //执行分页查询
		log.Infof("[data.userRepo.GetUserGroupId] 查询entity.User,userId:%v数据对象失败,err:%s", params.UserId, err)
		return nil, err
	}

	res := &user_v1.GroupIdResponse{
		GroupId: user.GroupID,
	}

	return res, nil
}
