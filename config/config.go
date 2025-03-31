package config

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
	SyncConfig           *SyncConfig           `mapstructure:"sync" yaml:"sync"`
	RateLimitConfig      *RateLimitConfig      `mapstructure:"rate_limit" yaml:"rate_limit"`
	GeneralConfig        *GeneralConfig        `mapstructure:"general" yaml:"general"`
	TestConfig           *TestConfig           `mapstructure:"test" yaml:"test"`
	RedisConfig          *RedisConfig          `mapstructure:",squash"` // Redis 配置提升到顶层
	LogConfig            *LogConfig            `mapstructure:"log" yaml:"log"`
	DatabaseConfig       *DatabaseConfig       `mapstructure:"database" yaml:"database"`
}

// GetConfig returns the global configuration instance.
func GetConfig() *Config {
	return Cfg
}
