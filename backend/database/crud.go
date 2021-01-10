package database

import (
	"image-repo/core"
)

// This file implements common CRUD operations for database models.

// GetUser returns the User in the database with the given username
func GetUser(username string) (User, error) {
	var user User
	result := DB.Model(&User{}).Where("username = ?", username).First(&user)
	return user, result.Error
}

// GetUserFromJWT returns the User in the database associated with the given
// JWT.
func GetUserFromJWT(authToken string) (User, error) {
	var user User
	username, err := core.GetTokenUser(authToken)
	if err != nil {
		return user, err
	}
	user, err = GetUser(username)
	return user, err
}

// GetImageFileStore returns the store filepath of the image with the given ID
// if it is owned by the given user ID or it is a public image.
func GetImageFileStore(id uint, user uint) (string, error) {
	var metadata ImageMetadata
	result := DB.Model(ImageMetadata{}).Where("id = ? AND user_id = ?", id, user).First(&metadata)
	return metadata.FileStore, result.Error
}

// SearchQueryImages returns the metadata for images matching the given
// search query string for the given user ID.
func SearchQueryImages(user uint, query string, public bool) ([]ImageMetadata, error) {
	var metadata []ImageMetadata
	var subQuery = DB.Model(ImageMetadata{})

	subQuery.Where("name LIKE ?", "%"+query+"%")

	// Include public images in the query if specified in the request
	result := DB.Table("(?) as sq", subQuery)
	if public {
		result.Where("user_id = ? OR private = ?", user, false)
	} else {
		result.Where("user_id = ?", user)
	}

	// Return matching image metadata
	result.Find(&metadata)
	return metadata, result.Error
}

// InsertImage inserts the given ImageMetadata model into the database.
func InsertImage(metadata *ImageMetadata) error {
	return DB.Model(ImageMetadata{}).Create(metadata).Error
}
