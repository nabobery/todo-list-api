// @title         Todo List API
// @version       1.0
// @description   A RESTful API for managing to-do items.
// @host          localhost:8080
// @BasePath      /
package main

import (
	"log"
	"os"
	"todo-list-api/config"
	"todo-list-api/routes"

	"github.com/gin-gonic/gin"

	// Import the generated Swagger docs package
	_ "todo-list-api/docs"
)

func main() {
	// Load configuration and connect to MongoDB
	config.LoadConfig()

	// Initialize Gin router
	router := gin.Default()

	// Register all routes / controllers (including auth and to-do endpoints)
	routes.RegisterRoutes(router)

	// Set default port if not specified in environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	log.Fatal(router.Run(":" + port))
}
