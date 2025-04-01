package config

import (
	"github.com/spf13/viper"
)

// setDefaultValues sets default configuration values.
func setDefaultValues() {
	viper.SetDefault("system.system_name", "One API")
	viper.SetDefault("server.address", "http://localhost")
	viper.SetDefault("server.port", 3000)
}
