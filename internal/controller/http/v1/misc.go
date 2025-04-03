package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/9688101/HX/internal/dyncfg"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/pkg/email"
	"github.com/9688101/HX/pkg/i18n"
	"github.com/9688101/HX/pkg/valid"
	"github.com/9688101/HX/pkg/verif"

	"github.com/gin-gonic/gin"
)

type MiscController struct {
}

func NewMiscController() *MiscController {
	return &MiscController{}
}

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

// 获取系统状态
func (mc MiscController) GetStatus(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			// "version":                     pkg.Version,
			// "start_time":                  pkg.StartTime,
			"email_verification":          dyncfg.EmailVerificationEnabled,
			"github_oauth":                dyncfg.GitHubOAuthEnabled,
			"github_client_id":            dyncfg.GitHubClientId,
			"lark_client_id":              dyncfg.LarkClientId,
			"system_name":                 dyncfg.SystemName,
			"logo":                        dyncfg.Logo,
			"footer_html":                 dyncfg.Footer,
			"wechat_qrcode":               dyncfg.WeChatAccountQRCodeImageURL,
			"wechat_login":                dyncfg.WeChatAuthEnabled,
			"server_address":              dyncfg.ServerAddress,
			"turnstile_check":             dyncfg.TurnstileCheckEnabled,
			"turnstile_site_key":          dyncfg.TurnstileSiteKey,
			"oidc":                        dyncfg.OidcEnabled,
			"oidc_client_id":              dyncfg.OidcClientId,
			"oidc_well_known":             dyncfg.OidcWellKnown,
			"oidc_authorization_endpoint": dyncfg.OidcAuthorizationEndpoint,
			"oidc_token_endpoint":         dyncfg.OidcTokenEndpoint,
			"oidc_userinfo_endpoint":      dyncfg.OidcUserinfoEndpoint,
		},
	})
	return
}

func (mc MiscController) GetNotice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    dyncfg.Notice,
	})
	return
}

func (mc MiscController) GetAbout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    dyncfg.About,
	})
	return
}

func (mc MiscController) GetHomePageContent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    dyncfg.HomePageContent,
	})
	return
}

func (mc MiscController) SendEmailVerification(c *gin.Context) {
	mes := c.Query("email")
	if err := valid.ValidateVar(mes, "required,email"); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": i18n.Translate(c, "invalid_parameter"),
		})
		return
	}
	if dyncfg.EmailDomainRestrictionEnabled {
		allowed := false
		for _, domain := range dyncfg.EmailDomainWhitelist {
			if strings.HasSuffix(mes, "@"+domain) {
				allowed = true
				break
			}
		}
		if !allowed {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "管理员启用了邮箱域名白名单，您的邮箱地址的域名不在白名单中",
			})
			return
		}
	}
	// if model.IsEmailAlreadyTaken(email) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"success": false,
	// 		"message": "邮箱地址已被占用",
	// 	})
	// 	return
	// }
	code := verif.GenerateVerificationCode(6)
	verif.RegisterVerificationCodeWithKey(mes, code, verif.EmailVerificationPurpose)
	subject := fmt.Sprintf("%s 邮箱验证邮件", dyncfg.SystemName)
	content := email.EmailTemplate(
		subject,
		fmt.Sprintf(`
			<p>您好！</p>
			<p>您正在进行 %s 邮箱验证。</p>
			<p>您的验证码为：</p>
			<p style="font-size: 24px; font-weight: bold; color: #333; background-color: #f8f8f8; padding: 10px; text-align: center; border-radius: 4px;">%s</p>
			<p style="color: #666;">验证码 %d 分钟内有效，如果不是本人操作，请忽略。</p>
		`, dyncfg.SystemName, code, verif.VerificationValidMinutes),
	)
	err := email.SendEmail(subject, mes, content)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}

func (mc MiscController) SendPasswordResetEmail(c *gin.Context) {
	mes := c.Query("email")
	if err := valid.ValidateVar(mes, "required,email"); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": i18n.Translate(c, "invalid_parameter"),
		})
		return
	}
	// if !model.IsEmailAlreadyTaken(email) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"success": false,
	// 		"message": "该邮箱地址未注册",
	// 	})
	// 	return
	// }
	code := verif.GenerateVerificationCode(0)
	verif.RegisterVerificationCodeWithKey(mes, code, verif.PasswordResetPurpose)
	link := fmt.Sprintf("%s/user/reset?email=%s&token=%s", dyncfg.ServerAddress, mes, code)
	subject := fmt.Sprintf("%s 密码重置", dyncfg.SystemName)
	content := email.EmailTemplate(
		subject,
		fmt.Sprintf(`
			<p>您好！</p>
			<p>您正在进行 %s 密码重置。</p>
			<p>请点击下面的按钮进行密码重置：</p>
			<p style="text-align: center; margin: 30px 0;">
				<a href="%s" style="background-color: #007bff; color: white; padding: 12px 24px; text-decoration: none; border-radius: 4px; display: inline-block;">重置密码</a>
			</p>
			<p style="color: #666;">如果按钮无法点击，请复制以下链接到浏览器中打开：</p>
			<p style="background-color: #f8f8f8; padding: 10px; border-radius: 4px; word-break: break-all;">%s</p>
			<p style="color: #666;">重置链接 %d 分钟内有效，如果不是本人操作，请忽略。</p>
		`, dyncfg.SystemName, link, link, verif.VerificationValidMinutes),
	)
	err := email.SendEmail(subject, mes, content)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": fmt.Sprintf("%s%s", i18n.Translate(c, "send_email_failed"), err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}
