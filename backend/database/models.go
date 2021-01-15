package database

import (
	"time"
)

// ImageMetadata Database Model
// An image in the database. The actual image is stored as a file
// in the /database/images folder. The file is referenced by the FileStore
// member specifying the image's filename in the images folder.
type ImageMetadata struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uint   // Foreign key to the image owner
	Name        string // Image name
	Format      string // Image format
	FileStore   string `gorm:"unique"` // Filename of image in database
	Description string // A description of the image
	Geolocation string // The image's geolocation
	MLTags      string // Image tags generated from ML model
	Private     bool   // Image visibility (public or private)
}

// ImageUploadMeta JSON Model
// The metadata supplied for an image when uploaded by a user.
type ImageUploadMeta struct {
	Name        string `json:"name" binding:"required"`
	Format      string `json:"format" binding:"required"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Private     bool   `json:"private" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

// ImageQuery JSON Model
// The data supplied when a user submits an image search.
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
// The data supplied when a user attempts to login.
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserSession Model
// The data returned to the user and stored in their browser to make
// requests to endpoints requiring authorization (and authentication for image
// image ownership).
type UserSession struct {
	Username  string `json:"username" binding:"required"`
	AuthToken string `json:"authToken" binding:"required"`
	ID        uint   `json:"id"`
}
