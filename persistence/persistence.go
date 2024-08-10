package persistence

import (
	"inventory-app/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func ConnectDb() bool {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Define the connection string.
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")

	dsn := "host= " + host + " user=" + user + " dbname=default_database sslmode=disable password=" + pass

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return false
	}

	// Auto Migrate
	Db.AutoMigrate(&models.Box{})

	log.Println("Database connection successful")

	return true
}
