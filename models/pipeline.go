package models

import (
	"gorm.io/gorm"
)

type Pipeline struct {
	gorm.Model
	ID             uint            `gorm:"primaryKey"`
	Name           string          `json:"name" binding:"required"`
	ClientID       uint            `json:"client_id" gorm:"not null;index"`
	PipelineStages []PipelineStage `gorm:"foreignKey:PipelineID"`
}

type PipelineStage struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	Name       string `json:"name" binding:"required"`
	Status     string `json:"status" binding:"required"`
	PipelineID uint   `json:"pipeline_id" gorm:"not null;index"`
}
