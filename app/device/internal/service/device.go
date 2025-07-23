package service

import (
	"context"
	pb "kratos_sample/api/device/v1"
	"kratos_sample/app/device/internal/biz"
)

type DeviceService struct {
	pb.UnimplementedDeviceServer
	uc *biz.DeviceUsecase
}

func NewDeviceService(uc *biz.DeviceUsecase) *DeviceService {
	return &DeviceService{uc: uc}
}

func (s *DeviceService) ListDevice(ctx context.Context, req *pb.ListDeviceRequest) (*pb.ListDeviceReply, error) {
	return s.uc.DeviceList(ctx, req)
}

func (s *DeviceService) DeviceGroupAllList(ctx context.Context, req *pb.IDRequest) (*pb.DeviceGroupListResponse, error) {
	return s.uc.DeviceGroupAllList(ctx, req)
}
