package database

import (
	"fmt"
	"path/filepath"
	"runtime"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"image-repo/core"
)

// DB is the database handle for backend CRUD functions.
var DB *gorm.DB

// populateTestData adds test data to the database
func populateTestData(db *gorm.DB) {
	err := CreateUser("admin", "password")
	if err != nil {
		fmt.Println(err.Error())
	}
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

// GetImagesDir returns the absolute path to the directory containing all of
// the image files in the image repository. It will create the directory if it
// does not already exist.
func GetImagesDir() string {
	imagesDir := filepath.Join(getDatabaseDir(), "images")
	core.EnsureDirExists(imagesDir)
	return imagesDir
}

// GetFileStoreFullPath returns the absolute path to the file in the filestore
// with the given filename parameter.
func GetFileStoreFullPath(filename string) string {
	imagesDir := GetImagesDir()
	return filepath.Join(imagesDir, filename)
}

// InitializeDatabase sets
func InitializeDatabase() {
	dbPath := getDatabaseFilepath()
	imagesDir := GetImagesDir()

	fmt.Println(dbPath)
	fmt.Println(imagesDir)

	// reset database state
	// core.DeleteAllFilesInDirectory(imagesDir)
	// core.DeleteFile(dbPath)

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	DB.AutoMigrate(&User{}, &ImageMetadata{})

	populateTestData(DB)
}
