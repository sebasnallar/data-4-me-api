package handlers

import (
	"data-4-me-api/auth"
	"data-4-me-api/models"
	"data-4-me-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateClientHandler(c *gin.Context) {
	var client models.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := auth.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client.UserID = user.ID

	err = services.CreateClient(&client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client": client})
}

func GetClientsHandler(c *gin.Context) {
	user, err := auth.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	clients, err := services.GetClients(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching clients"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"clients": clients})
}
