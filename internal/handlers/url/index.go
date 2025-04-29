package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/farpat/go-url-shortener/internal/models"
	urlRepository "github.com/farpat/go-url-shortener/internal/repositories"
)

type IndexResponse struct {
	Data []models.UrlListItem `json:"data"`
}

func Index(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	urls, err := urlRepository.All()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]string{
			"error": "Internal Server Error",
		})
		return
	}

	json.NewEncoder(response).Encode(IndexResponse{
		Data: urls,
	})
}
