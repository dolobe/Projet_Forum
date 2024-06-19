package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const (
	insertCommentQuery = `INSERT INTO Comments (username, comment_text) VALUES (?, ?)`
	insertReplyQuery   = `INSERT INTO Replies (comment_id, username, reply_text) VALUES (?, ?, ?)`
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

	// Take the categories
	categories, err := getCategories(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Take the subjects
	subjects, err := getSubjects(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// use the first subject name
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

	// Render the post page
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

// Take UserName by email
func GetUsernameByEmail(db *sql.DB, email string) (string, error) {
	var username string
	query := `SELECT name FROM Users WHERE email = ?`
	err := db.QueryRow(query, email).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

// Insert Comment
func HandleInsertComment(db *sql.DB, username string, commentText string) error {
	stmt, err := db.Prepare(insertCommentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, commentText)
	if err != nil {
		return err
	}

	return nil
}

// Insert Reply
func HandleInsertReply(db *sql.DB, commentID int, username string, replyText string) error {
	stmt, err := db.Prepare(insertReplyQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(commentID, username, replyText)
	if err != nil {
		return err
	}

	return nil
}

// HandlePostComment
func HandlePostComment(w http.ResponseWriter, r *http.Request) {
	var comment struct {
		Username    string `json:"username"`
		CommentText string `json:"commentText"`
	}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db, err := DatabasePath()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := HandleInsertComment(db, comment.Username, comment.CommentText); err != nil {
		http.Error(w, "Failed to insert comment", http.StatusInternalServerError)
		return
	}

	// Réponse de succès
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "Comment inserted successfully"}
	json.NewEncoder(w).Encode(response)
}

// HandlePostReply
func HandlePostReply(w http.ResponseWriter, r *http.Request) {
	var reply struct {
		CommentID int    `json:"commentID"`
		Username  string `json:"username"`
		ReplyText string `json:"replyText"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reply); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db, err := DatabasePath()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := HandleInsertReply(db, reply.CommentID, reply.Username, reply.ReplyText); err != nil {
		http.Error(w, "Failed to insert reply", http.StatusInternalServerError)
		return
	}

	// Reponse de succès
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Reply inserted successfully")
}

func HandleGetComments(w http.ResponseWriter, r *http.Request) {
	db, err := DatabasePath()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	comments, err := GetCommentsFromDB(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// converted to JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

// GetCommentsFromDBs
func GetCommentsFromDB(db *sql.DB) ([]Comment, error) {
	rows, err := db.Query(`SELECT id, username, comment_text FROM Comments`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.Username, &comment.CommentText); err != nil {
			return nil, err
		}

		// Take the replies for the comment
		comment.Replies, err = GetRepliesForComment(db, comment.ID)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

// GetRepliesForComment
func GetRepliesForComment(db *sql.DB, commentID int) ([]Reply, error) {
	rows, err := db.Query(`SELECT username, reply_text FROM Replies WHERE comment_id = ?`, commentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var replies []Reply
	for rows.Next() {
		var reply Reply
		if err := rows.Scan(&reply.Username, &reply.ReplyText); err != nil {
			return nil, err
		}
		replies = append(replies, reply)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return replies, nil
}
