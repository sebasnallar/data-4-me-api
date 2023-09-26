package handlers

import (
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

	err := services.CreateClient(&client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client": client})
}

func DeleteClientHandler(c *gin.Context) {

}
