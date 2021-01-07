package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"image-repo/core"
)

// routeHome handles get requests to '/'
func routeHome(c *gin.Context) {
	c.String(http.StatusOK, "image-repo backend running")
}

// routeCheckAuthToken handles get requests to '/check-auth' and responds
// indicating if the request was made with a valid JWT or not.
func routeCheckAuthToken(c *gin.Context) {
	valid, _ := core.RequestTokenValid(c)
	c.JSON(http.StatusOK, gin.H{"valid": valid})
}

// authMiddleware checks if incoming requests are authorized and will reject
// unauthorized requests.
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// return an unauthorized response if the request is not authorized
		valid, err := core.RequestTokenValid(c)
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
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
	app.GET("/check-auth", routeCheckAuthToken)
	app.POST("/login", routeLogin)

	// Authorization protected endpoints
	app.POST("/images", authMiddleware(), routeImages)
	app.GET("/image/:id", authMiddleware(), routeImage)
	app.POST("/upload", authMiddleware(), routeUpload)
}
