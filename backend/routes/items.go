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
	r.GET("/:id", getItem)
	r.POST("/", createItem)
	r.DELETE("/:id", deleteItem)
	r.PUT("/:id", updateItem)
	r.GET("/search", searchItems)
}

func getItem(c *gin.Context) {
	id := c.Params.ByName("id")

	var item models.Item
	if err := persistence.Db.Where("id = ?", id).First(&item).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, item)
	}
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

func searchItems(c *gin.Context) {
	query := c.DefaultQuery("query", "") // Default to empty string if no query parameter

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	var items []models.Item
	result := persistence.Db.Where("name ILIKE ?", "%"+query+"%").Find(&items)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
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

func deleteItem(c *gin.Context) {
	id := c.Params.ByName("id")

	var item models.Item
	persistence.Db.Delete(&item, id)
	c.JSON(http.StatusOK, gin.H{"id": id, "result": "deleted"})
}

func updateItem(c *gin.Context) {
	id := c.Params.ByName("id")

	var item models.Item
	if err := persistence.Db.Where("id = ?", id).First(&item).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		var newBox models.Item

		if err := c.BindJSON(&newBox); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		item.Name = newBox.Name

		if err := persistence.Db.Save(&item).Error; err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.Status(http.StatusNoContent)
	}
}
