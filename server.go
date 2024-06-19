package main

import (
	"Projet_Forum/src/handlers"
	"fmt"
	"net/http"
)

func main() {
	// Handle the assets
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Handle the pages
	http.HandleFunc("/", handlers.HandleHomePage)
	http.HandleFunc("/Login", handlers.HandleLoginPage)
	http.HandleFunc("/signup", handlers.HandleSignupPage)
	http.HandleFunc("/category", handlers.HandleCategoryPage)
	http.HandleFunc("/allCategory", handlers.HandleAllCategoryPage)

	http.HandleFunc("/post", handlers.HandlePostPage)
	http.HandleFunc("/googleSignup", handlers.HandleGoogleSignup)
	http.HandleFunc("/googleLogin", handlers.HandleGoogleLogin)
	http.HandleFunc("/google/callback", handlers.HandleGoogleCallback)
	http.HandleFunc("/pseudo", handlers.HandlePseudo)

	http.HandleFunc("/postComment", handlers.HandlePostComment)
	http.HandleFunc("/postReply", handlers.HandlePostReply)

	http.HandleFunc("/getComments", handlers.HandleGetComments)

	// Start the serverS
	fmt.Println("Démarrage du serveur sur le port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Échec du démarrage du serveur :", err)
		return
	}
}
