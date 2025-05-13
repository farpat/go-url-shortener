package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/farpat/go-url-shortener/internal/utils/jwt"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiredAt   string `json:"expired_at"`
}

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	tokenString, error := jwt.GenerateToken()

	if error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]string{
			"error": "Internal Server Error",
		})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(LoginResponse{
		AccessToken: tokenString,
		ExpiredAt:   time.Now().Add(time.Minute * 5).Format("2006-01-02 15:04:05"),
	})
}
