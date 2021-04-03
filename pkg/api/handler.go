package api

import (
	"net/http"

	"github.com/arh0329/supermarket-api/models"
	"github.com/gin-gonic/gin"
)

var empty = models.Item{}

func getAllProduce(c *gin.Context) {
	produce := models.GetAllProduce()
	c.JSON(http.StatusOK, gin.H{"produce": produce})
}

func addProduce(c *gin.Context) {
	var items []models.Item
	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, item := range items {
		if err := item.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		models.AddProduce(item)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Item(s) added"})
}

func getOneItem(c *gin.Context) {
	pc := c.Param("pc")

	item := models.GetOneItem(pc)
	if item != empty {
		c.JSON(http.StatusOK, item)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
	}
}

func deleteItem(c *gin.Context) {
	pc := c.Param("pc")
	item := models.DeleteProduce(pc)
	if item != empty {
		c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
	}
}
