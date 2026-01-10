package event

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type eventRegister struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	EventTime   string `json:"event_time" binding:"required"`
	Capacity    int    `json:"capacity" binding:"required,min=1"`
}

func CreateEvent(c *gin.Context) {
	var req eventRegister

	err := c.ShouldBindBodyWithJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request data",
		})
		return
	}

	eventTime, err := time.Parse(time.RFC3339, req.EventTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event_time format"})
		return
	}

	userID := c.GetInt64("userId")

	event := &Event{
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		EventTime:   eventTime,
		Capacity:    req.Capacity,
		CreatedBy:   userID,
	}

	err = Create(event)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create event"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func ListEvents(c *gin.Context) {
	
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	status := c.Query("status")

	var createdBy *int64
	if c.Query("createdBy") == "me" {
		uid := c.GetInt64("userId")
		createdBy = &uid

	}

	var fromDate *time.Time
	if dateStr := c.Query("from_date"); dateStr != "" {
		t, err := time.Parse("2006-01-02", dateStr)
		if err == nil {
			fromDate = &t
		}
	}

	searchQuery := c.Query("search")
	var search *string

	if searchQuery != "" {
		search = &searchQuery
	}

	events, err := GetWillFilter(page, limit, status, createdBy, fromDate, search)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not fetch events",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":   page,
		"limit": limit,
		"data":   events,
	})


	c.JSON(http.StatusOK, events)
}

func GetEvent(c *gin.Context){
	id, err := strconv.ParseInt(c.Param("id"),10,64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	event, err := GetByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	c.JSON(http.StatusOK, event)

}

func DeleteEvent(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"),10,64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	event, err := GetByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	userID := c.GetInt64("userId")
	role :=  c.GetString("role")

	if event.CreatedBy != userID && role != "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	err = Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event deleted"})
}