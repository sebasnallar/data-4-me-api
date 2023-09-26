package main

import (
	"data-4-me-api/pkg/db"
	"data-4-me-api/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Migrate(database); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	routes.SetupRoutes(r)

	if err := r.Run(); err != nil {
		fmt.Println("Server run failed: ", err)
	}
}
