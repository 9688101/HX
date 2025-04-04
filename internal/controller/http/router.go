package http

import (
	v1 "github.com/9688101/HX/internal/controller/http/v1"
	"github.com/9688101/HX/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(r *gin.Engine) {
	// 公开路由
	r.POST("/api/login/account", v1.Login)
	r.POST("/api/login/outLogin", v1.Logout)

	// 需要认证的路由
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/currentUser", v1.GetCurrentUser)
		// 在这里添加其他需要认证的路由
	}
}
