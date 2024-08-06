package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Box struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique"`
	Location string `json:"location"`
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
}

var (
	db  *gorm.DB
	err error
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Define the connection string.
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")

	dsn := "user=" + user + " dbname=default_database sslmode=disable password=" + pass

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

	r.POST("/box", createBox)

	log.Println("Database connection successful")

	r.Run()
}

func createBox(ctx *gin.Context) {
	var newBox Box

	if err := ctx.BindJSON(&newBox); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	result := db.Create(&newBox)

	fmt.Printf("Result: %#v\n", result)

	ctx.JSON(http.StatusOK, newBox)
}
