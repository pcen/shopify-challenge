package routes

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"image-repo/models"
	. "image-repo/models"
)

// routeImage handles get requests to '/image/<id>'
func routeImage(c* gin.Context) {
	requestedId := c.Param("id")
	fmt.Println(requestedId)
	imageFile := filepath.Join(GetImagesDir(), "rwUjTqdceEB1rSEgRZ5YWf6dc5MT5BcU.jpg")
	c.Status(http.StatusOK)
	c.File(imageFile)
}

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
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	// Get the requested image's metadata from the database
	metadata, err := GetImageMetadata(body.Image, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to retrieve images"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"images": map[uint]interface{}{
			metadata.ID: metadata,
		},
	})

}
