package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"image-repo/core"
)

// routeImages handles get requests to '/images'
func routeImages(c *gin.Context) {
	// Reject unauthorized requests
	if !core.RequestAuthorized(c) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"images": gin.H{
			"image 1": "image 1 data",
		},
	})
}
