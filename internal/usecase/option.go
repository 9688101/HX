package usecase

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/9688101/HX/internal/dyncfg"
// 	"github.com/9688101/HX/internal/entity"
// 	"github.com/9688101/HX/internal/repo"
// 	"github.com/9688101/HX/pkg/logger"
// )

// // 内部两个 map 存储动态配置，使用全局锁保证线程安全
// var (
// 	configSwitchMap = make(map[string]bool)
// 	configDataMap   = make(map[string]string)
// 	configMutex     sync.RWMutex
// )

// type OptionUsecase interface {
// 	// GetOptions 返回所有非敏感的配置（从数据库中读取）
// 	GetOptions() ([]*entity.Option, error)
// 	// UpdateOption 更新配置：先保存到数据库，再更新内部 map 和静态变量
// 	UpdateOption(key, value string) error
// 	// InitDynamicConfig 初始化动态配置：先用默认值初始化，再用数据库中的值覆盖
// 	InitDynamicConfig()
// 	// SyncOptions 定时从数据库同步配置到内部动态配置
// 	SyncOptions(frequency int)
// }

// type optionUsecase struct {
// 	repo repo.OptionRepository
// }

// func NewOptionUsecase(repo repo.OptionRepository) OptionUsecase {
// 	return &optionUsecase{repo: repo}
// }

// // GetOptions 返回所有非敏感配置选项（过滤掉以 Token 或 Secret 结尾的敏感数据）
// func (uc *optionUsecase) GetOptions() ([]*entity.Option, error) {
// 	var filtered []*entity.Option
// 	for k, v := range configDataMap {
// 		if strings.HasSuffix(v, "Token") || strings.HasSuffix(v, "Secret") {
// 			continue
// 		}
// 		filtered = append(filtered, &entity.Option{
// 			Key:   k,
// 			Value: v,
// 		})
// 	}
// 	for k, v := range configSwitchMap {
// 		filtered = append(filtered, &entity.Option{
// 			Key:   k,
// 			Value: strconv.FormatBool(v),
// 		})
// 	}
// 	return filtered, nil
// }

// // UpdateOption 更新数据库中的配置，然后安全地更新内部动态配置
// func (uc *optionUsecase) UpdateOption(key, value string) error {
// 	// 先更新数据库
// 	option, err := uc.repo.GetByKey(key)
// 	if err != nil {
// 		// 如果数据库中不存在，则创建新配置项
// 		option = &entity.Option{Key: key, Value: value}
// 	} else {
// 		option.Value = value
// 	}
// 	if err := uc.repo.Save(option); err != nil {
// 		return err
// 	}
// 	// 再更新内部存储与全局变量
// 	return updateDynamicOption(key, value)
// }

// // updateDynamicOption 安全地更新内部两个 map，并同步更新 dyncfg 中的全局变量
// func updateDynamicOption(key, value string) error {
// 	configMutex.Lock()
// 	defer configMutex.Unlock()

// 	// 若为开关变量（key 后缀 "Enabled"），转换为 bool 更新 map 与 dyncfg
// 	if strings.HasSuffix(key, "Enabled") {
// 		Value := strings.ToLower(value) == "true"
// 		configSwitchMap[key] = Value

// 		switch key {
// 		case "PasswordRegisterEnabled":
// 			dyncfg.PasswordRegisterEnabled = Value
// 		case "PasswordLoginEnabled":
// 			dyncfg.PasswordLoginEnabled = Value
// 		case "EmailVerificationEnabled":
// 			dyncfg.EmailVerificationEnabled = Value
// 		case "GitHubOAuthEnabled":
// 			dyncfg.GitHubOAuthEnabled = Value
// 		case "OidcEnabled":
// 			dyncfg.OidcEnabled = Value
// 		case "WeChatAuthEnabled":
// 			dyncfg.WeChatAuthEnabled = Value
// 		case "TurnstileCheckEnabled":
// 			dyncfg.TurnstileCheckEnabled = Value
// 		case "RegisterEnabled":
// 			dyncfg.RegisterEnabled = Value
// 		}
// 	} else {
// 		// 数据类型变量更新 map 与 dyncfg
// 		configDataMap[key] = value

