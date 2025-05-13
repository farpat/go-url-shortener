package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/farpat/go-url-shortener/internal/models"
	"github.com/farpat/go-url-shortener/internal/repositories"
	"github.com/farpat/go-url-shortener/internal/services/string_utils"
	"github.com/farpat/go-url-shortener/internal/validation"
	"github.com/go-playground/validator/v10"
)

type StoreUrlRequest struct {
	Url  string `json:"url" validate:"required,url,startswith=https://"`
	Slug string `json:"slug" validate:"required,unique_slug"`
}

type StoreResponse struct {
	Data models.UrlShowItem `json:"data"`
}

type StoreErrorResponse struct {
	Error    string            `json:"error"`
	Messages map[string]string `json:"messages"`
}

func Store(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var urlRequest StoreUrlRequest
	if err := json.NewDecoder(request.Body).Decode(&urlRequest); err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(response).Encode(StoreErrorResponse{
			Error:    "Malformed JSON",
			Messages: map[string]string{},
		})
		return
	}

	slug := string_utils.GenerateSlug(urlRequest.Url)
	urlRequest.Slug = slug
	if err := validation.GetValidate().Struct(urlRequest); err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(response).Encode(StoreErrorResponse{
			Error:    "Invalid data",
			Messages: validation.FormatErrors(err.(validator.ValidationErrors)),
		})
		return
	}

	err := repositories.NewUrlRepository().Create(models.UrlShowItem{
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

	url, _ := repositories.NewUrlRepository().Find(urlRequest.Slug)
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(StoreResponse{
		Data: url,
	})
}
