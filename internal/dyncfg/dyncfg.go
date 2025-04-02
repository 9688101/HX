package dyncfg

import (
	"strings"
)

var EmailDomainWhitelist = []string{
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

var Themes = map[string]bool{
	"default": true,
	"berry":   true,
	"air":     true,
}

// 开关
var PasswordLoginEnabled = true
var PasswordRegisterEnabled = true
var EmailVerificationEnabled = false
var GitHubOAuthEnabled = false
var OidcEnabled = false
var WeChatAuthEnabled = false
var TurnstileCheckEnabled = false
var RegisterEnabled = true
var EmailDomainRestrictionEnabled = true

// 数据
var SystemName = "晖雄 AI"
var ServerAddress = "http://localhost:3000"
var Footer = ""
var Logo = ""
var HomePageContent = ""
var About = ""
var Notice = ""
var Theme = "default"

var SMTPServer = ""
var SMTPPort = 587
var SMTPAccount = ""
var SMTPFrom = ""
var SMTPToken = ""

var GitHubClientId = ""
var GitHubClientSecret = ""

var LarkClientId = ""
var LarkClientSecret = ""

var OidcClientId = ""
var OidcClientSecret = ""
var OidcWellKnown = ""
var OidcAuthorizationEndpoint = ""
var OidcTokenEndpoint = ""
var OidcUserinfoEndpoint = ""

var WeChatServerAddress = ""
var WeChatServerToken = ""
var WeChatAccountQRCodeImageURL = ""

var MessagePusherAddress = ""
var MessagePusherToken = ""

var TurnstileSiteKey = ""
var TurnstileSecretKey = ""

var RootUserEmail = ""

func SetEmailDomainWhitelist(key, val string) {
	EmailDomainWhitelist = strings.Split(val, val)
}
func SetValidThemes(key string, val bool) {
	Themes[key] = val
}

// validateConfig performs custom configuration validation.
func ValidThemes(key string) bool {
	if !Themes[key] {
		return false
	}
	return true
}
