package config

import (
	"io"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/google/uuid"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

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
	PrivacyPolicyLink           string  `mapstructure:"privacy_policy_link" yaml:"privacy_policy_link"`                     // 新增隐私政策链接
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

// reids 配置
type RedisConfig struct {
	RedisConnString string `mapstructure:"redis_conn_string" yaml:"redis_conn_string"`
	RedisPassword   string `mapstructure:"redis_password" yaml:"redis_password"`
	RedisMasterName string `mapstructure:"redis_master_name" yaml:"redis_master_name"`
	SyncFrequency   int    `mapstructure:"sync_frequency" yaml:"sync_frequency"`
	Database        int    `mapstructure:"database" yaml:"database"` // 新增 Redis 数据库选择
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
	JwtSecret                     string   `mapstructure:"jwt_secret" yaml:"jwt_secret"`                                         // 新增 JWT Secret
	JwtExpiration                 int      `mapstructure:"jwt_expiration" yaml:"jwt_expiration"`                                 // 新增 JWT 过期时间 (秒)
	LoginFailuresLockoutEnabled   bool     `mapstructure:"login_failures_lockout_enabled" yaml:"login_failures_lockout_enabled"` // 新增登录失败锁定
	MaxLoginFailures              int      `mapstructure:"max_login_failures" yaml:"max_login_failures"`                         // 新增最大登录失败次数
	LockoutDuration               int      `mapstructure:"lockout_duration" yaml:"lockout_duration"`                             // 新增锁定持续时间 (分钟)
}

// 调试配置
type DebugConfig struct {
	DebugEnabled       bool `mapstructure:"debug_enabled" yaml:"debug_enabled"`
	DebugSQLEnabled    bool `mapstructure:"debug_sql_enabled" yaml:"debug_sql_enabled"`
	MemoryCacheEnabled bool `mapstructure:"memory_cache_enabled" yaml:"memory_cache_enabled"`
	EnableProfiling    bool `mapstructure:"enable_profiling" yaml:"enable_profiling"` // 新增性能分析开关
	ProfilingPort      int  `mapstructure:"profiling_port" yaml:"profiling_port"`     // 新增性能分析端口
}

// SMTP 配置
type SMTPConfig struct {
	SMTPServer    string `mapstructure:"smtp_server" yaml:"smtp_server"`
	SMTPPort      int    `mapstructure:"smtp_port" yaml:"smtp_port"`
	SMTPAccount   string `mapstructure:"smtp_account" yaml:"smtp_account"`
	SMTPFrom      string `mapstructure:"smtp_from" yaml:"smtp_from"`
	SMTPToken     string `mapstructure:"smtp_token" yaml:"smtp_token"`
	EnableTLS     bool   `mapstructure:"enable_tls" yaml:"enable_tls"`           // 新增启用 TLS
	TLSSkipVerify bool   `mapstructure:"tls_skip_verify" yaml:"tls_skip_verify"` // 新增跳过 TLS 验证
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
	GoogleClientId            string `mapstructure:"google_client_id" yaml:"google_client_id"`         // 新增 Google OAuth Client ID
	GoogleClientSecret        string `mapstructure:"google_client_secret" yaml:"google_client_secret"` // 新增 Google OAuth Client Secret
	GoogleRedirectURL         string `mapstructure:"google_redirect_url" yaml:"google_redirect_url"`   // 新增 Google OAuth Redirect URL
}

// 微信配置
type WeChatConfig struct {
	WeChatServerAddress         string `mapstructure:"wechat_server_address" yaml:"wechat_server_address"`
	WeChatServerToken           string `mapstructure:"wechat_server_token" yaml:"wechat_server_token"`
	WeChatAccountQRCodeImageURL string `mapstructure:"wechat_account_qr_code_image_url" yaml:"wechat_account_qr_code_image_url"`
	WeChatAppID                 string `mapstructure:"wechat_app_id" yaml:"wechat_app_id"`         // 新增微信 AppID
	WeChatAppSecret             string `mapstructure:"wechat_app_secret" yaml:"wechat_app_secret"` // 新增微信 AppSecret
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

// 验证码配置
type TurnstileConfig struct {
	TurnstileSiteKey   string `mapstructure:"turnstile_site_key" yaml:"turnstile_site_key"`
	TurnstileSecretKey string `mapstructure:"turnstile_secret_key" yaml:"turnstile_secret_key"`
}

// 配额配置
type QuotaConfig struct {
	QuotaForNewUser      int64 `mapstructure:"quota_for_new_user" yaml:"quota_for_new_user"`
	QuotaForInviter      int64 `mapstructure:"quota_for_inviter" yaml:"quota_for_inviter"`
	QuotaForInvitee      int64 `mapstructure:"quota_for_invitee" yaml:"quota_for_invitee"`
	QuotaRemindThreshold int64 `mapstructure:"quota_remind_threshold" yaml:"quota_remind_threshold"`
	PreConsumedQuota     int64 `mapstructure:"pre_consumed_quota" yaml:"pre_consumed_quota"`
	InviteCodeLength     int   `mapstructure:"invite_code_length" yaml:"invite_code_length"` // 新增邀请码长度
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

// 指标配置
type MetricConfig struct {
	EnableMetric               bool    `mapstructure:"enable_metric" yaml:"enable_metric"`
	MetricQueueSize            int     `mapstructure:"metric_queue_size" yaml:"metric_queue_size"`
	MetricSuccessRateThreshold float64 `mapstructure:"metric_success_rate_threshold" yaml:"metric_success_rate_threshold"`
	MetricSuccessChanSize      int     `mapstructure:"metric_success_chan_size" yaml:"metric_success_chan_size"`
	MetricFailChanSize         int     `mapstructure:"metric_fail_chan_size" yaml:"metric_fail_chan_size"`
	PrometheusEnabled          bool    `mapstructure:"prometheus_enabled" yaml:"prometheus_enabled"`               // 新增 Prometheus 指标开关
	PrometheusListenAddress    string  `mapstructure:"prometheus_listen_address" yaml:"prometheus_listen_address"` // 新增 Prometheus 监听地址
}

// 中继配置
type RelayConfig struct {
	RelayTimeout                   int                               `mapstructure:"relay_timeout" yaml:"relay_timeout"`
	RelayProxy                     string                            `mapstructure:"relay_proxy" yaml:"relay_proxy"`
	UserContentRequestProxy        string                            `mapstructure:"user_content_request_proxy" yaml:"user_content_request_proxy"`
	UserContentRequestTimeout      int                               `mapstructure:"user_content_request_timeout" yaml:"user_content_request_timeout"`
	ApiRequestTimeout              int                               `mapstructure:"api_request_timeout" yaml:"api_request_timeout"`                             // 新增 API 请求超时时间
	StreamRequestTimeout           int                               `mapstructure:"stream_request_timeout" yaml:"stream_request_timeout"`                       // 新增流式请求超时时间
	ModelMappingEnabled            bool                              `mapstructure:"model_mapping_enabled" yaml:"model_mapping_enabled"`                         // 新增模型映射开关
	ModelMapping                   map[string]string                 `mapstructure:"model_mapping" yaml:"model_mapping"`                                         // 新增模型映射配置
	ChannelMappingEnabled          bool                              `mapstructure:"channel_mapping_enabled" yaml:"channel_mapping_enabled"`                     // 新增渠道映射开关
	ChannelMapping                 map[string]string                 `mapstructure:"channel_mapping" yaml:"channel_mapping"`                                     // 新增渠道映射配置
	CompletionParamsMappingEnabled bool                              `mapstructure:"completion_params_mapping_enabled" yaml:"completion_params_mapping_enabled"` // 新增补全参数映射开关
	CompletionParamsMapping        map[string]map[string]interface{} `mapstructure:"completion_params_mapping" yaml:"completion_params_mapping"`                 // 新增补全参数映射
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

// 数据库配置
type DatabaseConfig struct {
	Type         string `mapstructure:"type" yaml:"type"`
	Host         string `mapstructure:"host" yaml:"host"`
	Port         int    `mapstructure:"port" yaml:"port"`
	Username     string `mapstructure:"username" yaml:"username"`
	Password     string `mapstructure:"password" yaml:"password"`
	Database     string `mapstructure:"database" yaml:"database"`
	Charset      string `mapstructure:"charset" yaml:"charset"`
	Options      string `mapstructure:"options" yaml:"options"` // 数据库连接参数
	IsMasterNode bool   `mapstructure:"is_master_node" yaml:"is_master_node"`
}

// 支付配置
type PaymentConfig struct {
	Enabled     bool   `mapstructure:"enabled" yaml:"enabled"`
	Type        string `mapstructure:"type" yaml:"type"` // 例如: alipay, stripe
	ApiKey      string `mapstructure:"api_key" yaml:"api_key"`
	SecretKey   string `mapstructure:"secret_key" yaml:"secret_key"`
	WebhookUrl  string `mapstructure:"webhook_url" yaml:"webhook_url"`
	ReturnUrl   string `mapstructure:"return_url" yaml:"return_url"`
	NotifyUrl   string `mapstructure:"notify_url" yaml:"notify_url"`
	ProductName string `mapstructure:"product_name" yaml:"product_name"`
}

// 测试配置
type TestConfig struct {
	TestPrompt string `mapstructure:"test_prompt" yaml:"test_prompt"`
}

// Config represents the overall application configuration.
type Config struct {
	SystemConfig         *SystemConfig         `mapstructure:"system" yaml:"system"`
	ServerConfig         *ServerConfig         `mapstructure:"server" yaml:"server"`
	SessionConfig        *SessionConfig        `mapstructure:"session" yaml:"session"`
	AuthenticationConfig *AuthenticationConfig `mapstructure:"authentication" yaml:"authentication"`
	DebugConfig          *DebugConfig          `mapstructure:"debug" yaml:"debug"`
	SMTPConfig           *SMTPConfig           `mapstructure:"smtp" yaml:"smtp"`
	OAuthConfig          *OAuthConfig          `mapstructure:"oauth" yaml:"oauth"`
	WeChatConfig         *WeChatConfig         `mapstructure:"wechat" yaml:"wechat"`
	MessageConfig        *MessageConfig        `mapstructure:"message" yaml:"message"`
	TurnstileConfig      *TurnstileConfig      `mapstructure:"turnstile" yaml:"turnstile"`
	QuotaConfig          *QuotaConfig          `mapstructure:"quota" yaml:"quota"`
	SyncConfig           *SyncConfig           `mapstructure:"sync" yaml:"sync"`
	RateLimitConfig      *RateLimitConfig      `mapstructure:"rate_limit" yaml:"rate_limit"`
	MetricConfig         *MetricConfig         `mapstructure:"metric" yaml:"metric"`
	RelayConfig          *RelayConfig          `mapstructure:"relay" yaml:"relay"`
	GeneralConfig        *GeneralConfig        `mapstructure:"general" yaml:"general"`
	TestConfig           *TestConfig           `mapstructure:"test" yaml:"test"`
	RedisConfig          *RedisConfig          `mapstructure:",squash"` // Redis 配置提升到顶层
	LogConfig            *LogConfig            `mapstructure:"log" yaml:"log"`
	DatabaseConfig       *DatabaseConfig       `mapstructure:"database" yaml:"database"`
	PaymentConfig        *PaymentConfig        `mapstructure:"payment" yaml:"payment"`
}

// Cfg is the global configuration instance.
var Cfg = new(Config)

// InitConfig initializes the configuration from various sources.
func InitConfig() error {
	// 1. Load from configuration file (config.yaml by default)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")   // Look for config in the working directory
	viper.AddConfigPath("/etc/app")   // Optionally add other paths
	viper.AddConfigPath("$HOME/.app") // Optionally add user's home directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error occurred
			return err
		}
		// Config file not found; proceed to load from other sources
	} else {
		println("Using config file:", viper.ConfigFileUsed())
	}

	// 2. Load from environment variables (can override config file)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Replace . with _ in env vars

	// 3. Load from command-line arguments (can override env vars and config file)
	if err := loadFromCommandLine(); err != nil {
		return err
	}

	// 4. Set default values (applied if not set in file, env, or CLI)
	setDefaultValues()

	// 5. Bind the loaded configuration to the Cfg struct
	if err := viper.Unmarshal(Cfg); err != nil {
		return err
	}

	// 6. Post-processing and validation
	if err := validateConfig(); err != nil {
		return err
	}

	// Generate default session secret if not provided
	if Cfg.SessionConfig.SessionSecret == "" {
		Cfg.SessionConfig.SessionSecret = uuid.New().String()
	}

	return nil
}

// loadFromCommandLine loads configuration from command-line arguments using pflag.
func loadFromCommandLine() error {
	pflag.String("server.address", "", "Server address")
	pflag.Int("server.port", 0, "Server port")
	pflag.String("redis_conn_string", "", "Redis connection string")
	pflag.String("redis_password", "", "Redis password")
	pflag.Int64("quota.quota_for_new_user", 0, "Quota for new user")
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return err
	}
	return nil
}

// setDefaultValues sets default configuration values.
func setDefaultValues() {
	viper.SetDefault("system.system_name", "One API")
	viper.SetDefault("server.address", "http://localhost")
	viper.SetDefault("server.port", 3000)
	viper.SetDefault("server.mode", "release")
	viper.SetDefault("quota.quota_for_new_user", 0)
	viper.SetDefault("rate_limit.global_api_rate_limit_num", 480)
	viper.SetDefault("session.cookie_name", "session_id")
	viper.SetDefault("session.cookie_path", "/")
	viper.SetDefault("session.cookie_secure", true)
	viper.SetDefault("session.cookie_http_only", true)
	viper.SetDefault("session.max_age", 3600) // 1 hour
	viper.SetDefault("redis.database", 0)
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
	viper.SetDefault("general.time_zone", "Asia/Singapore")
	viper.SetDefault("general.default_language", "en")
	viper.SetDefault("rate_limit.rate_limit_key_expiration_duration", time.Minute*5)
	viper.SetDefault("relay.relay_timeout", 30)
	viper.SetDefault("relay.user_content_request_timeout", 10)
	viper.SetDefault("relay.api_request_timeout", 60)
	viper.SetDefault("relay.stream_request_timeout", 300)
}

// validateConfig performs custom configuration validation.
func validateConfig() error {
	if Cfg.ServerConfig.Port <= 0 || Cfg.ServerConfig.Port > 65535 {
		return NewConfigError("server.port", "must be between 1 and 65535")
	}
	if Cfg.SessionConfig.SessionSecret == "" && Cfg.ServerConfig.Mode == "release" {
		println("Warning: Session secret is not set. Auto-generated in debug mode.")
	}
	if Cfg.QuotaConfig.InviteCodeLength <= 0 {
		Cfg.QuotaConfig.InviteCodeLength = 6 // Set a reasonable default for invite code length
	}
	if Cfg.AuthenticationConfig.JwtExpiration <= 0 {
		Cfg.AuthenticationConfig.JwtExpiration = 3600 // Default JWT expiration to 1 hour
	}
	if Cfg.RateLimitConfig.RateLimitKeyExpirationDuration <= 0 {
		Cfg.RateLimitConfig.RateLimitKeyExpirationDuration = time.Minute * 5 // Default rate limit key expiration
	}
	// Add more validation rules as needed
	return nil
}

// ConfigError is a custom error type for configuration issues.
type ConfigError struct {
	Field   string
	Message string
}

func NewConfigError(field, message string) *ConfigError {
	return &ConfigError{Field: field, Message: message}
}

func (e *ConfigError) Error() string {
	return "config error: field '" + e.Field + "' " + e.Message
}

// GetConfig returns the global configuration instance.
func GetConfig() *Config {
	return Cfg
}

// GetSystemConfig returns the system configuration.
func GetSystemConfig() *SystemConfig {
	return Cfg.SystemConfig
}

// GetServerConfig returns the server configuration.
func GetServerConfig() *ServerConfig {
	return Cfg.ServerConfig
}

// GetSessionConfig returns the session configuration.
func GetSessionConfig() *SessionConfig {
	return Cfg.SessionConfig
}

// GetRedisConfig returns the Redis configuration.
func GetRedisConfig() *RedisConfig {
	return Cfg.RedisConfig
}

// GetAuthenticationConfig returns the authentication configuration.
func GetAuthenticationConfig() *AuthenticationConfig {
	return Cfg.AuthenticationConfig
}

// GetDebugConfig returns the debug configuration.
func GetDebugConfig() *DebugConfig {
	return Cfg.DebugConfig
}

// GetSMTPConfig returns the SMTP configuration.
func GetSMTPConfig() *SMTPConfig {
	return Cfg.SMTPConfig
}

// GetOAuthConfig returns the OAuth configuration.
func GetOAuthConfig() *OAuthConfig {
	return Cfg.OAuthConfig
}

// GetWeChatConfig returns the WeChat configuration.
func GetWeChatConfig() *WeChatConfig {
	return Cfg.WeChatConfig
}

// GetMessageConfig returns the message configuration.
func GetMessageConfig() *MessageConfig {
	return Cfg.MessageConfig
}

// GetTurnstileConfig returns the turnstile configuration.
func GetTurnstileConfig() *TurnstileConfig {
	return Cfg.TurnstileConfig
}

// GetQuotaConfig returns the quota configuration.
func GetQuotaConfig() *QuotaConfig {
	return Cfg.QuotaConfig
}

// GetSyncConfig returns the sync configuration.
func GetSyncConfig() *SyncConfig {
	return Cfg.SyncConfig
}

// GetRateLimitConfig returns the rate limit configuration.
func GetRateLimitConfig() *RateLimitConfig {
	return Cfg.RateLimitConfig
}

// GetMetricConfig returns the metric configuration.
func GetMetricConfig() *MetricConfig {
	return Cfg.MetricConfig
}

// GetRelayConfig returns the relay configuration.
func GetRelayConfig() *RelayConfig {
	return Cfg.RelayConfig
}

// GetGeneralConfig returns the general configuration.
func GetGeneralConfig() *GeneralConfig {
	return Cfg.GeneralConfig
}

// GetTestConfig returns the test configuration.
func GetTestConfig() *TestConfig {
	return Cfg.TestConfig
}

// GetLogConfig returns the log configuration.
func GetLogConfig() *LogConfig {
	return Cfg.LogConfig
}

// GetDatabaseConfig returns the database configuration.
func GetDatabaseConfig() *DatabaseConfig {
	return Cfg.DatabaseConfig
}

// GetPaymentConfig returns the payment configuration.
func GetPaymentConfig() *PaymentConfig {
	return Cfg.PaymentConfig
}

// GetViperInstance returns the underlying viper instance for more advanced operations.
func GetViperInstance() *viper.Viper {
	return viper.GetViper()
}

// SetConfigValue programmatically sets a configuration value.
// Note: This will only affect the in-memory configuration.
func SetConfigValue(key string, value interface{}) {
	viper.Set(key, value)
	// Optionally update the global Cfg struct if needed
	if strings.Contains(key, "system") {
		viper.UnmarshalKey("system", &Cfg.SystemConfig)
	} else if strings.Contains(key, "server") {
		viper.UnmarshalKey("server", &Cfg.ServerConfig)
	} // Add more conditions for other config sections
}

// WatchConfigFile reloads the configuration when the config file changes.
// This is useful for development environments.
func WatchConfigFile() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		println("Config file changed:", e.Name)
		if err := viper.Unmarshal(Cfg); err != nil {
			println("Error unmarshalling config:", err)
		}
	})
}

// LoadConfigFromReader allows loading configuration from an io.Reader.
// This can be useful for loading config from a string or other sources.
func LoadConfigFromReader(configType string, reader io.Reader) error {
	v := viper.New()
	v.SetConfigType(configType)
	if err := v.ReadConfig(reader); err != nil {
		return err
	}
	return v.Unmarshal(Cfg)
}

// LoadConfigFromString allows loading configuration from a string.
func LoadConfigFromString(configType, configString string) error {
	v := viper.New()
	v.SetConfigType(configType)
	if err := v.ReadConfig(strings.NewReader(configString)); err != nil {
		return err
	}
	return v.Unmarshal(Cfg)
}

// LoadConfigFromMap allows loading configuration from a map.
func LoadConfigFromMap(configMap map[string]interface{}) error {
	v := viper.New()
	for key, value := range configMap {
		v.Set(key, value)
	}
	return v.Unmarshal(Cfg)
}
