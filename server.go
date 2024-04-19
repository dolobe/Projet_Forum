package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/Home.html")
	})

	http.HandleFunc("/test.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/css/Home.css")
	})

	fmt.Println("Serveur démarré sur le port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
