package config

import "os"

var JWTSecret string

func LoadConfig() {
	JWTSecret = os.Getenv("JWT_SECRET")
}
