package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	. "image-repo/models"
)

// getImageUploadMetadata extracts the uploaded image metadata from the
// request body and returns it as a list of ImageUploadMeta structs.
func getImageUploadMetadata(c *gin.Context) ([]ImageUploadMeta, error) {
	var meta []ImageUploadMeta
	err := json.Unmarshal([]byte(c.Request.FormValue("meta")), &meta)
	return meta, err
}

// saveImages extracts uploaded image files from the request body and saves
// them to the image repository. The metadata for each image is inserted into
// the metadata database.
func saveImages(c *gin.Context, meta []ImageUploadMeta) {
	for _, m := range meta {
		name := m.Name
		fmt.Println(name)
		image, _, err := c.Request.FormFile(name)
		if err != nil {
			panic(err.Error())
		}
		out, err := os.Create(filepath.Join(GetImagesDir(), name))
		if err != nil {
			panic(err.Error())
		}
		defer out.Close()
		_, err = io.Copy(out, image)
		if err != nil {
			panic(err.Error())
		}
	}
}

// routeUpload handles post requests to '/upload'
func routeUpload(c *gin.Context) {
	meta, err := getImageUploadMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid upload meta"})
		return
	}

	saveImages(c, meta)

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successful",
	})
}
