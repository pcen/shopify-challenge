package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"image-repo/database"
	"image-repo/routes"
)

// getPort returns the port to run the backend on based off of the port defined
// in either the .env file or as an environment variable.
func getPort() string {
	return fmt.Sprintf(":%s", os.Getenv("IMAGE_REPO_SERVER_PORT"))
}

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
	app.Run(getPort())
}
