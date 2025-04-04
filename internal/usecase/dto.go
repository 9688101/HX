package usecase

import (
	"github.com/9688101/HX/internal/repo"
)

type UserUseCase interface {
	// RegisterUser(ctx context.Context, req RegisterUserRequest) error
	// Login(ctx context.Context, username, password string) (*entity.User, error)
	// GetUserList(ctx context.Context, offset, limit int, order string) ([]*entity.User, error)
	// SearchUsers(ctx context.Context, keyword string) ([]*entity.User, error)
	// GetUser(ctx context.Context, id int, callerRole int) (*entity.User, error)
	// UpdateSelf(ctx context.Context, req UpdateSelfRequest, userID int) error
	// GetSelf(ctx context.Context, userID int) (*entity.User, error)
	// DeleteSelf(ctx context.Context, userID int) error
	// UpdateUser(ctx context.Context, req UpdateUserRequest, callerRole int) error
	// DeleteUser(ctx context.Context, id int, callerRole int) error
	// ManageUser(ctx context.Context, req ManageRequest, callerRole int) (*entity.User, error)
	// BindEmail(ctx context.Context, email, code string, userID int) error
	// GetAffCode(ctx context.Context, userID int) (string, error)
	// GenerateAccessToken(ctx context.Context, userID int) (string, error)
}

type userUseCase struct {
	repo repo.UserRepository
}

// NewUserUseCase 返回 UserUseCase 的实现
func NewUserUseCase(r repo.UserRepository) UserUseCase {
	return &userUseCase{
		repo: r,
	}
}

// 用户注册请求结构体
type RegisterUserRequest struct {
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	Email            string `json:"email"`             // 可选，若开启邮箱验证则必填
	VerificationCode string `json:"verification_code"` // 验证码，邮箱验证时必填
	AffCode          string `json:"aff_code"`          // 邀请码
}

// 用户登录请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UpdateSelfRequest 定义当前用户更新自己信息的请求结构体
type UpdateSelfRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password"`     // 可选，若为空则不更新密码
	DisplayName string `json:"display_name"` // 可选
	// 其他需要更新的字段
}

// UpdateUserRequest 定义管理员更新用户信息的请求 DTO
type UpdateUserRequest struct {
	Id          int    `json:"id" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password"` // 若为空或为 "$I_LOVE_U"，表示不更新密码
	DisplayName string `json:"display_name"`
	Role        int    `json:"role"` // 若需要更新用户角色
}

// ManageRequest 定义了管理员管理用户的请求结构体
type ManageRequest struct {
	name   string `json:"username" binding:"required"`
	Action string `json:"action" binding:"required"`
}

type PasswordResetRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
