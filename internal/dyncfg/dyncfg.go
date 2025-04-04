package dyncfg

import (
	"strings"
)

var EmailDomainWhitelist []string

var Themes map[string]bool

// 开关
var PasswordLoginEnabled bool
var PasswordRegisterEnabled bool
var EmailVerificationEnabled bool
var GitHubOAuthEnabled bool
var OidcEnabled bool
var WeChatAuthEnabled bool
var TurnstileCheckEnabled bool
var RegisterEnabled bool
var EmailDomainRestrictionEnabled bool

// 数据
var SystemName string
var ServerAddress string
var Footer string
var Logo string
var HomePageContent string
var About string
var Notice string
var Theme string

var SMTPServer string
var SMTPPort int
var SMTPAccount string
var SMTPFrom string
var SMTPToken string

var GitHubClientId string
var GitHubClientSecret string

var LarkClientId string
var LarkClientSecret string

var OidcClientId string
var OidcClientSecret string
var OidcWellKnown string
var OidcAuthorizationEndpoint string
var OidcTokenEndpoint string
var OidcUserinfoEndpoint string

var WeChatServerAddress string
var WeChatServerToken string
var WeChatAccountQRCodeImageURL string

var MessagePusherAddress string
var MessagePusherToken string

var TurnstileSiteKey string
var TurnstileSecretKey string

var RootUserEmail string

func SetEmailDomainWhitelist(val string) {
	EmailDomainWhitelist = strings.Split(val, ",")
}
func SetValidThemes(val string) {
	Themes[val] = true
}

// validateConfig performs custom configuration validation.
func ValidThemes(val string) bool {
	if _, ok := Themes[val]; !ok {
		return false
	}
	return true
}
