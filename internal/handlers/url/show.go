package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/farpat/go-url-shortener/internal/models"
	urlRepository "github.com/farpat/go-url-shortener/internal/repositories"
	"github.com/gorilla/mux"
)

type ShowResponse struct {
	Data models.UrlShowItem `json:"data"`
}

func Show(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)
	slug := vars["slug"]
	url, err := urlRepository.Find(slug)
	if err != nil {
		var jsonError map[string]string
		var notFoundError *urlRepository.NotFoundError

		if errors.As(err, &notFoundError) {
			response.WriteHeader(http.StatusNotFound)
			jsonError = map[string]string{"error": notFoundError.Error()}
		} else {
			response.WriteHeader(http.StatusInternalServerError)
			jsonError = map[string]string{"error": "Internal Server Error"}
		}

		json.NewEncoder(response).Encode(jsonError)
		return
	}

	json.NewEncoder(response).Encode(ShowResponse{
		Data: url,
	})
}
