package routes

import (
	"data-4-me-api/api/handlers"
	"data-4-me-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health-check", handlers.HealthCheckHandler)

		// Auth Routes
		v1.POST("/register", handlers.RegisterHandler)
		v1.POST("/log-in", handlers.LoginHandler)

		// Client Routes
		v1.POST("/create-client", middleware.AuthRequired(), handlers.CreateClientHandler)
	}
}
