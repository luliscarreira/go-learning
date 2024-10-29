package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/exercises/events-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not process request",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Events found",
		"events":  events,
	})
}

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})

		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not process request",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event found",
		"event":   event,
	})
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process request",
		})

		return
	}

	event.UserID = c.GetInt64("UserId")
	err = event.Save()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})

		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event by ID",
		})

		return
	}

	if event.UserID != c.GetInt64("UserId") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized to update this event",
		})

		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process request",
		})

		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})

		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event by ID",
		})

		return
	}

	if event.UserID != c.GetInt64("UserId") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized to delete this event",
		})

		return
	}

	err = event.Delete()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
