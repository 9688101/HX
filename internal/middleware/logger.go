package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/9688101/HX/pkg/helper"
	"github.com/9688101/HX/pkg/logger"
)

// GinLoggerMiddleware 是一个 Gin 中间件，记录请求日志并调用全局 Logger 输出日志信息
func GinLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		// 继续处理请求
		c.Next()

		latency := time.Since(startTime)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		// 尝试从 gin.Context 中获取请求 ID，通常在其他中间件中设置
		var requestID string
		if id, exists := c.Get(helper.RequestIdKey); exists {
			// 这里假设请求 ID 为字符串类型
			requestID = fmt.Sprintf("%v", id)
		}

		// 使用全局 Logger 记录请求日志
		logger.Info(c.Request.Context(), "HTTP request processed",
			zap.String("timestamp", startTime.Format("2006/01/02 - 15:04:05")),
			zap.Int("status", statusCode),
			zap.Duration("latency", latency),
			zap.String("clientIP", clientIP),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("requestID", requestID),
		)
	}
}
