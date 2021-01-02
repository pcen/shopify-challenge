package models

// This file implements common CRUD operations for database models.

// GetUser returns the User in the database with the given username
func GetUser(username string) (User, error) {
	var user User
	result := DB.Model(&User{}).Where("username = ?", username).First(&user)
	return user, result.Error
}

// InsertImageMetadata inserts the given ImageMetadata model into the database.
func InsertImageMetadata(metadata *ImageMetadata) error {
	result := DB.Model(ImageMetadata{}).Create(metadata)
	return result.Error
}
