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
	gorm.Model `json:"-"`
	Name       string `json:"name" gorm:"unique"`
	Location   string `json:"location"`
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
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
	r.GET("/box/all", getAllBoxes)
	r.GET("/box/:id", getBox)
	r.DELETE("/box/:id", deleteBox)

	log.Println("Database connection successful")

	r.Run()
}

func getBox(c *gin.Context) {
	id := c.Params.ByName("id")

	var box Box
	if err := db.Where("id = ?", id).First(&box).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, box)
	}
}

func createBox(c *gin.Context) {
	var newBox Box

	if err := c.BindJSON(&newBox); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := db.Create(&newBox)
	fmt.Printf("Result: %#v\n", result)

	if result.Error != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, newBox)
}

func getAllBoxes(c *gin.Context) {
	var boxes []Box
	if err := db.Find(&boxes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, boxes)
	}
}

func deleteBox(c *gin.Context) {
	id := c.Params.ByName("id")

	var box Box
	d := db.Delete(&box, id)
	fmt.Println(d)
	c.JSON(200, gin.H{"id": id, "result": "deleted"})
}
