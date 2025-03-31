package config

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

// 验证码配置
type TurnstileConfig struct {
	TurnstileSiteKey   string `mapstructure:"turnstile_site_key" yaml:"turnstile_site_key"`
	TurnstileSecretKey string `mapstructure:"turnstile_secret_key" yaml:"turnstile_secret_key"`
}
