package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dev-tools-425/social-profile/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// these should be in env vars in production, hardcoded here for simplicity
var googleConfig = &oauth2.Config{
	ClientID:     config.GoogleClientID,
	ClientSecret: config.GoogleClientSecret,
	RedirectURL:  config.GoogleCallbackURL,
	Scopes:       []string{"email", "profile"},
	Endpoint:     google.Endpoint,
}

// Redirect user to Google login
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleConfig.AuthCodeURL("randomstate")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Handle Google callback
func GoogleCallback(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := googleConfig.Client(context.Background(), token)

	resp, _ := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	var userInfo map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&userInfo)

	json.NewEncoder(w).Encode(userInfo)
}

// Microsoft OAuth2 configuration
var microsoftConfig = &oauth2.Config{
	ClientID:     config.MSClientID,
	ClientSecret: config.MSClientSecret,
	RedirectURL:  config.MSCallbackURL,
	Scopes:       []string{"User.Read"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
		TokenURL: "https://login.microsoftonline.com/common/oauth2/v2.0/token",
	},
}

// Redirect user to Microsoft login
func MicrosoftLogin(w http.ResponseWriter, r *http.Request) {
	url := microsoftConfig.AuthCodeURL("randomstate")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Handle Microsoft callback
func MicrosoftCallback(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")

	token, err := microsoftConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := microsoftConfig.Client(context.Background(), token)

	resp, _ := client.Get("https://graph.microsoft.com/v1.0/me")

	var userInfo map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&userInfo)

	json.NewEncoder(w).Encode(userInfo)
}
