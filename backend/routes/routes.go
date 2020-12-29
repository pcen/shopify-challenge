package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"image-repo/core"
)

// routeStatus always returns 200 if the application is running
func routeStatus(c *gin.Context) {
	c.String(http.StatusOK, "running")
}

// routeHome handles get requests to '/'
func routeHome(c *gin.Context) {
	c.String(http.StatusOK, "home")
}

// routeCheckAuthToken handles get requests to '/check-auth' and responds
// indicating if the request was made with a valid JWT or not.
func routeCheckAuthToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"valid": core.RequestTokenValid(c)})
}

// authMiddleware checks if incoming requests are authorized and will reject
// unauthorized requests.
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// return an unauthorized response if the request is not authorized
		if !core.RequestTokenValid(c) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalid",
			})
			c.Abort()
			return
		}
		// continue handling authorized requests
		c.Next()
	}
}

// AttachAll adds all route definitions to the passed gin Engine.
func AttachAll(app *gin.Engine) {
	// Public endpoints
	app.GET("/", routeHome)
	app.GET("/status", routeStatus)
	app.GET("/check-auth", routeCheckAuthToken)
	app.POST("/login", routeLogin)

	// Authorization protected endpoints
	app.GET("/images", authMiddleware(), routeImages)
}
