package http

import (
	"github.com/9688101/HX/internal/controller/http/middleware"
	v1 "github.com/9688101/HX/internal/controller/http/v1"

	"github.com/gin-gonic/gin"
)

func SetRelayRouter(router *gin.Engine) {
	router.Use(middleware.CORS())
	router.Use(middleware.GzipDecodeMiddleware())
	// https://platform.openai.com/docs/api-reference/introduction
	modelsRouter := router.Group("/v1/models")
	modelsRouter.Use(middleware.TokenAuth())
	{
		modelsRouter.GET("", v1.ListModels)
		modelsRouter.GET("/:model", v1.RetrieveModel)
	}
	relayV1Router := router.Group("/v1")
	relayV1Router.Use(middleware.RelayPanicRecover(), middleware.TokenAuth(), middleware.Distribute())
	{
		relayV1Router.Any("/oneapi/proxy/:channelid/*target", v1.Relay)
		relayV1Router.POST("/completions", v1.Relay)
		relayV1Router.POST("/chat/completions", v1.Relay)
		relayV1Router.POST("/edits", v1.Relay)
		relayV1Router.POST("/images/generations", v1.Relay)
		relayV1Router.POST("/images/edits", v1.RelayNotImplemented)
		relayV1Router.POST("/images/variations", v1.RelayNotImplemented)
		relayV1Router.POST("/embeddings", v1.Relay)
		relayV1Router.POST("/engines/:model/embeddings", v1.Relay)
		relayV1Router.POST("/audio/transcriptions", v1.Relay)
		relayV1Router.POST("/audio/translations", v1.Relay)
		relayV1Router.POST("/audio/speech", v1.Relay)
		relayV1Router.GET("/files", v1.RelayNotImplemented)
		relayV1Router.POST("/files", v1.RelayNotImplemented)
		relayV1Router.DELETE("/files/:id", v1.RelayNotImplemented)
		relayV1Router.GET("/files/:id", v1.RelayNotImplemented)
		relayV1Router.GET("/files/:id/content", v1.RelayNotImplemented)
		relayV1Router.POST("/fine_tuning/jobs", v1.RelayNotImplemented)
		relayV1Router.GET("/fine_tuning/jobs", v1.RelayNotImplemented)
		relayV1Router.GET("/fine_tuning/jobs/:id", v1.RelayNotImplemented)
		relayV1Router.POST("/fine_tuning/jobs/:id/cancel", v1.RelayNotImplemented)
		relayV1Router.GET("/fine_tuning/jobs/:id/events", v1.RelayNotImplemented)
		relayV1Router.DELETE("/models/:model", v1.RelayNotImplemented)
		relayV1Router.POST("/moderations", v1.Relay)
		relayV1Router.POST("/assistants", v1.RelayNotImplemented)
		relayV1Router.GET("/assistants/:id", v1.RelayNotImplemented)
		relayV1Router.POST("/assistants/:id", v1.RelayNotImplemented)
		relayV1Router.DELETE("/assistants/:id", v1.RelayNotImplemented)
		relayV1Router.GET("/assistants", v1.RelayNotImplemented)
		relayV1Router.POST("/assistants/:id/files", v1.RelayNotImplemented)
		relayV1Router.GET("/assistants/:id/files/:fileId", v1.RelayNotImplemented)
		relayV1Router.DELETE("/assistants/:id/files/:fileId", v1.RelayNotImplemented)
		relayV1Router.GET("/assistants/:id/files", v1.RelayNotImplemented)
		relayV1Router.POST("/threads", v1.RelayNotImplemented)
		relayV1Router.GET("/threads/:id", v1.RelayNotImplemented)
		relayV1Router.POST("/threads/:id", v1.RelayNotImplemented)
		relayV1Router.DELETE("/threads/:id", v1.RelayNotImplemented)
		relayV1Router.POST("/threads/:id/messages", v1.RelayNotImplemented)
		relayV1Router.GET("/threads/:id/messages/:messageId", v1.RelayNotImplemented)
		relayV1Router.POST("/threads/:id/messages/:messageId", v1.RelayNotImplemented)
		relayV1Router.GET("/threads/:id/messages/:messageId/files/:filesId", v1.RelayNotImplemented)
		relayV1Router.GET("/threads/:id/messages/:messageId/files", v1.RelayNotImplemented)
		relayV1Router.POST("/threads/:id/runs", v1.RelayNotImplemented)
		relayV1Router.GET("/threads/:id/runs/:runsId", v1.RelayNotImplemented)
		relayV1Router.POST("/threads/:id/runs/:runsId", v1.RelayNotImplemented)
		relayV1Router.GET("/threads/:id/runs", v1.RelayNotImplemented)
		relayV1Router.POST("/threads/:id/runs/:runsId/submit_tool_outputs", v1.RelayNotImplemented)
		relayV1Router.POST("/threads/:id/runs/:runsId/cancel", v1.RelayNotImplemented)
		relayV1Router.GET("/threads/:id/runs/:runsId/steps/:stepId", v1.RelayNotImplemented)
		relayV1Router.GET("/threads/:id/runs/:runsId/steps", v1.RelayNotImplemented)
	}
}
