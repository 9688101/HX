package http

import (
	v1 "github.com/9688101/HX/internal/controller/http/v1"
	"github.com/9688101/HX/internal/repo"
	"github.com/9688101/HX/internal/usecase"
	"github.com/9688101/HX/pkg/db"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetupUserRouter() *gin.Engine {
	r := gin.Default()

	userRepo := repo.NewUserRepository(db.GetDB())
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := v1.NewUserController(userUseCase)

	apiRouter := r.Group("/user")
	apiRouter.Use(gzip.Gzip(gzip.DefaultCompression))
	// apiRouter.Use(middleware.GlobalAPIRateLimit())
	{
		// 注册用户接口
		r.POST("/register", userController.RegisterUserHandler)
		r.POST("/login", userController.LoginHandler)
		r.POST("/logout", userController.LogoutHandler)
		r.GET("/users", userController.GetUserListHandler)
		r.GET("/users/search", userController.SearchUsersHandler)
		r.GET("/user/:id", userController.GetUserHandler)
		r.PUT("/user/self", userController.UpdateSelfHandler)
		r.GET("/user/self", userController.GetSelfHandler)
		r.DELETE("/user/self", userController.DeleteSelfHandler)
		r.PUT("/user/:id", userController.UpdateUserHandler)
		r.DELETE("/user/:id", userController.DeleteUserHandler)
		r.POST("/user/bind-email", userController.EmailBindHandler)
		r.POST("/user/manage", userController.ManageUserHandler)
		r.GET("/user/aff-code", userController.GetAffCodeHandler)
		r.POST("/user/generate-access-token", userController.GenerateAccessTokenHandler)

		return r
	}
}
