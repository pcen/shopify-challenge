package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"image-repo/models"
)

// validLogin returns true if the username password combination
// is a valid set of credentials
func validLogin(username string, password string) bool {
	return username == "user" && password == "password"
}

// routeLogin handles post requests to '/login'
func routeLogin(c *gin.Context) {
	var body models.Login
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid := validLogin(body.Username, body.Password)
	authToken := ""
	err := "Invalid credentials"

	if valid {
		// Create JWT
		authToken, err = newToken(body.Username)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": valid,
		"user": models.User{
			Username:  body.Username,
			AuthToken: authToken,
		},
		"error": err,
	})
}
