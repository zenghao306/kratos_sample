package ginplus

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/shopspring/decimal"
	"golang.org/x/text/language"
	"net/http"
)

const (
	prefix = "kratos_sample"
	// UserIDKey 存储上下文中的键(用户ID)
	UserIDKey = prefix + "/user_id"
	// TraceIDKey 存储上下文中的键(跟踪ID)
	TraceIDKey = prefix + "/trace_id"
	// ResBodyKey 存储上下文中的键(响应Body数据)
	ResBodyKey       = prefix + "/res_body"
	AllowedResultKey = prefix + "/allowed_result"

	StatusOK = 22000 // RFC 9110, 15.3.1
)

var bundle *i18n.Bundle

type FormatError struct {
	Code   int64  `json:"code"`
	Detail string `json:"detail"`
}

// API返回格式
type response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 错误返回
func ResponseError(c *gin.Context, err error) {
	formatError := FormatError{}
	_ = jsoniter.UnmarshalFromString(err.Error(), &formatError)
	if formatError.Detail == "" {
		JSON(c, http.StatusOK, &response{
			Code:    400,
			Message: err.Error(),
		})
	} else {
		JSON(c, http.StatusOK, &response{
			Code:    400,
			Message: formatError.Detail,
		})
	}
}

type HTTPList struct {
	List        interface{} `json:"list"`
	Total       int         `json:"total"`
	CurrentPage int         `json:"current_page"`
	PageSize    int         `json:"page_size"`
	LastPage    int64       `json:"last_page"`
}
type PaginationResult struct {
	Total    uint32 `json:"total"`
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
}

func ResponsePagePost(c *gin.Context, v interface{}, pr *PaginationResult) {
	list := HTTPList{
		List: v,
	}

	if pr != nil {
		list.Total = int(pr.Total)
		list.CurrentPage = int(pr.Page)
		if pr.PageSize > 0 {
			list.PageSize = int(pr.PageSize)
			list.LastPage = decimal.NewFromInt32(int32(pr.Total)).Div(decimal.NewFromInt32(int32(pr.PageSize))).IntPart()
			if pr.Total%(pr.PageSize) > 0 {
				list.LastPage += 1
			}
		}
	}

	JSON(c, http.StatusOK, &response{
		Code:    StatusOK,
		Message: "ok",
		Data:    list,
	})
}

// 返回结果
func ResJsonForGeneral(c *gin.Context, code int, msg string, data interface{}) {
	if msg == "" {
		msg = "success"
	}
	//localizer := getLocalizer(c)
	//message, _ := localizer.Localize(&i18n.LocalizeConfig{
	//	MessageID: msg,
	//})

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
	})
}

// JSONMarshalToString JSON编码为字节
func JSONMarshalToByte(v interface{}) ([]byte, error) {
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)

	s, err := jsoniter.Marshal(v)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// 成功返回
func ResponseSuccess(c *gin.Context, obj interface{}) {
	ResponseJSON(c, StatusOK, nil, obj)
}

func ResponseJSON(c *gin.Context, status int, err error, obj interface{}) {
	var message = "操作成功"
	if err != nil {
		message = err.Error()
	}

	if obj == nil {
		obj = struct{}{}
	}

	JSON(c, http.StatusOK, &response{
		Code:    status,
		Message: message,
		Data:    obj,
	})
}

func JSON(c *gin.Context, status int, obj interface{}) {
	buf, err := json.Marshal(obj)
	if err != nil {
		buf, _ = json.Marshal(&response{
			Code:    status,
			Message: "操作失败",
			Data:    struct{}{},
		})
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}

func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		log.Infof("[ParseJSON] 解析请求JSON，无效的请求参数，err：%s", err)
		return NewBadRequestError("无效的请求参数")
	}
	return nil
}

func ResJson(c *gin.Context, code int, msg string, data interface{}) {
	if msg == "" {
		msg = "操作成功"
	}
	res := map[string]interface{}{"code": code, "msg": msg, "data": data}

	buf, err := JSONMarshalToByte(res)
	if err != nil {
		log.Infof("[ginplus.ResJSON] 响应JSON数据时发生错误，err：%s", err)
		c.JSON(500, gin.H{"error": NewInternalServerError})
		return
	}
	c.Set(ResBodyKey, buf)
	c.Data(200, "application/json; charset=utf-8", buf)
	c.Abort()
}

func NewMessageError(parent error, msg ...string) error {
	if parent == nil {
		return nil
	}

	m := parent.Error()
	if len(msg) > 0 {
		m = msg[0]
	}
	return &MessageError{parent, m}
}

type MessageError struct {
	err error
	msg string
}

func (m *MessageError) Error() string {
	return m.msg
}

// Parent 父级错误
func (m *MessageError) Parent() error {
	return m.err
}

func getLocalizer(c *gin.Context) *i18n.Localizer {
	lang, exists := c.Get("lang")
	if !exists {
		lang = "en"
	}

	supportedTags := []language.Tag{
		language.English,
		language.SimplifiedChinese,  // 简体中文
		language.TraditionalChinese, //繁体中文
		language.French,             //法语
		language.Arabic,             //阿拉伯语
		language.German,             //德语
		language.Japanese,           //日语
	}

	matcher := language.NewMatcher(supportedTags)
	userLang, _, _ := language.ParseAcceptLanguage(lang.(string))
	tag, _, _ := matcher.Match(userLang...)

	return i18n.NewLocalizer(bundle, tag.String())
}
