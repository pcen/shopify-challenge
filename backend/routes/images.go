package routes

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"

	. "image-repo/database"
)

// invalidImageID returns an "invalid image id" error as JSON for the given
// gin context.
func invalidImageID(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "invalid image id",
	})
}

// routeImage handles get requests to '/image/<id>'
// Returns the image file for the image ID specified in the URL if the image is
// either owned by the requestee or is a public image.
func routeImage(c *gin.Context) {
	// Get the requested image ID from the endpoint
	imageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		invalidImageID(c)
		return
	}

	// Get the user associated with the request's JWT
	user, err := GetUserFromJWT(c.GetHeader("Authorization"))

	// Get the filename of the image in the database
	filename, err := GetImageFilepath(uint(imageID), user.ID)
	if err != nil {
		invalidImageID(c)
		return
	}

	// Return the image file
	imageFile := filepath.Join(GetImagesDir(), filename)
	c.File(imageFile)
	c.Status(http.StatusOK)
}

// routeImageChange handles post requests to '/image/<id>/edit'
// Edits the image metadata for the image ID specified in the URL based off of
// the request body content if the image is owned by the requestee.
func routeImageEdit(c *gin.Context) {
	// Get the requested image ID from the endpoint
	imageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		invalidImageID(c)
		return
	}

	// Get the requested changes to be made from the request body
	var body ImageMetadata
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// Get the user associated with the request's JWT
	user, _ := GetUserFromJWT(c.GetHeader("Authorization"))

	// Update the image metadata
	if err := UpdateImage(&body, user.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("cannot change image %d", imageID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("changed image %d", imageID),
	})
}

// routeImageDelete handles delete requests to '/image/<id>/delete'
// Returns a message if the image is successfully deleted, otherwise returns an
// error if the image failed to be deleted. A deletion attempt will fail if the
// requestee does not own the image specified in the URL.
func routeImageDelete(c *gin.Context) {
	imageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		invalidImageID(c)
		return
	}

	// Get the user associated with the request's JWT
	user, _ := GetUserFromJWT(c.GetHeader("Authorization"))

	// Get the path to the image file in the database
	imageFile, err := GetImageFilepath(uint(imageID), user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("cannot delete image %d", imageID),
		})
		return
	}

	// First check to see if the user may delete the image by attempting to
	// delete the image metadata.
	err = DeleteImage(uint(imageID), user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("cannot delete image %d", imageID),
		})
		return
	}

	// If the user was permitted to delete the image metadata, delete the image
	// from the image file store.
	DeleteImageFile(imageFile)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("deleted image %d", imageID),
	})
}

// routeImageTags handles get requests to '/image/<id>/tags'
// Returns the auto-generated tags for the requested image if the requestee
// owns the image or the image is public.
func routeImageTags(c *gin.Context) {
	imageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		invalidImageID(c)
		return
	}

	// Get the user associated with the request's JWT
	user, _ := GetUserFromJWT(c.GetHeader("Authorization"))

	// Get the specified image's metadata
	image, err := GetImage(uint(imageID), user.ID)

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"tags": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tags": image.MLTags,
	})
}

// routeImages handles post requests to '/images'
// Accepts a query string in the request body and returns a list of image
// metadata found in the database as a result of the given query string.
func routeImages(c *gin.Context) {
	var body ImageQuery
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body"},
		)
		return
	}

	// Get the user associated with the request's JWT
	user, err := GetUserFromJWT(c.GetHeader("Authorization"))

	// Get the requested image's metadata from the database
	metadata, err := SearchQueryImages(user.ID, body.Query, body.IncludePublic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to retrieve images"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"images": metadata},
	)
}
