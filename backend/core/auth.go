package core

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func getSigningKey() []byte {
	return []byte("secret")
}

// NewToken returns a new signed JWT string
func NewToken(username string) (string, string) {
	errorMsg := ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenStr, err := token.SignedString(getSigningKey())
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
		return getSigningKey(), nil
	})
	if err != nil {
		return false
	}
	_, claimsOk := token.Claims.(jwt.MapClaims)
	return err == nil && claimsOk && token.Valid
}
