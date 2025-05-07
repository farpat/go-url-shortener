package config

import "github.com/farpat/go-url-shortener/internal/utils/framework"

var App = map[string]string{
	"port":       framework.Env("PORT", "8080"),
	"jwt_secret": framework.Env("JWT_SECRET", "secret"),
}
