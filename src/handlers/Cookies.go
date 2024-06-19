package handlers

import (
	"net/http"
	"time"
)

func SetSessionCookies(w http.ResponseWriter, email string) {
	cookie := http.Cookie{
		Name:     "session",
		Value:    email,
		Expires:  time.Now().Add(12 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func GetSessionEmail(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func GetUserFromSession(r *http.Request) (string, string, error) {
	emailCookie, err := r.Cookie("email")
	if err != nil {
		return "", "", err
	}
	usernameCookie, err := r.Cookie("username")
	if err != nil {
		return "", "", err
	}
	return emailCookie.Value, usernameCookie.Value, nil
}
