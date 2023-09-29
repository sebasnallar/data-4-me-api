package services

import (
	"data-4-me-api/models"
	"data-4-me-api/pkg/db"
)

func CreateNote(note *models.Note) error {
	result := db.GlobalDB.Create(note)
	return result.Error
}
