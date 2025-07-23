package ginplus

import (
	"github.com/pkg/errors"
)

var (
	New               = errors.New
	Wrap              = errors.Wrap
	Wrapf             = errors.Wrapf
	ErrForbidden      = New("禁止访问")
	ErrNotFound       = New("资源不存在")
	ErrBadRequest     = New("请求无效")
	ErrUnauthorized   = New("未授权")
	ErrInternalServer = New("服务器错误")
)

// NewBadRequestError 创建请求无效错误
func NewBadRequestError(msg ...string) error {
	return NewMessageError(ErrBadRequest, msg...)
}

// NewInternalServerError 创建服务器错误
func NewInternalServerError(msg ...string) error {
	return NewMessageError(ErrInternalServer, msg...)
}
