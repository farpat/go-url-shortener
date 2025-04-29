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
	urls, err := urlRepository.All()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(IndexResponse{
		Data: urls,
	})
}
