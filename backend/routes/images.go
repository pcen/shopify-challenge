package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	. "image-repo/models"
)

// routeImages handles post requests to '/images'
func routeImages(c *gin.Context) {
	var body ImageDownloadMeta
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user associated with the request's JWT
	user, err := GetUserFromJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorization error"})
		return
	}

	// Get the requested image's metadata from the database
	metadata, err := GetImageMetadata(body.Image, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to retrieve images"})
		return
	}

	fmt.Println(metadata)

	c.JSON(http.StatusOK, gin.H{
		"images": map[uint]interface{}{
			metadata.ID: metadata,
		},
	})
}
