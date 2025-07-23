package server

import (
	v1 "kratos_sample/api/device/v1"
	"kratos_sample/app/device/internal/conf"
	"kratos_sample/app/device/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.DeviceService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.DeviceGrpc.Network != "" {
		opts = append(opts, grpc.Network(c.DeviceGrpc.Network))
	}
	if c.DeviceGrpc.Addr != "" {
		opts = append(opts, grpc.Address(c.DeviceGrpc.Addr))
	}
	if c.DeviceGrpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.DeviceGrpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterDeviceServer(srv, greeter)
	return srv
}
