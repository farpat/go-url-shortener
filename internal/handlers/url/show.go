package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/farpat/go-url-shortener/internal/models"
	urlRepository "github.com/farpat/go-url-shortener/internal/repositories"
	"github.com/gorilla/mux"
)

type ShowResponse struct {
	Data models.UrlShowItem `json:"data"`
}

func Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	slug := vars["slug"]
	url, err := urlRepository.Find(slug)
	if err != nil {
		errorJSON := map[string]string{
			"error": "URL linked to '" + slug + "' not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorJSON)
		return
	}

	json.NewEncoder(w).Encode(ShowResponse{
		Data: url,
	})
}
