package ginplus

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	common_proto "kratos_sample/api/common/v1"
	"net/http"
	"strings"
)

// SetUserID 设定用户ID
func SetUserID(c *gin.Context, userID common_proto.TokenClaims) {
	c.Set(UserIDKey, userID)
}

func GetUserID(c *gin.Context) *common_proto.TokenClaims {
	value, e := c.Get(UserIDKey)
	if e == false {
		return nil
	}
	p, ok := (value).(common_proto.TokenClaims)
	if ok == false {
		return nil
	}
	return &p
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		//if method == "OPTIONS" {
		//    c.JSON(http.StatusOK, "Options Request!")
		//}

		// 特别处理WebSocket升级请求
		if c.Request.Header.Get("Upgrade") == "websocket" {
			c.Writer.Header().Set("Connection", "Upgrade")
			c.Writer.Header().Set("Upgrade", "websocket")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		// 处理请求
		c.Next() //  处理请求
	}
}

type SkipperFunc func(*gin.Context) bool

// 检验token
func CheckSessionToken(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		if status, err := checkSessionToken(c); err != nil {
			ResJson(c, status, err.Error(), "")
			c.Abort()
			return
		}
	}
}

func checkSessionToken(c *gin.Context) (int, error) {
	sessionToken := c.GetHeader("authorization")
	if sessionToken == "" {
		return http.StatusUnauthorized, errors.New("无效的Session Token")
	}
	sessionToken = strings.Replace(sessionToken, "bearer ", "", 1)

	// 分割JWT的三个部分
	parts := strings.Split(sessionToken, ".")
	if len(parts) != 3 {
		log.Errorf("[gatewany checkSessionToken] 解析错误的token:%v", sessionToken)
		return http.StatusUnauthorized, errors.New("无效的Session Token")
	}

	// 解码payload部分 (第二部分)
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		log.Errorf("[gatewany checkSessionToken] Failed to decode payload:%v", err)
		return http.StatusUnauthorized, errors.New("无效的Session Token")
	}

	// 解析JSON
	var claims common_proto.TokenClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		log.Errorf("Failed to unmarshal claims: %v", err)
		return http.StatusUnauthorized, errors.New("Token 解析错误")
	}

	////异步刷新token过期时间
	//tokenCh <- UserToken{
	//	UserId: claims.User.UserId,
	//	Token:  sessionToken,
	//}

	SetUserID(c, claims)

	return http.StatusOK, nil
}
