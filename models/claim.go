package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Claims struct {
	gorm.Model

	Email string `json:"email"`
	jwt.RegisteredClaims
}
