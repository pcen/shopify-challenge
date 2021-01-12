package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"image-repo/database"
	"image-repo/routes"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()

	// Initialize the database
	database.InitializeDatabase()

	// Initialize gin Engine
	app := gin.Default()

	// Initialize endpoints
	routes.AttachAll(app)

	// Run web application on port 8000
	app.Run(os.Getenv("IMAGE_REPO_SERVER_PORT"))
}
