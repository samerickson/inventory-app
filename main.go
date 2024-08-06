package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Box struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Location string
	Id       uint `gorm:"primaryKey;autoIncrement"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Define the connection string.
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")

	dsn := "user=" + user + " dbname=default_database sslmode=disable password=" + pass

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auto Migrate
	db.AutoMigrate(&Box{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	log.Println("Database connection successful")

	r.Run()
}
