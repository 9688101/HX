package db

// import (
// 	"fmt"
// 	"hx-helper/config"
// 	"hx-helper/pkg/logger"
// 	"strings"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// func chooseDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
// 	dsn := cfg.DSN
// 	switch {
// 	case strings.HasPrefix(dsn, "postgres://"):
// 		// Use PostgreSQL
// 		return openPostgreSQL(cfg)
// 	case dsn != "":
// 		// Use MySQL
// 		return openMySQL(cfg)
// 	default:
// 		// Use SQLite
// 		return openSQLite(cfg)
// 	}
// }

// func openMySQL(cfg *config.DatabaseConfig) (*gorm.DB, error) {
// 	logger.SysLog("using MySQL as database")
// 	cfg.UsingSql.MySQL = true
// 	return gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{
// 		PrepareStmt: true, // precompile SQL
// 	})
// }

// func openPostgreSQL(cfg *config.DatabaseConfig) (*gorm.DB, error) {
// 	logger.SysLog("using PostgreSQL as database")
// 	cfg.UsingSql.PostgreSQL = true
// 	return gorm.Open(postgres.New(postgres.Config{
// 		DSN:                  cfg.DSN,
// 		PreferSimpleProtocol: true, // disables implicit prepared statement usage
// 	}), &gorm.Config{
// 		PrepareStmt: true, // precompile SQL
// 	})
// }

// func openSQLite(cfg *config.DatabaseConfig) (*gorm.DB, error) {
// 	logger.SysLog("SQL_DSN not set, using SQLite as database")
// 	cfg.UsingSql.SQLite = true
// 	dsn := fmt.Sprintf("%s?_busy_timeout=%d", cfg.Sqllite.Path, cfg.Sqllite.BusyTimeout)
// 	return gorm.Open(sqlite.Open(dsn), &gorm.Config{
// 		PrepareStmt: true, // precompile SQL
// 	})
// }

// // func migrateDB() error {
// // 	var err error
// // 	if err = DB.AutoMigrate(&Channel{}); err != nil {
// // 		return err
// // 	}
// // 	if err = DB.AutoMigrate(&Token{}); err != nil {
// // 		return err
// // 	}
// // 	if err = DB.AutoMigrate(&User{}); err != nil {
// // 		return err
// // 	}
// // 	if err = DB.AutoMigrate(&Option{}); err != nil {
// // 		return err
// // 	}
// // 	if err = DB.AutoMigrate(&Redemption{}); err != nil {
// // 		return err
// // 	}
// // 	if err = DB.AutoMigrate(&Ability{}); err != nil {
// // 		return err
// // 	}
// // 	if err = DB.AutoMigrate(&Log{}); err != nil {
// // 		return err
// // 	}
// // 	if err = DB.AutoMigrate(&Channel{}); err != nil {
// // 		return err
// // 	}
// // 	return nil
// // }

// func migrateLOGDB() error {
// 	// var err error
// 	// if err = LOG_DB.AutoMigrate(&Log{}); err != nil {
// 	// 	return err
// 	// }
// 	return nil
// }
