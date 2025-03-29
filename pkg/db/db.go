package db

import "gorm.io/gorm"

// import (
// 	"database/sql"
// 	"hx-helper/config"
// 	"hx-helper/global"
// 	"hx-helper/pkg/env"
// 	"hx-helper/pkg/logger"
// 	"time"

// 	"gorm.io/gorm"
// )

var DB *gorm.DB

// var LOG_DB *gorm.DB

// func GetDB() *gorm.DB {
// 	return DB
// }
// func GetLOGDB() *gorm.DB {
// 	return LOG_DB
// }
// func InitDB(DB *gorm.DB) {
// 	cfg := config.GetConfig()
// 	var err error
// 	DB, err = chooseDB(cfg.Db)
// 	if err != nil {
// 		logger.FatalLog("failed to initialize database: " + err.Error())
// 		return
// 	}

// 	sqlDB := setDBConns(DB)

// 	if !global.IsMasterNode {
// 		return
// 	}

// 	if cfg.Db.UsingSql.MySQL {
// 		_, _ = sqlDB.Exec("DROP INDEX idx_channels_key ON channels;") // TODO: delete this line when most users have upgraded
// 	}

// 	logger.SysLog("database migration started")
// 	// if err = migrateDB(); err != nil {
// 	// 	logger.FatalLog("failed to migrate database: " + err.Error())
// 	// 	return
// 	// }
// 	logger.SysLog("database migrated")
// }

// func InitLogDB() {
// 	cfg := config.GetConfig().Db
// 	dsn := cfg.DSN
// 	if dsn == "" {
// 		LOG_DB = DB
// 		return
// 	}

// 	logger.SysLog("using secondary database for table logs")
// 	var err error
// 	LOG_DB, err = chooseDB(cfg)
// 	if err != nil {
// 		logger.FatalLog("failed to initialize secondary database: " + err.Error())
// 		return
// 	}

// 	setDBConns(LOG_DB)

// 	if !cfg.IsMasterNode {
// 		return
// 	}

// 	logger.SysLog("secondary database migration started")
// 	err = migrateLOGDB()
// 	if err != nil {
// 		logger.FatalLog("failed to migrate secondary database: " + err.Error())
// 		return
// 	}
// 	logger.SysLog("secondary database migrated")
// }
// func closeDB(db *gorm.DB) error {
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return err
// 	}
// 	err = sqlDB.Close()
// 	return err
// }

// func CloseDB(LOG_DB *gorm.DB) error {
// 	if LOG_DB != DB {
// 		err := closeDB(LOG_DB)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return closeDB(DB)
// }
// func setDBConns(db *gorm.DB) *sql.DB {
// 	if global.DebugSQLEnabled {
// 		db = db.Debug()
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		logger.FatalLog("failed to connect database: " + err.Error())
// 		return nil
// 	}

// 	sqlDB.SetMaxIdleConns(env.Int("SQL_MAX_IDLE_CONNS", 100))
// 	sqlDB.SetMaxOpenConns(env.Int("SQL_MAX_OPEN_CONNS", 1000))
// 	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(env.Int("SQL_MAX_LIFETIME", 60)))
// 	return sqlDB
// }
