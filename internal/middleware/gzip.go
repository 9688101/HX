package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/9688101/HX/pkg/logger"
	"github.com/gin-gonic/gin"
)

// GzipDecodeMiddleware 解码请求体中 gzip 压缩数据，失败时返回 400
func GzipDecodeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.EqualFold(c.GetHeader("Content-Encoding"), "gzip") {
			gzipReader, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				logger.Error(c.Request.Context(), "failed to create gzip reader", logger.ToZapField("error", err))
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			// 替换请求 Body 为解压后的数据；注意：此处不再 defer 关闭 gzipReader，
			// 因为后续处理时会在读取完毕后自动关闭底层流（可参考 io.NopCloser 实现）
			c.Request.Body = io.NopCloser(gzipReader)
		}
		c.Next()
	}
}
