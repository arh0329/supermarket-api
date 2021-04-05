package api

import (
	"net/http"
	"sync"

	"github.com/arh0329/supermarket-api/models"
	logger "github.com/arh0329/supermarket-api/pkg/logging"
	"github.com/gin-gonic/gin"
)

var empty = models.Item{}

type response struct {
	Message string `json:"message"`
}

func getAllProduce(c *gin.Context) {
	produce := models.GetAllProduce()
	c.JSON(http.StatusOK, produce)
}

func addProduce(c *gin.Context) {
	var items []models.Item
	if err := c.ShouldBindJSON(&items); err != nil {
		logger.Log().WithError(err).Error("Error occurred getting request body")
		c.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	chanErrs := []string{}
	chanItems := []string{}

	var mutex = &sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(len(items))
	errChan := make(chan error, len(items))
	itemChan := make(chan string, len(items))
	for _, item := range items {

		go func(it models.Item) {
			if err := it.Validate(); err != nil {
				logger.Log().WithError(err).Error("Error validating item")
				errChan <- err
			} else {
				mutex.Lock()
				if err := models.AddProduce(it); err != nil {
					errChan <- err
				} else {
					itemChan <- it.Name
				}
				mutex.Unlock()
			}
			wg.Done()
		}(item)
	}
	wg.Wait()
	close(errChan)
	close(itemChan)

	if len(errChan) != 0 {
		for err := range errChan {
			if err != nil {
				chanErrs = append(chanErrs, err.Error())
			}
		}
	}

	if len(itemChan) != 0 {
		for item := range itemChan {
			chanItems = append(chanItems, item)
		}
	}

	if len(chanItems) == 0 {
		resp := gin.H{"message": "No items added", "added": chanItems, "errors": chanErrs}
		logger.Log().Info(resp)
		c.JSON(http.StatusBadRequest, resp)
	} else {
		resp := gin.H{"message": "Item(s) added", "added": chanItems, "errors": chanErrs}
		logger.Log().Info(resp)
		c.JSON(http.StatusCreated, resp)
	}

}

func getOneItem(c *gin.Context) {
	pc := c.Param("pc")

	item := models.GetOneItem(pc)
	if item != empty {
		c.JSON(http.StatusOK, item)
	} else {
		logger.Log().WithField("productCode", pc).Info("item not found")
		c.JSON(http.StatusNotFound, response{Message: "Item not found"})
	}
}

func deleteItem(c *gin.Context) {
	pc := c.Param("pc")
	item := models.DeleteProduce(pc)
	if item != empty {
		logger.Log().WithField("item", item).Info("item deleted")
		c.JSON(http.StatusOK, response{Message: "Item deleted"})
	} else {
		logger.Log().WithField("productCode", pc).Info("item not found")
		c.JSON(http.StatusNotFound, response{Message: "Item not found"})
	}
}
