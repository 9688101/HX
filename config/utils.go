package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// SetConfigValue programmatically sets a configuration value.
// Note: This will only affect the in-memory configuration.
func SetConfigValue(key string, value interface{}) {
	viper.Set(key, value)
	// Optionally update the global Cfg struct if needed
	if strings.Contains(key, "server") {
		viper.UnmarshalKey("server", &GlobalConfig.ServerConfig)
	}
	if strings.Contains(key, "redis") {
		viper.UnmarshalKey("redis", &GlobalConfig.RedisConfig)
	}
	if strings.Contains(key, "database") {
		viper.UnmarshalKey("database", &GlobalConfig.DatabaseConfig)
	}
	if strings.Contains(key, "debug") {
		viper.UnmarshalKey("debug", &GlobalConfig.DebugConfig)
	}
	if strings.Contains(key, "general") {
		viper.UnmarshalKey("general", &GlobalConfig.GeneralConfig)
	}
	if strings.Contains(key, "rate_limit") {
		viper.UnmarshalKey("rate_limit", &GlobalConfig.RateLimitConfig)
	}
}

// GetViperInstance returns the underlying viper instance for more advanced operations.
func GetViperInstance() *viper.Viper {
	return viper.GetViper()
}

// deepPrintValue 是一个递归函数，用于打印包含指针的任意类型的值
func deepPrintValue(val interface{}, depth int) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}

	v := reflect.ValueOf(val)

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			fmt.Printf("<nil>")
			return
		}
		deepPrintValue(v.Elem().Interface(), depth+1) // 递归调用，处理指针指向的值
		return
	}

	switch v.Kind() {
	case reflect.Struct:
		fmt.Println("{")
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := v.Type().Field(i)
			fmt.Printf("%s%s: ", indent+"  ", fieldType.Name)
			deepPrintValue(field.Interface(), depth+1)
			fmt.Println()
		}
		fmt.Printf("%s}", indent)
	case reflect.Map:
		fmt.Println("{")
		iter := v.MapRange()
		for iter.Next() {
			key := iter.Key()
			value := iter.Value()
			fmt.Printf("%s%v: ", indent+"  ", key.Interface())
			deepPrintValue(value.Interface(), depth+1)
			fmt.Println()
		}
		fmt.Printf("%s}", indent)
	default:
		fmt.Printf("%+v", val)
	}
}