// 		switch key {
// 		case "EmailDomainWhitelist":
// 			// 前端传入以逗号分隔的字符串
// 			dyncfg.EmailDomainWhitelist = strings.Split(value, ",")
// 		case "SystemName":
// 			dyncfg.SystemName = value
// 		case "ServerAddress":
// 			dyncfg.ServerAddress = value
// 		case "Footer":
// 			dyncfg.Footer = value
// 		case "Logo":
// 			dyncfg.Logo = value
// 		case "SMTPServer":
// 			dyncfg.SMTPServer = value
// 		case "SMTPPort":
// 			if port, err := strconv.Atoi(value); err == nil {
// 				dyncfg.SMTPPort = port
// 			}
// 		case "SMTPAccount":
// 			dyncfg.SMTPAccount = value
// 		case "SMTPFrom":
// 			dyncfg.SMTPFrom = value
// 		case "SMTPToken":
// 			dyncfg.SMTPToken = value
// 		case "GitHubClientId":
// 			dyncfg.GitHubClientId = value
// 		case "GitHubClientSecret":
// 			dyncfg.GitHubClientSecret = value
// 		case "LarkClientId":
// 			dyncfg.LarkClientId = value
// 		case "LarkClientSecret":
// 			dyncfg.LarkClientSecret = value
// 		case "OidcClientId":
// 			dyncfg.OidcClientId = value
// 		case "OidcClientSecret":
// 			dyncfg.OidcClientSecret = value
// 		case "OidcWellKnown":
// 			dyncfg.OidcWellKnown = value
// 		case "OidcAuthorizationEndpoint":
// 			dyncfg.OidcAuthorizationEndpoint = value
// 		case "OidcTokenEndpoint":
// 			dyncfg.OidcTokenEndpoint = value
// 		case "OidcUserinfoEndpoint":
// 			dyncfg.OidcUserinfoEndpoint = value
// 		case "WeChatServerAddress":
// 			dyncfg.WeChatServerAddress = value
// 		case "WeChatServerToken":
// 			dyncfg.WeChatServerToken = value
// 		case "WeChatAccountQRCodeImageURL":
// 			dyncfg.WeChatAccountQRCodeImageURL = value
// 		case "MessagePusherAddress":
// 			dyncfg.MessagePusherAddress = value
// 		case "MessagePusherToken":
// 			dyncfg.MessagePusherToken = value
// 		case "TurnstileSiteKey":
// 			dyncfg.TurnstileSiteKey = value
// 		case "TurnstileSecretKey":
// 			dyncfg.TurnstileSecretKey = value
// 		case "RootUserEmail":
// 			dyncfg.RootUserEmail = value
// 		}
// 	}
// 	return nil
// }

// // InitDynamicConfig 初始化内部 map：先设置默认值，再加载数据库中的配置覆盖默认值
// func (uc *optionUsecase) InitDynamicConfig() {
// 	// 设置默认值（文件中的初始化配置）
// 	configMutex.Lock()
// 	// 默认开关变量
// 	configSwitchMap["PasswordLoginEnabled"] = true
// 	configSwitchMap["PasswordRegisterEnabled"] = true
// 	configSwitchMap["EmailVerificationEnabled"] = false
// 	configSwitchMap["GitHubOAuthEnabled"] = false
// 	configSwitchMap["OidcEnabled"] = false
// 	configSwitchMap["WeChatAuthEnabled"] = false
// 	configSwitchMap["TurnstileCheckEnabled"] = false
// 	configSwitchMap["RegisterEnabled"] = true

// 	// 默认数据变量（均以字符串存储）
// 	configDataMap["EmailDomainWhitelist"] = ""
// 	configDataMap["SystemName"] = "晖雄 AI"
// 	configDataMap["ServerAddress"] = "http://localhost:3000"
// 	configDataMap["Footer"] = ""
// 	configDataMap["Logo"] = ""
// 	configDataMap["SMTPServer"] = ""
// 	configDataMap["SMTPPort"] = "587"
// 	configDataMap["SMTPAccount"] = ""
// 	configDataMap["SMTPFrom"] = ""
// 	configDataMap["SMTPToken"] = ""
// 	configDataMap["GitHubClientId"] = ""
// 	configDataMap["GitHubClientSecret"] = ""
// 	configDataMap["LarkClientId"] = ""
// 	configDataMap["LarkClientSecret"] = ""
// 	configDataMap["OidcClientId"] = ""
// 	configDataMap["OidcClientSecret"] = ""
// 	configDataMap["OidcWellKnown"] = ""
// 	configDataMap["OidcAuthorizationEndpoint"] = ""
// 	configDataMap["OidcTokenEndpoint"] = ""
// 	configDataMap["OidcUserinfoEndpoint"] = ""
// 	configDataMap["WeChatServerAddress"] = ""
// 	configDataMap["WeChatServerToken"] = ""
// 	configDataMap["WeChatAccountQRCodeImageURL"] = ""
// 	configDataMap["MessagePusherAddress"] = ""
// 	configDataMap["MessagePusherToken"] = ""
// 	configDataMap["TurnstileSiteKey"] = ""
// 	configDataMap["TurnstileSecretKey"] = ""
// 	configDataMap["RootUserEmail"] = ""
// 	configMutex.Unlock()

// 	// 将默认值同步更新到 dyncfg
// 	applyConfigToDyncfg()

