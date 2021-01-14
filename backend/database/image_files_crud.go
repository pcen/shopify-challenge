package database

import (
	"path/filepath"
	"mime/multipart"

	"image-repo/core"
)

// CRUD operations for image files

// DeleteImageFile deletes the given filename parameter from the image
// file storage. Unlike metadata CRUD operations, this does not check if the
// requestee of this operation owns the file, and permission to delete an image
// should be established first by deleting the image metadata.
func DeleteImageFile(filename string) error {
	path := GetFileStoreFullPath(filename)
	return core.DeleteFile(path)
}

// WriteImageFile creates a file with the given filename inside of the image
// file store and writes the given multipart file.
func WriteImageFile(filename string, file multipart.File) error {
	fullPath := filepath.Join(GetImagesDir(), filename)
	return core.WriteFile(fullPath, file)
}

// GetImageFilepath returns the store filepath of the image with the given ID
// if it is owned by the given user ID or it is a public image.
func GetImageFilepath(id uint, user uint) (string, error) {
	var metadata ImageMetadata
	metadata, err := GetImage(id, user)
	return metadata.FileStore, err
}
