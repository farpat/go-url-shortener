package main

import (
	"net/http"

	"github.com/farpat/go-url-shortener/internal/config"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/urls", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).Methods("GET")

	port := config.App["port"]
	if port == "" {
		port = "8080"
	}

	// Start the server
	http.ListenAndServeTLS(":"+port, "certs/cert.pem", "certs/key.pem", router)
}
