package main

import (
	"event-booking-api/internal/routes"
	"event-booking-api/pkg/config"
	"event-booking-api/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Load()
	database.Connect()

	router := gin.Default()

	routes.Register(router)
	
	port := config.Get("APP_PORT")
	router.Run(":" + port)
}
