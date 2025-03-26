package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/9688101/HX/common"
	"github.com/9688101/HX/common/config"
	"github.com/9688101/HX/common/ctxkey"
	"github.com/9688101/HX/common/i18n"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/internal/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UseCase *usecase.UserUseCase
}

func NewUserController(uc *usecase.UserUseCase) *UserController {
	return &UserController{
		UseCase: uc,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	ctx := c.Request.Context()
	// rep := newUserResponse()
	if !config.RegisterEnabled {
		c.JSON(http.StatusOK, BaseResponse{
			Message: "管理员关闭了新用户注册",
			Success: false,
		})
		return
	}
	if !config.PasswordRegisterEnabled {
		c.JSON(http.StatusOK, BaseResponse{
			Message: "管理员关闭了通过密码进行注册，请使用第三方账户验证的形式进行注册",
			Success: false,
		})
		return
	}
	req := new(RegisterRequest)
	err := json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			Success: false,
			Message: i18n.Translate(c, "invalid_parameter"),
		})
		return
	}
	if err := common.Validate.Struct(req); err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			Success: false,
			Message: i18n.Translate(c, "invalid_input"),
		})
		return
	}
	if config.EmailVerificationEnabled {
		if req.Email == "" || req.VerificationCode == "" {
			c.JSON(http.StatusOK, BaseResponse{
				Success: false,
				Message: "管理员开启了邮箱验证，请输入邮箱地址和验证码",
			})
			return
		}
		if !common.VerifyCodeWithKey(req.Email, req.VerificationCode, common.EmailVerificationPurpose) {
			c.JSON(http.StatusOK, BaseResponse{
				Success: false,
				Message: "验证码错误或已过期",
			})
			return
		}
	}

	affCode := req.AffCode // this code is the inviter's code, not the user's own code
	inviterId, _ := entity.GetUserIdByAffCode(affCode)
	cleanUser := entity.User{
		Username:    req.Username,
		Password:    req.Password,
		DisplayName: req.Username,
		InviterId:   inviterId,
	}
	if config.EmailVerificationEnabled {
		cleanUser.Email = req.Email
	}
	if err := uc.UseCase.Register(ctx, &cleanUser); err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, BaseResponse{
		Success: true,
		Message: "ok",
	})
	return
}
func (uc *UserController) Login(c *gin.Context) {
	if !config.PasswordLoginEnabled {
		c.JSON(http.StatusOK, BaseResponse{
			Message: "管理员关闭了密码登录",
			Success: false,
		})
		return
	}
	var loginRequest LoginRequest
	err := json.NewDecoder(c.Request.Body).Decode(&loginRequest)
	if err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			Message: i18n.Translate(c, "invalid_parameter"),
			Success: false,
		})
		return
	}
	username := loginRequest.Username
	password := loginRequest.Password
	if username == "" || password == "" {
		c.JSON(http.StatusOK, BaseResponse{
			Message: i18n.Translate(c, "invalid_parameter"),
			Success: false,
		})
		return
	}
	user := entity.User{
		Username: username,
		Password: password,
	}
	err = uc.UseCase.ValidateAndFill(&user)
	if err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			Message: err.Error(),
			Success: false,
		})
		return
	}
	SetupLogin(&user, c)
}
func (uc *UserController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			Message: err.Error(),
			Success: false,
		})
		return
	}
	c.JSON(http.StatusOK, BaseResponse{
		Message: "",
		Success: true,
	})
}

// setup session & cookies and then return user info
func (uc *UserController) SetupLogin(user *entity.User, c *gin.Context) {
	session := sessions.Default(c)
	session.Set("id", user.Id)
	session.Set("username", user.Username)
	session.Set("role", user.Role)
	session.Set("status", user.Status)
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "无法保存会话信息，请重试",
			"success": false,
		})
		return
	}
	cleanUser := entity.User{
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
func (uc *UserController) GetUserDashboard(c *gin.Context) {
	id := c.GetInt(ctxkey.Id)
	now := time.Now()
	startOfDay := now.Truncate(24*time.Hour).AddDate(0, 0, -6).Unix()
	endOfDay := now.Truncate(24 * time.Hour).Add(24*time.Hour - time.Second).Unix()

	dashboards, err := entity.SearchLogsByDayAndModel(id, int(startOfDay), int(endOfDay))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无法获取统计信息",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    dashboards,
	})
	return
}
func (uc *UserController) GetSelf(c *gin.Context) {
}
func (uc *UserController) UpdateSelf(c *gin.Context) {
}
func (uc *UserController) DeleteSelf(c *gin.Context) {
}
func (uc *UserController) GenerateAccessToken(c *gin.Context) {
}
func (uc *UserController) GetAffCode(c *gin.Context) {
}
func (uc *UserController) TopUp(c *gin.Context) {
}
func (uc *UserController) GetUserAvailableModels(c *gin.Context) {
}
func (uc *UserController) GetAllUsers(c *gin.Context) {
}
func (uc *UserController) SearchUsers(c *gin.Context) {
}
func (uc *UserController) CreateUser(c *gin.Context) {
}
func (uc *UserController) ManageUser(c *gin.Context) {
}
func (uc *UserController) UpdateUser(c *gin.Context) {
}
func (uc *UserController) DeleteUser(c *gin.Context) {
}
