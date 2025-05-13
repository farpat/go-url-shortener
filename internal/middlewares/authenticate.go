package middlewares

import (
	"net/http"
	"strings"

	"github.com/farpat/go-url-shortener/internal/utils/jwt"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerContent := r.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(headerContent, "Bearer ")
		if tokenString == "" {
			returnUnauthorized(w)
			return
		}

		if _, err := jwt.ValidateToken(tokenString); err != nil {
			returnUnauthorized(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func returnUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Unauthorized"))
}
