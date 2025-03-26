package repo

import (
	"github.com/9688101/HX/common/helper"
	"github.com/9688101/HX/common/random"
	"github.com/9688101/HX/internal/entity"
	"gorm.io/gorm"
)

type TokenRepo struct {
	DB *gorm.DB
}

var tr = new(TokenRepo)

// NewUserRepo 返回 UserRepo 的实现
func GetTokenRepo() *TokenRepo {
	return tr
}

func (tr *TokenRepo) Insert(id int) error {
	// create default token
	t := entity.Token{
		UserId:         id,
		Name:           "default",
		Key:            random.GenerateKey(),
		CreatedTime:    helper.GetTimestamp(),
		AccessedTime:   helper.GetTimestamp(),
		ExpiredTime:    -1,
		RemainQuota:    -1,
		UnlimitedQuota: true,
	}
	var err error
	err = tr.DB.Create(t).Error
	return err
}
