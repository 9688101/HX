package entity

import "gorm.io/gorm"

func MigrateDB(db *gorm.DB) error {
	var err error

	if err = db.AutoMigrate(&User{}); err != nil {
		return err
	}

	return nil
}
