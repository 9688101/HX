package repo

import (
	"github.com/9688101/HX/internal/entity"
	"gorm.io/gorm"
)

type OptionRepository interface {
	GetAll() ([]*entity.Option, error)
	GetByKey(key string) (*entity.Option, error)
	Save(option *entity.Option) error
}

type optionRepository struct {
	db *gorm.DB
}

func NewOptionRepository(db *gorm.DB) OptionRepository {
	return &optionRepository{db: db}
}

func (r *optionRepository) GetAll() ([]*entity.Option, error) {
	var options []*entity.Option
	if err := r.db.Find(&options).Error; err != nil {
		return nil, err
	}
	return options, nil
}
func (r *optionRepository) GetByKey(key string) (*entity.Option, error) {
	var option entity.Option
	if err := r.db.First(&option, "key = ?", key).Error; err != nil {
		return nil, err
	}
	return &option, nil
}

func (r *optionRepository) Save(option *entity.Option) error {
	return r.db.Save(option).Error
}
