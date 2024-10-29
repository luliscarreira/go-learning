package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/exercises/events-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("UserId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})

		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event by ID",
		})

		return
	}

	err = event.Register(userId)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not register for event",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully registered for event",
	})
}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("UserId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})

		return
	}

	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not cancel registration for event",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully cancelled registration for event",
	})
}
