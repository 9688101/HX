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
}

type userUseCase struct {
	repo repo.UserRepository
}

// NewUserUseCase 返回 UserUseCase 的实现
func NewUserUseCase(r repo.UserRepository) *userUseCase {
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
