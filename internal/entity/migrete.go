package entity

import "gorm.io/gorm"

func MigrateDB(db *gorm.DB) error {
	var err error

	if err = db.AutoMigrate(&User{}); err != nil {
		return err
	}

	return nil
}

type Error struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    any    `json:"code"`
}
