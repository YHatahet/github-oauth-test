package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
