package repo

import (
	"context"

	"github.com/9688101/HX/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *entity.User, inviterId int) error
	GetUserIdByAffCode(ctx context.Context, affCode string) (int, error)
	GetUserByUsername(ctx context.Context, username string, includeSensitive bool) (*entity.User, error)
	GetUserList(ctx context.Context, offset, limit int, order string) ([]*entity.User, error)
	SearchUsers(ctx context.Context, keyword string) ([]*entity.User, error)
	GetUserByID(ctx context.Context, id int, includeSensitive bool) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUserByID(ctx context.Context, id int) error
}

type userRepo struct {
	db *gorm.DB
}

// NewUserRepo 返回 UserRepo 的实现
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}
