package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Claims struct {
	gorm.Model

	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}
