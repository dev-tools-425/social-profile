package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleConfig = &oauth2.Config{
	ClientID:     "GOOGLE_CLIENT_ID",
	ClientSecret: "GOOGLE_CLIENT_SECRET",
	RedirectURL:  "http://localhost:8080/auth/google/callback",
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
