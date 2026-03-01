package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("supersecret")

// GenerateJWT creates token for authenticated user
func GenerateJWT(userID string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
