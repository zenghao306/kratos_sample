package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	device_pb "kratos_sample/api/device/v1"
	"kratos_sample/app/gateway/internal/router/device"
	"kratos_sample/app/gateway/internal/router/ginplus"
)

var ProviderSet = wire.NewSet(
	NewDiscovery,
	NewRegistrar,
	NewRouter,
	NewDeviceServiceClient, //设备服的发现
	NewUserServiceClient,   //用户服的发现
	NewServiceClients,
	device.NewDeviceHandler,
)

type Router struct {
	Engine         *gin.Engine
	serviceClients *ServiceClients       // 你的业务服务
	deviceHandler  *device.DeviceHandler // 注入 DeviceHandler
}

// NewRouter .
func NewRouter(
	logger log.Logger,
	serviceClients *ServiceClients, // 你的业务服务
	deviceHandler *device.DeviceHandler, // 注入 DeviceHandler

) (*Router, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Router{
		Engine:         gin.New(),
		serviceClients: serviceClients,
		deviceHandler:  deviceHandler,
	}, cleanup, nil
}

func (r *Router) SetupRoutes() {
	g := r.Engine.Group("/service")
	g.Use(ginplus.Cors()) //跨域
	v1 := g.Group("/v1")
	// 注册 device服务 的路由
	r.deviceHandler.RegisterRoutes(v1)

	// 注册路由
	v1.GET("/users/:id", r.getUser)
}

func (r *Router) getUser(c *gin.Context) {
	// 调用Kratos服务
	reply, err := r.serviceClients.deviceClient.ListDevice(c.Request.Context(), &device_pb.ListDeviceRequest{})
	if err != nil {
		c.JSON(500, gin.H{"error": "test"})
		return
	}
	c.JSON(200, reply)
}

func loggingMiddleware(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.NewHelper(logger).Infof(
			"[GIN] %s %s",
			c.Request.Method,
			c.Request.URL.Path,
		)
		c.Next()
	}
}
