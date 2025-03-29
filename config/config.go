package config

import (
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type Config struct {
	SystemConfig         *SystemConfig         `mapstructure:"system" yaml:"system"`
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
	RedisConfig          *RedisConfig          `mapstructure:",squash"`
}

var Cfg = new(Config)

// 初始化配置
func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	// 加载配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 使用默认值，如果未在配置文件中设置
	viper.SetDefault("system.system_name", "One API")
	viper.SetDefault("server.address", "http://localhost:3000")
	viper.SetDefault("quota.quota_for_new_user", 0)
	viper.SetDefault("rate_limit.global_api_rate_limit_num", 480)

	// 将配置绑定到结构体中
	if err := viper.Unmarshal(Cfg); err != nil {
		return err
	}

	// 如果没有设置 sessionSecret, 生成一个默认值
	if Cfg.SessionConfig.SessionSecret == "" {
		Cfg.SessionConfig.SessionSecret = uuid.New().String()
	}

	// 返回nil表示成功初始化
	return nil
}

// 获取配置
func GetConfig() *Config {
	return Cfg
}
