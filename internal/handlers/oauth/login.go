package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/farpat/go-url-shortener/internal/utils/jwt"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
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
	})
}
