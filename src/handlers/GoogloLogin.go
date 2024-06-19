package handlers

import (
	"net/http"

	"golang.org/x/oauth2"
)

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, googleOAuthConfig.AuthCodeURL("state-login", oauth2.AccessTypeOffline), http.StatusTemporaryRedirect)
}
