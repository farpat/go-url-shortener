package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/farpat/go-url-shortener/internal/models"
	urlRepository "github.com/farpat/go-url-shortener/internal/repositories"
)

type ShowResponse struct {
	Data models.UrlShowItem `json:"data"`
}

func Show(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	slug := request.Context().Value("slug").(string)
	url, err := urlRepository.Find(slug)
	if err != nil {
		errorJSON := map[string]string{
			"error": "URL linked to '" + slug + "' not found",
		}
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errorJSON)
		return
	}

	json.NewEncoder(response).Encode(ShowResponse{
		Data: url,
	})
}
