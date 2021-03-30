package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	logger "github.com/arh0329/supermarket-api/pkg/logging"
)

func main() {
	r := gin.New()
	r.Use(logMiddleware())
	r.Use(gin.Recovery())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	err := r.Run(":8000")
	if err != nil {
		logger.Log().WithError(err).Fatal("Error occurred running http server")
	}
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
