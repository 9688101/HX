package config

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

// GetSyncConfig returns the sync configuration.
func GetSyncConfig() *SyncConfig {
	return Cfg.SyncConfig
}

// GetRateLimitConfig returns the rate limit configuration.
func GetRateLimitConfig() *RateLimitConfig {
	return Cfg.RateLimitConfig
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
