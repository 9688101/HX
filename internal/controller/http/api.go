package http

import (
	"github.com/9688101/HX/internal/controller/http/V1/auth"
	"github.com/9688101/HX/internal/controller/http/middleware"
	v1 "github.com/9688101/HX/internal/controller/http/v1"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	apiRouter := router.Group("/api")
	apiRouter.Use(gzip.Gzip(gzip.DefaultCompression))
	apiRouter.Use(middleware.GlobalAPIRateLimit())
	{
		apiRouter.GET("/status", v1.GetStatus)
		apiRouter.GET("/models", middleware.UserAuth(), v1.DashboardListModels)
		apiRouter.GET("/notice", v1.GetNotice)
		apiRouter.GET("/about", v1.GetAbout)
		apiRouter.GET("/home_page_content", v1.GetHomePageContent)
		apiRouter.GET("/verification", middleware.CriticalRateLimit(), middleware.TurnstileCheck(), v1.SendEmailVerification)
		apiRouter.GET("/reset_password", middleware.CriticalRateLimit(), middleware.TurnstileCheck(), v1.SendPasswordResetEmail)
		apiRouter.POST("/user/reset", middleware.CriticalRateLimit(), v1.ResetPassword)
		apiRouter.GET("/oauth/github", middleware.CriticalRateLimit(), auth.GitHubOAuth)
		apiRouter.GET("/oauth/oidc", middleware.CriticalRateLimit(), auth.OidcAuth)
		apiRouter.GET("/oauth/lark", middleware.CriticalRateLimit(), auth.LarkOAuth)
		apiRouter.GET("/oauth/state", middleware.CriticalRateLimit(), auth.GenerateOAuthCode)
		apiRouter.GET("/oauth/wechat", middleware.CriticalRateLimit(), auth.WeChatAuth)
		apiRouter.GET("/oauth/wechat/bind", middleware.CriticalRateLimit(), middleware.UserAuth(), auth.WeChatBind)
		apiRouter.GET("/oauth/email/bind", middleware.CriticalRateLimit(), middleware.UserAuth(), v1.EmailBind)
		apiRouter.POST("/topup", middleware.AdminAuth(), v1.AdminTopUp)

		userRoute := apiRouter.Group("/user")
		{
			userRoute.POST("/register", middleware.CriticalRateLimit(), middleware.TurnstileCheck(), v1.Register)
			userRoute.POST("/login", middleware.CriticalRateLimit(), v1.Login)
			userRoute.GET("/logout", v1.Logout)

			selfRoute := userRoute.Group("/")
			selfRoute.Use(middleware.UserAuth())
			{
				selfRoute.GET("/dashboard", v1.GetUserDashboard)
				selfRoute.GET("/self", v1.GetSelf)
				selfRoute.PUT("/self", v1.UpdateSelf)
				selfRoute.DELETE("/self", v1.DeleteSelf)
				selfRoute.GET("/token", v1.GenerateAccessToken)
				selfRoute.GET("/aff", v1.GetAffCode)
				selfRoute.POST("/topup", v1.TopUp)
				selfRoute.GET("/available_models", v1.GetUserAvailableModels)
			}

			adminRoute := userRoute.Group("/")
			adminRoute.Use(middleware.AdminAuth())
			{
				adminRoute.GET("/", v1.GetAllUsers)
				adminRoute.GET("/search", v1.SearchUsers)
				adminRoute.GET("/:id", v1.GetUser)
				adminRoute.POST("/", v1.CreateUser)
				adminRoute.POST("/manage", v1.ManageUser)
				adminRoute.PUT("/", v1.UpdateUser)
				adminRoute.DELETE("/:id", v1.DeleteUser)
			}
		}
		optionRoute := apiRouter.Group("/option")
		optionRoute.Use(middleware.RootAuth())
		{
			optionRoute.GET("/", v1.GetOptions)
			optionRoute.PUT("/", v1.UpdateOption)
		}
		channelRoute := apiRouter.Group("/channel")
		channelRoute.Use(middleware.AdminAuth())
		{
			channelRoute.GET("/", v1.GetAllChannels)
			channelRoute.GET("/search", v1.SearchChannels)
			channelRoute.GET("/models", v1.ListAllModels)
			channelRoute.GET("/:id", v1.GetChannel)
			channelRoute.GET("/test", v1.TestChannels)
			channelRoute.GET("/test/:id", v1.TestChannel)
			channelRoute.GET("/update_balance", v1.UpdateAllChannelsBalance)
			channelRoute.GET("/update_balance/:id", v1.UpdateChannelBalance)
			channelRoute.POST("/", v1.AddChannel)
			channelRoute.PUT("/", v1.UpdateChannel)
			channelRoute.DELETE("/disabled", v1.DeleteDisabledChannel)
			channelRoute.DELETE("/:id", v1.DeleteChannel)
		}
		tokenRoute := apiRouter.Group("/token")
		tokenRoute.Use(middleware.UserAuth())
		{
			tokenRoute.GET("/", v1.GetAllTokens)
			tokenRoute.GET("/search", v1.SearchTokens)
			tokenRoute.GET("/:id", v1.GetToken)
			tokenRoute.POST("/", v1.AddToken)
			tokenRoute.PUT("/", v1.UpdateToken)
			tokenRoute.DELETE("/:id", v1.DeleteToken)
		}
		redemptionRoute := apiRouter.Group("/redemption")
		redemptionRoute.Use(middleware.AdminAuth())
		{
			redemptionRoute.GET("/", v1.GetAllRedemptions)
			redemptionRoute.GET("/search", v1.SearchRedemptions)
			redemptionRoute.GET("/:id", v1.GetRedemption)
			redemptionRoute.POST("/", v1.AddRedemption)
			redemptionRoute.PUT("/", v1.UpdateRedemption)
			redemptionRoute.DELETE("/:id", v1.DeleteRedemption)
		}
		logRoute := apiRouter.Group("/log")
		logRoute.GET("/", middleware.AdminAuth(), v1.GetAllLogs)
		logRoute.DELETE("/", middleware.AdminAuth(), v1.DeleteHistoryLogs)
		logRoute.GET("/stat", middleware.AdminAuth(), v1.GetLogsStat)
		logRoute.GET("/self/stat", middleware.UserAuth(), v1.GetLogsSelfStat)
		logRoute.GET("/search", middleware.AdminAuth(), v1.SearchAllLogs)
		logRoute.GET("/self", middleware.UserAuth(), v1.GetUserLogs)
		logRoute.GET("/self/search", middleware.UserAuth(), v1.SearchUserLogs)
		groupRoute := apiRouter.Group("/group")
		groupRoute.Use(middleware.AdminAuth())
		{
			groupRoute.GET("/", v1.GetGroups)
		}
	}
}
