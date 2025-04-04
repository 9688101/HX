package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/9688101/HX/pkg/i18n"
)

// Language 中间件：根据 Accept-Language 设置语言环境，默认 en 或 zh-CN
func Language() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		lang = strings.TrimSpace(strings.ToLower(lang))
		if lang == "" {
			lang = "en"
		} else if strings.HasPrefix(lang, "zh") {
			lang = "zh-CN"
		} else {
			lang = "en"
		}
		c.Set(i18n.ContextKey, lang)
		c.Next()
	}
}
