package v1

import (
	"net/http"

	"github.com/9688101/HX/internal/usecase"
	"github.com/9688101/HX/pkg"
	"github.com/9688101/HX/pkg/config"
	"github.com/9688101/HX/pkg/i18n"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase usecase.UserUseCase
}

// 创建 UserController 实例
func NewUserController(uc usecase.UserUseCase) *UserController {
	return &UserController{
		usecase: uc,
	}
}

// RegisterUserHandler 处理用户注册请求
func (uc *UserController) RegisterUserHandler(c *gin.Context) {
	// 检查系统配置是否允许注册
	if !config.RegisterEnabled {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "管理员关闭了新用户注册",
		})
		return
	}
	if !config.PasswordRegisterEnabled {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "管理员关闭了通过密码进行注册，请使用第三方账户验证的形式进行注册",
		})
		return
	}

	// 解析 JSON 请求体到 DTO
	var req usecase.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": i18n.Translate(c, "invalid_parameter"),
		})
		return
	}

	// 进行输入验证
	if err := pkg.Validate.Struct(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": i18n.Translate(c, "invalid_input"),
		})
		return
	}

	// 调用 UseCase 层处理注册逻辑
	if err := uc.usecase.RegisterUser(c.Request.Context(), req); err != nil {
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
