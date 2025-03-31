package config

import "time"

// 系统配置
type SystemConfig struct {
	SystemName                  string  `mapstructure:"system_name" yaml:"system_name"`
	ServerAddress               string  `mapstructure:"server_address" yaml:"server_address"`
	Footer                      string  `mapstructure:"footer" yaml:"footer"`
	Logo                        string  `mapstructure:"logo" yaml:"logo"`
	TopUpLink                   string  `mapstructure:"top_up_link" yaml:"top_up_link"`
	ChatLink                    string  `mapstructure:"chat_link" yaml:"chat_link"`
	QuotaPerUnit                float64 `mapstructure:"quota_per_unit" yaml:"quota_per_unit"`
	DisplayInCurrencyEnabled    bool    `mapstructure:"display_in_currency_enabled" yaml:"display_in_currency_enabled"`
	DisplayTokenStatEnabled     bool    `mapstructure:"display_token_stat_enabled" yaml:"display_token_stat_enabled"`
	MaintenanceModeEnabled      bool    `mapstructure:"maintenance_mode_enabled" yaml:"maintenance_mode_enabled"`           // 新增维护模式
	MaintenanceModeMessage      string  `mapstructure:"maintenance_mode_message" yaml:"maintenance_mode_message"`           // 新增维护模式消息
	AnnouncementEnabled         bool    `mapstructure:"announcement_enabled" yaml:"announcement_enabled"`                   // 新增公告开关
	AnnouncementContent         string  `mapstructure:"announcement_content" yaml:"announcement_content"`                   // 新增公告内容
	AnnouncementBackgroundColor string  `mapstructure:"announcement_background_color" yaml:"announcement_background_color"` // 新增公告背景色
	AnnouncementTextColor       string  `mapstructure:"announcement_text_color" yaml:"announcement_text_color"`             // 新增公告文字颜色
	TermsOfServiceLink          string  `mapstructure:"terms_of_service_link" yaml:"terms_of_service_link"`                 // 新增服务条款链接
	PrivacyPolicyLink           string  `mapstructure:"privacy_policy_link" yaml:"privacy_policy_link"`
	Theme                       string  `mapstructure:"theme" yaml:"theme"`
}

// 调试配置
type DebugConfig struct {
	DebugEnabled       bool `mapstructure:"debug_enabled" yaml:"debug_enabled"`
	DebugSQLEnabled    bool `mapstructure:"debug_sql_enabled" yaml:"debug_sql_enabled"`
	MemoryCacheEnabled bool `mapstructure:"memory_cache_enabled" yaml:"memory_cache_enabled"`
	EnableProfiling    bool `mapstructure:"enable_profiling" yaml:"enable_profiling"` // 新增性能分析开关
	ProfilingPort      int  `mapstructure:"profiling_port" yaml:"profiling_port"`     // 新增性能分析端口
}

// 通用配置
type GeneralConfig struct {
	InitialRootToken       string `mapstructure:"initial_root_token" yaml:"initial_root_token"`
	InitialRootAccessToken string `mapstructure:"initial_root_access_token" yaml:"initial_root_access_token"`
	GeminiVersion          string `mapstructure:"gemini_version" yaml:"gemini_version"`
	OnlyOneLogFile         bool   `mapstructure:"only_one_log_file" yaml:"only_one_log_file"`
	RootUserEmail          string `mapstructure:"root_user_email" yaml:"root_user_email"`
	TimeZone               string `mapstructure:"time_zone" yaml:"time_zone"`               // 新增时区配置
	DefaultLanguage        string `mapstructure:"default_language" yaml:"default_language"` // 新增默认语言
}

// 服务器配置
type ServerConfig struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
	Mode    string `mapstructure:"mode" yaml:"mode"` // 新增运行模式 (debug, release)
}

// 会话配置
type SessionConfig struct {
	SessionSecret  string `mapstructure:"session_secret" yaml:"session_secret"`
	CookieName     string `mapstructure:"cookie_name" yaml:"cookie_name"`           // 新增 Cookie 名称
	CookieDomain   string `mapstructure:"cookie_domain" yaml:"cookie_domain"`       // 新增 Cookie 域名
	CookiePath     string `mapstructure:"cookie_path" yaml:"cookie_path"`           // 新增 Cookie 路径
	CookieSecure   bool   `mapstructure:"cookie_secure" yaml:"cookie_secure"`       // 新增 Cookie Secure 属性
	CookieHttpOnly bool   `mapstructure:"cookie_http_only" yaml:"cookie_http_only"` // 新增 Cookie HttpOnly 属性
	MaxAge         int    `mapstructure:"max_age" yaml:"max_age"`                   // 新增 Session 最大生存时间
}

