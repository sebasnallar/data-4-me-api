package services

import (
	"data-4-me-api/models"
	"data-4-me-api/pkg/db"
)

func CreateUser(user *models.User) error {
	err := user.SetPassword(user.Password)

	if err != nil {
		return err
	}

	result := db.GlobalDB.Create(user)
	return result.Error
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := db.GlobalDB.Find(&users)
	return users, result.Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := db.GlobalDB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
