package routes

import (
	"example.com/exercises/events-api/middlwares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)

	authenticated := server.Group("/events")
	authenticated.Use(middlwares.Authenticate)
	authenticated.POST("/", createEvent)
	authenticated.PUT("/:id", updateEvent)
	authenticated.DELETE("/:id", deleteEvent)
	authenticated.POST("/:id/register", registerForEvent)
	authenticated.DELETE("/:id/register", cancelRegistration)

}
