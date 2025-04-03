package app

import (
	"fmt"
	"runtime"
	"time"

	"github.com/fatih/color"
	"gorm.io/gorm"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/pkg/helper"
	"github.com/9688101/HX/pkg/logger"
	"github.com/9688101/HX/pkg/random"
)

// 定义常量
const (
	appName    = "HX"     // 应用程序名称
	appVersion = "v1.0.0" // 版本号
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
func start() {
	// 获取当前时间
	startTime := time.Now().Format("2006-01-02 15:04:05")

	// 获取运行时环境信息
	goVersion := runtime.Version() // Go 版本
	os := runtime.GOOS             // 操作系统
	arch := runtime.GOARCH         // 架构

	// 模拟配置信息和服务端口
	configPath := "/path/to/config.yaml" // 配置文件路径
	listenPort := 8080                   // 服务监听端口

	// 显示启动动画
	fmt.Println(color.CyanString("正在启动 %s %s...", appName, appVersion))
	for i := 0; i < 20; i++ {
		fmt.Print(color.YellowString("█"))
		time.Sleep(200 * time.Millisecond) // 每隔200毫秒显示一个方块
	}
	fmt.Println()

	// 显示启动信息
	color.Green("╔════════════════════════════════════════════════╗")
	color.Green("║          %s %s            ║", appName, appVersion)
	color.Green("╚════════════════════════════════════════════════╝")
	color.Blue("启动时间: %s", startTime)
	color.Blue("运行环境:")
	color.Magenta("  - 操作系统: %s", os)
	color.Magenta("  - 架构: %s", arch)
	color.Magenta("  - Go 版本: %s", goVersion)
	color.Blue("配置文件: %s", configPath)
	color.Blue("服务监听端口: %d", listenPort)
	color.Green("══════════════════════════════════════════════════")
	color.Cyan("欢迎使用 %s！", appName)
}

func init() {
	// 启用颜色输出
	color.NoColor = false
}
