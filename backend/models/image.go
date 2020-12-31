package models

import (
	"mime/multipart"

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

// ImageMultipart Form Model
type ImageMultipart struct {
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	Description string                `form:"description"`
	Location    string                `form:"location"`
	Private     bool                  `form:"private" binding:"required"`
}
