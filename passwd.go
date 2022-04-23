package main

import (
	"golang.org/x/crypto/bcrypt"
)

// hash password with bcrypt
func hashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

// test our password for match with bcrypt
func doPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currPassword))
	return err == nil
}
