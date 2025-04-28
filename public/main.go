package main

import (
	"net/http"

	"github.com/farpat/go-url-shortener/internal/config"
	"github.com/farpat/go-url-shortener/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/urls", handlers.ListUrls).Methods("GET")

	http.ListenAndServeTLS(":"+config.App["port"], "certs/cert.pem", "certs/key.pem", router)
}
