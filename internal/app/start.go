package app

import (
	"fmt"
	"runtime"
	"time"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/pkg/helper"
	"github.com/9688101/HX/pkg/logger"
	"github.com/9688101/HX/pkg/random"
)

func CreateRootAccountIfNeed(db *gorm.DB) error {
	var user entity.User
	//if user.Status != util.UserStatusEnabled {
	if err := db.First(&user).Error; err != nil {
		logger.SysLog("no user exists, creating a root user for you: username is root, password is 123456")
		hashedPassword, err := helper.Password2Hash("123456")
		if err != nil {
			return err
		}
		accessToken := random.GetUUID()
		if config.GetGeneralConfig().InitialRootAccessToken != "" {
			accessToken = config.GetGeneralConfig().InitialRootAccessToken
		}
		rootUser := entity.User{
			Username:    "root",
			Password:    hashedPassword,
			Role:        entity.RoleRootUser,
			Status:      entity.UserStatusEnabled,
			DisplayName: "Root User",
			AccessToken: accessToken,
			AffCode:     random.GetRandomString(4),
		}
		db.Create(&rootUser)
	}
	return nil
}

var (
	appName    = "HUIXIONG AI" // 替换为你的应用名称
	appVersion = "v1.0.0"      // 替换为你的应用版本
)

// start 启动帮助函数 (重新设计)
func start(cfg *config.Config) {
	// 获取当前时间
	startTime := time.Now().Format("2006-01-02 15:04:05")

	// 获取运行时环境信息
	goVersion := runtime.Version() // Go 版本
	osInfo := runtime.GOOS         // 操作系统
	arch := runtime.GOARCH         // 架构

	// 显示启动动画
	logger.SysInfo("正在启动 %s %s...", zap.String("appName", appName), zap.String("appVersion", appVersion))
	for i := 0; i < 20; i++ {
		fmt.Print(color.YellowString("█"))
		time.Sleep(200 * time.Millisecond) // 每隔200毫秒显示一个方块
	}
	fmt.Println()

	// 显示启动信息
	color.Green("╔════════════════════════════════════════════════╗")
	color.Green("║          %s %s            ║", appName, appVersion)
	color.Green("╚════════════════════════════════════════════════╝")
	color.Blue("启动时间: %s", startTime)
	color.Blue("运行环境:")
	color.Magenta("  - 操作系统: %s", osInfo)
	color.Magenta("  - 架构: %s", arch)
	color.Magenta("  - Go 版本: %s", goVersion)
	color.Blue("配置文件: %s", cfg.LoggerConfig.LogDir)
	color.Blue("服务监听端口: %d", cfg.ServerConfig.Port)
	color.Green("══════════════════════════════════════════════════")
	color.Cyan("欢迎使用 %s！", appName)
}

func init() {
	// 启用颜色输出
	color.NoColor = false
}
