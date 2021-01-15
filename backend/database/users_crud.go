package database

import (
	"image-repo/core"
)

// CRUD operations for users

// CreateUser adds a new user to the database.
func CreateUser(username string, password string) error {
	result := DB.Create(&User{
		Username: username,
		PasswordHash: core.PasswordSaltAndHash(password),
		Role: RegularUser,
	})
	return result.Error
}

// GetUserFromUsername returns the User in the database with the given
// username.
func GetUserFromUsername(username string) (User, error) {
	var user User
	result := DB.Model(&User{}).Where("username = ?", username).First(&user)
	return user, result.Error
}

// GetUserFromJWT returns the User in the database associated with the given
// JWT.
func GetUserFromJWT(authToken string) (User, error) {
	var user User
	username, err := core.GetTokenUser(authToken)
	if err != nil {
		return user, err
	}
	user, err = GetUserFromUsername(username)
	return user, err
}

// UsernameExists returns true if the given username is taken, and returns
// false otherwise.
func UsernameExists(username string) bool {
	_, err := GetUserFromUsername(username)
	// If there is no error getting a user with the given username, then the
	// username is taken.
	return err == nil
}
