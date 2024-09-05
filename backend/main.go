package main

import (
	"inventory-app/persistence"
	"inventory-app/routes"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Box struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" gorm:"unique"`
	Location   string `json:"location"`
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
}

func main() {
	if persistence.ConnectDb() {
		r := gin.Default()

		r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})

		v1 := r.Group("/v1")
		routes.AddBoxes(v1)
		routes.AddItems(v1)

		r.Run()
	}
}
