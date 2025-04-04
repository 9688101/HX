package config

import "time"

// StartTime 可用于记录应用启动时间
var StartTime time.Time

// Config 定义了系统所有的配置项
type Config struct {
	ServerConfig    *ServerConfig    `mapstructure:"server" yaml:"server"`
	RateLimitConfig *RateLimitConfig `mapstructure:"rate_limit" yaml:"rate_limit"`
	GeneralConfig   *GeneralConfig   `mapstructure:"general" yaml:"general"`
	RedisConfig     *RedisConfig     `mapstructure:"redis" yaml:"redis"`
	DatabaseConfig  *DatabaseConfig  `mapstructure:"database" yaml:"database"`
	MailConfig      *MailConfig      `mapstructure:"mail" yaml:"mail"`
	LoggerConfig    *LoggerConfig    `mapstructure:"logger" yaml:"logger"`
}

// LoggerConfig 定义日志模块配置
type LoggerConfig struct {
	LogDir         string `mapstructure:"log_dir" yaml:"log_dir"`                     // 日志文件存储目录（不为空时写入文件）
	OnlyOneLogFile bool   `mapstructure:"only_one_log_file" yaml:"only_one_log_file"` // 是否使用单一日志文件
	DebugEnabled   bool   `mapstructure:"debug_enabled" yaml:"debug_enabled"`         // 是否启用 Debug 级别日志
	UseZap         bool   `mapstructure:"use_zap" yaml:"use_zap"`                     // 是否使用 zap 日志库
}

// MailConfig 定义邮件发送配置
type MailConfig struct {
	Provider    string `mapstructure:"provider" yaml:"provider"`         // 邮件服务提供商标识，如 "smtp"
	SMTPAccount string `mapstructure:"smtp_account" yaml:"smtp_account"` // SMTP 账号
	SMTPToken   string `mapstructure:"smtp_token" yaml:"smtp_token"`     // SMTP 密码或令牌
	SMTPServer  string `mapstructure:"smtp_server" yaml:"smtp_server"`   // SMTP 服务器地址
	SMTPPort    int    `mapstructure:"smtp_port" yaml:"smtp_port"`       // SMTP 服务器端口
	SMTPFrom    string `mapstructure:"smtp_from" yaml:"smtp_from"`       // 发件人邮箱地址
}

// ServerConfig 定义服务器相关配置
type ServerConfig struct {
	Port            string `mapstructure:"port" yaml:"port"`                           // 服务器监听端口
	SessionSecret   string `mapstructure:"session_secret" yaml:"session_secret"`       // 会话加密密钥
	SystemName      string `mapstructure:"system_name" yaml:"system_name"`             // 系统名称
	GINMode         string `mapstructure:"gin_mode" yaml:"gin_mode"`                   // Gin 运行模式（如 "debug", "release"）
	JWTSecret       string `mapstructure:"jwt_secret" yaml:"jwt_secret"`               // JWT 密钥
	FrontendBaseUrl string `mapstructure:"frontend_base_url" yaml:"frontend_base_url"` // 前端基础 URL
	Theme           string `mapstructure:"theme" yaml:"theme"`                         // 主题
}

// RedisConfig 定义 Redis 相关配置
type RedisConfig struct {
	RedisConnString string `mapstructure:"redis_conn_string" yaml:"redis_conn_string"` // Redis 连接字符串
	RedisPassword   string `mapstructure:"redis_password" yaml:"redis_password"`       // Redis 密码
	Database        int    `mapstructure:"database" yaml:"database"`                   // Redis 数据库索引
	RedisMasterName string `mapstructure:"redis_master_name" yaml:"redis_master_name"` // Redis 主节点名称（集群模式）
	SyncFrequency   string `mapstructure:"sync_frequency" yaml:"sync_frequency"`       // 同步频率（如 "5m"）
}

