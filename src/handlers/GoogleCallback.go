package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
)

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
		Email      string `json:"email"`
		Name       string `json:"name"`
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		http.Error(w, "Échec de la lecture des informations de profil", http.StatusInternalServerError)
		return
	}

	basededonnes, err := DatabasePath()
	if err != nil {
		http.Error(w, "Échec de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer basededonnes.Close()

	http.Redirect(w, r, "/pseudo?email="+profile.Email+"&name="+profile.Name, http.StatusSeeOther)
}

func HandlePseudo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		email := r.FormValue("email")
		pseudo := r.FormValue("pseudo")

		basededonnes, err := DatabasePath()
		if err != nil {
			http.Error(w, "Échec de la connexion à la base de données", http.StatusInternalServerError)
			return
		}
		defer basededonnes.Close()

		_, err = basededonnes.Exec("UPDATE Users SET pseudo = ? WHERE email = ?", pseudo, email)
		if err != nil {
			http.Error(w, "Échec de la mise à jour du pseudo", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/allCategory", http.StatusSeeOther)
		return
	}

	email := r.URL.Query().Get("email")

	tmpl := template.Must(template.ParseFiles("templates/pseudo.html"))
	err := tmpl.Execute(w, map[string]string{
		"Email": email,
	})
	if err != nil {
		http.Error(w, "Échec du rendu de la page", http.StatusInternalServerError)
	}
}
