package usecase

import (
	"strconv"
	"strings"
	"time"

	"github.com/9688101/HX/internal/dynamicconfig"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/internal/repo"
	"github.com/9688101/HX/pkg/logger"
)

type OptionUsecase interface {
	GetOptions() ([]*entity.Option, error)
	UpdateOption(key, value string) error
	InitDynamicConfig()
	SyncOptions(frequency int)
}

type optionUsecase struct {
	repo repo.OptionRepository
}

func NewOptionUsecase(repo repo.OptionRepository) OptionUsecase {
	return &optionUsecase{repo: repo}
}

// GetOptions 返回所有非敏感配置选项（数据库中存储的动态配置）
func (uc *optionUsecase) GetOptions() ([]*entity.Option, error) {
	options, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	// 过滤掉包含 Token 或 Secret 后缀的敏感数据
	var filtered []*entity.Option
	for _, opt := range options {
		if strings.HasSuffix(opt.Key, "Token") || strings.HasSuffix(opt.Key, "Secret") {
			continue
		}
		filtered = append(filtered, opt)
	}
	return filtered, nil
}

// UpdateOption 更新数据库与内部动态配置
func (uc *optionUsecase) UpdateOption(key, value string) error {
	option, err := uc.repo.GetByKey(key)
	if err != nil {
		// 不存在则创建
		option = &entity.Option{Key: key, Value: value}
	} else {
		option.Value = value
	}
	if err := uc.repo.Save(option); err != nil {
		return err
	}
	return updateDynamicOption(key, value)
}

// updateDynamicOption 更新内部动态配置，同时在必要时同步到静态配置
func updateDynamicOption(key, value string) error {
	// 更新动态配置内部存储
	dynamicconfig.Set(key, value)
	// 若配置项为开关类型，同步更新静态配置
	// 其它配置项若有必要更新到静态配置，同样处理
	return nil
}

// InitDynamicConfig 初始化内部动态配置，将静态默认值写入内部存储，并加载数据库中已有的配置
func (uc *optionUsecase) InitDynamicConfig() {
	// 初始化默认配置值（仅写入动态配置内部存储，不对外暴露）
	dynamicconfig.Set("PasswordLoginEnabled", strconv.FormatBool(true))
	dynamicconfig.Set("PasswordRegisterEnabled", strconv.FormatBool(true))
	dynamicconfig.Set("EmailVerificationEnabled", strconv.FormatBool(false))
	dynamicconfig.Set("GitHubOAuthEnabled", strconv.FormatBool(false))
	dynamicconfig.Set("OidcEnabled", strconv.FormatBool(false))
	dynamicconfig.Set("WeChatAuthEnabled", strconv.FormatBool(false))
	dynamicconfig.Set("TurnstileCheckEnabled", strconv.FormatBool(false))
	dynamicconfig.Set("RegisterEnabled", strconv.FormatBool(true))
	dynamicconfig.Set("EmailDomainRestrictionEnabled", strconv.FormatBool(true))
	EmailDomainWhitelist := []string{
		"gmail.com",
		"163.com",
		"126.com",
		"qq.com",
		"outlook.com",
		"hotmail.com",
		"icloud.com",
		"yahoo.com",
		"foxmail.com",
	}
	dynamicconfig.Set("EmailDomainWhitelist", strings.Join(EmailDomainWhitelist, ","))

	// 加载数据库中已有配置，并覆盖默认值
	_, err := uc.repo.GetAll()
	if err != nil {
		logger.SysError("failed to load options: " + err.Error())
		return
	}
}

// SyncOptions 定时从数据库同步配置到内部动态配置
func (uc *optionUsecase) SyncOptions(frequency int) {
	go func() {
		for {
			time.Sleep(time.Duration(frequency) * time.Second)
			logger.SysLog("syncing options from database")
			options, err := uc.repo.GetAll()
			if err != nil {
				logger.SysError("failed to sync options: " + err.Error())
				continue
			}
			for _, option := range options {
				_ = updateDynamicOption(option.Key, option.Value)
			}
		}
	}()
}
