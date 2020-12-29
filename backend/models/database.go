package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// populateTestData adds test data to the database
func populateTestData(db *gorm.DB) {
	db.Create(&User{Username: "Paul", PasswordHash: "12345", Role: Admin})
	db.Create(
		&ImageMetadata{
			UserID: 1,
			Filepath: "file.jpg",
			Description: "This is an image description.",
			Geolocation: "Paris, France",
			OCRText: "This is OCR text.",
			Public: true,
			AverageHash: 12345,
			DifferenceHash: 54321,
		},
	)
}

const databaseName = "images.db"

// GetDatabaseHandle returns the database handle
func GetDatabaseHandle() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&User{}, &ImageMetadata{})

	return db
}
