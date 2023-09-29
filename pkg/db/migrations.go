package db

import (
	"data-4-me-api/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Client{},
		&models.Product{},
		&models.Note{},
	)
}
