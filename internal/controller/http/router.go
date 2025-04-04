package http

import (
	v1 "github.com/9688101/HX/internal/controller/http/v1"
	"github.com/9688101/HX/internal/middleware"
	"github.com/9688101/HX/internal/repo"
	"github.com/9688101/HX/internal/usecase"
	"github.com/9688101/HX/pkg/database"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	// 初始化仓储、用例、控制器
	db := database.GetDB()
	userRepo := repo.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := v1.NewUserController(userUseCase)

	optionRepo := repo.NewOptionRepository(db)
	optionUseCase := usecase.NewOptionUsecase(optionRepo)
	optionController := v1.NewOptionController(optionUseCase)

	miscController := v1.NewMiscController()

	apiRouter := router.Group("/api")
	apiRouter.Use(gzip.Gzip(gzip.DefaultCompression))
	apiRouter.Use(middleware.GlobalAPIRateLimit())

	{
		apiRouter.GET("/status", miscController.GetStatus)

		userRoute := apiRouter.Group("/user")
		{
			userRoute.POST("/register", middleware.CriticalRateLimit(), middleware.TurnstileCheck(), userController.RegisterUserHandler)
			userRoute.POST("/login", middleware.CriticalRateLimit(), userController.LoginHandler)
			userRoute.GET("/logout", middleware.UserAuth(), userController.LogoutHandler)

			// 用户自身操作
			selfRoute := userRoute.Group("/self")
			selfRoute.Use(middleware.UserAuth())
			{
				selfRoute.GET("/", userController.GetSelfHandler)
				selfRoute.PUT("/", userController.UpdateSelfHandler)
				selfRoute.DELETE("/", userController.DeleteSelfHandler)
				selfRoute.GET("/token", userController.GenerateAccessTokenHandler)
				selfRoute.GET("/aff", userController.GetAffCodeHandler)
			}

			// 管理员操作
			adminRoute := userRoute.Group("/admin")
			adminRoute.Use(middleware.AdminAuth())
			{
				adminRoute.GET("/", userController.GetUserListHandler)
				adminRoute.GET("/search", userController.SearchUsersHandler)
				adminRoute.GET("/:id", userController.GetUserHandler)
				adminRoute.POST("/manage", userController.ManageUserHandler)
				adminRoute.PUT("/", userController.UpdateUserHandler)
				adminRoute.DELETE("/:id", userController.DeleteUserHandler)
			}
		}

		optionRoute := apiRouter.Group("/option")
		optionRoute.Use(middleware.RootAuth())
		{
			optionRoute.GET("/", optionController.GetOptions)
			optionRoute.PUT("/", optionController.UpdateOption)
		}
	}
}
