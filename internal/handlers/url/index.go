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

func Index(w http.ResponseWriter, r *http.Request) {
	urls, err := urlRepository.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(IndexResponse{
		Data: urls,
	})
}
