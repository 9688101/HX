package http

import (
	"embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/internal/middleware"
	"github.com/9688101/HX/pkg/embfs"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func SetWebRouter(router *gin.Engine, buildFS embed.FS) {
	indexPageData, _ := buildFS.ReadFile(fmt.Sprintf("web/build/%s/index.html", config.GetServerConfig().Theme))
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	// router.Use(middleware.GlobalWebRateLimit())
	router.Use(middleware.Cache())
	router.Use(static.Serve("/", embfs.EmbedFolder(buildFS, fmt.Sprintf("web/build/%s", config.GetServerConfig().Theme))))
	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/v1") || strings.HasPrefix(c.Request.RequestURI, "/api") {
			// v1.RelayNotFound(c)
			return
		}
		c.Header("Cache-Control", "no-cache")
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexPageData)
	})
}
