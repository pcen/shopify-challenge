package models

// UserLogin Model
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	PasswordHash string `json:"password" binding:"required"`
}

// UserSession Model
type UserSession struct {
	Username string `json:"username" binding:"required"`
	AuthToken string `json:"authToken" binding:"required"`
}
