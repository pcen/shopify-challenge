package models

import (
	"gorm.io/gorm"
)

// ImageMetadata Model
type ImageMetadata struct {
	gorm.Model

	UserID         uint   // Foreign key to the image owner
	Name           string // Image name
	Filepath       string `gorm:"unique"` // Path to the image file
	Description    string // A description of the image
	Geolocation    string // The image's geolocation
	OCRText        string // OCR text from image
	Public         bool   // Image visibility (public or private)
	AverageHash    uint64 // Perceptual hash
	DifferenceHash uint64 // Perceptual hash
}

// ImageUploadMeta Model
type ImageUploadMeta struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Private     bool   `json:"private" binding:"required"`
	Type        string `json:"type" binding:"required"`
}
