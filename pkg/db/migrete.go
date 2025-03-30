package db

import (
	"fmt"
	"strings"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/pkg"
	"github.com/9688101/HX/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ChooseDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := cfg.Options
	switch {
	case strings.HasPrefix(dsn, "postgres://"):
		// Use PostgreSQL
		return OpenPostgreSQL(cfg)
	case dsn != "":
		// Use MySQL
		return OpenMySQL(cfg)
	default:
		// Use SQLite
		return OpenSQLite(cfg)
	}
}

func OpenMySQL(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	logger.SysLog("using MySQL as database")
	pkg.UsingMySQL = true
	return gorm.Open(mysql.Open(cfg.Options), &gorm.Config{
		PrepareStmt: true, // precompile SQL
	})
}

func OpenPostgreSQL(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	logger.SysLog("using PostgreSQL as database")
	pkg.UsingPostgreSQL = true
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.Options,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		PrepareStmt: true, // precompile SQL
	})
}

func OpenSQLite(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	logger.SysLog("SQL_DSN not set, using SQLite as database")
	pkg.UsingSQLite = true
	dsn := fmt.Sprintf("%s?_busy_timeout=%d", pkg.SQLitePath, pkg.SQLiteBusyTimeout)
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{
		PrepareStmt: true, // precompile SQL
	})
}
