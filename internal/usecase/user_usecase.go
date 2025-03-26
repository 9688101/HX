package usecase

import (
	"context"
	"errors"

	"github.com/9688101/HX/common"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/internal/repo"
)

// type UserUseCase interface {
// 	// RegisterUser(ctx context.Context, user *entity.User) error
// 	// GetUser(ctx context.Context, id int) (*entity.User, error)
// }

type UserUseCase struct {
	userRepo *repo.UserRepo
}

// NewUserUseCase 返回 UserUseCase 的实现
func NewUserUseCase(r *repo.UserRepo) *UserUseCase {
	return &UserUseCase{
		userRepo: r,
	}
}

func (uu *UserUseCase) Register(ctx context.Context, u *entity.User) error {
	affCode := u.AffCode // this code is the inviter's code, not the user's own code
	inviterId, _ := uu.userRepo.GetUserIdByAffCode(affCode)
	// cleanUser := entity.User{
	// 	Username:    u.Username,
	// 	Password:    u.Password,
	// 	DisplayName: u.Username,
	// 	InviterId:   inviterId,
	// }
	u.InviterId = inviterId

	return uu.userRepo.Insert(ctx, u)
}

// ValidateAndFill check password & user status
func (uu *UserUseCase) ValidateAndFill(u *entity.User) (err error) {
	// When querying with struct, GORM will only query with non-zero fields,
	// that means if your field’s value is 0, '', false or other zero values,
	// it won’t be used to build query conditions
	password := u.Password
	if u.Username == "" || password == "" {
		return errors.New("用户名或密码为空")
	}
	err = uu.userRepo.ValidateByName(u)
	if err != nil {
		err = uu.userRepo.ValidateByEmail(u)
		if err != nil {
			return errors.New("用户名或密码错误，或用户已被封禁")
		}
	}
	okay := common.ValidatePasswordAndHash(password, u.Password)
	if !okay || u.Status != entity.UserStatusEnabled {
		return errors.New("用户名或密码错误，或用户已被封禁")
	}
	return nil
}
