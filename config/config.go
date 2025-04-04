package config

// Cfg is the global configuration instance.
var GlobalConfig = new(Config)

// GetConfig returns the global configuration instance.
func GetConfig() *Config {
	return GlobalConfig
}

// GetRateLimitConfig returns the rate limit configuration.
func GetRateLimitConfig() *RateLimitConfig {
	return GlobalConfig.RateLimitConfig
}

// GetGeneralConfig returns the general configuration.
func GetGeneralConfig() *GeneralConfig {
	return GlobalConfig.GeneralConfig
}

// GetDatabaseConfig returns the database configuration.
func GetDatabaseConfig() *DatabaseConfig {
	return GlobalConfig.DatabaseConfig
}

// GetServerConfig returns the server configuration.
func GetServerConfig() *ServerConfig {
	return GlobalConfig.ServerConfig
}

// GetRedisConfig returns the Redis configuration.
func GetRedisConfig() *RedisConfig {
	return GlobalConfig.RedisConfig
}

func GetMailConfig() *MailConfig {
	return GlobalConfig.MailConfig
}

// GetSystemConfig returns the system configuration.
// func GetSystemConfig() *SystemConfig {
// 	return GlobalConfig.SystemConfig
// }

// GetAuthenticationConfig returns the authentication configuration.
// func GetAuthenticationConfig() *AuthenticationConfig {
// 	return GlobalConfig.AuthenticationConfig
// }

// GetSMTPConfig returns the SMTP configuration.
// func GetSMTPConfig() *SMTPConfig {
// 	return GlobalConfig.SMTPConfig
// 	}

// GetOAuthConfig returns the OAuth configuration.
// func GetOAuthConfig() *OAuthConfig {
// 	return GlobalConfig.OAuthConfig
// }

// GetWeChatConfig returns the WeChat configuration.
// func GetWeChatConfig() *WeChatConfig {
// 	return GlobalConfig.WeChatConfig
// }

// GetMessageConfig returns the message configuration.
// func GetMessageConfig() *MessageConfig {
// 	return GlobalConfig.MessageConfig
// }

// GetTurnstileConfig returns the turnstile configuration.
// func GetTurnstileConfig() *TurnstileConfig {
// 	return GlobalConfig.TurnstileConfig
// }
