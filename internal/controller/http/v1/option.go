package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/9688101/HX/internal/dyncfg"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/internal/usecase"
	"github.com/9688101/HX/pkg/i18n"
	"github.com/gin-gonic/gin"
)

type OptionController struct {
	ou usecase.OptionUsecase
}

func NewOptionController(optionUsecase usecase.OptionUsecase) *OptionController {
	return &OptionController{ou: optionUsecase}
}

// GetOptions 获取所有非敏感配置项
func (ctrl *OptionController) GetOptions(c *gin.Context) {
	options, err := ctrl.ou.GetOptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	var resp []entity.Option
	for _, opt := range options {
		resp = append(resp, *opt)
		fmt.Println(opt.Key, opt.Value)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    resp,
	})
}

// UpdateOption 更新配置项（含参数校验）
func (ctrl *OptionController) UpdateOption(c *gin.Context) {
	var option entity.Option
	if err := json.NewDecoder(c.Request.Body).Decode(&option); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": i18n.Translate(c, "invalid_parameter"),
		})
		return
	}
	// 业务验证
	switch option.Key {
	case "Theme":
		if !dyncfg.ValidThemes(option.Value) {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "无效的主题",
			})
			return
		}
	case "GitHubOAuthEnabled":
		if option.Value == "true" && dyncfg.GitHubClientId == "" {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "无法启用 GitHub OAuth，请先填入 GitHub Client Id 以及 GitHub Client Secret！",
			})
			return
		}
	case "EmailDomainRestrictionEnabled":
		if option.Value == "true" && len(dyncfg.EmailDomainWhitelist) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "无法启用邮箱域名限制，请先填入限制的邮箱域名！",
			})
			return
		}
	case "WeChatAuthEnabled":
		if option.Value == "true" && dyncfg.WeChatServerAddress == "" {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "无法启用微信登录，请先填入微信登录相关配置信息！",
			})
			return
		}
	case "TurnstileCheckEnabled":
		if option.Value == "true" && dyncfg.TurnstileSiteKey == "" {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "无法启用 Turnstile 校验，请先填入 Turnstile 校验相关配置信息！",
			})
			return
		}
	}
	if err := ctrl.ou.UpdateOption(option.Key, option.Value); err != nil {
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
}
