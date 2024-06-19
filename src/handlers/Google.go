package handlers

import (
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
