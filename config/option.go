package config

// import (
// 	"errors"
// 	"strconv"
// 	"strings"

// 	"sync"

// 	"github.com/spf13/viper"
// 	// 其他依赖
// )

// // OptionMap 用于动态配置展示（key->string）
// var OptionMap map[string]string

// // OptionMapMutex 保护 OptionMap 的读写
// var OptionMapMutex sync.RWMutex

// // UpdateConfigValue 根据 key 更新 GlobalConfig 和 OptionMap（这里只处理部分示例配置）
// func UpdateConfigValue(key, value string) error {
// 	// 这里根据 key 对应的配置类型进行转换，并更新 GlobalConfig
// 	switch key {
// 	case "systemName":
// 		GlobalConfig.SystemConfig.SystemName = value
// 	case "serverAddress":
// 		GlobalConfig.ServerConfig.Address = value
// 	case "passwordLoginEnabled":
// 		boolVal, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return errors.New("passwordLoginEnabled 转换错误")
// 		}
// 		GlobalConfig.AuthenticationConfig.PasswordLoginEnabled = boolVal
// 	case "passwordRegisterEnabled":
// 		boolVal, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return errors.New("passwordRegisterEnabled 转换错误")
// 		}
// 		GlobalConfig.AuthenticationConfig.PasswordRegisterEnabled = boolVal
// 	case "emailVerificationEnabled":
// 		boolVal, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return errors.New("emailVerificationEnabled 转换错误")
// 		}
// 		GlobalConfig.AuthenticationConfig.EmailVerificationEnabled = boolVal
// 	case "gitHubOAuthEnabled":
// 		boolVal, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return errors.New("gitHubOAuthEnabled 转换错误")
// 		}
// 		GlobalConfig.AuthenticationConfig.GitHubOAuthEnabled = boolVal
// 	case "oidcEnabled":
// 		boolVal, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return errors.New("oidcEnabled 转换错误")
// 		}
// 		GlobalConfig.OAuthConfig.OidcEnabled = boolVal
// 	case "weChatAuthEnabled":
// 		boolVal, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return errors.New("weChatAuthEnabled 转换错误")
// 		}
// 		GlobalConfig.WeChatConfig.WeChatAuthEnabled = boolVal
// 	case "turnstileCheckEnabled":
// 		boolVal, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return errors.New("turnstileCheckEnabled 转换错误")
// 		}
// 		GlobalConfig.TurnstileConfig.TurnstileCheckEnabled = boolVal
// 	case "registerEnabled":
// 		boolVal, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return errors.New("registerEnabled 转换错误")
// 		}
// 		GlobalConfig.AuthenticationConfig.RegisterEnabled = boolVal
// 	case "emailDomainRestrictionEnabled":
// 		boolVal, err := strconv.ParseBool(value)
// 		if err != nil {
// 			return errors.New("emailDomainRestrictionEnabled 转换错误")
// 		}
// 		GlobalConfig.AuthenticationConfig.EmailDomainRestrictionEnabled = boolVal
// 	case "theme":
// 		// 校验主题是否有效
// 		if !GlobalConfig.ValidThemes[value] {
// 			return errors.New("无效的主题")
// 		}
// 		GlobalConfig.SystemConfig.Theme = value
// 	case "emailDomainWhitelist":
// 		// 多个邮箱域名以逗号分隔
// 		GlobalConfig.AuthenticationConfig.EmailDomainWhitelist = strings.Split(value, ",")
// 	// 其它配置项……
// 	default:
// 		// 如果 key 未处理，则可以选择返回错误或忽略
// 		return errors.New("不支持的配置项")
// 	}

// 	// 更新 OptionMap
// 	OptionMapMutex.Lock()
// 	defer OptionMapMutex.Unlock()
// 	OptionMap[key] = value

// 	// 同步更新 viper 内部配置
// 	viper.Set(key, value)
// 	// 如需要持久化，可调用 viper.WriteConfig()
// 	return nil
// }
