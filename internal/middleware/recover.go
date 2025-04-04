package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/9688101/HX/pkg/ginutil"
	"github.com/9688101/HX/pkg/logger"
	"github.com/gin-gonic/gin"
)

// RelayPanicRecover 中间件：捕获 panic，记录详细日志并返回 500 错误
func RelayPanicRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx := c.Request.Context()
				// 记录 panic 信息与堆栈
				logger.Error(ctx, fmt.Sprintf("panic detected: %v", err),
					logger.ToZapField("stacktrace", string(debug.Stack())))
				// 记录请求关键信息
				logger.Error(ctx, fmt.Sprintf("request: %s %s", c.Request.Method, c.Request.URL.Path))
				if body, e := ginutil.GetRequestBody(c); e == nil {
					logger.Error(ctx, fmt.Sprintf("request body: %s", string(body)))
				}
				// 返回错误响应
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": gin.H{
						"message": fmt.Sprintf("Panic detected: %v. Please submit an issue with the related log at: https://github.com/songquanpeng/one-api", err),
						"type":    "one_api_panic",
					},
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
