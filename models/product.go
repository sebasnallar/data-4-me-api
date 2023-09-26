package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID      uint   `gorm:"not null;index"`
	Name        string `gorm:"not null"`
	Description string
	Price       float64 `gorm:"not null"`
	SKU         string  `gorm:"uniqueIndex"`
	Category    string
	ImageURL    string
	Status      string
}
