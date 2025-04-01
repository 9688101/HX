package config

import (
	"io"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/google/uuid"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// InitConfig initializes the configuration from various sources.
func InitConfig() error {
	// 1. Load from configuration file (config.yaml by default)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // Look for config in the working directory
	viper.AddConfigPath("/")        // Optionally add other paths
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error occurred
			return err
		}
		// Config file not found; proceed to load from other sources
	} else {
		println("Using config file:", viper.ConfigFileUsed())
	}

	// 2. Load from environment variables (can override config file)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Replace . with _ in env vars

	// 3. Load from command-line arguments (can override env vars and config file)
	if err := loadFromCommandLine(); err != nil {
		return err
	}

	// 4. Set default values (applied if not set in file, env, or CLI)
	setDefaultValues()

	// 5. Bind the loaded configuration to the Cfg struct
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		return err
	}

	// 6. Post-processing and validation
	if err := validateConfig(); err != nil {
		return err
	}

	// Generate default session secret if not provided
	if GlobalConfig.AuthenticationConfig.SessionSecret == "" {
		GlobalConfig.AuthenticationConfig.SessionSecret = uuid.New().String()
	}
	// configData := viper.AllSettings()

	// fmt.Println("Viper Configuration Data (Deep Print):")
	// for key, value := range configData {
	// 	fmt.Printf("%s: ", key)
	// 	deepPrintValue(value, 0)
	// 	fmt.Println()
	// }
	return nil
}

// loadFromCommandLine loads configuration from command-line arguments using pflag.
func loadFromCommandLine() error {
	pflag.String("server.address", "", "Server address")
	pflag.Int("server.port", 0, "Server port")
	pflag.String("redis_conn_string", "", "Redis connection string")
	pflag.String("redis_password", "", "Redis password")
	pflag.Int64("quota.quota_for_new_user", 0, "Quota for new user")
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return err
	}
	return nil
}

// WatchConfigFile reloads the configuration when the config file changes.
// This is useful for development environments.
func WatchConfigFile() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		println("Config file changed:", e.Name)
		if err := viper.Unmarshal(GlobalConfig); err != nil {
			println("Error unmarshalling config:", err)
		}
	})
}

// LoadConfigFromReader allows loading configuration from an io.Reader.
// This can be useful for loading config from a string or other sources.
func LoadConfigFromReader(configType string, reader io.Reader) error {
	v := viper.New()
	v.SetConfigType(configType)
	if err := v.ReadConfig(reader); err != nil {
		return err
	}
	return v.Unmarshal(GlobalConfig)
}

// LoadConfigFromString allows loading configuration from a string.
func LoadConfigFromString(configType, configString string) error {
	v := viper.New()
	v.SetConfigType(configType)
	if err := v.ReadConfig(strings.NewReader(configString)); err != nil {
		return err
	}
	return v.Unmarshal(GlobalConfig)
}

// LoadConfigFromMap allows loading configuration from a map.
func LoadConfigFromMap(configMap map[string]interface{}) error {
	v := viper.New()
	for key, value := range configMap {
		v.Set(key, value)
	}
	return v.Unmarshal(GlobalConfig)
}
