package models

import (
	"gorm.io/gorm"
)

// ImageMetadata Database Model
type ImageMetadata struct {
	gorm.Model

	UserID         uint   // Foreign key to the image owner
	Name           string // Image name
	Format         string // Image format
	FileStore      string `gorm:"unique"` // Filename of image in database
	Description    string // A description of the image
	Geolocation    string // The image's geolocation
	OCRText        string // OCR text from image
	Private        bool   // Image visibility (public or private)
	AverageHash    uint64 // Perceptual hash
	DifferenceHash uint64 // Perceptual hash
}

// ImageUploadMeta JSON Model
type ImageUploadMeta struct {
	Name        string `json:"name" binding:"required"`
	Format      string `json:"format" binding:"required"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Private     bool   `json:"private" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

// ImageDownloadMeta JSON Model
type ImageDownloadMeta struct {
	Image       uint `json:"image" binding:"required"`
}
