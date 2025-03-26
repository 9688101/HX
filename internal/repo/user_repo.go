package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/9688101/HX/common"
	"github.com/9688101/HX/common/config"
	"github.com/9688101/HX/common/logger"
	"github.com/9688101/HX/common/random"
	"github.com/9688101/HX/internal/entity"
	"gorm.io/gorm"
)

// type UserRepo interface {
//     // CreateUser(ctx context.Context, user *entity.User) error
//     // GetUserByID(ctx context.Context, id int) (*entity.User, error)
// }

type UserRepo struct {
	DB *gorm.DB
}

// NewUserRepo 返回 UserRepo 的实现
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (uu *UserRepo) GetUserIdByAffCode(affCode string) (int, error) {
	if affCode == "" {
		return 0, errors.New("affCode 为空！")
	}
	u := entity.NewUser()
	err := uu.DB.Select("id").First(u, "aff_code = ?", affCode).Error
	return u.Id, err
}

func (uu *UserRepo) Insert(ctx context.Context, u *entity.User) error {
	var err error
	if u.Password != "" {
		u.Password, err = common.Password2Hash(u.Password)
		if err != nil {
			return err
		}
	}
	u.Quota = config.QuotaForNewUser
	u.AccessToken = random.GetUUID()
	u.AffCode = random.GetRandomString(4)
	result := uu.DB.Create(u)
	if result.Error != nil {
		return result.Error
	}
	// if config.QuotaForNewUser > 0 {
	// 	RecordLog(ctx, user.Id, LogTypeSystem, fmt.Sprintf("新用户注册赠送 %s", common.LogQuota(config.QuotaForNewUser)))
	// }
	// if inviterId != 0 {
	// 	if config.QuotaForInvitee > 0 {
	// 		_ = IncreaseUserQuota(user.Id, config.QuotaForInvitee)
	// 		RecordLog(ctx, user.Id, LogTypeSystem, fmt.Sprintf("使用邀请码赠送 %s", common.LogQuota(config.QuotaForInvitee)))
	// 	}
	// 	if config.QuotaForInviter > 0 {
	// 		_ = IncreaseUserQuota(inviterId, config.QuotaForInviter)
	// 		RecordLog(ctx, inviterId, LogTypeSystem, fmt.Sprintf("邀请用户赠送 %s", common.LogQuota(config.QuotaForInviter)))
	// 	}
	// }
	// create default token
	// cleanToken := entity.Token{
	// 	UserId:         u.Id,
	// 	Name:           "default",
	// 	Key:            random.GenerateKey(),
	// 	CreatedTime:    helper.GetTimestamp(),
	// 	AccessedTime:   helper.GetTimestamp(),
	// 	ExpiredTime:    -1,
	// 	RemainQuota:    -1,
	// 	UnlimitedQuota: true,
	// }
	result.Error = tr.Insert(u.Id)
	if result.Error != nil {
		// do not block
		logger.SysError(fmt.Sprintf("create default token for user %d failed: %s", u.Id, result.Error.Error()))
	}
	return nil
}

func (uu *UserRepo) ValidateAndFill() error {

	return nil
}

// ValidateAndFill check password & user status
func (uu *UserRepo) ValidateByName(u *entity.User) error {
	return uu.DB.Where("username = ?", u.Username).First(u).Error
}

// we must make sure check username firstly
// consider this case: a malicious user set his username as other's email
func (uu *UserRepo) ValidateByEmail(u *entity.User) error {
	return uu.DB.Where("email = ?", u.Username).First(u).Error

}
