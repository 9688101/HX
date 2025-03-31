package config

import "time"

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
	if Cfg.ServerConfig.Port <= 0 || Cfg.ServerConfig.Port > 65535 {
		return NewConfigError("server.port", "must be between 1 and 65535")
	}
	if Cfg.SessionConfig.SessionSecret == "" && Cfg.ServerConfig.Mode == "release" {
		println("Warning: Session secret is not set. Auto-generated in debug mode.")
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
