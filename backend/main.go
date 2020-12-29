package main

import (
	// "github.com/gin-gonic/gin"

	// "image-repo/routes"
	"image-repo/models"
)

func main() {
	db := models.GetDatabaseHandle()
	if db == nil {
		panic("db handle is nil")
	}
	return

	// Initialize gin Engine
	// app := gin.Default()

	// Initialize endpoints
	// routes.AttachAll(app)

	// Run web application on port 8000
	// app.Run(":8000")
}
