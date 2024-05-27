package handlers

import (
	"database/sql"
	"io/ioutil"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	basededonnees, err := DatabasePath()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer basededonnees.Close()

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" {
			http.Error(w, "Tout les champs sont obligatoires, veuillez les remplir s'il vous plaît", http.StatusBadRequest)
			return
		}

		TruePassword, err := CheckUser(basededonnees, email, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if TruePassword {
			http.Redirect(w, r, "/category", http.StatusSeeOther)
			return
		}

		emailIsHere, err := CheckEmail(basededonnees, email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if emailIsHere {
			http.Error(w, "Votre email & mot de passe est incorrect", http.StatusBadRequest)
		} else {
			http.Error(w, "Email non trouvé", http.StatusBadRequest)
		}
		return
	}

	htmlfile, err := ioutil.ReadFile("templates/LoginPage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write(htmlfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CheckUser(basededonnees *sql.DB, email, password string) (bool, error) {
	var TablePassword string
	query := `SELECT password FROM Users WHERE email = ?`
	err := basededonnees.QueryRow(query, email).Scan(&TablePassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return TablePassword == password, nil
}

func CheckEmail(basededonnees *sql.DB, email string) (bool, error) {
	query := `SELECT email FROM Users WHERE email = ?`
	err := basededonnees.QueryRow(query, email).Scan(&email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
