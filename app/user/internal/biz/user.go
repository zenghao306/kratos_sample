package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos_sample/api/user/v1"
)

// User is a User model.
type User struct {
}

// UserRepo is a Greater repo.
type UserRepo interface {
	GetUserGroupId(ctx context.Context, req *pb.GetUserGroupRequest) (*pb.GroupIdResponse, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) GetUserGroupId(ctx context.Context, req *pb.GetUserGroupRequest) (*pb.GroupIdResponse, error) {
	return uc.repo.GetUserGroupId(ctx, req)
}
