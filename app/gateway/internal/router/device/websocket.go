package device

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"kratos_sample/app/gateway/internal/router/ginplus"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	EveryCliMaxWebSoketCount = 5
)

// Client 表示一个WebSocket连接
type Client struct {
	ID             string
	IP             string
	Conn           *websocket.Conn
	Send           chan []byte
	Done           chan struct{}
	IsClosed       bool
	Mu             sync.Mutex
	UserAgent      string
	AcceptLanguage string
	Platform       string
	Fingerprint    string
	LastActivity   time.Time
}

// RequestData 客户端请求数据结构
type RequestData struct {
	Business    string `json:"business"`
	Fingerprint string `json:"fingerprint"`
}

// ResponseData 服务端响应数据结构
type ResponseData struct {
	Code     uint64   `json:"code"`
	Data     BindInfo `json:"data"`
	ConCount int      `json:"con_count"`
	Business string   `json:"business"`
}

type ReqData struct { //请求排期
	TvID              string `json:"tv_id"`              // 电视 ID
	PlayerFingerprint string `json:"player_fingerprint"` // 播放器指纹
	Type              string `json:"type"`               // 类型
	Value             string `json:"value"`              // 值
}

type SendMsg struct {
	Code     int         `json:"code"`
	Business string      `json:"business"`
	Data     interface{} `json:"data"`
}

type BindInfo struct {
	Status           uint32 `json:"status"`
	Serial           string `json:"serial"`
	Name             string `json:"name"`
	DeviceID         string `json:"device_id"`
	ID               uint64 `json:"id"`
	GroupID          int64  `json:"group_id"`
	GroupName        string `json:"group_name"`
	DisplayRatioType int32  `json:"display_ratio_type"`
	ClearCache       int32  `json:"clear_cache"`
}

// ClientManager 管理所有客户端连接
type ClientManager struct {
	Clients       map[string]*Client
	Register      chan *Client
	Unregister    chan *Client
	Broadcast     chan []byte
	Mu            sync.RWMutex
	EveryConCount sync.Map //每个客户端建立的链接数
}

var manager = ClientManager{
	Clients:    make(map[string]*Client),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Broadcast:  make(chan []byte),
}

// 生成唯一客户端ID
func generateClientID(fingerprint string) string {
	//return fmt.Sprintf("client_%s_%s", fingerprint, time.Now().Format("20060102150405"))
	return fmt.Sprintf("client_%s", fingerprint)
}

// 根据用户ip+fingerprint获取链接数量
func (manager *ClientManager) getUserConCount(ip, fingerprint string) int {
	key := fmt.Sprintf("%s", fingerprint)
	if value, ok := manager.EveryConCount.Load(key); ok {
		return value.(int)
	} else {
		return 0
	}
}

func (manager *ClientManager) GetClientById(clientID string) *Client {
	defer manager.Mu.RUnlock()
	manager.Mu.RLock()
	client, ok := manager.Clients[clientID]
	if ok {
		return client
	}
	return nil
}

func (manager *ClientManager) addUserConsCount(ip, fingerprint string) {
	key := fmt.Sprintf("%s", fingerprint)
	if value, ok := manager.EveryConCount.Load(key); ok {
		manager.EveryConCount.Store(key, value.(int)+1)
	} else {
		manager.EveryConCount.Store(key, 1)
	}
}
func (manager *ClientManager) subUserConsCount(ip, fingerprint string) {
	key := fmt.Sprintf("%s", fingerprint)
	if value, ok := manager.EveryConCount.Load(key); ok {
		manager.EveryConCount.Store(key, value.(int)-1)
	} else {
		manager.EveryConCount.Store(key, 0)
	}
}

