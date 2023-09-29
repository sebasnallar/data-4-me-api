package services

import (
	"data-4-me-api/models"
	"data-4-me-api/pkg/db"
)

func CreateNote(note *models.Note) error {
	result := db.GlobalDB.Create(note)
	return result.Error
}

func GetNotes(userID uint) ([]models.Note, error) {
	var notes []models.Note
	if err := db.GlobalDB.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}
