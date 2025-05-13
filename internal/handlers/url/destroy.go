package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	internalErrors "github.com/farpat/go-url-shortener/internal/errors"
	"github.com/farpat/go-url-shortener/internal/repositories"
	"github.com/gorilla/mux"
)

func Destroy(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	slug := mux.Vars(request)["slug"]
	err := repositories.NewUrlRepository().Delete(slug)
	if err != nil {
		var errorJSON map[string]string
		var notFoundError *internalErrors.NotFoundError

		if errors.As(err, &notFoundError) {
			response.WriteHeader(http.StatusNotFound)
			errorJSON = map[string]string{"error": notFoundError.Error()}
		} else {
			response.WriteHeader(http.StatusInternalServerError)
			errorJSON = map[string]string{"error": "Internal Server Error"}
		}

		json.NewEncoder(response).Encode(errorJSON)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}
