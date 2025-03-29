package usecase

import (
	"context"
	"errors"

	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/internal/repo"
	"github.com/9688101/HX/pkg"
	"github.com/9688101/HX/pkg/config"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, req RegisterUserRequest) error
	Login(ctx context.Context, username, password string) (*entity.User, error)
	// GetUserList 获取分页用户列表
	GetUserList(ctx context.Context, offset, limit int, order string) ([]*entity.User, error)
	// SearchUsers 根据关键词搜索用户
	SearchUsers(ctx context.Context, keyword string) ([]*entity.User, error)
	// GetUser 获取单个用户信息，并根据调用者角色过滤同级或更高级别用户
	GetUser(ctx context.Context, id int, callerRole int) (*entity.User, error)
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

func (uc *userUseCase) RegisterUser(ctx context.Context, req RegisterUserRequest) error {
	// 如果开启邮箱验证，则校验邮箱和验证码
	if config.EmailVerificationEnabled {
		if req.Email == "" || req.VerificationCode == "" {
			return errors.New("管理员开启了邮箱验证，请输入邮箱地址和验证码")
		}
		if !pkg.VerifyCodeWithKey(req.Email, req.VerificationCode, pkg.EmailVerificationPurpose) {
			return errors.New("验证码错误或已过期")
		}
	}

	// 根据邀请码获取邀请人ID（邀请码可能为空，返回0表示无邀请人）
	inviterId, _ := uc.repo.GetUserIdByAffCode(req.AffCode)

	// 如果 DisplayName 为空，则使用 Username
	displayName := req.Username

	// 构造待注册的用户对象（这里暂不对密码做处理，后续加密）
	newUser := entity.User{
		Username:    req.Username,
		Password:    req.Password,
		DisplayName: displayName,
		InviterId:   inviterId,
	}
	// 如果开启邮箱验证，则设置邮箱
	if config.EmailVerificationEnabled {
		newUser.Email = req.Email
	}

	// 调用 Repository 层插入用户记录
	if err := uc.repo.InsertUser(ctx, &newUser, inviterId); err != nil {
		return err
	}
	return nil
}

func (uc *userUseCase) Login(ctx context.Context, username, password string) (*entity.User, error) {
	// 从 Repo 层查询用户记录（假设按用户名唯一查询）
	user, err := uc.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 使用 common 包中的方法验证密码（假设 user.Password 为加密后密码）
	if !pkg.ValidatePasswordAndHash(password, user.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 检查用户状态
	if user.Status != entity.UserStatusEnabled {
		return nil, errors.New("用户已被封禁")
	}

	return user, nil
}

func (uc *userUseCase) GetUserList(ctx context.Context, offset, limit int, order string) ([]*entity.User, error) {
	return uc.repo.GetAllUsers(ctx, offset, limit, order)
}

func (uc *userUseCase) SearchUsers(ctx context.Context, keyword string) ([]*entity.User, error) {
	return uc.repo.SearchUsers(ctx, keyword)
}

func (uc *userUseCase) GetUser(ctx context.Context, id int, callerRole int) (*entity.User, error) {
	user, err := uc.repo.GetUserByID(ctx, id, false)
	if err != nil {
		return nil, err
	}
	// 权限校验：调用者的角色必须大于目标用户的角色，除非调用者为超级管理员
	if callerRole <= user.Role && callerRole != entity.RoleRootUser {
		return nil, errors.New("无权获取同级或更高权限等级用户的信息")
	}
	return user, nil
}
