package http

import (
	v1 "github.com/9688101/HX/internal/controller/http/v1"
	"github.com/9688101/HX/internal/repo"
	"github.com/9688101/HX/internal/usecase"
	"github.com/9688101/HX/pkg/db"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetupUserRouter(r *gin.Engine) *gin.Engine {

	userRepo := repo.NewUserRepository(db.GetDB())
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := v1.NewUserController(userUseCase)

	apiRouter := r.Group("/api")
	apiRouter.Use(gzip.Gzip(gzip.DefaultCompression))
	// apiRouter.Use(middleware.GlobalAPIRateLimit())
	{
		userRouter := apiRouter.Group("/user")
		{

			// 注册用户接口
			userRouter.POST("/register", userController.RegisterUserHandler)
			userRouter.POST("/login", userController.LoginHandler)
			userRouter.POST("/logout", userController.LogoutHandler)
			userRouter.GET("/users", userController.GetUserListHandler)
			userRouter.GET("/users/search", userController.SearchUsersHandler)
			userRouter.GET("/user/:id", userController.GetUserHandler)
			userRouter.PUT("/user/self", userController.UpdateSelfHandler)
			userRouter.GET("/user/self", userController.GetSelfHandler)
			userRouter.DELETE("/user/self", userController.DeleteSelfHandler)
			userRouter.PUT("/user/:id", userController.UpdateUserHandler)
			userRouter.DELETE("/user/:id", userController.DeleteUserHandler)
			userRouter.POST("/user/bind-email", userController.EmailBindHandler)
			userRouter.POST("/user/manage", userController.ManageUserHandler)
			userRouter.GET("/user/aff", userController.GetAffCodeHandler)
			userRouter.POST("/user/generate-access-token", userController.GenerateAccessTokenHandler)
		}
		return r
	}
}
