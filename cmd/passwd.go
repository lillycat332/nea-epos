package main

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password (type string) with bcrypt and returns a string containing the hashed password, error
func HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

// checkPasswordMatch tests a hashedPassword (string) and an unhashed password (string) with bcrypt, returning true if they match.
func checkPasswordMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currPassword))
	return err == nil
}
