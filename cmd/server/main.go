package main

import (
	"event-booking-api/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Load()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	port := config.Get("APP_PORT")

	r.Run(":" + port)
}
