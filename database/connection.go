package database

import (
	"github.com/joho/godotenv"
	"go-jwt-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Can not connect to DB. %v", err)
	}
	dsn := os.Getenv("DB_CONNECTION_STR")

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("can't connect to DB %v", err)
	}

	DB = connection

	if err := connection.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("can't migrate to DB: %v", err)
	}

}
