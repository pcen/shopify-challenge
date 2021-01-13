package database

import (
	"fmt"
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

// InsertImage inserts the given ImageMetadata model into the database.
func InsertImage(metadata *ImageMetadata) error {
	return DB.Model(ImageMetadata{}).Create(metadata).Error
}

// ownsImage returns true if the given user ID owns the given image ID.
func ownsImage(image uint, user uint) bool {
	var metadata ImageMetadata
	result := DB.Model(ImageMetadata{}).Where("id = ? AND user_id = ?", image, user).First(&metadata)
	return result.Error == nil
}

// imageIsPublic returns true if the given image ID is a public image.
func imageIsPublic(image uint) bool {
	var metadata ImageMetadata
	result := DB.Model(ImageMetadata{}).Where("id = ?", image).First(&metadata)
	return (result.Error == nil) && (!metadata.Private)
}

// GetImage gets the ImageMetada corresponding to the given id if the requestee
// owns the image, or the image visibility is public.
func GetImage(id uint, user uint) (ImageMetadata, error) {
	var metadata ImageMetadata
	result := DB.Model(ImageMetadata{}).Where("id = ? AND (user_id = ? OR private = ?)", id, user, false).First(&metadata)
	return metadata, result.Error
}

// UpdateImage updates the given image to match the metadata parameter if the
// requestee owns the image.
func UpdateImage(metadata *ImageMetadata, user uint) error {
	if !ownsImage(metadata.ID, user) {
		return fmt.Errorf("requestee does not own image %d", metadata.ID)
	}
	result := DB.Model(ImageMetadata{}).Where("id = ?", metadata.ID).Updates(metadata)
	return result.Error
}

// DeleteImage deletes the given ImageMetadata model from the database. Returns
// an error if the requestee does not own the requested image.
func DeleteImage(id uint, user uint) error {
	if !ownsImage(id, user) {
		return fmt.Errorf("requestee does not own image %d", id)
	}
	metadata, _ := GetImage(id, user)
	return DB.Model(ImageMetadata{}).Delete(&metadata).Error
}

// DeleteImageFileStore deletes the given filename parameter from the image
// file storage.
func DeleteImageFileStore(filename string) error {
	path := GetFileStoreFullPath(filename)
	return core.DeleteFile(path)
}

// GetImageFileStore returns the store filepath of the image with the given ID
// if it is owned by the given user ID or it is a public image.
func GetImageFileStore(id uint, user uint) (string, error) {
	var metadata ImageMetadata
	metadata, err := GetImage(id, user)
	return metadata.FileStore, err
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
