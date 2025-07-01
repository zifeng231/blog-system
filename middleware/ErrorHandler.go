package middleware

import (
	"blog-system/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("panic recovered:", zap.Any("error", err))
				context.JSON(500, gin.H{"error": "服务器内部错误"})
				context.Abort()
			}
		}()
		context.Next()
		for _, err := range context.Errors {
			switch e := err.Err.(type) {
			case *utils.AppError:
				context.AbortWithStatusJSON(e.Code, gin.H{
					"code":    e.Code,
					"message": e.Message,
				})
			default:
				context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "未知错误",
				})
			}
		}
	}
}
