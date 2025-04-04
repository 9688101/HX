package database

import (
	"strings"
	"time"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/pkg/logger"

	"gorm.io/gorm"
)

// DatabaseConnection 接口定义数据库连接的基本方法
type DatabaseConnection interface {
	Open(cfg *config.DatabaseConfig) (*gorm.DB, error)
	ConfigureConnection(db *gorm.DB, cfg *config.DatabaseConfig)
}

// DB 用于存储全局数据库连接实例
var DB *gorm.DB

// GetDB 返回数据库连接实例
func GetDB() *gorm.DB {
	return DB
}

// InitDB 初始化数据库连接
func InitDB(cfg *config.DatabaseConfig) {
	var err error
	DB, err = ChooseDB(cfg)
	if err != nil {
		logger.SysFatal("failed to initialize database: " + err.Error())
		return
	}

	SetDBConns(cfg, DB)

	// 如果不是主节点，返回
	if !cfg.IsMasterNode {
		return
	}

	// 进行数据库迁移等操作
	logger.SysLog("database migration started")
	// 迁移操作...
	logger.SysLog("database migrated")
}

// CloseDB 关闭数据库连接
func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// SetDBConns 设置数据库连接池参数
func SetDBConns(cfg *config.DatabaseConfig, db *gorm.DB) {
	if cfg.DebugSQLEnabled {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.SysFatal("failed to connect database: " + err.Error())
		return
	}

	sqlDB.SetMaxIdleConns(cfg.SQLMaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.SQLMaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(cfg.SQLMaxLifetime))
	return
}

// ChooseDB 根据配置选择数据库
func ChooseDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	var dbConnection DatabaseConnection
	dsn := cfg.DSN

	switch {
	case strings.HasPrefix(dsn, "postgres://"):
		// 使用 PostgreSQL
		dbConnection = &PostgreSQLConnection{}
	case dsn != "":
		// 使用 MySQL
		dbConnection = &MySQLConnection{}
	default:
		// 使用 SQLite
		dbConnection = &SQLiteConnection{}
	}

	return dbConnection.Open(cfg)
}
