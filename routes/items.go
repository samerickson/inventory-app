package routes

import (
	"fmt"
	"inventory-app/models"
	"inventory-app/persistence"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddItems(rg *gin.RouterGroup) {
	r := rg.Group("/item")

	r.GET("/", getAllItems)
	r.POST("/", createItem)
}

func getAllItems(c *gin.Context) {
	var items []models.Item
	if err := persistence.Db.Find(&items).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, items)
	}
}

func createItem(c *gin.Context) {
	var newItem models.Item

	if err := c.BindJSON(&newItem); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := persistence.Db.Create(&newItem)
	fmt.Printf("Result: %#v\n", result)

	if result.Error != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, newItem)
}
