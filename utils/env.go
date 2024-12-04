package utils

import (
	"log"
	"os"
)

// GetGithubClientID fetches the GitHub client ID from the environment
func GetGithubClientID() string {
	clientID, exists := os.LookupEnv("CLIENT_ID")
	if !exists {
		log.Fatal("GitHub Client ID not defined in .env file")
	}
	return clientID
}

// GetGithubClientSecret fetches the GitHub client secret from the environment
func GetGithubClientSecret() string {
	clientSecret, exists := os.LookupEnv("CLIENT_SECRET")
	if !exists {
		log.Fatal("GitHub Client Secret not defined in .env file")
	}
	return clientSecret
}
