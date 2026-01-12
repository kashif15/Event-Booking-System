package routes

import (
	"event-booking-api/internal/auth"
	"event-booking-api/internal/booking"
	"event-booking-api/internal/event"
	"event-booking-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.POST("/auth/register", auth.RegisterHandler)
	router.POST("/auth/login", auth.LoginHandler)

	router.POST("/auth/refresh", auth.RefreshHandler)
	router.POST("/auth/logout", auth.LogoutHandler)


	auth := router.Group("/")
	auth.Use(middleware.Authenticate())

	// Events
	auth.GET("/events", event.ListEvents)
	auth.GET("/events/:id", event.GetEvent)
	auth.POST("/events", event.CreateEvent)
	auth.DELETE("/events/:id", event.DeleteEvent)

	// Bookings
	auth.POST("/events/:id/book", booking.BookEvent)
	auth.DELETE("/events/:id/book", booking.CancelBooking)
	auth.GET("/bookings", booking.MyBookings)

}