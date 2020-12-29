package core

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// signingKey is the secret key used to sign JWTs.
const signingKey = "secret"

// GetSigningKey returns the signing key for JWTs.
func GetSigningKey() []byte {
	return []byte(signingKey)
}

// NewToken returns a new signed JWT string
func NewToken(username string) (string, string) {
	errorMsg := ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenStr, err := token.SignedString(GetSigningKey())
	if err != nil {
		fmt.Println(err)
		errorMsg = err.Error()
	}
	return tokenStr, errorMsg
}

// TokenValid returns true if the passed JWT string is valid
func TokenValid(authToken string) bool {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method %v", token.Header["alg"])
		}
		return GetSigningKey(), nil
	})
	if err != nil {
		return false
	}
	_, claimsOk := token.Claims.(jwt.MapClaims)
	return err == nil && claimsOk && token.Valid
}
