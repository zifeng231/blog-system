package utils

import "net/http"

var (
	ErrUnauthorized = NewError(http.StatusUnauthorized, "未授权")
	ErrForbidden    = NewError(http.StatusForbidden, "无权限")
	ErrNotFound     = NewError(http.StatusNotFound, "资源不存在")
	ErrBadRequest   = NewError(http.StatusBadRequest, "请求参数错误")
	ErrInternal     = NewError(http.StatusInternalServerError, "系统内部错误")
)

func NewError(code int, message string) error {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}
