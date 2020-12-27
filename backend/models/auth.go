package models

// Login data
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// User session data
type User struct {
	Username string `json:"username" binding:"required"`
	AuthToken string `json:"authToken" binding:"required"`
}
