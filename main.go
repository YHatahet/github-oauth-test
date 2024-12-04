package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yhatahet/github-oauth-test/handlers"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	// Define routes
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/login/github/", handlers.GithubLoginHandler)
	http.HandleFunc("/login/github/callback", handlers.GithubCallbackHandler)
	http.HandleFunc("/loggedin", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoggedInHandler(w, "")
	})

	// Start server
	fmt.Println("[ UP ON PORT 3000 ]")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
