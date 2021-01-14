package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

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
func saveImages(c *gin.Context, meta []ImageUploadMeta, user *User) []uint {
	var savedImages []uint
	for _, m := range meta {
		// Generate the name of the image file in the file registry
		store := core.RandomAlphanumericString(32) + core.FileExtensionFromFormat(m.Format)

		// Insert image metadata into SQL database
		imageMetadata := ImageMetadata{
			UserID:      user.ID,
			Name:        m.Name,
			Format:      m.Format,
			FileStore:   store,
			Description: m.Description,
			Geolocation: m.Location,
			MLTags:      "",
			Private:     m.Private,
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

		savedImages = append(savedImages, imageMetadata.ID)
	}

	// Return a list of all the saved image IDs
	// These IDs are used in order to identify which images to tag.
	return savedImages
}

// tagImages tags the given image IDs. The user ID is passed to make sure that
// the image files are uploaded to the tagging API only if the requestee owns
// the images.
func tagImages(images []uint, user uint) {
	for index, id := range images {
		// Get the path to the image file in the images directory
		imageFile, _ := GetImageFilepath(id, user)
		fullPath := filepath.Join(GetImagesDir(), imageFile)
		// Get image tags
		tags, err := core.GetImageTags(fullPath)
		if err == nil {
			// Update image tags in metadata database
			SetImageTags(id, tags)
		} else {
			fmt.Println(err.Error())
		}
		// Wait 1 second between tagging each image since API is rate limited
		if index != len(images)-1 {
			time.Sleep(1 * time.Second)
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
	savedImages := saveImages(c, meta, &user)

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successful"},
	)

	// Tag images once upload is completed
	// Tagging images makes use of a rate limited API so it must be done after
	// returning a reply to the client, otherwise the rate limited API requests
	// will force the client to wait longer than necessary for a responce upon
	// uploading.
	go tagImages(savedImages, user.ID)
}
