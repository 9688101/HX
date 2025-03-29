package config

import "time"

// 系统配置
type SystemConfig struct {
	SystemName               string  `mapstructure:"system_name" yaml:"system_name"`
	ServerAddress            string  `mapstructure:"server_address" yaml:"server_address"`
	Footer                   string  `mapstructure:"footer" yaml:"footer"`
	Logo                     string  `mapstructure:"logo" yaml:"logo"`
	TopUpLink                string  `mapstructure:"top_up_link" yaml:"top_up_link"`
	ChatLink                 string  `mapstructure:"chat_link" yaml:"chat_link"`
	QuotaPerUnit             float64 `mapstructure:"quota_per_unit" yaml:"quota_per_unit"`
	DisplayInCurrencyEnabled bool    `mapstructure:"display_in_currency_enabled" yaml:"display_in_currency_enabled"`
	DisplayTokenStatEnabled  bool    `mapstructure:"display_token_stat_enabled" yaml:"display_token_stat_enabled"`
}

// 服务器配置
type ServerConfig struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

// 会话配置
type SessionConfig struct {
	SessionSecret string `mapstructure:"session_secret" yaml:"session_secret"`
}

// reids 配置
type RedisConfig struct {
	RedisConnString string `mapstructure:"REDIS_CONN_STRING" yaml:"REDIS_CONN_STRING"`
	RedisPassword   string `mapstructure:"REDIS_PASSWORD" yaml:"REDIS_PASSWORD"`
	RedisMasterName string `mapstructure:"REDIS_MASTER_NAME" yaml:"REDIS_MASTER_NAME"`
	SyncFrequency   string `mapstructure:"SYNC_FREQUENCY" yaml:"SYNC_FREQUENCY"`
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
}

// 调试配置
type DebugConfig struct {
	DebugEnabled       bool `mapstructure:"debug_enabled" yaml:"debug_enabled"`
	DebugSQLEnabled    bool `mapstructure:"debug_sql_enabled" yaml:"debug_sql_enabled"`
	MemoryCacheEnabled bool `mapstructure:"memory_cache_enabled" yaml:"memory_cache_enabled"`
}

// SMTP 配置
type SMTPConfig struct {
	SMTPServer  string `mapstructure:"smtp_server" yaml:"smtp_server"`
	SMTPPort    int    `mapstructure:"smtp_port" yaml:"smtp_port"`
	SMTPAccount string `mapstructure:"smtp_account" yaml:"smtp_account"`
	SMTPFrom    string `mapstructure:"smtp_from" yaml:"smtp_from"`
	SMTPToken   string `mapstructure:"smtp_token" yaml:"smtp_token"`
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

// 消息配置
type MessageConfig struct {
	MessagePusherAddress string `mapstructure:"message_pusher_address" yaml:"message_pusher_address"`
	MessagePusherToken   string `mapstructure:"message_pusher_token" yaml:"message_pusher_token"`
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
}

// 指标配置
type MetricConfig struct {
	EnableMetric               bool    `mapstructure:"enable_metric" yaml:"enable_metric"`
	MetricQueueSize            int     `mapstructure:"metric_queue_size" yaml:"metric_queue_size"`
	MetricSuccessRateThreshold float64 `mapstructure:"metric_success_rate_threshold" yaml:"metric_success_rate_threshold"`
	MetricSuccessChanSize      int     `mapstructure:"metric_success_chan_size" yaml:"metric_success_chan_size"`
	MetricFailChanSize         int     `mapstructure:"metric_fail_chan_size" yaml:"metric_fail_chan_size"`
}

// 中继配置
type RelayConfig struct {
	RelayTimeout              int    `mapstructure:"relay_timeout" yaml:"relay_timeout"`
	RelayProxy                string `mapstructure:"relay_proxy" yaml:"relay_proxy"`
	UserContentRequestProxy   string `mapstructure:"user_content_request_proxy" yaml:"user_content_request_proxy"`
	UserContentRequestTimeout int    `mapstructure:"user_content_request_timeout" yaml:"user_content_request_timeout"`
}

// 通用配置
type GeneralConfig struct {
	InitialRootToken       string `mapstructure:"initial_root_token" yaml:"initial_root_token"`
	InitialRootAccessToken string `mapstructure:"initial_root_access_token" yaml:"initial_root_access_token"`
	GeminiVersion          string `mapstructure:"gemini_version" yaml:"gemini_version"`
	OnlyOneLogFile         bool   `mapstructure:"only_one_log_file" yaml:"only_one_log_file"`
	RootUserEmail          string `mapstructure:"root_user_email" yaml:"root_user_email"`
	IsMasterNode           bool   `mapstructure:"is_master_node" yaml:"is_master_node"`
}

// 测试配置
type TestConfig struct {
	TestPrompt string `mapstructure:"test_prompt" yaml:"test_prompt"`
}
