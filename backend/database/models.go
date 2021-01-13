package database

import (
	"time"
)

// ImageMetadata Database Model
type ImageMetadata struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
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

// ImageQuery JSON Model
type ImageQuery struct {
	Query         string `json:"query"`
	IncludePublic bool   `json:"includePublic"`
}

// UserRole Enumeration
type UserRole int

const (
	None UserRole = iota
	RegularUser
	Admin
)

// User Model
type User struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Username     string `gorm:"unique"`
	PasswordHash string
	Role         UserRole
	// User has many ImageMetadata
	ImageMetadatas []ImageMetadata
}

// UserLogin Model
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserSession Model
type UserSession struct {
	Username  string `json:"username" binding:"required"`
	AuthToken string `json:"authToken" binding:"required"`
}
