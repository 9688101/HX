package config

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

// validateConfig performs custom configuration validation.
func validateConfig() error {
	if GlobalConfig.ServerConfig.Port <= 0 || GlobalConfig.ServerConfig.Port > 65535 {
		return NewConfigError("server.port", "must be between 1 and 65535")
	}
	// if GlobalConfig.RateLimitConfig.RateLimitKeyExpirationDuration <= 0 {
	// 	GlobalConfig.RateLimitConfig.RateLimitKeyExpirationDuration = time.Minute * 5 // Default rate limit key expiration
	// }
	// Add more validation rules as needed
	return nil
}
