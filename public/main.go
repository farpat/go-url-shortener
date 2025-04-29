package main

import (
	"net/http"

	"github.com/farpat/go-url-shortener/internal/config"
	urlHandler "github.com/farpat/go-url-shortener/internal/handlers/url"
	"github.com/farpat/go-url-shortener/internal/utils"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/urls", urlHandler.Index).Methods("GET")
	router.HandleFunc("/api/urls/{slug:[a-z0-9]+}", urlHandler.Show).Methods("GET")
	router.HandleFunc("/api/urls/{slug:[a-z0-9]+}", urlHandler.Destroy).Methods("DELETE")

	http.ListenAndServeTLS(
		":"+config.App["port"],
		utils.ProjectPath("certs/cert.pem"),
		utils.ProjectPath("certs/key.pem"),
		router,
	)
}
