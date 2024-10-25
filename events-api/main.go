package main

import (
	"example.com/exercises/events-api/db"
	"example.com/exercises/events-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run()
}
