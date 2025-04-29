package config

import "github.com/farpat/go-url-shortener/internal/utils"

var App = map[string]string{
	"port": utils.Env("PORT", "8080"),
}
