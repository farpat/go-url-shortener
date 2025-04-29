package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/farpat/go-url-shortener/internal/models"
	urlRepository "github.com/farpat/go-url-shortener/internal/repositories"
	"github.com/farpat/go-url-shortener/internal/requests"
	"github.com/farpat/go-url-shortener/internal/services"
	"github.com/go-playground/validator/v10"
)

type StoreResponse struct {
	Data models.UrlShowItem `json:"data"`
}

type StoreErrorResponse struct {
	Error    string            `json:"error"`
	Messages map[string]string `json:"messages"`
}

func Store(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var urlRequest requests.StoreUrlRequest
	if err := json.NewDecoder(request.Body).Decode(&urlRequest); err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(response).Encode(StoreErrorResponse{
			Error:    "Malformed JSON",
			Messages: map[string]string{},
		})
		return
	}

	slug := services.GenerateSlug(urlRequest.Url)
	urlRequest.Slug = slug
	if err := makeValidator().Struct(urlRequest); err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(response).Encode(StoreErrorResponse{
			Error:    "Invalid data",
			Messages: formatValidationErrors(err.(validator.ValidationErrors)),
		})
		return
	}

	err := urlRepository.Create(models.UrlShowItem{
		Url:  urlRequest.Url,
		Slug: urlRequest.Slug,
	})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(StoreErrorResponse{
			Error:    "Internal Server Error",
			Messages: map[string]string{},
		})
		return
	}

	url, _ := urlRepository.Find(urlRequest.Slug)

	json.NewEncoder(response).Encode(StoreResponse{
		Data: url,
	})
}

func makeValidator() *validator.Validate {
	validate := validator.New()

	validate.RegisterValidation("unique_slug", func(fl validator.FieldLevel) bool {
		_, err := urlRepository.Find(fl.Field().String())
		var notFoundError *urlRepository.NotFoundError

		// error is "NotFoundError" => slug does not exist => good
		return errors.As(err, &notFoundError)
	})

	return validate
}

func formatValidationErrors(validationErrors validator.ValidationErrors) map[string]string {
	var messages = map[string]string{}
	for _, e := range validationErrors {
		messages[e.Field()] = e.Tag()
	}
	return messages
}
