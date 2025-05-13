package router

import (
	OauthHandlers "github.com/farpat/go-url-shortener/internal/handlers/oauth"
	UrlHandlers "github.com/farpat/go-url-shortener/internal/handlers/url"
	"github.com/farpat/go-url-shortener/internal/middlewares"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// API routes
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middlewares.Authenticate)
	apiRouter.HandleFunc("/urls", UrlHandlers.Index).Methods("GET")
	apiRouter.HandleFunc("/urls/{slug}", UrlHandlers.Show).Methods("GET")
	apiRouter.HandleFunc("/urls/{slug}", UrlHandlers.Destroy).Methods("DELETE")
	apiRouter.HandleFunc("/urls", UrlHandlers.Store).Methods("POST")

	// OAuth routes
	oauthRouter := router.PathPrefix("/oauth").Subrouter()
	oauthRouter.HandleFunc("/login", OauthHandlers.Login).Methods("POST")

	return router
}
