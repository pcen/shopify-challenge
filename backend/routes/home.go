package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// routeHome handles get requests to '/'
func routeHome(c *gin.Context) {
	c.String(http.StatusOK, "home")
}
