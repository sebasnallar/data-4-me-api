package routes

import (
	"data-4-me-api/api/handlers"
	"data-4-me-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Auth Required Routes
	router.GET("/", middleware.AuthRequired(), handlers.HomeHandler)

	// Auth Routes
	router.POST("/register", handlers.RegisterHandler)

	// Users routes
	router.POST("/create-user", handlers.CreateUserHandler)
	router.GET("/users", handlers.GetUsersHandler)
}
