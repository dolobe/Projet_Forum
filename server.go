package main

import (
	"Projet_Forum/src/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

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

	fmt.Println("Démarrage du serveur sur le port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Échec du démarrage du serveur :", err)
		return
	}
}
