package routes

import (
	"data-4-me-api/api/handlers"
	"data-4-me-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", middleware.AuthRequired(), handlers.HomeHandler)
		v1.POST("/create-client", middleware.AuthRequired(), handlers.CreateClientHandler)

		v1.POST("/register", handlers.RegisterHandler)
		v1.POST("/log-in", handlers.LoginHandler)

		v1.POST("/create-user", handlers.CreateUserHandler)
		v1.GET("/users", handlers.GetUsersHandler)
	}
}
