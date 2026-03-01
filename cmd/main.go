package main

import (
	"log"
	"net/http"

	"github.com/dev-tools-425/social-profile/config"
	"github.com/dev-tools-425/social-profile/routes"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	config.LoadConfig()
	routes.RegisterRoutes()

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
