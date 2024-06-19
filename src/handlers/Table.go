package handlers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type PageData struct {
	Categories []Category
}

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
	if err != nil {
		return err
	}

	createTableCategoryQuery := `CREATE TABLE IF NOT EXISTS Category (
		id TEXT PRIMARY KEY,
		CategoryName TEXT UNIQUE,
		SubjectName TEXT
	)`
	_, err = basededonnees.Exec(createTableCategoryQuery)
	if err != nil {
		return err
	}
	return nil
}

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
