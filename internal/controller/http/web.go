package http

import (
	"embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/9688101/HX/config"
	v1 "github.com/9688101/HX/internal/controller/http/v1"
	"github.com/9688101/HX/internal/middleware"
	"github.com/9688101/HX/pkg"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func SetWebRouter(router *gin.Engine, buildFS embed.FS) {
	cfg := config.GetSystemConfig()
	indexPageData, _ := buildFS.ReadFile(fmt.Sprintf("web/build/%s/index.html", cfg.Theme))
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	// router.Use(middleware.GlobalWebRateLimit())
	router.Use(middleware.Cache())
	router.Use(static.Serve("/", pkg.EmbedFolder(buildFS, fmt.Sprintf("web/build/%s", cfg.Theme))))
	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/v1") || strings.HasPrefix(c.Request.RequestURI, "/api") {
			v1.RelayNotFound(c)
			return
		}
		c.Header("Cache-Control", "no-cache")
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexPageData)
	})
}
