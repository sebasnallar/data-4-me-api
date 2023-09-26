package db

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func init() {
	godotenv.Load()
}

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var db *gorm.DB
	var err error

	// Bypass to wait for DB to start
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Failed to connect to database. Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	if db == nil {
		return nil, fmt.Errorf("failed to connect to database after 10 retries")
	}

	GlobalDB = db

	return db, nil
}
