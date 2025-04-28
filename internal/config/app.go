package config

import (
	"github.com/farpat/go-url-shortener/internal/services"
)

var App = map[string]string{
	"port": services.Env("PORT", "8080"),
}
