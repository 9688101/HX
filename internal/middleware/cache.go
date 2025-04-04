package middleware

import (
	"github.com/gin-gonic/gin"
)

// Cache 中间件：根据请求 URI 设置 Cache-Control 头
func Cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		if uri == "" || uri == "/" {
			c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		} else {
			// 设置为一周缓存，同时注意可加入 ETag 或 Last-Modified 等
			c.Header("Cache-Control", "public, max-age=604800")
		}
		c.Next()
	}
}
