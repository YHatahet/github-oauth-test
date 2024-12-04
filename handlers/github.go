package handlers

import (
	"net/http"

	"github.com/yhatahet/github-oauth-test/services"
	"github.com/yhatahet/github-oauth-test/utils"
)

// GithubLoginHandler handles GitHub login
func GithubLoginHandler(w http.ResponseWriter, r *http.Request) {
	clientID := utils.GetGithubClientID()
	redirectURL := "http://localhost:3000/login/github/callback"

	http.Redirect(w, r, "https://github.com/login/oauth/authorize?client_id="+clientID+"&redirect_uri="+redirectURL, http.StatusMovedPermanently)
}

// GithubCallbackHandler handles the GitHub OAuth callback
func GithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	accessToken := services.GetGithubAccessToken(code)
	githubData := services.GetGithubData(accessToken)
	LoggedInHandler(w, githubData)
}
