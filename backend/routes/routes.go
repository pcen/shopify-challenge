package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"image-repo/core"
	"image-repo/database"
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
// unauthorized requests. It will also reject requests for which the JWT
// contains a user that is not registered in the database.
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// return an unauthorized response if the request is not authorized
		valid, errStr := core.RequestTokenValid(c)
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": errStr,
			})
			c.Abort()
			return
		}
		// return a bad request if the requestee is not a user in the database
		_, err := database.GetUserFromJWT(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid user",
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
	app.POST("/create-user", routeCreateUser)

	// Authorization protected endpoints
	app.POST("/images", authMiddleware(), routeImages)
	app.GET("/image/:id", authMiddleware(), routeImage)
	app.GET("/image/:id/tags", authMiddleware(), routeImageTags)
	app.POST("/image/:id/edit", authMiddleware(), routeImageEdit)
	app.DELETE("/image/:id/delete", authMiddleware(), routeImageDelete)
	app.POST("/upload", authMiddleware(), routeUpload)
}
