package db

import (
	"database/sql"
	"time"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/pkg/env"
	"github.com/9688101/HX/pkg/logger"

	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	return err
}
func SetDBConns(db *gorm.DB) *sql.DB {
	if config.GetDebugConfig().DebugSQLEnabled {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.FatalLog("failed to connect database: " + err.Error())
		return nil
	}

	sqlDB.SetMaxIdleConns(env.Int("SQL_MAX_IDLE_CONNS", 100))
	sqlDB.SetMaxOpenConns(env.Int("SQL_MAX_OPEN_CONNS", 1000))
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(env.Int("SQL_MAX_LIFETIME", 60)))
	return sqlDB
}
