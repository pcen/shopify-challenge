package core

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// PasswordSaltAndHash creates a salted hash of the given password.
func PasswordSaltAndHash(password string) string {
	passwordBytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}

// PasswordEqualsHashed returns true if the plain text password yields the same
// salted hash as the given bcrypt hash.
func PasswordEqualsHashed(password string, hash string) bool {
	passwordBytes := []byte(password)
	hashBytes := []byte(hash)
	err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)
	if err != nil {
		fmt.Printf(
			"Password %s does not match bcrypt hash %s",
			password,
			hash,
		)
		return false
	}
	return true
}
