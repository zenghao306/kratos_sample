package device

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"kratos_sample/app/gateway/internal/router/ginplus"
)

// RedisSubscriber 用于订阅 Redis 消息
type RedisSubscriber struct {
	client *redis.Client
	logger log.Logger
}

// NewRedisSubscriber 创建 RedisSubscriber
func NewRedisSubscriber(redisAddr string, redisPassword string, logger log.Logger) *RedisSubscriber {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})
	return &RedisSubscriber{
		client: client,
		logger: logger,
	}
}

// Subscribe 订阅 Redis 消息
func (r *RedisSubscriber) Subscribe(channel string) {
	ctx := context.Background()
	pubsub := r.client.Subscribe(ctx, channel)

	// 等待订阅确认
	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.NewHelper(r.logger).Errorf("Failed to subscribe to channel %s: %v", channel, err)
		return
	}

	// 启动消息处理协程
	go func() {
		ch := pubsub.Channel()
		for msg := range ch {
			log.NewHelper(r.logger).Infof("Received message from channel %s: %s", channel, msg.Payload)

			// 处理消息，广播到 WebSocket 客户端
			r.handleMessage(msg.Payload)
		}
	}()
}

// handleMessage 处理接收到的 Redis 消息
// handleMessage 处理接收到的 Redis 消息
func (r *RedisSubscriber) handleMessage(message string) {
	var reqData ReqData
	err := json.Unmarshal([]byte(message), &reqData)
	if err != nil {
		log.NewHelper(r.logger).Errorf("Failed to parse message: %v", err)
		return
	}

	// 根据 player_fingerprint 找到对应的 WebSocket 客户端
	clientID := generateClientID(reqData.PlayerFingerprint)
	client := manager.GetClientById(clientID)
	if client == nil {
		log.NewHelper(r.logger).Warnf("Client with ID %s not found", clientID)
		return
	}

	// 构建发送消息
	msg := &SendMsg{
		Code:     ginplus.StatusOK,
		Business: reqData.Type,
		Data:     reqData,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.NewHelper(r.logger).Errorf("Failed to marshal message: %v", err)
		return
	}

	// 发送消息到 WebSocket 客户端
	select {
	case client.Send <- msgBytes:
		log.NewHelper(r.logger).Infof("Message sent to client %s", clientID)
	default:
		log.NewHelper(r.logger).Warnf("Client %s send channel full, message dropped", clientID)
	}
}

// Close 关闭 Redis 客户端
func (r *RedisSubscriber) Close() {
	_ = r.client.Close()
}
