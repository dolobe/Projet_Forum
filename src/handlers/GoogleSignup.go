package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOAuthConfig = &oauth2.Config{
	ClientID:     "781588257796-guiaanpesar28qi2st467mn4ip5ta5m5.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "http://localhost:8080/google/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

func HandleGoogleSignup(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, googleOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline), http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := googleOAuthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Échec de l'échange de code", http.StatusInternalServerError)
		return
	}

	client := googleOAuthConfig.Client(r.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		http.Error(w, "Échec de la récupération des informations de profil", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var profile struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		http.Error(w, "Échec de la lecture des informations de profil", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Email: %s, Name: %s", profile.Email, profile.Name)
}
