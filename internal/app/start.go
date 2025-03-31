package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/pkg/db"
	"github.com/9688101/HX/pkg/logger"
	"github.com/9688101/HX/pkg/random"
	"github.com/9688101/HX/pkg/utils"
)

func CreateRootAccountIfNeed() error {
	var user entity.User
	//if user.Status != util.UserStatusEnabled {
	if err := db.DB.First(&user).Error; err != nil {
		logger.SysLog("no user exists, creating a root user for you: username is root, password is 123456")
		hashedPassword, err := utils.Password2Hash("123456")
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
			Quota:       500000000000000,
		}
		db.DB.Create(&rootUser)
		// if config.GetGeneralConfig().InitialRootToken != "" {
		// 	logger.SysLog("creating initial root token as requested")
		// 	token := entity.Token{
		// 		Id:             1,
		// 		UserId:         rootUser.Id,
		// 		Key:            config.InitialRootToken,
		// 		Status:         entity.TokenStatusEnabled,
		// 		Name:           "Initial Root Token",
		// 		CreatedTime:    helper.GetTimestamp(),
		// 		AccessedTime:   helper.GetTimestamp(),
		// 		ExpiredTime:    -1,
		// 		RemainQuota:    500000000000000,
		// 		UnlimitedQuota: true,
		// 	}
		// 	db.DB.Create(&token)
	}
	// }
	return nil
}

func start() {
	flag.Parse()

	if *PrintVersion {
		os.Exit(0)
	}

	if *PrintHelp {
		printHelp()
		os.Exit(0)
	}

	if os.Getenv("SESSION_SECRET") != "" {
		if os.Getenv("SESSION_SECRET") == "random_string" {
			logger.SysError("SESSION_SECRET is set to an example value, please change it to a random string.")
		} else {
			config.GetSessionConfig().SessionSecret = os.Getenv("SESSION_SECRET")
		}
	}
	if *LogDir != "" {
		var err error
		*LogDir, err = filepath.Abs(*LogDir)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := os.Stat(*LogDir); os.IsNotExist(err) {
			err = os.Mkdir(*LogDir, 0777)
			if err != nil {
				log.Fatal(err)
			}
		}
		logger.LogDir = *LogDir
	}
}

var (
	Port         = flag.Int("port", 3000, "the listening port")
	PrintVersion = flag.Bool("version", false, "print version and exit")
	PrintHelp    = flag.Bool("help", false, "print help and exit")
	LogDir       = flag.String("log-dir", "./logs", "specify the log directory")
)

func printHelp() {
	fmt.Println("Copyright (C) 2023 JustSong. All rights reserved.")
	fmt.Println("GitHub: https://github.com/songquanpeng/one-api")
	fmt.Println("Usage: one-api [--port <port>] [--log-dir <log directory>] [--version] [--help]")
}
