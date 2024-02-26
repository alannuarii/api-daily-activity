package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectionDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := "user=" + username + " password=" + password + " host=" + host + " port=" + port + " dbname=" + name + " sslmode=" + sslmode

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	DB = database

	// Menutup koneksi secara otomatis setelah selesai menggunakan koneksi
	// defer DB.Close()
}
