package routes

import (
	"fmt"
	"inventory-app/models"
	"inventory-app/persistence"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBoxes(rg *gin.RouterGroup) {
	r := rg.Group("/box")

	r.POST("/", createBox)
	r.GET("/all", getAllBoxes)
	r.GET("/:id", getBox)
	r.DELETE("/:id", deleteBox)
}

func getBox(c *gin.Context) {
	id := c.Params.ByName("id")

	var box models.Box
	if err := persistence.Db.Where("id = ?", id).First(&box).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, box)
	}
}

func createBox(c *gin.Context) {
	var newBox models.Box

	if err := c.BindJSON(&newBox); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	result := persistence.Db.Create(&newBox)
	fmt.Printf("Result: %#v\n", result)

	if result.Error != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, newBox)
}

func getAllBoxes(c *gin.Context) {
	var boxes []models.Box
	if err := persistence.Db.Find(&boxes).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, boxes)
	}
}

func deleteBox(c *gin.Context) {
	id := c.Params.ByName("id")

	var box models.Box
	d := persistence.Db.Delete(&box, id)
	fmt.Println(d)
	c.JSON(http.StatusOK, gin.H{"id": id, "result": "deleted"})
}
