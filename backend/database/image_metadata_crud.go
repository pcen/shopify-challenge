package database

import (
	"fmt"
	"image-repo/core"
	"strings"
)

// CRUD operations for image metadata

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

// InsertImage inserts the given ImageMetadata model into the database.
func InsertImage(metadata *ImageMetadata) error {
	return DB.Model(ImageMetadata{}).Create(metadata).Error
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

// SetImageTags sets the tags for the given image ID based off of the list of
// given tag structs.
func SetImageTags(id uint, tags []core.ImageTag) error {
	var sb strings.Builder
	for _, tag := range tags {
		sb.WriteString(tag.Tag.Value)
		sb.WriteString(" ")
	}
	result := DB.Model(ImageMetadata{}).Where("id = ?", id).Update("ml_tags", sb.String())
	return result.Error
}

// SearchQueryImages returns the metadata for images matching the given
// search query string for the given user ID.
func SearchQueryImages(user uint, query string, public bool) ([]ImageMetadata, error) {
	var metadata []ImageMetadata
	var subQuery = DB.Model(ImageMetadata{})

	fuzzyQuery := fmt.Sprintf("%s%s%s", "%", query, "%")

	subQuery.Where("name LIKE ?", fuzzyQuery)
	subQuery.Or("geolocation LIKE ?", fuzzyQuery)

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