// 同步配置
type SyncConfig struct {
	SyncFrequency int `mapstructure:"sync_frequency" yaml:"sync_frequency"`
}

// 限流配置
type RateLimitConfig struct {
	GlobalApiRateLimitNum          int           `mapstructure:"global_api_rate_limit_num" yaml:"global_api_rate_limit_num"`
	GlobalApiRateLimitDuration     int64         `mapstructure:"global_api_rate_limit_duration" yaml:"global_api_rate_limit_duration"`
	GlobalWebRateLimitNum          int           `mapstructure:"global_web_rate_limit_num" yaml:"global_web_rate_limit_num"`
	GlobalWebRateLimitDuration     int64         `mapstructure:"global_web_rate_limit_duration" yaml:"global_web_rate_limit_duration"`
	UploadRateLimitNum             int           `mapstructure:"upload_rate_limit_num" yaml:"upload_rate_limit_num"`
	UploadRateLimitDuration        int64         `mapstructure:"upload_rate_limit_duration" yaml:"upload_rate_limit_duration"`
	DownloadRateLimitNum           int           `mapstructure:"download_rate_limit_num" yaml:"download_rate_limit_num"`
	DownloadRateLimitDuration      int64         `mapstructure:"download_rate_limit_duration" yaml:"download_rate_limit_duration"`
	CriticalRateLimitNum           int           `mapstructure:"critical_rate_limit_num" yaml:"critical_rate_limit_num"`
	CriticalRateLimitDuration      int64         `mapstructure:"critical_rate_limit_duration" yaml:"critical_rate_limit_duration"`
	RateLimitKeyExpirationDuration time.Duration `mapstructure:"rate_limit_key_expiration_duration" yaml:"rate_limit_key_expiration_duration"`
	IpRateLimitEnabled             bool          `mapstructure:"ip_rate_limit_enabled" yaml:"ip_rate_limit_enabled"`   // 新增 IP 限流开关
	IpRateLimitNum                 int           `mapstructure:"ip_rate_limit_num" yaml:"ip_rate_limit_num"`           // 新增 IP 限流次数
	IpRateLimitDuration            int64         `mapstructure:"ip_rate_limit_duration" yaml:"ip_rate_limit_duration"` // 新增 IP 限流持续时间
}

// 测试配置
type TestConfig struct {
	TestPrompt string `mapstructure:"test_prompt" yaml:"test_prompt"`
}

// 日志配置
type LogConfig struct {
	LogLevel      string `mapstructure:"log_level" yaml:"log_level"`
	LogPath       string `mapstructure:"log_path" yaml:"log_path"`
	LogFilename   string `mapstructure:"log_filename" yaml:"log_filename"`
	MaxSize       int    `mapstructure:"max_size" yaml:"max_size"` // 单位 MB
	MaxBackups    int    `mapstructure:"max_backups" yaml:"max_backups"`
	MaxAge        int    `mapstructure:"max_age" yaml:"max_age"` // 单位天
	Compress      bool   `mapstructure:"compress" yaml:"compress"`
	EnableConsole bool   `mapstructure:"enable_console" yaml:"enable_console"` // 新增控制台输出开关
}

// 消息配置
type MessageConfig struct {
	MessagePusherAddress string `mapstructure:"message_pusher_address" yaml:"message_pusher_address"`
	MessagePusherToken   string `mapstructure:"message_pusher_token" yaml:"message_pusher_token"`
	TelegramBotToken     string `mapstructure:"telegram_bot_token" yaml:"telegram_bot_token"`   // 新增 Telegram Bot Token
	TelegramChatID       string `mapstructure:"telegram_chat_id" yaml:"telegram_chat_id"`       // 新增 Telegram Chat ID
	SlackWebhookURL      string `mapstructure:"slack_webhook_url" yaml:"slack_webhook_url"`     // 新增 Slack Webhook URL
	DiscordWebhookURL    string `mapstructure:"discord_webhook_url" yaml:"discord_webhook_url"` // 新增 Discord Webhook URL
}
