package repo

import (
	"context"

	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/pkg/utils"
)

// InsertUser 插入用户
func (r *userRepo) InsertUser(ctx context.Context, user *entity.User, inviterId int) error {
	// 此处可以进行密码加密（如果尚未加密），例如：
	hashedPwd, err := utils.Password2Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPwd

	// 其他默认值设置、日志记录等也可在此处执行

	return r.db.WithContext(ctx).Create(user).Error
}

// GetUserIdByAffCode 根据邀请码获取用户ID
func (r *userRepo) GetUserIdByAffCode(ctx context.Context, affCode string) (int, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("aff_code = ?", affCode).First(&user).Error; err != nil {
		return 0, err
	}
	return user.Id, nil
}

// // GetUserByUsername 根据用户名获取用户
func (r *userRepo) GetUserByUsername(ctx context.Context, username string, includeSensitive bool) (*entity.User, error) {
	var user entity.User
	query := r.db.WithContext(ctx).Where("username = ?", username)
	if !includeSensitive {
		query = query.Omit("password", "access_token")
	}
	if err := query.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// // GetAllUsers 分页查询用户列表，按 order 排序，过滤掉状态为删除的用户
func (r *userRepo) GetUserList(ctx context.Context, offset, limit int, order string) ([]*entity.User, error) {
	var users []*entity.User
	query := r.db.WithContext(ctx).Model(&entity.User{}).Where("status != ?", entity.UserStatusDeleted)
	switch order {
	case "quota":
		query = query.Order("quota desc")
	case "used_quota":
		query = query.Order("used_quota desc")
	case "request_count":
		query = query.Order("request_count desc")
	default:
		query = query.Order("id desc")
	}
	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// // SearchUsers 根据关键词模糊搜索用户
func (r *userRepo) SearchUsers(ctx context.Context, keyword string) ([]*entity.User, error) {
	var users []*entity.User
	// 此处对 username, email, display_name 进行模糊查询
	if err := r.db.WithContext(ctx).Omit("password").Where(
		"username LIKE ? OR email LIKE ? OR display_name LIKE ?",
		keyword+"%", keyword+"%", keyword+"%",
	).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// // GetUserByID 根据 id 查询用户（includeSensitive 控制是否返回敏感信息）
func (r *userRepo) GetUserByID(ctx context.Context, id int, includeSensitive bool) (*entity.User, error) {
	var user entity.User
	query := r.db.WithContext(ctx).Where("id = ?", id)
	if !includeSensitive {
		query = query.Omit("password", "access_token")
	}
	if err := query.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", user.Id).Updates(user).Error
}

func (r *userRepo) DeleteUserByID(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.User{}).Error
}
