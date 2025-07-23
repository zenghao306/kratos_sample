package router

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	grpcx "google.golang.org/grpc"
	device_v1 "kratos_sample/api/device/v1"
	user_v1 "kratos_sample/api/user/v1"
	"kratos_sample/app/gateway/internal/conf"
	"time"
)

type ServiceClients struct {
	deviceClient device_v1.DeviceClient // 你的业务服务
}

func NewServiceClients(dc device_v1.DeviceClient) *ServiceClients {
	return &ServiceClients{deviceClient: dc}
}

func NewRegistrar(etcdpoint *conf.Etcd) (registry.Registrar, func(), error) {
	// ETCD源地址
	endpoint := []string{etcdpoint.Address}

	// ETCD配置信息
	etcdCfg := clientv3.Config{
		Endpoints:   endpoint,
		DialTimeout: time.Second,
		DialOptions: []grpcx.DialOption{grpcx.WithBlock()},
	}

	// 创建ETCD客户端
	client, err := clientv3.New(etcdCfg)
	if err != nil {
		panic(err)
	}
	clean := func() {
		_ = client.Close()
	}

	// 创建服务注册 reg
	regi := etcd.New(client)

	return regi, clean, nil
}

func NewDiscovery(etcdpoint *conf.Etcd) registry.Discovery {
	// ETCD源地址
	endpoint := []string{etcdpoint.Address}

	// ETCD配置信息
	etcdCfg := clientv3.Config{
		Endpoints:   endpoint,
		DialTimeout: time.Second,
		DialOptions: []grpcx.DialOption{grpcx.WithBlock()},
	}

	// 创建ETCD客户端
	client, err := clientv3.New(etcdCfg)
	if err != nil {
		panic(err)
	}

	// new dis with etcd client
	dis := etcd.New(client)
	return dis

}

func NewDeviceServiceClient(covery registry.Discovery, s *conf.Service) device_v1.DeviceClient {
	// ETCD源地址 discovery:///device.service
	endpoint := s.Device.Endpoint
	conn, err := grpc.DialInsecure( //不使用TLS
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(covery),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(5*time.Second),
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)
	if err != nil {
		panic(err)
	}
	c := device_v1.NewDeviceClient(conn)
	return c
}

func NewUserServiceClient(covery registry.Discovery, s *conf.Service) user_v1.UserClient {
	// ETCD源地址 discovery:///user.service
	endpoint := s.User.Endpoint
	conn, err := grpc.DialInsecure( //不使用TLS
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(covery),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(5*time.Second),
		grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)
	if err != nil {
		panic(err)
	}
	c := user_v1.NewUserClient(conn)
	return c
}
