package services

import (
	"data-4-me-api/models"
	"data-4-me-api/pkg/db"
)

func CreateClient(client *models.Client) error {
	result := db.GlobalDB.Create(client)
	return result.Error
}
