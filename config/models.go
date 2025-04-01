package config

import "time"

// Config 配置.
type Config struct {
	SystemConfig         *SystemConfig         `mapstructure:"system" yaml:"system"`
	ServerConfig         *ServerConfig         `mapstructure:"server" yaml:"server"`
	AuthenticationConfig *AuthenticationConfig `mapstructure:"authentication" yaml:"authentication"`
	DebugConfig          *DebugConfig          `mapstructure:"debug" yaml:"debug"`
	SMTPConfig           *SMTPConfig           `mapstructure:"smtp" yaml:"smtp"`
	OAuthConfig          *OAuthConfig          `mapstructure:"oauth" yaml:"oauth"`
	WeChatConfig         *WeChatConfig         `mapstructure:"wechat" yaml:"wechat"`
	MessageConfig        *MessageConfig        `mapstructure:"message" yaml:"message"`
	TurnstileConfig      *TurnstileConfig      `mapstructure:"turnstile" yaml:"turnstile"`
	RateLimitConfig      *RateLimitConfig      `mapstructure:"rate_limit" yaml:"rate_limit"`
	GeneralConfig        *GeneralConfig        `mapstructure:"general" yaml:"general"`
	RedisConfig          *RedisConfig          `mapstructure:"redis" yaml:"redis"` // Redis 配置保持为指针类型，并使用 "redis" 标签
	DatabaseConfig       *DatabaseConfig       `mapstructure:"database" yaml:"database"`
}

// 系统配置
type SystemConfig struct {
	SystemName      string `mapstructure:"system_name" yaml:"system_name"`
	Notice          string `mapstructure:"notice" yaml:"notice"`
	About           string `mapstructure:"about" yaml:"about"`
	HomePageContent string `mapstructure:"home_page_content" yaml:"home_page_content"`
	Footer          string `mapstructure:"footer" yaml:"footer"`
	Logo            string `mapstructure:"logo" yaml:"logo"`
	Theme           string `mapstructure:"theme" yaml:"theme"`
	Version         string `mapstructure:"version" yaml:"version"`
}

// 服务器配置
type ServerConfig struct {
	Address        string `mapstructure:"address"`
	Port           int    `mapstructure:"port"`
	StartTime      time.Time
	OnlyOneLogFile bool `mapstructure:"only_one_log_file"`
}

// reids 配置
type RedisConfig struct {
	RedisConnString string `mapstructure:"redis_conn_string" yaml:"redis_conn_string"`
	RedisPassword   string `mapstructure:"redis_password" yaml:"redis_password"`
	Database        int    `mapstructure:"database" yaml:"database"`
	RedisMasterName string `mapstructure:"redis_master_name" yaml:"redis_master_name"`
	SyncFrequency   string `mapstructure:"sync_frequency" yaml:"sync_frequency"`
}

// 数据库配置
type DatabaseConfig struct {
	Options           string `mapstructure:"options" yaml:"options"` // 数据库连接参数
	IsMasterNode      bool   `mapstructure:"is_master_node" yaml:"is_master_node"`
	DebugSQLEnabled   bool   `mapstructure:"debug_sql_enabled" yaml:"debug_sql_enabled"`
	SQLitePath        string `mapstructure:"sqlite_path" yaml:"sqlite_path"`
	SQLiteBusyTimeout int    `mapstructure:"sqlite_busy_timeout" yaml:"sqlite_busy_timeout"`
	UsingSQLite       bool   `mapstructure:"using_sqlite" yaml:"using_sqlite"`
	UsingPostgreSQL   bool   `mapstructure:"using_postgresql" yaml:"using_postgresql"`
	UsingMySQL        bool   `mapstructure:"using_mysql" yaml:"using_mysql"`
}

// 调试配置
type DebugConfig struct {
	DebugEnabled       bool `mapstructure:"debug_enabled" yaml:"debug_enabled"`
	MemoryCacheEnabled bool `mapstructure:"memory_cache_enabled" yaml:"memory_cache_enabled"`
}

// 通用配置
type GeneralConfig struct {
	InitialRootToken       string `mapstructure:"initial_root_token" yaml:"initial_root_token"`
	InitialRootAccessToken string `mapstructure:"initial_root_access_token" yaml:"initial_root_access_token"`
	RootUserEmail          string `mapstructure:"root_user_email" yaml:"root_user_email"`
	SyncFrequency          int    `mapstructure:"sync_frequency" yaml:"sync_frequency"`
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
}

