package v1

import (
	"fmt"
	"net/http"

	"github.com/9688101/HX/internal/entity"
	"github.com/gin-gonic/gin"
)

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/9688101/HX/config"
// 	"github.com/9688101/HX/internal/entity"
// 	"github.com/9688101/HX/internal/usecase"
// 	"github.com/gin-gonic/gin"
// )

// type MiscController struct {
// 	cfgUsecase *usecase.ConfigUsecase
// }

// func NewMiscController(cfgUsecase *usecase.ConfigUsecase) *MiscController {
// 	return &MiscController{
// 		cfgUsecase: cfgUsecase,
// 	}
// }

func RelayNotFound(c *gin.Context) {
	err := entity.Error{
		Message: fmt.Sprintf("Invalid URL (%s %s)", c.Request.Method, c.Request.URL.Path),
		Type:    "invalid_request_error",
		Param:   "",
		Code:    "",
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": err,
	})
}

// // 获取系统状态
// func (mc MiscController) GetStatus(c *gin.Context) {
// 	sysCfg, _ := mc.cfgUsecase.GetConfig("system")
// 	wxCfg, _ := mc.cfgUsecase.GetConfig("wechat")
// 	msgCfg, _ := mc.cfgUsecase.GetConfig("message")
// 	authCfg, _ := mc.cfgUsecase.GetConfig("authentication")
// 	TurnstileCfg, _ := mc.cfgUsecase.GetConfig("turnstile")
// 	cfg, _ := mc.cfgUsecase.GetConfig("system")
// 	OAuthCfg, _ := mc.cfgUsecase.GetConfig("oauth")

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data": gin.H{
// 			// "version":                     pkg.Version,
// 			// "start_time":                  pkg.StartTime,
// 			"email_verification":          cfg.AuthenticationConfig.EmailVerificationEnabled,
// 			"github_oauth":                cfg.OAuthConfig.GitHubOAuthEnabled,
// 			"github_client_id":            cfg.OAuthConfig.GitHubClientId,
// 			"lark_client_id":              cfg.OAuthConfig.LarkClientId,
// 			"system_name":                 sysCfg.SystemName,
// 			"logo":                        cfg.SystemConfig.Logo,
// 			"footer_html":                 cfg.SystemConfig.FooterHTML,
// 			"wechat_qrcode":               cfg.WeChatConfig.WeChatAccountQRCodeImageURL,
// 			"wechat_login":                cfg.WeChatConfig.WeChatAuthEnabled,
// 			"server_address":              cfg.ServerConfig.ServerAddress,
// 			"turnstile_check":             cfg.AuthenticationConfig.TurnstileCheckEnabled,
// 			"turnstile_site_key":          cfg.TurnstileConfig.TurnstileSiteKey,
// 			"oidc":                        cfg.AuthenticationConfig.OidcEnabled,
// 			"oidc_client_id":              cfg.OAuthConfig.OidcClientId,
// 			"oidc_well_known":             cfg.OAuthConfig.OidcWellKnown,
// 			"oidc_authorization_endpoint": cfg.OAuthConfig.OidcAuthorizationEndpoint,
// 			"oidc_token_endpoint":         cfg.OAuthConfig.OidcTokenEndpoint,
// 			"oidc_userinfo_endpoint":      cfg.OAuthConfig.OidcUserinfoEndpoint,
// 		},
// 	})
// 	return
// }

// func (mc MiscController) GetNotice(c *gin.Context) {
// 	config.OptionMapRWMutex.RLock()
// 	defer config.OptionMapRWMutex.RUnlock()
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    config.OptionMap["Notice"],
// 	})
// 	return
// }

// func (mc MiscController) GetAbout(c *gin.Context) {
// 	config.OptionMapRWMutex.RLock()
// 	defer config.OptionMapRWMutex.RUnlock()
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    config.OptionMap["About"],
// 	})
// 	return
// }

// func (mc MiscController) GetHomePageContent(c *gin.Context) {
// 	config.OptionMapRWMutex.RLock()
// 	defer config.OptionMapRWMutex.RUnlock()
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    config.OptionMap["HomePageContent"],
// 	})
// 	return
// }

