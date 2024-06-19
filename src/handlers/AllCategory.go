package handlers

import (
	"html/template"
	"net/http"
)

func HandleAllCategoryPage(w http.ResponseWriter, r *http.Request) {
	basededonnees, err := DatabasePath()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer basededonnees.Close()

	email, err := GetSessionEmail(r)
	if err != nil {
		http.Redirect(w, r, "/Login", http.StatusSeeOther)
		return
	}

	username, err := getUsernameEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := fetchCategories(basededonnees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/AllCategory.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	donnees := struct {
		Username   string
		Categories []Category
	}{
		Username:   username,
		Categories: categories,
	}

	err = tmpl.Execute(w, donnees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
