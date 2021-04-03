package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateRouter() *gin.Engine {
	r := gin.New()
	r.Use(logMiddleware())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.POST("/produce", addProduce)
	r.GET("/produce", getAllProduce)
	r.GET("/produce/:pc", getOneItem)
	r.DELETE("/produce/:pc", deleteItem)

	return r
}

func logMiddleware() gin.HandlerFunc {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.InfoLevel)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		logger.WithFields(logrus.Fields{
			"statusCode":  statusCode,
			"latencyTime": fmt.Sprintf("%13v", latencyTime),
			"reqMethod":   reqMethod,
			"reqURI":      reqUri,
		}).Info()
	}
}
