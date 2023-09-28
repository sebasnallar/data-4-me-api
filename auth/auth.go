package auth

import (
	"data-4-me-api/models"
	"data-4-me-api/pkg/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) (models.User, error) {
	var user models.User
	email, exists := c.Get("email")
	if !exists {
		return user, gin.Error{
			Err:  fmt.Errorf("could not retrieve email from context"),
			Type: gin.ErrorTypePublic,
		}
	}

	if err := db.GlobalDB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, gin.Error{
			Err:  fmt.Errorf("error fetching user: %v", err),
			Type: gin.ErrorTypePublic,
		}
	}

	return user, nil
}
