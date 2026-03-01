package routes

import (
	"net/http"

	"github.com/dev-tools-425/social-profile/handlers"
)

func RegisterRoutes() {

	http.HandleFunc("/auth/google", handlers.GoogleLogin)
	http.HandleFunc("/auth/google/callback", handlers.GoogleCallback)

	http.HandleFunc("/auth/microsoft", handlers.MicrosoftLogin)
	http.HandleFunc("/auth/microsoft/callback", handlers.MicrosoftCallback)

	http.HandleFunc("/profile/update", handlers.UpdateProfile)
	http.HandleFunc("/profile/upload", handlers.UploadProfileImage)
}
