package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"image-repo/core"
	. "image-repo/database"
)

// validLogin returns true, along with the user's ID if the username password
// combination is a valid set of credentials.
func validLogin(username string, password string) (bool, uint) {
	user, err := GetUserFromUsername(username)
	if err != nil {
		fmt.Println("Login user query failed with error", err.Error())
		return false, 0
	}
	fmt.Printf(
		"Login user query succeeded, comparing DB password %s to login password %s\n",
		user.PasswordHash,
		password,
	)
	if core.PasswordEqualsHashed(password, user.PasswordHash) {
		return true, user.ID
	}
	return false, 0
}

// routeLogin handles post requests to '/login'
// Returns a JWT if the credentials in the request body are valid.
func routeLogin(c *gin.Context) {
	var body UserLogin
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	valid, userID := validLogin(body.Username, body.Password)
	authToken := ""
	err := "invalid credentials"

	if valid {
		// Create JWT
		authToken, err = core.NewToken(body.Username)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": valid,
		"user": UserSession{
			Username:  body.Username,
			AuthToken: authToken,
			ID:        userID,
		},
		"error": err,
	})
}

// routeCreateUser handles post requests to '/create-user'
// Creates a new user with the credentials given in the request body.
func routeCreateUser(c *gin.Context) {
	var body UserLogin
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// Return an error message if the given username already exists
	if UsernameExists(body.Username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username and password combination",
		})
		return
	}

	// Create a new user with the given credentials
	CreateUser(body.Username, body.Password)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("user %s created", body.Username),
	})
}
