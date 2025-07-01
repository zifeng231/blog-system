package utils

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel) // 设置日志级别为 Debug
	config.Encoding = "console"                         // 开发时建议用 console 更易读
	return config.Build()
}
