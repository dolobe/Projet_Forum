package handlers

import (
	"database/sql"
)

type PageData struct {
	Categories []Category
	Username   string
}

type User struct {
	ID       string
	Name     string
	LastName string
	Pseudo   string
	Email    string
	Password string
}

// Category is a struct that represents a category in the database
func DatabasePath() (*sql.DB, error) {
	basededonnees, err := sql.Open("sqlite3", "./src/DataBase/forum.db")
	if err != nil {
		return nil, err
	}
	if err := createTables(basededonnees); err != nil {
		return nil, err
	}
	return basededonnees, nil
}

// CreateTables User
func createTables(db *sql.DB) error {
	createTableUsersQuery := `CREATE TABLE IF NOT EXISTS Users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT UNIQUE,
        last_name TEXT UNIQUE,
        pseudo TEXT UNIQUE,
        email TEXT UNIQUE,
        password TEXT
    )`

	// CreateTable Category
	createTableCategoryQuery := `CREATE TABLE IF NOT EXISTS Category (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        CategoryName TEXT UNIQUE,
        SubjectName TEXT
    )`

	_, err := db.Exec(createTableUsersQuery)
	if err != nil {
		return err
	}

	_, err = db.Exec(createTableCategoryQuery)
	if err != nil {
		return err
	}

	return nil
}

// Get Gategories
func getCategories(basededonnees *sql.DB) ([]Category, error) {
	rows, err := basededonnees.Query(`SELECT CategoryName, SubjectName FROM Category`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.CategoryName, &category.SubjectName); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

// GetSubjects
func getSubjects(db *sql.DB) ([]Category, error) {
	rows, err := db.Query(`SELECT CategoryName, SubjectName FROM Category`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.CategoryName, &category.SubjectName); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
