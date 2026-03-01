package handlers

import (
	"encoding/json"
	"net/http"
)

// UpdateProfile updates user profile
func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	var profile map[string]string

	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here update DB (mock response)
	profile["message"] = "Profile updated successfully"

	json.NewEncoder(w).Encode(profile)
}

func UploadProfileImage(w http.ResponseWriter, r *http.Request) {

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save file locally (example only)
	w.Write([]byte("Image uploaded successfully"))
}