// Start 启动客户端管理器
func (manager *ClientManager) Start() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case client := <-manager.Register:
			manager.Mu.Lock()
			manager.Clients[client.ID] = client
			manager.Mu.Unlock()
			log.Infof("Client %s connected", client.ID)

		case client := <-manager.Unregister:
			manager.Mu.Lock()
			if _, ok := manager.Clients[client.ID]; ok {
				client.Close()
				delete(manager.Clients, client.ID)
				log.Infof("Client %s disconnected and cleaned up", client.ID)
			}
			manager.Mu.Unlock()

		case message := <-manager.Broadcast:
			manager.Mu.RLock()
			for _, client := range manager.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(manager.Clients, client.ID)
				}
			}
			manager.Mu.RUnlock()

		case <-ticker.C:
			// 定期检查不活跃连接
			manager.Mu.Lock()
			for id, client := range manager.Clients {
				if time.Since(client.LastActivity) > 2*time.Minute {
					log.Infof("Client %s inactive for too long, disconnecting", id)
					client.Close()
					delete(manager.Clients, id)
				}
			}
			manager.Mu.Unlock()
		}
	}
}

// Close 安全关闭客户端连接
func (c *Client) Close() {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	manager.subUserConsCount(c.IP, c.Fingerprint)

	if !c.IsClosed {
		close(c.Done)
		close(c.Send)
		c.Conn.Close()
		c.IsClosed = true
	}
}

// calculateHash 计算请求信息的哈希值
func (c *Client) calculateHash() string {
	combined := strings.Join([]string{c.UserAgent, c.AcceptLanguage, c.Platform, c.Fingerprint}, "|")
	hash := sha256.Sum256([]byte(combined))
	return hex.EncodeToString(hash[:])
}

// ReadPump 从客户端读取消息
func (c *Client) ReadPump() {
	defer func() {
		manager.Unregister <- c
	}()

	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		c.LastActivity = time.Now()
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Log(log.LevelInfo, "Client %s read error: %v", c.ID, err)
			}
			break
		}

		c.LastActivity = time.Now()

		// 处理接收到的消息
		var reqData RequestData
		if err := json.Unmarshal(message, &reqData); err != nil {
			log.Info("Client %s JSON unmarshal error: %v", c.ID, err)
			continue
		}

		msg := &SendMsg{
			Code:     ginplus.StatusOK,
			Business: "heartbeat ",
			Data: &ReqData{
				TvID:              "",
				PlayerFingerprint: reqData.Fingerprint, // 播放器指纹
				Type:              "heartbeat ",        // 类型
				Value:             "",                  // 值
			},
		}
		respBytes, _ := json.Marshal(msg)
		if respBytes != nil {
			c.Send <- respBytes
		}
	}
}

type SerialInfo struct {
	UserAgent      string `json:"user_agent"`
	AcceptLanguage string `json:"accept_language"`
	Platform       string `json:"platform"`
	Client         string `json:"client"`
	Code           string `json:"code"`
	Fingerprint    string `json:"fingerprint"`
}

// WritePump 向客户端发送消息
func (c *Client) WritePump() {
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
		manager.Unregister <- c
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(40 * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(40 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		case <-c.Done:
			return
		}
	}
}

// handleWebSocket 处理WebSocket连接
func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Infof("WebSocket upgrade error: %v", err)
		ginplus.ResponseError(c, err)
		return
	}
	// 获取请求头信息
	userAgent := c.Request.Header.Get("User-Agent")
	acceptLanguage := c.Request.Header.Get("accept-language")
	platform := c.Request.Header.Get("sec-ch-ua-platform")
	fingerprint := c.Request.Header.Get("fingerprint")

	clientIP := c.ClientIP()
	count := manager.getUserConCount(clientIP, fingerprint)
	if count >= EveryCliMaxWebSoketCount { //websoket连接数达到最大数量返回错误
		ginplus.ResponseError(c, errors.New("websocket too large"))
		return
	}

	// 创建客户端
	client := &Client{
		ID:             generateClientID(fingerprint),
		IP:             clientIP,
		Conn:           conn,
		Send:           make(chan []byte, 256),
		Done:           make(chan struct{}),
		UserAgent:      userAgent,
		AcceptLanguage: acceptLanguage,
		Platform:       platform,
		Fingerprint:    fingerprint,
		LastActivity:   time.Now(),
	}

	// 注册客户端
	manager.Register <- client
	manager.addUserConsCount(clientIP, fingerprint)

	// 启动读写协程
	go client.WritePump()
	go client.ReadPump()

	////停顿2秒主动发送一条消息给客户端
	//time.Sleep(2 * time.Second)
	//playerFingerprint := client.calculateHash()
	//client.GenerateClientMsg(playerFingerprint)
}

