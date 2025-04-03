package config

import "time"

var StartTime time.Time

// Config 配置.
type Config struct {
	ServerConfig    *ServerConfig    `mapstructure:"server" yaml:"server"`         // 服务器配置
	RateLimitConfig *RateLimitConfig `mapstructure:"rate_limit" yaml:"rate_limit"` // 限流配置
	GeneralConfig   *GeneralConfig   `mapstructure:"general" yaml:"general"`       // 通用配置
	RedisConfig     *RedisConfig     `mapstructure:"redis" yaml:"redis"`           // Redis 配置保持为指针类型，并使用 "redis" 标签
	DatabaseConfig  *DatabaseConfig  `mapstructure:"database" yaml:"database"`     // 数据库配置
	MailConfig      *MailConfig      `mapstructure:"mail" yaml:"mail"`             // 邮件配置
}

// MailConfig 定义邮件发送的配置
type MailConfig struct {
	Provider    string `mapstructure:"provider"`     // 邮件服务提供商标识，如 "smtp"，未来可扩展其他如 sendgrid、aliyun 等
	SMTPAccount string `mapstructure:"smtp_account"` // SMTP 账号
	SMTPToken   string `mapstructure:"smtp_token"`   // SMTP 口令/Token
	SMTPServer  string `mapstructure:"smtp_server"`  // SMTP 服务器地址
	SMTPPort    int    `mapstructure:"smtp_port"`    // SMTP 服务器端口
	SMTPFrom    string `mapstructure:"smtp_from"`    // 发件人邮箱地址
}

// 服务器配置
type ServerConfig struct {
	Port          int    `mapstructure:"port"`           // 服务器端口
	SessionSecret string `mapstructure:"session_secret"` // 会话密钥
	SystemName    string `mapstructure:"system_name"`    // 系统名称
	GINMode       string `mapstructure:"gin_mode"`       // Gin 模式
}

// reids 配置
type RedisConfig struct {
	RedisConnString string `mapstructure:"redis_conn_string" yaml:"redis_conn_string"` // Redis 连接字符串
	RedisPassword   string `mapstructure:"redis_password" yaml:"redis_password"`       // Redis 密码
	Database        int    `mapstructure:"database" yaml:"database"`                   // Redis 数据库索引
	RedisMasterName string `mapstructure:"redis_master_name" yaml:"redis_master_name"` // Redis 主节点名称（集群模式）
	SyncFrequency   string `mapstructure:"sync_frequency" yaml:"sync_frequency"`       // 同步频率

}

// 数据库配置
type DatabaseConfig struct {
	DSN               string `mapstructure:"dsn" yaml:"dsn"`                                 // 数据库连接参数
	SQLMaxIdleConns   int    `mapstructure:"sql_max_idle_conns" yaml:"sql_max_idle_conns"`   // 最大空闲连接数
	SQLMaxOpenConns   int    `mapstructure:"sql_max_open_conns" yaml:"sql_max_open_conns"`   // 最大打开连接数
	SQLMaxLifetime    int    `mapstructure:"sql_max_lifetime" yaml:"sql_max_lifetime"`       // 连接最大存活时间
	IsMasterNode      bool   `mapstructure:"is_master_node" yaml:"is_master_node"`           // 是否是主节点
	DebugSQLEnabled   bool   `mapstructure:"debug_sql_enabled" yaml:"debug_sql_enabled"`     // 是否启用SQL调试
	SQLitePath        string `mapstructure:"sqlite_path" yaml:"sqlite_path"`                 // SQLite数据库路径
	SQLiteBusyTimeout int    `mapstructure:"sqlite_busy_timeout" yaml:"sqlite_busy_timeout"` // SQLite数据库繁忙超时时间
	UsingSQLite       bool   `mapstructure:"using_sqlite" yaml:"using_sqlite"`               // 是否使用SQLite数据库
	UsingPostgreSQL   bool   `mapstructure:"using_postgresql" yaml:"using_postgresql"`       // 是否使用PostgreSQL数据库
	UsingMySQL        bool   `mapstructure:"using_mysql" yaml:"using_mysql"`                 // 是否使用MySQL数据库
}

// 通用配置
type GeneralConfig struct {
	InitialRootToken       string `mapstructure:"initial_root_token"`        // 初始化根用户令牌
	InitialRootAccessToken string `mapstructure:"initial_root_access_token"` // 初始化根用户访问令牌
	MemoryCacheEnabled     bool   `mapstructure:"memory_cache_enabled"`      // 是否启用内存缓存
	DebugEnabled           bool   `mapstructure:"debug_enabled"`             // 是否启用调试
	OnlyOneLogFile         bool   `mapstructure:"only_one_log_file"`         // 是否只使用一个日志文件

}

// 限流配置
type RateLimitConfig struct {
	GlobalApiRateLimitNum          int           `mapstructure:"global_api_rate_limit_num" yaml:"global_api_rate_limit_num"`                   // 全局API限流次数
	GlobalApiRateLimitDuration     int64         `mapstructure:"global_api_rate_limit_duration" yaml:"global_api_rate_limit_duration"`         // 全局API限流时长
	GlobalWebRateLimitNum          int           `mapstructure:"global_web_rate_limit_num" yaml:"global_web_rate_limit_num"`                   // 全局Web限流次数
	GlobalWebRateLimitDuration     int64         `mapstructure:"global_web_rate_limit_duration" yaml:"global_web_rate_limit_duration"`         // 全局Web限流时长
	UploadRateLimitNum             int           `mapstructure:"upload_rate_limit_num" yaml:"upload_rate_limit_num"`                           // 上传限流次数
	UploadRateLimitDuration        int64         `mapstructure:"upload_rate_limit_duration" yaml:"upload_rate_limit_duration"`                 // 上传限流时长
	DownloadRateLimitNum           int           `mapstructure:"download_rate_limit_num" yaml:"download_rate_limit_num"`                       // 下载限流次数
	DownloadRateLimitDuration      int64         `mapstructure:"download_rate_limit_duration" yaml:"download_rate_limit_duration"`             // 下载限流时长
	CriticalRateLimitNum           int           `mapstructure:"critical_rate_limit_num" yaml:"critical_rate_limit_num"`                       // 关键操作限流次数
	CriticalRateLimitDuration      int64         `mapstructure:"critical_rate_limit_duration" yaml:"critical_rate_limit_duration"`             // 关键操作限流时长
	RateLimitKeyExpirationDuration time.Duration `mapstructure:"rate_limit_key_expiration_duration" yaml:"rate_limit_key_expiration_duration"` // 限流键过期时长
}
