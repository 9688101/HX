package repo

import (
	"context"

	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/pkg"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *entity.User, inviterId int) error
	GetUserIdByAffCode(affCode string) (int, error)
}

type userRepo struct {
	db *gorm.DB
}

// NewUserRepo 返回 UserRepo 的实现
func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) InsertUser(ctx context.Context, user *entity.User, inviterId int) error {
	// 此处可以进行密码加密（如果尚未加密），例如：
	hashedPwd, err := pkg.Password2Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPwd

	// 其他默认值设置、日志记录等也可在此处执行

	return r.db.WithContext(ctx).Create(user).Error
}
