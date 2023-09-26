package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
	UserID   uint   `json:"user_id" gorm:"not null;index"`
}
