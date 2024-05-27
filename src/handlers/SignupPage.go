package handlers

import (
	"database/sql"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func DatabasePath() (*sql.DB, error) {
	basededonnees, err := sql.Open("sqlite3", "./src/DataBase/forum.db")
	if err != nil {
		return nil, err
	}
	if err := createTableUsers(basededonnees); err != nil {
		return nil, err
	}
	return basededonnees, nil
}

func createTableUsers(basededonnees *sql.DB) error {
	createTableUsersQuery := `CREATE TABLE IF NOT EXISTS Users (
		id TEXT PRIMARY KEY,
		name TEXT UNIQUE,
		last_name TEXT UNIQUE,
		pseudo TEXT UNIQUE,
		email TEXT UNIQUE,
		password TEXT
	)`
	_, err := basededonnees.Exec(createTableUsersQuery)
	return err
}

func HandleSignupPage(w http.ResponseWriter, r *http.Request) {
	basededonnees, err := DatabasePath()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer basededonnees.Close()

	if r.Method == http.MethodPost {
		name := r.FormValue("nom")
		lastName := r.FormValue("prénom")
		pseudo := r.FormValue("pseudo")
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirmPassword")

		if password != confirmPassword {
			http.Error(w, "Votre mot de passe n'est pas identique", http.StatusBadRequest)
			return
		}

		if name == "" || lastName == "" || pseudo == "" || email == "" || password == "" {
			http.Error(w, "Tout les champs sont obligatoires, veuillez les remplir s'il vous plaît", http.StatusBadRequest)
			return
		}

		if err := insertUser(basededonnees, name, lastName, pseudo, email, password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	htmlfile, err := ioutil.ReadFile("templates/signupPage.html")
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

func insertUser(basededonnees *sql.DB, name, lastName, pseudo, email, password string) error {
	id := uuid.New().String()
	insertUserQuery := `INSERT INTO Users (id, name, last_name, pseudo, email, password) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := basededonnees.Exec(insertUserQuery, id, name, lastName, pseudo, email, password)
	return err
}