// sendMessageToClient 从其他handler发送消息到指定客户端
func sendMessageToClient(c *gin.Context) {
	clientID := c.Query("client_id")
	message := c.Query("message")

	manager.Mu.RLock()
	client, ok := manager.Clients[clientID]
	manager.Mu.RUnlock()

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
		return
	}

	// 构建消息
	msg := map[string]interface{}{
		"timestamp": time.Now().Unix(),
		"message":   message,
		"client_info": map[string]string{
			"user_agent":      client.UserAgent,
			"accept_language": client.AcceptLanguage,
			"platform":        client.Platform,
			"fingerprint":     client.Fingerprint,
		},
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal message"})
		return
	}

	// 发送消息
	select {
	case client.Send <- msgBytes:
		c.JSON(http.StatusOK, gin.H{
			"status":    "message sent",
			"client_id": clientID,
			"hash":      client.calculateHash(),
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "client send channel full"})
	}
}

// broadcastMessage 广播消息到所有客户端
func broadcastMessage(c *gin.Context) {
	message := c.Query("message")

	manager.Mu.RLock()
	clientsCount := len(manager.Clients)
	manager.Mu.RUnlock()

	if clientsCount == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "no clients connected"})
		return
	}

	msg := map[string]interface{}{
		"timestamp": time.Now().Unix(),
		"message":   message,
		"broadcast": true,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal message"})
		return
	}

	manager.Broadcast <- msgBytes

	c.JSON(http.StatusOK, gin.H{
		"status":  "broadcast sent",
		"clients": clientsCount,
	})
}

// listClients 列出所有连接的客户端
func listClients(c *gin.Context) {
	manager.Mu.RLock()
	defer manager.Mu.RUnlock()

	clientsInfo := make([]map[string]interface{}, 0, len(manager.Clients))
	for _, client := range manager.Clients {
		clientsInfo = append(clientsInfo, map[string]interface{}{
			"id":              client.ID,
			"user_agent":      client.UserAgent,
			"accept_language": client.AcceptLanguage,
			"platform":        client.Platform,
			"fingerprint":     client.Fingerprint,
			"last_activity":   client.LastActivity.Format(time.RFC3339),
			"hash":            client.calculateHash(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"clients":       clientsInfo,
		"total_clients": len(clientsInfo),
	})
}

//func main222() {
//	// 启动客户端管理器
//	go manager.Start()
//
//	r := gin.Default()
//
//	// WebSocket路由
//	r.GET("/ws", handleWebSocket)
//
//	// 其他handler示例
//	r.GET("/send", sendMessageToClient)
//	r.GET("/broadcast", broadcastMessage)
//	r.GET("/clients", listClients)
//
//	// 启动服务器
//	if err := r.Run(":8080"); err != nil {
//		log.Info("Failed to start server: %v", err)
//	}
//}

func (a *DeviceHandler) WebSocketSend(c *gin.Context) {
	var params ReqData
	err := ginplus.ParseJSON(c, &params)
	if err != nil {
		ginplus.ResponseError(c, err)
		return
	}

	clientID := generateClientID(params.PlayerFingerprint)

	client := manager.GetClientById(clientID)
	if client != nil {
		msg := &SendMsg{
			Code:     ginplus.StatusOK,
			Business: params.Type,
			Data:     params,
		}
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal message"})
			return
		}
		// 发送消息
		select {
		case client.Send <- msgBytes:
			c.JSON(http.StatusOK, gin.H{
				"code":      ginplus.StatusOK,
				"status":    "message sent",
				"client_id": clientID,
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "client empty"})
			return
		}
	}
	ginplus.ResponseSuccess(c, nil)
}
