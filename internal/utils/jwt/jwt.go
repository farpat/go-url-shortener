package jwt

import (
	"time"

	"github.com/farpat/go-url-shortener/internal/config"
	jwtlib "github.com/golang-jwt/jwt/v5"
)

// GenerateToken crée un nouveau token JWT valide pour 5 minutes
func GenerateToken() (string, error) {
	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, jwtlib.MapClaims{
		"iss": "go-url-shortener",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	return token.SignedString([]byte(config.App["jwt_secret"]))
}
