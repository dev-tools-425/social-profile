package middleware

import (
	"errors"
	"time"

	"github.com/dev-tools-425/social-profile/config"
	"github.com/golang-jwt/jwt/v5"
)

// JWT middleware for generating and validating tokens
var jwtSecret = []byte(config.JWTSecret)

// GenerateJWT creates token for authenticated user
func GenerateJWT(userID string) (string, error) {

	if string(jwtSecret) == "" {
		return "", errors.New("JWT_SECRET not set")
	}
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
