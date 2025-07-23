package device

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	device_v1 "kratos_sample/api/device/v1"
	user_v1 "kratos_sample/api/user/v1"
	"kratos_sample/app/gateway/internal/conf"
	"kratos_sample/app/gateway/internal/router/ginplus"
)

type DeviceHandler struct {
	deviceClient    device_v1.DeviceClient // 设备服务
	userClient      user_v1.UserClient     // 用户服务
	redisSubscriber *RedisSubscriber       // Redis 消息订阅器
}

// NewDeviceHandler 创建 DeviceHandler 的构造函数
func NewDeviceHandler(
	deviceClient device_v1.DeviceClient,
	userClient user_v1.UserClient,
	c *conf.Data,
	logger log.Logger,
) *DeviceHandler {
	// 启动websoket客户端管理器
	go manager.Start()

	// 初始化 Redis 消息订阅器
	redisSubscriber := NewRedisSubscriber(c.Redis.Addr, c.Redis.Password, logger)
	// 启动 Redis 消息订阅
	redisSubscriber.Subscribe(c.Redis.TvPlayerTopic) // 从PHP接收的Redis 频道名称

	return &DeviceHandler{
		deviceClient:    deviceClient,
		userClient:      userClient,
		redisSubscriber: redisSubscriber,
	}
}

// RegisterRoutes 注册路由
func (r *DeviceHandler) RegisterRoutes(group *gin.RouterGroup) {
	deviceGroup := group.Group("/device")
	deviceGroup.Use(ginplus.CheckSessionToken()) //校验token并将token写入gin.Context

	deviceGroup.POST("/list", r.ListDevice)
	deviceGroup.GET("/ws", handleWebSocket)              //websocket长链接
	deviceGroup.POST("/webSocket/send", r.WebSocketSend) //发送消息给websocket
	deviceGroup.GET("/group/all-list", r.AllList)        //用户设备组

}
