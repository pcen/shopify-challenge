package database

import (
	"image-repo/core"
)

// CRUD operations for users

// GetUserFromUsername returns the User in the database with the given username
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
