package main

import (
	"github.com/gin-gonic/gin"

	"myapp/dbnya"
	"myapp/handlers"
)

func main() {
	router := gin.Default()

	// Connect to the database
	dbnya.Connect()

	// Routes
	router.GET("/get-electronic", handlers.GetElektronikBebasNamanya)

	// Start the server
	router.Run("localhost:8080")
}
