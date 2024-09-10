package main

import (
	"fmt"
	"inventory-app/persistence"
	"inventory-app/routes"

	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Box struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" gorm:"unique"`
	Location   string `json:"location"`
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("HERE +++++++++++++++++++++++++++++++++++++++")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	if persistence.ConnectDb() {
		router := gin.Default()

		// Default allows all origins
		router.Use(cors.Default())

		v1 := router.Group("/v1", cors.Default())
		routes.AddBoxes(v1)
		routes.AddItems(v1)

		router.Run()
	}
}
