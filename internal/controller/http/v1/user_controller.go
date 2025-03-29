package v1

import (
	"net/http"
	"strconv"

	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/internal/usecase"
	"github.com/9688101/HX/pkg"
	"github.com/9688101/HX/pkg/config"
	"github.com/9688101/HX/pkg/ctxkey"
	"github.com/9688101/HX/pkg/i18n"
	"github.com/gin-contrib/sessions"
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
		c.JSON(http.StatusOK, BaseResponse{
			success: false,
			message: "管理员关闭了新用户注册",
		})
		return
	}
	if !config.PasswordRegisterEnabled {
		c.JSON(http.StatusOK, BaseResponse{
			success: false,
			message: "管理员关闭了通过密码进行注册，请使用第三方账户验证的形式进行注册",
		})
		return
	}

	// 解析 JSON 请求体到 DTO
	var req usecase.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{
			success: false,
			message: i18n.Translate(c, "invalid_parameter"),
		})
		return
	}

	// 进行输入验证
	if err := pkg.Validate.Struct(&req); err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			success: false,
			message: i18n.Translate(c, "invalid_input"),
		})
		return
	}

	// 调用 UseCase 层处理注册逻辑
	if err := uc.usecase.RegisterUser(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			success: false,
			message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, BaseResponse{
		success: true,
		message: "",
	})
}

// LoginHandler 处理用户密码登录请求
func (uc *UserController) LoginHandler(c *gin.Context) {
	if !config.PasswordLoginEnabled {
		c.JSON(http.StatusOK, BaseResponse{
			message: "管理员关闭了密码登录",
			success: false,
		})
		return
	}

	var req usecase.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{
			message: i18n.Translate(c, "invalid_parameter"),
			success: false,
		})
		return
	}

	// 校验参数非空（binding 已确保必填，但可以额外检查）
	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusOK, BaseResponse{
			message: i18n.Translate(c, "invalid_parameter"),
			success: false,
		})
		return
	}

	// 调用 UseCase 层处理登录逻辑
	user, err := uc.usecase.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			message: err.Error(),
			success: false,
		})
		return
	}

	// 设置登录会话
	setupLogin(user, c)
}

// setupLogin 设置 session，并返回清理后的用户信息
func setupLogin(user *entity.User, c *gin.Context) {
	session := sessions.Default(c)
	session.Set("id", user.Id)
	session.Set("username", user.Username)
	session.Set("role", user.Role)
	session.Set("status", user.Status)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			message: "无法保存会话信息，请重试",
			success: false,
		})
		return
	}

	// 返回清理后的用户信息，去除敏感字段
	cleanUser := struct {
		Id          int    `json:"id"`
		Username    string `json:"username"`
		DisplayName string `json:"display_name"`
		Role        int    `json:"role"`
		Status      int    `json:"status"`
	}{
		Id:          user.Id,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Role:        user.Role,
		Status:      user.Status,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"success": true,
		"data":    cleanUser,
	})
}

// LogoutHandler 处理用户注销请求，清除 session
func (uc *UserController) LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"success": true,
	})
}

// GetAllUsersHandler 处理获取用户列表请求
func (uc *UserController) GetAllUsersHandler(c *gin.Context) {
	p, _ := strconv.Atoi(c.Query("p"))
	if p < 0 {
		p = 0
	}
	order := c.DefaultQuery("order", "")
	offset := p * config.ItemsPerPage
	limit := config.ItemsPerPage

	users, err := uc.usecase.GetUserList(c.Request.Context(), offset, limit, order)
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
		"data":    users,
	})
}

// SearchUsersHandler 处理用户搜索请求
func (uc *UserController) SearchUsersHandler(c *gin.Context) {
	keyword := c.Query("keyword")
	users, err := uc.usecase.SearchUsers(c.Request.Context(), keyword)
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
		"data":    users,
	})
}

// GetUserHandler 处理根据 ID 获取单个用户信息的请求
func (uc *UserController) GetUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	callerRole := c.GetInt(ctxkey.Role)
	user, err := uc.usecase.GetUser(c.Request.Context(), id, callerRole)
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
		"data":    user,
	})
}
