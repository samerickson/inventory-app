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
	r.GET("/", getAllBoxes)
	r.GET("/:id", getBox)
	r.DELETE("/:id", deleteBox)
	r.PUT("/:id", updateBox)
	r.GET("search", searchItems)
}

func getBox(c *gin.Context) {
	id := c.Params.ByName("id")

	var box models.Box
	if err := persistence.Db.Preload("Items").Where("id = ?", id).First(&box).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, box)
	}
}

func searchBox(c *gin.Context) {
	query := c.DefaultQuery("query", "") // Default to empty string if no query parameter

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	var boxes []models.Box
	result := persistence.Db.Where("name ILIKE ?", "%"+query+"%").Find(&boxes)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, boxes)
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
	if err := persistence.Db.Preload("Items").Find(&boxes).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, boxes)
	}
}

func deleteBox(c *gin.Context) {
	id := c.Params.ByName("id")

	var box models.Box
	persistence.Db.Delete(&box, id)
	c.JSON(http.StatusOK, gin.H{"id": id, "result": "deleted"})
}

func updateBox(c *gin.Context) {
	id := c.Params.ByName("id")

	var box models.Box
	if err := persistence.Db.Where("id = ?", id).First(&box).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		var newBox models.Box

		if err := c.BindJSON(&newBox); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		box.Location = newBox.Location
		box.Name = newBox.Name

		if err := persistence.Db.Save(&box).Error; err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.Status(http.StatusNoContent)
	}
}
