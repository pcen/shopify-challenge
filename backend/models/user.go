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
