package handlers

import (
	"html/template"
	"net/http"
)

func HandleCategoryPage(w http.ResponseWriter, r *http.Request) {
	email, err := GetSessionEmail(r)
	if err != nil {
		http.Redirect(w, r, "/Login", http.StatusSeeOther)
		return
	}

	username, err := getUsernameByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/category.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	donnees := struct {
		Username string
	}{
		Username: username,
	}

	err = tmpl.Execute(w, donnees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getUsernameByEmail(email string) (string, error) {
	basededonnees, err := DatabasePath()
	if err != nil {
		return "", err
	}
	defer basededonnees.Close()

	var username string
	query := `SELECT pseudo FROM Users WHERE email = ?`
	err = basededonnees.QueryRow(query, email).Scan(&username)
	if err != nil {
		return "", err
	}

	return username, nil
}
