package http

import (
	"embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/pkg/logger"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine, buildFS embed.FS) {
	SetApiRouter(r)
	frontendBaseUrl := config.GetServerConfig().FrontendBaseUrl
	if config.GetDatabaseConfig().IsMasterNode && frontendBaseUrl != "" {
		frontendBaseUrl = ""
		logger.SysLog("FRONTEND_BASE_URL is ignored on master node")
	}
	if frontendBaseUrl == "" {
		SetWebRouter(r, buildFS)
	} else {
		frontendBaseUrl = strings.TrimSuffix(frontendBaseUrl, "/")
		r.NoRoute(func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%s%s", frontendBaseUrl, c.Request.RequestURI))
		})
	}
}
