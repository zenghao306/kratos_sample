package server

import (
	v1 "kratos_sample/api/user/v1"
	"kratos_sample/app/user/internal/conf"
	"kratos_sample/app/user/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.UserService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.UserGrpc.Network != "" {
		opts = append(opts, grpc.Network(c.UserGrpc.Network))
	}
	if c.UserGrpc.Addr != "" {
		opts = append(opts, grpc.Address(c.UserGrpc.Addr))
	}
	if c.UserGrpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.UserGrpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServer(srv, greeter)
	return srv
}
