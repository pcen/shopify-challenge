package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// routeImages handles get requests to '/images'
func routeImages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"images": gin.H{
			"image 1": "image 1 data",
		},
	})
}
