package config

import "errors"

type ConfigService interface {
	GetOption(key string) (string, error)
	UpdateOption(key, value string) error
}

type configService struct {
	// 内部可以引用 GlobalConfig 或 Viper 实例
}

func NewConfigService() ConfigService {
	return &configService{}
}

func (s *configService) GetOption(key string) (string, error) {
	someValue, ok := GlobalConfig.OptionMap[key]
	if !ok {
		return "", errors.New("配置项不存在")
	}
	return someValue, nil
}

func (s *configService) UpdateOption(key, value string) error {
	// 调用上面实现的 UpdateConfigValue 或 updateOptionMap 方法
	return UpdateConfigValue(key, value)
}
