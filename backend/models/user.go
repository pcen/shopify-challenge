package models

import (
	"gorm.io/gorm"
)

// UserRole Enumeration
type UserRole int

const (
	None UserRole = iota
	RegularUser
	Admin
)

// User Model
type User struct {
	gorm.Model

	Username       string `gorm:"unique"`
	PasswordHash   string
	Role           UserRole
	// User has many ImageMetadata
	ImageMetadatas []ImageMetadata
}
