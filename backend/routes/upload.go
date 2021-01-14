package routes

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"image-repo/core"
	. "image-repo/database"
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
			MLTags:         "",
			Private:        m.Private,
		}

		if err := InsertImage(&imageMetadata); err != nil {
			panic(err.Error())
		}

		// Save image file to file registry
		image, _, err := c.Request.FormFile(m.Name)
		if err != nil {
			panic(err.Error())
		}

		if err = WriteImageFile(store, image); err != nil {
			panic(err.Error())
		}

		// Get image tags and update the image's metadata. If the tagging
		// fails, the metadata is unmodified and no tags are applied.
		fullPath := filepath.Join(GetImagesDir(), store)
		tags, err := core.GetImageTags(fullPath)
		if err == nil {
			// imageMetadata ID is set upon inserting into database
			SetImageTags(imageMetadata.ID, tags)
		} else {
			panic(err.Error())
		}
	}
}

// routeUpload handles post requests to '/upload'
func routeUpload(c *gin.Context) {

	// Get the user associated with the request's JWT
	user, _ := GetUserFromJWT(c.GetHeader("Authorization"))

	// Get the metadata of the uploaded images from the request body
	meta, err := getImageUploadMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid upload metadata"},
		)
		return
	}

	// Save the uploaded images to the database
	saveImages(c, meta, &user)

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successful"},
	)
}
