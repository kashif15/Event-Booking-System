package booking

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookEvent(c *gin.Context) {

	eventID, err := strconv.ParseInt(c.Param("id"),10,64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	userID := c.GetInt64("userId")
	err = Create(userID, eventID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "event booked successfully",
	})
}

func MyBookings(c *gin.Context) {

	userID := c.GetInt64("userId")

	bookings, err := GetByUser(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func CancelBooking(c *gin.Context) {

	eventID, err := strconv.ParseInt(c.Param("id"),10,64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	userID := c.GetInt64("userId")

	err = Cancel(userID, eventID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "booking cancelled",
	})
}