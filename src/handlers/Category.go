package handlers

import (
	"database/sql"
	"fmt"
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

// HandleCategoryPage gère les requêtes vers la page des catégories
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
			http.Error(w, "Les champs 'cat' et 'sub' doivent être remplis", http.StatusBadRequest)
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

	tmpl, err := template.ParseFiles("templates/Category.html")
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

// getUsernameEmail renvoie le pseudo de l'utilisateur à partir de son email
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

// insertCategory insère une nouvelle catégorie dans la base de données
func insertCategory(basededonnees *sql.DB, categoryName string, subjectName string) error {
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM Category WHERE categoryName = ?)`
	err := basededonnees.QueryRow(checkQuery, categoryName).Scan(&exists)
	if err != nil {
		log.Printf("Erreur lors de la vérification de la catégorie: %v\n", err)
		return err
	}

	log.Printf("Vérification de l'existence de la catégorie '%s': %v\n", categoryName, exists)

	if exists {
		return fmt.Errorf("la catégorie '%s' existe déjà", categoryName)
	}

	id := uuid.New().String()
	insertCategoryQuery := `INSERT INTO Category (id, categoryName, subjectName) VALUES (?, ?, ?)`
	_, err = basededonnees.Exec(insertCategoryQuery, id, categoryName, subjectName)
	if err != nil {
		log.Printf("Erreur lors de l'insertion de la catégorie: %v\n", err)
	}
	return err
}

// fetchCategories récupère toutes les catégories de la base de données
func fetchCategories(basededonnees *sql.DB) ([]Category, error) {
	rows, err := basededonnees.Query(`SELECT id, categoryName, subjectName FROM Category`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.CategoryName, &category.SubjectName); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
