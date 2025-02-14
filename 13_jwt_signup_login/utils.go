package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(user *User) (string, error) {
	secret := []byte("super-secret-key")
	method := jwt.SigningMethodHS256
	claims := jwt.MapClaims{
		"userId":   user.ID,
		"userName": user.Username,
		"exp":      time.Now().Add(time.Hour * 168).Unix(), //7 day
	}
	token, err := jwt.NewWithClaims(method, claims).SignedString(secret)
	if err != nil {
		return "", err
	}
	return token, nil
}
