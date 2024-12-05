package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// RootHandler handles the root route
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<a href="/login/github/">LOGIN</a>`)
}

// LoggedInHandler handles the logged-in route
func LoggedInHandler(w http.ResponseWriter, githubData string) {
	if githubData == "" {
		fmt.Fprintf(w, "UNAUTHORIZED!")
		return
	}

	// Set return type JSON
	w.Header().Set("Content-Type", "application/json")

	// Prettify JSON
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(githubData), "", "\t"); err != nil {
		log.Panic("JSON parse error")
	}

	// Write response
	w.Write(prettyJSON.Bytes())
}

func isValidEmail(email string) bool {
	// Define a regular expression for validating an email.
	// This regex is simple and covers most valid email cases.
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// Check if the email matches the regex
	return emailRegex.MatchString(email)
}

func LoggedInHandlerEmail(w http.ResponseWriter, githubEmail string) {
	if githubEmail == "" || !isValidEmail(githubEmail) {
		fmt.Fprintf(w, "UNAUTHORIZED!")
		return
	}

	// Set return type JSON
	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(githubEmail))
	// Write response
	// w.Write(prettyJSON.Bytes())
}
