package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	ClientID uint   `json:"client_id" gorm:"not null;index"`
	UserID   uint   `json:"user_id" gorm:"not null;index"`
}
