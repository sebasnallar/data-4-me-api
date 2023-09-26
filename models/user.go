package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           int    `gorm:"primaryKey"`
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required" gorm:"unique"`
	PasswordHash string `gorm:"column:password_hash"`
	Password     string `json:"password" binding:"required" gorm:"-"`
}

func (user *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hash)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}
