package app

import (
	"embed"
	"fmt"

	"github.com/9688101/HX/config"
	v1 "github.com/9688101/HX/internal/controller/http"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/internal/middleware"
	"github.com/9688101/HX/pkg/database"
	"github.com/9688101/HX/pkg/i18n"
	"github.com/9688101/HX/pkg/logger"
	"github.com/9688101/HX/pkg/redis"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Run 启动应用程序
func Run(buildFS embed.FS) {
	config.InitConfig()
	cfg := config.GetConfig()
	logger.InitLogger(cfg.LoggerConfig)

	if cfg.ServerConfig.GINMode != gin.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	if cfg.GeneralConfig.DebugEnabled {
		logger.SysInfo("running in debug mode")
	}

	// Initialize SQL Database
	database.InitDB(cfg.DatabaseConfig)
	db := database.GetDB()
	entity.MigrateDB(db)
	var err error
	err = CreateRootAccountIfNeed(db)
	if err != nil {
		logger.SysFatal("database init error: " + err.Error())
	}
	defer func() {
		err := database.CloseDB(db)
		if err != nil {
			logger.SysFatal("failed to close database: " + err.Error())
		}
	}()

	// Initialize Redis
	err = redis.InitRedisClient(cfg.RedisConfig)
	if err != nil {
		logger.SysFatal("failed to initialize Redis: " + err.Error())
	}
	// client.Init()
	// Initialize i18n
	if err := i18n.Init(); err != nil {
		logger.SysFatal("failed to initialize i18n: " + err.Error())
	}

	// Initialize HTTP server
	server := gin.Default()
	// server.Use(gin.Recovery())
	// This will cause SSE not to work!!!
	//server.Use(gzip.Gzip(gzip.DefaultCompression))
	// server.Use(middleware.RequestId())
	// server.Use(middleware.Language())
	// server.Use(middleware.GinLoggerMiddleware())
	// server.Use(middleware.CORS())
	// server.Use(middleware.Cache())
	// server.Use(middleware.TurnstileCheck())
	server.Use(middleware.JWTAuth())
	// Initialize session store
	store := cookie.NewStore([]byte(cfg.ServerConfig.SessionSecret))
	server.Use(sessions.Sessions("session", store))
	v1.SetRouter(server, buildFS)
	var port = cfg.ServerConfig.Port
	logger.SysInfo("server started on http://localhost:" + port)
	for _, route := range server.Routes() {
		fmt.Printf("%-6s -> %-40s  Handler: %s\n", route.Method, route.Path, route.Handler)
	}

	start(cfg) // 调用重新设计的 start 函数，并传递配置和日志器
	err = server.Run(":" + port)
	if err != nil {
		logger.SysFatal("failed to start HTTP server: " + err.Error())
	}
}
