package routes

import (
	"fmt"
	"net/http"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	. "image-repo/models"
)

// routeImages handles get requests to '/images'
func routeImages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"images": gin.H{
			"image 1": "image 1 data",
		},
	})
}

// routeUpload handles post requests to '/upload'
func routeUpload(c *gin.Context) {
	var imageForm ImageMultipart
	if err := c.Bind(&imageForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get image data
	name := imageForm.Image.Filename
	image, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(name)
	fmt.Println(imageForm.Description)

	// Save image to disk
	out, err := os.Create(GetImagesDir() + "/" + name)
	if err != nil {
		panic(err.Error())
	}

	defer out.Close()
	_, err = io.Copy(out, image)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("uploaded %s successfully", name),
	})
}
