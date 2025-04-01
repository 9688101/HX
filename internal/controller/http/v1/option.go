package v1

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/9688101/HX/config"
// 	"github.com/9688101/HX/pkg/helper"
// 	"github.com/9688101/HX/pkg/i18n"

// 	"github.com/gin-gonic/gin"
// )

// // OptionDTO 定义配置选项的结构体
// type OptionDTO struct {
// 	Key   string `json:"key"`
// 	Value string `json:"value"`
// }

// // GetOptionsHandler 返回所有非敏感配置选项
// func GetOptionsHandler(c *gin.Context) {
// 	config.OptionMapMutex.RLock()
// 	defer config.OptionMapMutex.RUnlock()

// 	var options []OptionDTO
// 	for k, v := range config.OptionMap {
// 		// 过滤掉包含 Token 或 Secret 的配置
// 		if strings.Contains(k, "Token") || strings.Contains(k, "Secret") {
// 			continue
// 		}
// 		options = append(options, OptionDTO{
// 			Key:   k,
// 			Value: helper.Interface2String(v),
// 		})
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "",
// 		"data":    options,
// 	})
// }

// // UpdateOptionRequest 定义更新配置的请求体
// type UpdateOptionRequest struct {
// 	Key   string `json:"key" binding:"required"`
// 	Value string `json:"value" binding:"required"`
// }

// // UpdateOptionHandler 处理配置更新请求
// func UpdateOptionHandler(c *gin.Context) {
// 	var req UpdateOptionRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"message": i18n.Translate(c, "invalid_parameter"),
// 		})
// 		return
// 	}

// 	// 针对特定配置项做额外校验
// 	switch req.Key {
// 	case "Theme":
// 		if !config.GlobalConfig.ValidThemes[req.Value] {
// 			c.JSON(http.StatusOK, gin.H{
// 				"success": false,
// 				"message": "无效的主题",
// 			})
// 			return
// 		}
// 	case "GitHubOAuthEnabled":
// 		if req.Value == "true" && config.GlobalConfig.GitHubClientId == "" {
// 			c.JSON(http.StatusOK, gin.H{
// 				"success": false,
// 				"message": "无法启用 GitHub OAuth，请先填写 GitHub Client Id 以及 GitHub Client Secret！",
// 			})
// 			return
// 		}
// 	case "EmailDomainRestrictionEnabled":
// 		if req.Value == "true" && len(config.GlobalConfig.EmailDomainWhitelist) == 0 {
// 			c.JSON(http.StatusOK, gin.H{
// 				"success": false,
// 				"message": "无法启用邮箱域名限制，请先填写限制的邮箱域名！",
// 			})
// 			return
// 		}
// 	case "WeChatAuthEnabled":
// 		if req.Value == "true" && config.GlobalConfig.WeChatServerAddress == "" {
// 			c.JSON(http.StatusOK, gin.H{
// 				"success": false,
// 				"message": "无法启用微信登录，请先填写微信登录相关配置信息！",
// 			})
// 			return
// 		}
// 	case "TurnstileCheckEnabled":
// 		if req.Value == "true" && config.GlobalConfig.TurnstileSiteKey == "" {
// 			c.JSON(http.StatusOK, gin.H{
// 				"success": false,
// 				"message": "无法启用 Turnstile 校验，请先填写 Turnstile 校验相关配置信息！",
// 			})
// 			return
// 		}
// 	}

// 	if err := config.UpdateConfigValue(req.Key, req.Value); err != nil {
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
// }
