package database

import (
	"fmt"

	"github.com/9688101/HX/config"
	"github.com/9688101/HX/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// MySQLConnection 用于 MySQL 数据库的连接
type MySQLConnection struct{}

func (m *MySQLConnection) Open(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	logger.SysLog("using MySQL as database")
	cfg.UsingMySQL = true
	return gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{
		PrepareStmt: true, // 预编译 SQL
	})
}

func (m *MySQLConnection) ConfigureConnection(db *gorm.DB, cfg *config.DatabaseConfig) {
	SetDBConns(cfg, db)
	return
}

// PostgreSQLConnection 用于 PostgreSQL 数据库的连接
type PostgreSQLConnection struct{}

func (p *PostgreSQLConnection) Open(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	logger.SysLog("using PostgreSQL as database")
	cfg.UsingPostgreSQL = true
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.DSN,
		PreferSimpleProtocol: true, // 禁用隐式预处理语句
	}), &gorm.Config{
		PrepareStmt: true, // 预编译 SQL
	})
}

func (p *PostgreSQLConnection) ConfigureConnection(db *gorm.DB, cfg *config.DatabaseConfig) {
	SetDBConns(cfg, db)
	return
}

// SQLiteConnection 用于 SQLite 数据库的连接
type SQLiteConnection struct{}

func (s *SQLiteConnection) Open(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	logger.SysLog("using SQLite as database")
	dsn := fmt.Sprintf("%s?_busy_timeout=%d", "hx.db", cfg.SQLiteBusyTimeout)
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{
		PrepareStmt: true, // 预编译 SQL
	})
}

func (s *SQLiteConnection) ConfigureConnection(db *gorm.DB, cfg *config.DatabaseConfig) {
	SetDBConns(cfg, db)
	return
}
