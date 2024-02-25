package db

import (
	"log"
	"os"
	"api-daily-activity/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase(){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatalf("Error loading .env file: %v", err)
	}
	
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := "user=" + username + " password=" + password + " host=" + host + " port=" + port + " dbname=" + name + " sslmode=" + sslmode

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatalf("Error connecting to database: %v", err)
	}

    err = database.AutoMigrate(&models.Activity{})
    if err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }

	err = database.AutoMigrate(&models.Photo{})
    if err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }

    DB = database
}