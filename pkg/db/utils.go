package db

import (
	"fmt"
	"strings"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ChooseDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	options := cfg.Options
	switch {
	case strings.HasPrefix(options, "postgres://"):
		// Use PostgreSQL
		return OpenPostgreSQL(cfg)
	case options != "":
		// Use MySQL
		return OpenMySQL(cfg)
	default:
		// Use SQLite
		return OpenSQLite(cfg)
	}
}

func OpenMySQL(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	logger.SysLog("using MySQL as database")
	cfg.UsingMySQL = true
	return gorm.Open(mysql.Open(cfg.Options), &gorm.Config{
		PrepareStmt: true, // precompile SQL
	})
}

func OpenPostgreSQL(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	logger.SysLog("using PostgreSQL as database")
	cfg.UsingPostgreSQL = true
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.Options,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		PrepareStmt: true, // precompile SQL
	})
}

func OpenSQLite(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	var path string = "hx.db"
	logger.SysLog("SQL_DSN not set, using SQLite as database")
	dsn := fmt.Sprintf("%s?_busy_timeout=%d", path, cfg.SQLiteBusyTimeout)
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{
		PrepareStmt: true, // precompile SQL
	})
}
