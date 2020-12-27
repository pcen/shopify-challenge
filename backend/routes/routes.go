package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// routeStatus always returns 200 if the application is running
func routeStatus(c *gin.Context) {
	c.String(http.StatusOK, "running")
}

// AttachAll adds all route definitions to the passed gin Engine.
func AttachAll(app *gin.Engine) {
	app.GET("/", routeHome)
	app.GET("/status", routeStatus)
	app.POST("/login", routeLogin)
	app.GET("/images", routeImages)
}
