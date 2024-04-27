package handlers

import (
	"io/ioutil"
	"net/http"
)

func HandleLoginPage(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := ioutil.ReadFile("templates/LoginPage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	_, err = w.Write(htmlFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
