package handlers

import (
	"database/sql"
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Category struct {
	ID           string
	CategoryName string
	SubjectName  string
}

func HandleCategoryPage(w http.ResponseWriter, r *http.Request) {
	basededonnees, err := DatabasePath()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer basededonnees.Close()

	if r.Method == http.MethodPost {
		categoryName := r.FormValue("cat")
		subjectName := r.FormValue("sub")

		if categoryName == "" || subjectName == "" {
			http.Error(w, "Category and subject name cannot be empty", http.StatusBadRequest)
			return
		}

		if err := insertCategory(basededonnees, categoryName, subjectName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/category", http.StatusSeeOther)
		return
	}

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

	tmpl, err := template.ParseFiles("templates/category.html")
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

func getUsernameEmail(email string) (string, error) {
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

func insertCategory(basededonnees *sql.DB, categoryName, subjectName string) error {
	var here bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM Categories WHERE categoryName = ?)`
	err := basededonnees.QueryRow(checkQuery, categoryName).Scan(&here)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(here)

	if here {
		return errors.New("Category already exists")
	}

	id := uuid.New().String()
	insertCategoryQuery := `INSERT INTO Categories (id, categoryName, subjectName) VALUES (?, ?, ?)`
	_, err = basededonnees.Exec(insertCategoryQuery, id, categoryName, subjectName)
	if err != nil {
		log.Println(err)
	}
	return err
}

func fetchCategories(basededonnees *sql.DB) ([]Category, error) {
	rows, err := basededonnees.Query(`SELECT id, categoryName, subjectName FROM Categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.CategoryName, &category.SubjectName)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
