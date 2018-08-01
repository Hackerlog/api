package main

import (
	"golang.org/x/crypto/bcrypt"
)

// GenericResponse This is an all purpose, generice response used throughout the app
type GenericResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// HashPassword Hashes a password and returns the hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword Check that the user has entered the correct password
func CheckPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
