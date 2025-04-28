package main

import (
	"net/http"

	"github.com/farpat/go-url-shortener/internal/config"
	urlHandler "github.com/farpat/go-url-shortener/internal/handlers/url"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/urls", urlHandler.Index).Methods("GET")
	router.HandleFunc("/api/urls/{slug}", urlHandler.Show).Methods("GET")

	http.ListenAndServeTLS(":"+config.App["port"], "certs/cert.pem", "certs/key.pem", router)
}
