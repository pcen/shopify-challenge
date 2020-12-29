package main

import (
	"github.com/gin-gonic/gin"

	"image-repo/models"
	"image-repo/routes"
)

func main() {
	// Initialize the database
	models.InitializeDatabase()

	// Initialize gin Engine
	app := gin.Default()

	// Initialize endpoints
	routes.AttachAll(app)

	// Run web application on port 8000
	app.Run(":8000")
}
