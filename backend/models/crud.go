package models

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

// SearchQueryImageMetadata returns the metadata for images matching the given
// search query string for the given user ID.
func SearchQueryImageMetadata(user uint, query string, public bool) ([]ImageMetadata, error) {
	var metadata []ImageMetadata
	if len(query) == 0 {
		result := DB.Model(ImageMetadata{}).Where("user_id = ?", user).Find(&metadata)
		return metadata, result.Error
	} else {
		result := DB.Model(ImageMetadata{}).Where("user_id = ?", user).Find(&metadata)
		return metadata, result.Error
	}
}

// InsertImageMetadata inserts the given ImageMetadata model into the database.
func InsertImageMetadata(metadata *ImageMetadata) error {
	result := DB.Model(ImageMetadata{}).Create(metadata)
	return result.Error
}
