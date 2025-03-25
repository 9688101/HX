package http

import (
	"github.com/9688101/HX/internal/controller/http/middleware"
	v1 "github.com/9688101/HX/internal/controller/http/v1"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetDashboardRouter(router *gin.Engine) {
	apiRouter := router.Group("/")
	apiRouter.Use(middleware.CORS())
	apiRouter.Use(gzip.Gzip(gzip.DefaultCompression))
	apiRouter.Use(middleware.GlobalAPIRateLimit())
	apiRouter.Use(middleware.TokenAuth())
	{
		apiRouter.GET("/dashboard/billing/subscription", v1.GetSubscription)
		apiRouter.GET("/v1/dashboard/billing/subscription", v1.GetSubscription)
		apiRouter.GET("/dashboard/billing/usage", v1.GetUsage)
		apiRouter.GET("/v1/dashboard/billing/usage", v1.GetUsage)
	}
}
