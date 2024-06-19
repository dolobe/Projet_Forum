package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func HandlePostPage(w http.ResponseWriter, r *http.Request) {
	db, err := DatabasePath()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Take the email from the session
	email, err := GetSessionEmail(r)
	if err != nil {
		http.Error(w, "Utilisateur non connecté", http.StatusUnauthorized)
		return
	}

	// Take the username from the database
	username, err := GetUsernameByEmail(db, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Take the categories in the database
	categories, err := getCategories(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupérer le SubjectName
	subjects, err := getSubjects(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Take the first subject name
	var subjectName string
	if len(subjects) > 0 {
		subjectName = subjects[0].SubjectName
	} else {
		subjectName = "Aucun sujet trouvé"
	}

	data := struct {
		Categories  []Category
		Username    string
		SubjectName string
	}{
		Categories:  categories,
		Username:    username,
		SubjectName: subjectName,
	}

	tmpl, err := template.ParseFiles("templates/post.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUsernameByEmail(db *sql.DB, email string) (string, error) {
	var username string
	query := `SELECT name FROM Users WHERE email = ?`
	err := db.QueryRow(query, email).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
