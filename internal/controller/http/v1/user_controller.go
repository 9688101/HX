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

// RegisterUserHandler 处理用户注册请求
func (uc *UserController) RegisterUserHandler(c *gin.Context) {
	// 检查系统配置是否允许注册
	if !config.RegisterEnabled {
		c.JSON(http.StatusOK, BaseResponse{
			Success: false,
			Message: "管理员关闭了新用户注册",
		})
		return
	}
	if !config.PasswordRegisterEnabled {
		c.JSON(http.StatusOK, BaseResponse{
			Success: false,
			Message: "管理员关闭了通过密码进行注册，请使用第三方账户验证的形式进行注册",
		})
		return
	}

	// 解析 JSON 请求体到 DTO
	var req usecase.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{
			Success: false,
			Message: i18n.Translate(c, "invalid_parameter"),
		})
		return
	}

	// 进行输入验证
	if err := pkg.Validate.Struct(&req); err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			Success: false,
			Message: i18n.Translate(c, "invalid_input"),
		})
		return
	}

	// 调用 UseCase 层处理注册逻辑
	if err := uc.usecase.RegisterUser(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, BaseResponse{
		Success: true,
		Message: "",
	})
}

// LoginHandler 处理用户密码登录请求
func (uc *UserController) LoginHandler(c *gin.Context) {
	if !config.PasswordLoginEnabled {
		c.JSON(http.StatusOK, BaseResponse{
			Message: "管理员关闭了密码登录",
			Success: false,
		})
		return
	}

	var req usecase.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{
			Message: i18n.Translate(c, "invalid_parameter"),
			Success: false,
		})
		return
	}

	// 校验参数非空（binding 已确保必填，但可以额外检查）
	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusOK, BaseResponse{
			Message: i18n.Translate(c, "invalid_parameter"),
			Success: false,
		})
		return
	}

	// 调用 UseCase 层处理登录逻辑
	user, err := uc.usecase.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			Message: err.Error(),
			Success: false,
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
			Message: "无法保存会话信息，请重试",
			Success: false,
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
	c.JSON(http.StatusOK, Response{
		BaseResponse: BaseResponse{
			Message: "",
			Success: true,
		},
		Data: cleanUser,
	})
}

// // LogoutHandler 处理用户注销请求，清除 session
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

// // GetAllUsersHandler 处理获取用户列表请求
func (uc *UserController) GetUserListHandler(c *gin.Context) {
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

// // SearchUsersHandler 处理用户搜索请求
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

// // GetUserHandler 处理根据 ID 获取单个用户信息的请求
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

// // UpdateSelfHandler 处理当前用户更新自己信息的请求
func (uc *UserController) UpdateSelfHandler(c *gin.Context) {
	var req usecase.UpdateSelfRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": i18n.Translate(c, "invalid_parameter"),
		})
		return
	}

	// 获取当前用户 ID，从上下文中（例如 ctxkey.Id）
	userID := c.GetInt(ctxkey.Id)
	if userID == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "无法获取当前用户信息",
		})
		return
	}

	if err := uc.usecase.UpdateSelf(c.Request.Context(), req, userID); err != nil {
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

// // GetSelfHandler 处理获取当前用户信息的请求
func (uc *UserController) GetSelfHandler(c *gin.Context) {
	userID := c.GetInt(ctxkey.Id)
	if userID == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "无法获取当前用户信息",
		})
		return
	}

	user, err := uc.usecase.GetSelf(c.Request.Context(), userID)
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

// // DeleteSelfHandler 处理当前用户自删除的请求
func (uc *UserController) DeleteSelfHandler(c *gin.Context) {
	userID := c.GetInt("id")
	if userID == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "无法获取当前用户信息",
		})
		return
	}
	if err := uc.usecase.DeleteSelf(c.Request.Context(), userID); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "账户已删除",
	})
}

// UpdateUserHandler 处理管理员更新用户的请求
func (uc *UserController) UpdateUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var req usecase.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": i18n.Translate(c, "invalid_parameter"),
		})
		return
	}

	callerRole := c.GetInt(ctxkey.Role)
	if callerRole == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "无法获取调用者权限",
		})
		return
	}

	if err := uc.usecase.UpdateUser(ctx, req, callerRole); err != nil {
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

// DeleteUserHandler 处理管理员删除用户的请求
func (uc *UserController) DeleteUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的用户 ID",
		})
		return
	}

	callerRole := c.GetInt(ctxkey.Role)
	if callerRole == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "无法获取调用者权限",
		})
		return
	}

	if err := uc.usecase.DeleteUser(ctx, id, callerRole); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "用户删除成功",
	})
}

// ManageUserHandler 处理管理员管理用户的请求
func (uc *UserController) ManageUserHandler(c *gin.Context) {
	var req usecase.ManageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": i18n.Translate(c, "invalid_parameter"),
		})
		return
	}

	callerRole := c.GetInt("role")
	if callerRole == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "无法获取调用者权限",
		})
		return
	}

	updatedUser, err := uc.usecase.ManageUser(c.Request.Context(), req, callerRole)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	// 如果删除操作成功，updatedUser 可能为 nil，此处按更新处理
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    updatedUser,
	})
}

// EmailBindHandler 处理邮箱绑定请求
func (uc *UserController) EmailBindHandler(c *gin.Context) {
	email := c.Query("email")
	code := c.Query("code")
	if email == "" || code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": i18n.Translate(c, "invalid_parameter"),
		})
		return
	}
	userID := c.GetInt("id")
	if userID == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "无法获取当前用户信息",
		})
		return
	}
	if err := uc.usecase.BindEmail(c.Request.Context(), email, code, userID); err != nil {
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

// GenerateAccessTokenHandler 处理生成访问令牌的请求
func (uc *UserController) GenerateAccessTokenHandler(c *gin.Context) {
	userID := c.GetInt(ctxkey.Id)
	if userID == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "无法获取当前用户信息",
		})
		return
	}
	token, err := uc.usecase.GenerateAccessToken(c.Request.Context(), userID)
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
		"data":    token,
	})
}

// GetAffCodeHandler 处理获取邀请码的请求
func (uc *UserController) GetAffCodeHandler(c *gin.Context) {
	userID := c.GetInt(ctxkey.Id)
	if userID == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "无法获取当前用户信息",
		})
		return
	}
	affCode, err := uc.usecase.GetAffCode(c.Request.Context(), userID)
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
		"data":    affCode,
	})
}
