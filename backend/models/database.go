package models

import (
	"fmt"
	"path/filepath"
	"runtime"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"image-repo/core"
)

// populateTestData adds test data to the database
func populateTestData(db *gorm.DB) {
	db.Create(&User{Username: "Paul", PasswordHash: "12345", Role: Admin})
	db.Create(
		&ImageMetadata{
			UserID:         1,
			Filepath:       "file.jpg",
			Description:    "This is an image description.",
			Geolocation:    "Paris, France",
			OCRText:        "This is OCR text.",
			Public:         true,
			AverageHash:    12345,
			DifferenceHash: 54321,
		},
	)
}

// getDatabaseDir returns the absolute path to the database directory. It will
// create the directory if it does not already exist.
func getDatabaseDir() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get caller information")
	}
	// the project's root directory is 3 levels up from files inside the models
	// directory
	rootDir := core.ParentDir(file, 3)
	dbDir := filepath.Join(rootDir, "database")
	core.EnsureDirExists(dbDir)
	return dbDir
}

// getDatabaseFilepath returns the absolute path to the SQL database containing
// user data and image metadata. If the database file does not exist, it will
// be created when the gorm connection to the database is opened.
func getDatabaseFilepath() string {
	return filepath.Join(getDatabaseDir(), "metadata.db")
}

// getImagesDir returns the absolute path to the directory containing all of
// the image files in the image repository. It will create the directory if it
// does not already exist.
func getImagesDir() string {
	imagesDir := filepath.Join(getDatabaseDir(), "images")
	core.EnsureDirExists(imagesDir)
	return imagesDir
}

// GetDatabaseHandle returns the database handle.
func GetDatabaseHandle() *gorm.DB {
	dbPath := getDatabaseFilepath()
	imagesDir := getImagesDir()

	fmt.Println(dbPath)
	fmt.Println(imagesDir)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&User{}, &ImageMetadata{})

	populateTestData(db)

	return db
}
