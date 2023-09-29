package handlers

import (
	"data-4-me-api/auth"
	"data-4-me-api/models"
	"data-4-me-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNoteHandler(c *gin.Context) {
	var note models.Note

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := auth.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	clientBelongs, err := services.VerifyClientBelongsToUser(note.ClientID, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !clientBelongs {
		c.JSON(http.StatusForbidden, gin.H{"error": "Client does not belong to the user"})
		return
	}

	err = services.CreateNote(&note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"note": note})
}
