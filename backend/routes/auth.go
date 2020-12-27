package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"image-repo/models"
)

// routeLogin handles post requests to '/login'
func routeLogin(c *gin.Context) {
	var body models.Login
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid := body.Username == "user" && body.Password == "password"

	c.JSON(http.StatusOK, gin.H{
		"success": valid,
		"message": fmt.Sprintf("Recieved login: Username: %s, Password: %s", body.Username, body.Password),
	})
}