// func (mc MiscController) SendEmailVerification(c *gin.Context) {
// 	email := c.Query("email")
// 	if err := common.Validate.Var(email, "required,email"); err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": i18n.Translate(c, "invalid_parameter"),
// 		})
// 		return
// 	}
// 	if config.EmailDomainRestrictionEnabled {
// 		allowed := false
// 		for _, domain := range config.EmailDomainWhitelist {
// 			if strings.HasSuffix(email, "@"+domain) {
// 				allowed = true
// 				break
// 			}
// 		}
// 		if !allowed {
// 			c.JSON(http.StatusOK, gin.H{
// 				"success": false,
// 				"message": "管理员启用了邮箱域名白名单，您的邮箱地址的域名不在白名单中",
// 			})
// 			return
// 		}
// 	}
// 	if model.IsEmailAlreadyTaken(email) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": "邮箱地址已被占用",
// 		})
// 		return
// 	}
// 	code := pkg.GenerateVerificationCode(6)
// 	pkg.RegisterVerificationCodeWithKey(email, code, pkg.EmailVerificationPurpose)
// 	subject := fmt.Sprintf("%s 邮箱验证邮件", config.SystemName)
// 	content := message.EmailTemplate(
// 		subject,
// 		fmt.Sprintf(`
// 			<p>您好！</p>
// 			<p>您正在进行 %s 邮箱验证。</p>
// 			<p>您的验证码为：</p>
// 			<p style="font-size: 24px; font-weight: bold; color: #333; background-color: #f8f8f8; padding: 10px; text-align: center; border-radius: 4px;">%s</p>
// 			<p style="color: #666;">验证码 %d 分钟内有效，如果不是本人操作，请忽略。</p>
// 		`, config.SystemName, code, pkg.VerificationValidMinutes),
// 	)
// 	err := message.SendEmail(subject, email, content)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 	})
// 	return
// }

// func (mc MiscController) SendPasswordResetEmail(c *gin.Context) {
// 	email := c.Query("email")
// 	if err := pkg.Validate.Var(email, "required,email"); err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": i18n.Translate(c, "invalid_parameter"),
// 		})
// 		return
// 	}
// 	if !model.IsEmailAlreadyTaken(email) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": "该邮箱地址未注册",
// 		})
// 		return
// 	}
// 	code := pkg.GenerateVerificationCode(0)
// 	pkg.RegisterVerificationCodeWithKey(email, code, pkg.PasswordResetPurpose)
// 	link := fmt.Sprintf("%s/user/reset?email=%s&token=%s", config.ServerAddress, email, code)
// 	subject := fmt.Sprintf("%s 密码重置", config.SystemName)
// 	content := message.EmailTemplate(
// 		subject,
// 		fmt.Sprintf(`
// 			<p>您好！</p>
// 			<p>您正在进行 %s 密码重置。</p>
// 			<p>请点击下面的按钮进行密码重置：</p>
// 			<p style="text-align: center; margin: 30px 0;">
// 				<a href="%s" style="background-color: #007bff; color: white; padding: 12px 24px; text-decoration: none; border-radius: 4px; display: inline-block;">重置密码</a>
// 			</p>
// 			<p style="color: #666;">如果按钮无法点击，请复制以下链接到浏览器中打开：</p>
// 			<p style="background-color: #f8f8f8; padding: 10px; border-radius: 4px; word-break: break-all;">%s</p>
// 			<p style="color: #666;">重置链接 %d 分钟内有效，如果不是本人操作，请忽略。</p>
// 		`, config.SystemName, link, link, common.VerificationValidMinutes),
// 	)
// 	err := message.SendEmail(subject, email, content)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": fmt.Sprintf("%s%s", i18n.Translate(c, "send_email_failed"), err.Error()),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 	})
// 	return
// }

// type (mc MiscController)PasswordResetRequest struct {
// 	Email string `json:"email"`
// 	Token string `json:"token"`
// }

// func (mc MiscController)ResetPassword(c *gin.Context) {
// 	var req PasswordResetRequest
// 	err := json.NewDecoder(c.Request.Body).Decode(&req)
// 	if req.Email == "" || req.Token == "" {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": i18n.Translate(c, "invalid_parameter"),
// 		})
// 		return
// 	}
// 	if !pkg.VerifyCodeWithKey(req.Email, req.Token, pkg.PasswordResetPurpose) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": "重置链接非法或已过期",
// 		})
// 		return
// 	}
// 	password := pkg.GenerateVerificationCode(12)
// 	err = model.ResetUserPasswordByEmail(req.Email, password)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	pkg.DeleteKey(req.Email, pkg.PasswordResetPurpose)
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    password,
// 	})
// 	return
// }
