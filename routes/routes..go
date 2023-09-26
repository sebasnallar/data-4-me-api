package routes

import (
	"data-4-me-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", handlers.HomeHandler)
	router.POST("/create-user", handlers.CreateUserHandler)
	router.GET("/users", handlers.GetUsersHandler)
}