// 	// 尝试从数据库加载已有配置并覆盖默认值
// 	options, err := uc.repo.GetAll()
// 	if err != nil {
// 		logger.SysError("failed to load options from db: " + err.Error())
// 		return
// 	}
// 	for _, opt := range options {
// 		_ = updateDynamicOption(opt.Key, opt.Value)
// 		fmt.Println(opt.Key, opt.Value)
// 	}
// }

// // applyConfigToDyncfg 将当前内部 map 中的配置同步到 dyncfg 包的全局变量
// func applyConfigToDyncfg() {
// 	configMutex.RLock()
// 	defer configMutex.RUnlock()

// 	// 同步开关类型
// 	dyncfg.PasswordLoginEnabled = configSwitchMap["PasswordLoginEnabled"]
// 	dyncfg.PasswordRegisterEnabled = configSwitchMap["PasswordRegisterEnabled"]
// 	dyncfg.EmailVerificationEnabled = configSwitchMap["EmailVerificationEnabled"]
// 	dyncfg.GitHubOAuthEnabled = configSwitchMap["GitHubOAuthEnabled"]
// 	dyncfg.OidcEnabled = configSwitchMap["OidcEnabled"]
// 	dyncfg.WeChatAuthEnabled = configSwitchMap["WeChatAuthEnabled"]
// 	dyncfg.TurnstileCheckEnabled = configSwitchMap["TurnstileCheckEnabled"]
// 	dyncfg.RegisterEnabled = configSwitchMap["RegisterEnabled"]

// 	// 同步数据类型
// 	dyncfg.EmailDomainWhitelist = strings.Split(configDataMap["EmailDomainWhitelist"], ",")
// 	dyncfg.SystemName = configDataMap["SystemName"]
// 	dyncfg.ServerAddress = configDataMap["ServerAddress"]
// 	dyncfg.Footer = configDataMap["Footer"]
// 	dyncfg.Logo = configDataMap["Logo"]
// 	dyncfg.SMTPServer = configDataMap["SMTPServer"]
// 	if port, err := strconv.Atoi(configDataMap["SMTPPort"]); err == nil {
// 		dyncfg.SMTPPort = port
// 	}
// 	dyncfg.SMTPAccount = configDataMap["SMTPAccount"]
// 	dyncfg.SMTPFrom = configDataMap["SMTPFrom"]
// 	dyncfg.SMTPToken = configDataMap["SMTPToken"]
// 	dyncfg.GitHubClientId = configDataMap["GitHubClientId"]
// 	dyncfg.GitHubClientSecret = configDataMap["GitHubClientSecret"]
// 	dyncfg.LarkClientId = configDataMap["LarkClientId"]
// 	dyncfg.LarkClientSecret = configDataMap["LarkClientSecret"]
// 	dyncfg.OidcClientId = configDataMap["OidcClientId"]
// 	dyncfg.OidcClientSecret = configDataMap["OidcClientSecret"]
// 	dyncfg.OidcWellKnown = configDataMap["OidcWellKnown"]
// 	dyncfg.OidcAuthorizationEndpoint = configDataMap["OidcAuthorizationEndpoint"]
// 	dyncfg.OidcTokenEndpoint = configDataMap["OidcTokenEndpoint"]
// 	dyncfg.OidcUserinfoEndpoint = configDataMap["OidcUserinfoEndpoint"]
// 	dyncfg.WeChatServerAddress = configDataMap["WeChatServerAddress"]
// 	dyncfg.WeChatServerToken = configDataMap["WeChatServerToken"]
// 	dyncfg.WeChatAccountQRCodeImageURL = configDataMap["WeChatAccountQRCodeImageURL"]
// 	dyncfg.MessagePusherAddress = configDataMap["MessagePusherAddress"]
// 	dyncfg.MessagePusherToken = configDataMap["MessagePusherToken"]
// 	dyncfg.TurnstileSiteKey = configDataMap["TurnstileSiteKey"]
// 	dyncfg.TurnstileSecretKey = configDataMap["TurnstileSecretKey"]
// 	dyncfg.RootUserEmail = configDataMap["RootUserEmail"]
// }

// // SyncOptions 定时从数据库同步配置到内部动态配置，安全更新两个 map 以及 dyncfg 中的全局变量
// func (uc *optionUsecase) SyncOptions(frequency int) {
// 	for {
// 		time.Sleep(time.Duration(frequency) * time.Second)
// 		logger.SysLog("syncing options from database")
// 		options, err := uc.repo.GetAll()
// 		if err != nil {
// 			logger.SysError("failed to sync options: " + err.Error())
// 			continue
// 		}
// 		for _, option := range options {
// 			_ = updateDynamicOption(option.Key, option.Value)
// 		}
// 	}
// }
