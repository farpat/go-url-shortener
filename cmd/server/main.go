package main

import (
	"log"
	"net/http"

	"github.com/farpat/go-url-shortener/internal/config"
	oauthHandlers "github.com/farpat/go-url-shortener/internal/handlers/oauth"
	urlHandlers "github.com/farpat/go-url-shortener/internal/handlers/url"
	"github.com/farpat/go-url-shortener/internal/middlewares"
	"github.com/farpat/go-url-shortener/internal/utils/framework"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// API routes
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middlewares.Authenticate)
	apiRouter.HandleFunc("/urls", urlHandlers.Index).Methods("GET")
	apiRouter.HandleFunc("/urls/{slug}", urlHandlers.Show).Methods("GET")
	apiRouter.HandleFunc("/urls/{slug}", urlHandlers.Destroy).Methods("DELETE")
	apiRouter.HandleFunc("/urls", urlHandlers.Store).Methods("POST")

	// OAuth routes
	oauthRouter := router.PathPrefix("/oauth").Subrouter()
	oauthRouter.HandleFunc("/login", oauthHandlers.Login).Methods("POST")

	server := &http.Server{
		Addr:    ":" + config.App["port"],
		Handler: router,
	}

	log.Fatal(server.ListenAndServeTLS(
		framework.ProjectPath("certs/cert.pem"),
		framework.ProjectPath("certs/key.pem"),
	))
}
