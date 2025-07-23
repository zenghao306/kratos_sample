package service

import (
	"context"
	pb "kratos_sample/api/user/v1"
	"kratos_sample/app/user/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) GetUserGroupId(ctx context.Context, req *pb.GetUserGroupRequest) (*pb.GroupIdResponse, error) {
	return s.uc.GetUserGroupId(ctx, req)
}
