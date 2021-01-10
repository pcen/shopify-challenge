package routes

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"

	. "image-repo/models"
)

// routeImage handles get requests to '/image/<id>'
func routeImage(c *gin.Context) {
	// Get the requested image ID from the endpoint
	imageID, err := strconv.Atoi(c.Param("id"))
	if err != nil || imageID < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid image id"})
		return
	}

	// Get the user associated with the request's JWT
	user, err := GetUserFromJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	// Get the filename of the image in the database
	filename, err := GetImageFileStore(uint(imageID), user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image file not found"})
		return
	}

	// Return the image file
	imageFile := filepath.Join(GetImagesDir(), filename)
	c.File(imageFile)
	c.Status(http.StatusOK)
}

// routeImages handles post requests to '/images'
func routeImages(c *gin.Context) {
	var body ImageDownloadMeta
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Get the user associated with the request's JWT
	user, err := GetUserFromJWT(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	// Get the requested image's metadata from the database
	metadata, err := SearchQueryImageMetadata(user.ID, "", false)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to retrieve images"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"images": metadata,
	})
}
