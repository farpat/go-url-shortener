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

	http.ListenAndServeTLS(":"+config.App["port"], "certs/cert.pem", "certs/key.pem", router)
}
