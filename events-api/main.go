package main

import (
	"net/http"

	"example.com/exercises/events-api/db"
	"example.com/exercises/events-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()
	r.GET("/events", getEvents)
	r.POST("/events", createEvent)
	r.Run()
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not process request",
			"error":   err,
		})
	}

	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process request",
			"error":   err,
		})

		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}
