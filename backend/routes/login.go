package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"image-repo/core"
	. "image-repo/database"
)

// validLogin returns true if the username password combination
// is a valid set of credentials
func validLogin(username string, password string) bool {
	user, err := GetUserFromUsername(username)
	if err != nil {
		fmt.Println("Login user query failed with error", err.Error())
		return false
	}
	fmt.Printf(
		"Login user query succeeded, comparing DB password %s to login password %s\n",
		user.PasswordHash,
		password,
	)
	return core.PasswordEqualsHashed(password, user.PasswordHash)
}

// routeLogin handles post requests to '/login'
func routeLogin(c *gin.Context) {
	var body UserLogin
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid := validLogin(body.Username, body.Password)
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
		},
		"error": err,
	})
}
