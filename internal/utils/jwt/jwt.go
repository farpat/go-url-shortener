package jwt

import (
	"time"

	"github.com/farpat/go-url-shortener/internal/config"
	jwtlib "github.com/golang-jwt/jwt/v5"
)

var secret []byte

func init() {
	secret = []byte(config.App["jwt_secret"])
}

// GenerateToken generates a new JWT token valid for 5 minutes
func GenerateToken() (string, error) {
	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, jwtlib.MapClaims{
		"iss": "go-url-shortener",
		"sub": "user",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	return token.SignedString(secret)
}

func ValidateToken(tokenString string) (*jwtlib.Token, error) {
	token, error := jwtlib.ParseWithClaims(tokenString, &jwtlib.MapClaims{}, func(token *jwtlib.Token) (interface{}, error) {
		return secret, nil
	})

	return token, error
}
