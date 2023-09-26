package handlers

import (
	"data-4-me-api/models"
	"data-4-me-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	dummyData := map[string]string{
		"message": "Welcome to the Data 4 Me API!",
		"status":  "success",
	}

	c.JSON(http.StatusOK, dummyData)
}

func GetUser(c *gin.Context) {

}

func CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func GetUsersHandler(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
