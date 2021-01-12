package core

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetSigningKey returns the signing key for JWTs.
func GetSigningKey() []byte {
	return []byte(os.Getenv("IMAGE_REPO_SIGNING_KEY"))
}

// NewToken returns a new signed JWT string
func NewToken(username string) (string, string) {
	errorMsg := ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenStr, err := token.SignedString(GetSigningKey())
	if err != nil {
		fmt.Println(err)
		errorMsg = err.Error()
	}
	return tokenStr, errorMsg
}

// parseToken returns a JWT from the given token string.
func parseToken(authToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method %v", token.Header["alg"])
		}
		return GetSigningKey(), nil
	})
	return token, err
}

// TokenValid returns true if the passed JWT string is valid
func TokenValid(authToken string) bool {
	token, err := parseToken(authToken)
	if err != nil {
		return false
	}
	_, claimsOk := token.Claims.(jwt.MapClaims)
	return err == nil && claimsOk && token.Valid
}

// RequestTokenValid returns true if the request in the passed gin Context
// supplies a valid authorization token, and returns false if it does not.
// This method is only used to check validity of JWTs, not if the requester
// has sufficient permissions to access a resource.
func RequestTokenValid(c *gin.Context) (bool, string) {
	if c.GetHeader("Authorization") == "" {
		return false, "token missing"
	}
	valid := TokenValid(c.GetHeader("Authorization"))
	if !valid {
		return valid, "token invalid"
	}
	return valid, ""
}

// GetTokenUser returns the username of the user to whom a JWT was issued.
func GetTokenUser(authToken string) (string, error) {
	token, err := parseToken(authToken)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		return claims["sub"].(string), nil
	}
	return "", fmt.Errorf("failed to extract claims from jwt")
}
