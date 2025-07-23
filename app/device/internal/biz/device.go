package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos_sample/api/device/v1"
)

//var (
//	// ErrUserNotFound is user not found.
//	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
//)

// Device is a Device model.
type Device struct {
	Hello string
}

// DeviceRepo is a Greater repo.
type DeviceRepo interface {
	ListDevice(ctx context.Context, in *pb.ListDeviceRequest) (*pb.ListDeviceReply, error)
	DeviceGroupAllList(ctx context.Context, in *pb.IDRequest) (*pb.DeviceGroupListResponse, error)
}

//ListDevice(ctx context.Context, in *ListDeviceRequest, opts ...grpc.CallOption) (*ListDeviceReply, error)

// DeviceUsecase is a Device usecase.
type DeviceUsecase struct {
	repo DeviceRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Device usecase.
func NewDeviceUsecase(repo DeviceRepo, logger log.Logger) *DeviceUsecase {
	return &DeviceUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DeviceUsecase) DeviceList(ctx context.Context, req *pb.ListDeviceRequest) (*pb.ListDeviceReply, error) {
	return uc.repo.ListDevice(ctx, req)
}

func (uc *DeviceUsecase) DeviceGroupAllList(ctx context.Context, req *pb.IDRequest) (*pb.DeviceGroupListResponse, error) {
	return uc.repo.DeviceGroupAllList(ctx, req)
}
