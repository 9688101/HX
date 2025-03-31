package config

import (
	"github.com/spf13/viper"
)

// setDefaultValues sets default configuration values.
func setDefaultValues() {
	viper.SetDefault("system.system_name", "One API")
	viper.SetDefault("server.address", "http://localhost")
	viper.SetDefault("server.port", 3000)
	viper.SetDefault("server.mode", "release")
	viper.SetDefault("quota.quota_for_new_user", 0)
	viper.SetDefault("rate_limit.global_api_rate_limit_num", 480)
	viper.SetDefault("rate_limit.global_api_rate_limit_duration", 60)
	viper.SetDefault("rate_limit.global_web_rate_limit_num", 480)
	viper.SetDefault("rate_limit.global_web_rate_limit_duration", 60)
	viper.SetDefault("session.cookie_name", "session_id")
	viper.SetDefault("session.cookie_path", "/")
	viper.SetDefault("session.cookie_secure", true)
	viper.SetDefault("session.cookie_http_only", true)
	viper.SetDefault("session.max_age", 3600) // 1 hour
	viper.SetDefault("redis.database", 8)
	viper.SetDefault("authentication.jwt_expiration", 3600) // 1 hour
	viper.SetDefault("debug.log_level", "info")
	viper.SetDefault("log.log_level", "info")
	viper.SetDefault("log.log_path", "./logs")
	viper.SetDefault("log.log_filename", "app.log")
	viper.SetDefault("log.max_size", 100)
	viper.SetDefault("log.max_backups", 5)
	viper.SetDefault("log.max_age", 7)
	viper.SetDefault("log.compress", true)
	viper.SetDefault("log.enable_console", true)
	viper.SetDefault("database.type", "sqlite3")
	viper.SetDefault("database.database", "app.db")
	viper.SetDefault("general.time_zone", "Asia/Shanghai")
	viper.SetDefault("general.default_language", "zh")
}
