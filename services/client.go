package services

import (
	"data-4-me-api/models"
	"data-4-me-api/pkg/db"
)

func CreateClient(client *models.Client) error {
	result := db.GlobalDB.Create(client)
	return result.Error
}

func GetClients(userID uint) ([]models.Client, error) {
	var clients []models.Client
	if err := db.GlobalDB.Where("user_id = ?", userID).Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func GetclientById(clientID uint) (*models.Client, error) {
	client := &models.Client{}
	if err := db.GlobalDB.First(client, clientID).Error; err != nil {
		return nil, err
	}
	return client, nil
}

func VerifyClientBelongsToUser(clientID, userID uint) (bool, error) {
	var client models.Client
	err := db.GlobalDB.First(&client, clientID).Error
	if err != nil {
		return false, err
	}
	return client.UserID == userID, nil
}
