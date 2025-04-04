package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/9688101/HX/pkg/helper"
)

// RequestId 中间件：生成请求 ID，存入 Context 与响应头，方便日志追踪
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := helper.GenRequestID()
		c.Set(helper.RequestIdKey, id)
		ctx := helper.SetRequestID(c.Request.Context(), id)
		c.Request = c.Request.WithContext(ctx)
		c.Header(helper.RequestIdKey, id)
		c.Next()
	}
}
