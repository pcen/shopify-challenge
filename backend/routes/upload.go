package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"image-repo/core"
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
func saveImages(c *gin.Context, meta []ImageUploadMeta, user *User) {

	for _, m := range meta {
		// Generate the name of the image file in the file registry
		store := core.RandomAlphanumericString(32) + core.FileExtensionFromFormat(m.Format)

		// Insert image metadata into SQL database
		imageMetadata := ImageMetadata{
			UserID:         user.ID,
			Name:           m.Name,
			Format:         m.Format,
			FileStore:      store,
			Description:    m.Description,
			Geolocation:    m.Location,
			OCRText:        "ocr text",
			Private:        m.Private,
			AverageHash:    0,
			DifferenceHash: 0,
		}

		err := InsertImageMetadata(&imageMetadata)
		if err != nil {
			fmt.Println(err.Error())
		}

		// Save image file to file registry
		image, _, err := c.Request.FormFile(m.Name)
		if err != nil {
			panic(err.Error())
		}

		err = core.WriteFile(filepath.Join(GetImagesDir(), store), image)
		if err != nil {
			panic(err.Error())
		}

	}
}

// routeUpload handles post requests to '/upload'
func routeUpload(c *gin.Context) {

	username, err := core.GetTokenUser(c.GetHeader("Authorization"))
	if err != nil {
		panic(err.Error())
	}

	user, err := GetUser(username)
	if err != nil {
		panic(err.Error())
	}

	meta, err := getImageUploadMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid upload meta"})
		return
	}

	saveImages(c, meta, &user)

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successful",
	})
}