// 消息配置
type MessageConfig struct {
	MessagePusherAddress string `mapstructure:"message_pusher_address" yaml:"message_pusher_address"`
	MessagePusherToken   string `mapstructure:"message_pusher_token" yaml:"message_pusher_token"`
}

// 认证配置
type AuthenticationConfig struct {
	PasswordLoginEnabled          bool     `mapstructure:"password_login_enabled" yaml:"password_login_enabled"`
	PasswordRegisterEnabled       bool     `mapstructure:"password_register_enabled" yaml:"password_register_enabled"`
	EmailVerificationEnabled      bool     `mapstructure:"email_verification_enabled" yaml:"email_verification_enabled"`
	GitHubOAuthEnabled            bool     `mapstructure:"github_oauth_enabled" yaml:"github_oauth_enabled"`
	OidcEnabled                   bool     `mapstructure:"oidc_enabled" yaml:"oidc_enabled"`
	WeChatAuthEnabled             bool     `mapstructure:"wechat_auth_enabled" yaml:"wechat_auth_enabled"`
	TurnstileCheckEnabled         bool     `mapstructure:"turnstile_check_enabled" yaml:"turnstile_check_enabled"`
	RegisterEnabled               bool     `mapstructure:"register_enabled" yaml:"register_enabled"`
	EmailDomainRestrictionEnabled bool     `mapstructure:"email_domain_restriction_enabled" yaml:"email_domain_restriction_enabled"`
	EmailDomainWhitelist          []string `mapstructure:"email_domain_whitelist" yaml:"email_domain_whitelist"`
	SessionSecret                 string   `mapstructure:"session_secret" yaml:"session_secret"`
}

// OAuth 配置
type OAuthConfig struct {
	GitHubClientId            string `mapstructure:"github_client_id" yaml:"github_client_id"`
	GitHubClientSecret        string `mapstructure:"github_client_secret" yaml:"github_client_secret"`
	LarkClientId              string `mapstructure:"lark_client_id" yaml:"lark_client_id"`
	LarkClientSecret          string `mapstructure:"lark_client_secret" yaml:"lark_client_secret"`
	OidcClientId              string `mapstructure:"oidc_client_id" yaml:"oidc_client_id"`
	OidcClientSecret          string `mapstructure:"oidc_client_secret" yaml:"oidc_client_secret"`
	OidcWellKnown             string `mapstructure:"oidc_well_known" yaml:"oidc_well_known"`
	OidcAuthorizationEndpoint string `mapstructure:"oidc_authorization_endpoint" yaml:"oidc_authorization_endpoint"`
	OidcTokenEndpoint         string `mapstructure:"oidc_token_endpoint" yaml:"oidc_token_endpoint"`
	OidcUserinfoEndpoint      string `mapstructure:"oidc_userinfo_endpoint" yaml:"oidc_userinfo_endpoint"`
}

// 微信配置
type WeChatConfig struct {
	WeChatServerAddress         string `mapstructure:"wechat_server_address" yaml:"wechat_server_address"`
	WeChatServerToken           string `mapstructure:"wechat_server_token" yaml:"wechat_server_token"`
	WeChatAccountQRCodeImageURL string `mapstructure:"wechat_account_qr_code_image_url" yaml:"wechat_account_qr_code_image_url"`
}

// SMTP 配置
type SMTPConfig struct {
	SMTPServer  string `mapstructure:"smtp_server" yaml:"smtp_server"`
	SMTPPort    int    `mapstructure:"smtp_port" yaml:"smtp_port"`
	SMTPAccount string `mapstructure:"smtp_account" yaml:"smtp_account"`
	SMTPFrom    string `mapstructure:"smtp_from" yaml:"smtp_from"`
	SMTPToken   string `mapstructure:"smtp_token" yaml:"smtp_token"`
}

// 验证码配置
type TurnstileConfig struct {
	TurnstileSiteKey   string `mapstructure:"turnstile_site_key" yaml:"turnstile_site_key"`
	TurnstileSecretKey string `mapstructure:"turnstile_secret_key" yaml:"turnstile_secret_key"`
}
