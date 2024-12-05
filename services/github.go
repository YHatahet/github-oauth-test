package services

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/yhatahet/github-oauth-test/utils"
)

// AccessTokenRequestBody represents the request body for GitHub access token
type AccessTokenRequestBody struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

// GithubAccessTokenResponse represents the response from GitHub access token API
type GithubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

// GetGithubAccessToken fetches the access token from GitHub
func GetGithubAccessToken(code string) string {
	clientID := utils.GetGithubClientID()
	clientSecret := utils.GetGithubClientSecret()

	requestBody := AccessTokenRequestBody{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Code:         code,
	}

	requestJSON, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(requestJSON))
	if err != nil {
		log.Panic("Request creation failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("Request failed")
	}

	respBody, _ := io.ReadAll(resp.Body)
	var tokenResponse GithubAccessTokenResponse
	json.Unmarshal(respBody, &tokenResponse)

	return tokenResponse.AccessToken
}

// GetGithubData fetches user data from GitHub
func GetGithubData(accessToken string) string {
	req, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		log.Panic("Request creation failed")
	}
	req.Header.Set("Authorization", "token "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic("Request failed")
	}

	respBody, _ := io.ReadAll(resp.Body)
	return string(respBody)
}