// DatabaseConfig 定义数据库相关配置
type DatabaseConfig struct {
	DSN               string `mapstructure:"dsn" yaml:"dsn"`                                 // 数据库连接参数
	SQLMaxIdleConns   int    `mapstructure:"sql_max_idle_conns" yaml:"sql_max_idle_conns"`   // 最大空闲连接数
	SQLMaxOpenConns   int    `mapstructure:"sql_max_open_conns" yaml:"sql_max_open_conns"`   // 最大打开连接数
	SQLMaxLifetime    int    `mapstructure:"sql_max_lifetime" yaml:"sql_max_lifetime"`       // 连接最大存活时间（秒）
	IsMasterNode      bool   `mapstructure:"is_master_node" yaml:"is_master_node"`           // 是否为主节点
	DebugSQLEnabled   bool   `mapstructure:"debug_sql_enabled" yaml:"debug_sql_enabled"`     // 是否启用 SQL 调试
	SQLitePath        string `mapstructure:"sqlite_path" yaml:"sqlite_path"`                 // SQLite 数据库路径
	SQLiteBusyTimeout int    `mapstructure:"sqlite_busy_timeout" yaml:"sqlite_busy_timeout"` // SQLite 繁忙超时（毫秒）
	UsingSQLite       bool   `mapstructure:"using_sqlite" yaml:"using_sqlite"`               // 是否使用 SQLite
	UsingPostgreSQL   bool   `mapstructure:"using_postgresql" yaml:"using_postgresql"`       // 是否使用 PostgreSQL
	UsingMySQL        bool   `mapstructure:"using_mysql" yaml:"using_mysql"`                 // 是否使用 MySQL
}

// GeneralConfig 定义通用配置
type GeneralConfig struct {
	InitialRootToken       string `mapstructure:"initial_root_token" yaml:"initial_root_token"`               // 初始化根用户令牌
	InitialRootAccessToken string `mapstructure:"initial_root_access_token" yaml:"initial_root_access_token"` // 初始化根用户访问令牌
	MemoryCacheEnabled     bool   `mapstructure:"memory_cache_enabled" yaml:"memory_cache_enabled"`           // 是否启用内存缓存
	DebugEnabled           bool   `mapstructure:"debug_enabled" yaml:"debug_enabled"`                         // 是否启用调试
	OnlyOneLogFile         bool   `mapstructure:"only_one_log_file" yaml:"only_one_log_file"`                 // 是否只使用一个日志文件
}

// RateLimitConfig 定义限流配置
type RateLimitConfig struct {
	GlobalApiRateLimitNum          int           `mapstructure:"global_api_rate_limit_num" yaml:"global_api_rate_limit_num"`                   // 全局 API 限流次数
	GlobalApiRateLimitDuration     int64         `mapstructure:"global_api_rate_limit_duration" yaml:"global_api_rate_limit_duration"`         // 全局 API 限流时长（秒）
	GlobalWebRateLimitNum          int           `mapstructure:"global_web_rate_limit_num" yaml:"global_web_rate_limit_num"`                   // 全局 Web 限流次数
	GlobalWebRateLimitDuration     int64         `mapstructure:"global_web_rate_limit_duration" yaml:"global_web_rate_limit_duration"`         // 全局 Web 限流时长（秒）
	UploadRateLimitNum             int           `mapstructure:"upload_rate_limit_num" yaml:"upload_rate_limit_num"`                           // 上传限流次数
	UploadRateLimitDuration        int64         `mapstructure:"upload_rate_limit_duration" yaml:"upload_rate_limit_duration"`                 // 上传限流时长（秒）
	DownloadRateLimitNum           int           `mapstructure:"download_rate_limit_num" yaml:"download_rate_limit_num"`                       // 下载限流次数
	DownloadRateLimitDuration      int64         `mapstructure:"download_rate_limit_duration" yaml:"download_rate_limit_duration"`             // 下载限流时长（秒）
	CriticalRateLimitNum           int           `mapstructure:"critical_rate_limit_num" yaml:"critical_rate_limit_num"`                       // 关键操作限流次数
	CriticalRateLimitDuration      int64         `mapstructure:"critical_rate_limit_duration" yaml:"critical_rate_limit_duration"`             // 关键操作限流时长（秒）
	RateLimitKeyExpirationDuration time.Duration `mapstructure:"rate_limit_key_expiration_duration" yaml:"rate_limit_key_expiration_duration"` // 限流键过期时长（如 "1m"）
}
