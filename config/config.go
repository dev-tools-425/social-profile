package config

import "os"

var JWTSecret string
var GoogleClientID string
var GoogleClientSecret string
var GoogleCallbackURL string
var MSClientID string
var MSClientSecret string
var MSCallbackURL string

func LoadConfig() {
	JWTSecret = os.Getenv("JWT_SECRET")
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	GoogleCallbackURL = os.Getenv("GOOGLE_CALLBACK_URL")
	MSClientID = os.Getenv("MS_CLIENT_ID")
	MSClientSecret = os.Getenv("MS_CLIENT_SECRET")
	MSCallbackURL = os.Getenv("MS_CALLBACK_URL")
}
