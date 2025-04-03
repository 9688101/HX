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
	userRepo := repo.NewUserRepository(database.GetDB())
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := v1.NewUserController(userUseCase)
	optionRepo := repo.NewOptionRepository(database.GetDB())
	optionUseCase := usecase.NewOptionUsecase(optionRepo)
	optionController := v1.NewOptionController(optionUseCase)
	miscController := v1.NewMiscController()

	apiRouter := router.Group("/api")
	apiRouter.Use(gzip.Gzip(gzip.DefaultCompression))
	apiRouter.Use(middleware.GlobalAPIRateLimit())
	{
		apiRouter.GET("/status", miscController.GetStatus)
		// apiRouter.GET("/models", middleware.UserAuth(), controller.DashboardListModels)
		// apiRouter.GET("/notice", controller.GetNotice)
		// apiRouter.GET("/about", controller.GetAbout)
		// apiRouter.GET("/home_page_content", controller.GetHomePageContent)
		// apiRouter.GET("/verification", middleware.CriticalRateLimit(), middleware.TurnstileCheck(), controller.SendEmailVerification)
		// apiRouter.GET("/reset_password", middleware.CriticalRateLimit(), middleware.TurnstileCheck(), controller.SendPasswordResetEmail)
		// apiRouter.POST("/user/reset", middleware.CriticalRateLimit(), controller.ResetPassword)
		// apiRouter.GET("/oauth/github", middleware.CriticalRateLimit(), auth.GitHubOAuth)
		// apiRouter.GET("/oauth/oidc", middleware.CriticalRateLimit(), auth.OidcAuth)
		// apiRouter.GET("/oauth/lark", middleware.CriticalRateLimit(), auth.LarkOAuth)
		// apiRouter.GET("/oauth/state", middleware.CriticalRateLimit(), auth.GenerateOAuthCode)
		// apiRouter.GET("/oauth/wechat", middleware.CriticalRateLimit(), auth.WeChatAuth)
		// apiRouter.GET("/oauth/wechat/bind", middleware.CriticalRateLimit(), middleware.UserAuth(), auth.WeChatBind)
		// apiRouter.GET("/oauth/email/bind", middleware.CriticalRateLimit(), middleware.UserAuth(), controller.EmailBind)

		userRoute := apiRouter.Group("/user")
		{
			userRoute.POST("/register", middleware.CriticalRateLimit(), middleware.TurnstileCheck(), userController.RegisterUserHandler)
			userRoute.POST("/login", middleware.CriticalRateLimit(), userController.LoginHandler)
			userRoute.GET("/logout", userController.LogoutHandler)

			selfRoute := userRoute.Group("/")
			selfRoute.Use(middleware.UserAuth())
			{
				// selfRoute.GET("/dashboard", controller.GetUserDashboard)
				selfRoute.GET("/self", userController.GetSelfHandler)
				selfRoute.PUT("/self", userController.UpdateSelfHandler)
				selfRoute.DELETE("/self", userController.DeleteSelfHandler)
				selfRoute.GET("/token", userController.GenerateAccessTokenHandler)
				selfRoute.GET("/aff", userController.GetAffCodeHandler)
				// selfRoute.POST("/topup", userController.TopUpHandler)
				// selfRoute.GET("/available_models", controller.GetUserAvailableModels)
			}

			adminRoute := userRoute.Group("/")
			adminRoute.Use(middleware.AdminAuth())
			{
				adminRoute.GET("/", userController.GetUserListHandler)
				adminRoute.GET("/search", userController.SearchUsersHandler)
				adminRoute.GET("/:id", userController.GetUserHandler)
				// adminRoute.POST("/", userController.CreateUserHandler)
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
		// channelRoute := apiRouter.Group("/channel")
		// channelRoute.Use(middleware.AdminAuth())
		// {
		// 	channelRoute.GET("/", controller.GetAllChannels)
		// 	channelRoute.GET("/search", controller.SearchChannels)
		// 	channelRoute.GET("/models", controller.ListAllModels)
		// 	channelRoute.GET("/:id", controller.GetChannel)
		// 	channelRoute.GET("/test", controller.TestChannels)
		// 	channelRoute.GET("/test/:id", controller.TestChannel)
		// 	channelRoute.GET("/update_balance", controller.UpdateAllChannelsBalance)
		// 	channelRoute.GET("/update_balance/:id", controller.UpdateChannelBalance)
		// 	channelRoute.POST("/", controller.AddChannel)
		// 	channelRoute.PUT("/", controller.UpdateChannel)
		// 	channelRoute.DELETE("/disabled", controller.DeleteDisabledChannel)
		// 	channelRoute.DELETE("/:id", controller.DeleteChannel)
		// }
		// tokenRoute := apiRouter.Group("/token")
		// tokenRoute.Use(middleware.UserAuth())
		// {
		// 	tokenRoute.GET("/", controller.GetAllTokens)
		// 	tokenRoute.GET("/search", controller.SearchTokens)
		// 	tokenRoute.GET("/:id", controller.GetToken)
		// 	tokenRoute.POST("/", controller.AddToken)
		// 	tokenRoute.PUT("/", controller.UpdateToken)
		// 	tokenRoute.DELETE("/:id", controller.DeleteToken)
		// }
		// redemptionRoute := apiRouter.Group("/redemption")
		// redemptionRoute.Use(middleware.AdminAuth())
		// {
		// 	redemptionRoute.GET("/", controller.GetAllRedemptions)
		// 	redemptionRoute.GET("/search", controller.SearchRedemptions)
		// 	redemptionRoute.GET("/:id", controller.GetRedemption)
		// 	redemptionRoute.POST("/", controller.AddRedemption)
		// 	redemptionRoute.PUT("/", controller.UpdateRedemption)
		// 	redemptionRoute.DELETE("/:id", controller.DeleteRedemption)
		// }
		// logRoute := apiRouter.Group("/log")
		// logRoute.GET("/", middleware.AdminAuth(), controller.GetAllLogs)
		// logRoute.DELETE("/", middleware.AdminAuth(), controller.DeleteHistoryLogs)
		// logRoute.GET("/stat", middleware.AdminAuth(), controller.GetLogsStat)
		// logRoute.GET("/self/stat", middleware.UserAuth(), controller.GetLogsSelfStat)
		// logRoute.GET("/search", middleware.AdminAuth(), controller.SearchAllLogs)
		// logRoute.GET("/self", middleware.UserAuth(), controller.GetUserLogs)
		// logRoute.GET("/self/search", middleware.UserAuth(), controller.SearchUserLogs)
		// groupRoute := apiRouter.Group("/group")
		// groupRoute.Use(middleware.AdminAuth())
		// {
		// 	groupRoute.GET("/", controller.GetGroups)
		// }
	}
}
