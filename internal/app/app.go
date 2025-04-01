package app

import (
	"embed"
	"os"
	"strconv"

	"github.com/9688101/HX/config"
	v1 "github.com/9688101/HX/internal/controller/http"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/internal/middleware"
	"github.com/9688101/HX/pkg/db"
	"github.com/9688101/HX/pkg/i18n"
	"github.com/9688101/HX/pkg/logger"
	"github.com/9688101/HX/pkg/redis"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Run(buildFS embed.FS) {
	config.InitConfig()
	cfg := config.GetConfig()
	start()
	logger.SetupLogger()
	logger.SysLogf("HX %s started", "1.0.0")

	if os.Getenv("GIN_MODE") != gin.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	if cfg.DebugConfig.DebugEnabled {
		logger.SysLog("running in debug mode")
	}

	// Initialize SQL Database
	db.InitDB()
	entity.MigrateDB(db.DB)
	var err error
	err = CreateRootAccountIfNeed()
	if err != nil {
		logger.FatalLog("database init error: " + err.Error())
	}
	defer func() {
		err := db.CloseDB(db.DB)
		if err != nil {
			logger.FatalLog("failed to close database: " + err.Error())
		}
	}()

	// Initialize Redis
	err = redis.InitRedisClient()
	if err != nil {
		logger.FatalLog("failed to initialize Redis: " + err.Error())
	}

	// if config.MemoryCacheEnabled {
	// 	logger.SysLog("memory cache enabled")
	// 	logger.SysLog(fmt.Sprintf("sync frequency: %d seconds", config.SyncFrequency))
	// 	model.InitChannelCache()
	// }
	// if config.MemoryCacheEnabled {
	// 	go model.SyncOptions(config.SyncFrequency)
	// 	go model.SyncChannelCache(config.SyncFrequency)
	// }
	// if os.Getenv("CHANNEL_TEST_FREQUENCY") != "" {
	// 	frequency, err := strconv.Atoi(os.Getenv("CHANNEL_TEST_FREQUENCY"))
	// 	if err != nil {
	// 		logger.FatalLog("failed to parse CHANNEL_TEST_FREQUENCY: " + err.Error())
	// 	}
	// go controller.AutomaticallyTestChannels(frequency)
	// }
	// if os.Getenv("BATCH_UPDATE_ENABLED") == "true" {
	// 	config.BatchUpdateEnabled = true
	// 	logger.SysLog("batch update enabled with interval " + strconv.Itoa(config.BatchUpdateInterval) + "s")
	// 	model.InitBatchUpdater()
	// }
	// if config.EnableMetric {
	// 	logger.SysLog("metric enabled, will disable channel if too much request failed")
	// }
	// openai.InitTokenEncoders()
	// client.Init()

	// Initialize i18n
	if err := i18n.Init(); err != nil {
		logger.FatalLog("failed to initialize i18n: " + err.Error())
	}

	// Initialize HTTP server
	server := gin.New()
	server.Use(gin.Recovery())
	// This will cause SSE not to work!!!
	//server.Use(gzip.Gzip(gzip.DefaultCompression))
	server.Use(middleware.RequestId())
	server.Use(middleware.Language())
	middleware.SetUpLogger(server)
	// Initialize session store
	store := cookie.NewStore([]byte(config.GetAuthenticationConfig().SessionSecret))
	server.Use(sessions.Sessions("session", store))
	v1.SetRouter(server, buildFS)
	var port = os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(config.GetServerConfig().Port)
	}
	logger.SysLogf("server started on http://localhost:%s", port)
	err = server.Run(":" + port)
	if err != nil {
		logger.FatalLog("failed to start HTTP server: " + err.Error())
	}
}
