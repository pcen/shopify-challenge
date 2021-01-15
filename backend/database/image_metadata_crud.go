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
	result := DB.Model(ImageMetadata{}).Where("id = ?", metadata.ID)
	result.Update("description", metadata.Description)
	result.Update("geolocation", metadata.Geolocation)
	result.Update("private", metadata.Private)
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
	for index, tag := range tags {
		sb.WriteString(tag.Tag.Value)
		if index != len(tags)-1 {
			sb.WriteString(",")
		}
	}
	result := DB.Model(ImageMetadata{}).Where("id = ?", id).Update("ml_tags", sb.String())
	return result.Error
}

// SearchQueryImages returns the metadata for images matching the given
// search query string for the given user ID.
func SearchQueryImages(user uint, query string, public bool) ([]ImageMetadata, error) {
	var metadata []ImageMetadata
	var subQuery = DB.Model(ImageMetadata{})
	fuzzyQuery := fmt.Sprintf("%%%s%%", query)

	// Query for matching image name
	subQuery.Where("name LIKE ?", fuzzyQuery)
	// Query for matching image location
	subQuery.Or("geolocation LIKE ?", fuzzyQuery)
	// Query for single matching image tag
	subQuery.Or("ml_tags LIKE ?", fuzzyQuery)
	// Query for matching description
	subQuery.Or("description LIKE ?", fuzzyQuery)

	// Query metadata tags by a comma separated list of query tags
	tagList := strings.Split(query, ",")
	for _, tag := range tagList {
		fuzzyTag := fmt.Sprintf("%%%s%%", strings.TrimSpace(tag))
		subQuery.Or("ml_tags LIKE ?", fuzzyTag)
	}

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
