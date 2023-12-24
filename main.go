package main

import (
	"github.com/gin-gonic/gin"
	"seasonal_produce/db"
	"seasonal_produce/routes"
)

func main() {
	// Connect to the database
	db.ConnectDatabase()

	// Set up the router
	router := routes.SetupRouter()
	// Example route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the seasonal produce app!",
		})
	})
	// Run the server
	router.Run(":8080")
}
