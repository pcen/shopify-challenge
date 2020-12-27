package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// routeImages handles get requests to '/images'
func routeImages(c *gin.Context) {
	authorized := tokenValid(c.GetHeader("Authorization"))
	// Reject unauthorized requests
	if !authorized {
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