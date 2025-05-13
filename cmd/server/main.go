package main

import (
	"log"
	"net/http"

	"github.com/farpat/go-url-shortener/internal/config"
	"github.com/farpat/go-url-shortener/internal/router"
	"github.com/farpat/go-url-shortener/internal/utils/framework"
)

func main() {
	server := &http.Server{
		Addr:    ":" + config.App["port"],
		Handler: router.SetupRouter(),
	}

	log.Fatal(server.ListenAndServeTLS(
		framework.ProjectPath("certs/cert.pem"),
		framework.ProjectPath("certs/key.pem"),
	))
}
