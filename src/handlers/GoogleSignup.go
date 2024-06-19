package handlers

import (
	"net/http"

	"golang.org/x/oauth2"
)

func HandleGoogleSignup(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, googleOAuthConfig.AuthCodeURL("state-signup", oauth2.AccessTypeOffline), http.StatusTemporaryRedirect)
}
