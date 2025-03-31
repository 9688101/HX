package config

import (
	"strings"

	"github.com/spf13/viper"
)

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

// GetViperInstance returns the underlying viper instance for more advanced operations.
func GetViperInstance() *viper.Viper {
	return viper.GetViper()
}
