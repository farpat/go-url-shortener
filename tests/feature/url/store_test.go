package url_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/farpat/go-url-shortener/internal/models"
	"github.com/farpat/go-url-shortener/internal/router"
	"github.com/farpat/go-url-shortener/internal/utils/jwt"
	"github.com/farpat/go-url-shortener/tests"
	"github.com/stretchr/testify/assert"
)

const storeUrl = "/api/urls"

func TestStoreUrl(t *testing.T) {
	// ARRANGE
	router := router.SetupRouter()
	teardown, _ := tests.SetupTestDB()
	defer teardown()

	urlData := models.UrlShowItem{Url: "https://example.com"}
	jsonData, _ := json.Marshal(urlData)

	// ACT
	request := httptest.NewRequest(http.MethodPost, storeUrl, bytes.NewBuffer(jsonData))
	tokenString, _ := jwt.GenerateToken()
	request.Header.Set("Authorization", "Bearer "+tokenString)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, http.StatusCreated, response.Code)
	var storeResponse map[string]any
	json.NewDecoder(response.Body).Decode(&storeResponse)
	storeResponseData := storeResponse["data"].(map[string]any)
	assert.Equal(t, "https://example.com", storeResponseData["url"])
	assert.NotEmpty(t, storeResponseData["slug"])
	assert.NotEmpty(t, storeResponseData["created_at"])
}

func TestStoreUrlDoesntWorksIfDataAreInvalid(t *testing.T) {
	// ARRANGE
	router := router.SetupRouter()
	teardown, _ := tests.SetupTestDB()
	defer teardown()

	urlData := models.UrlShowItem{}
	jsonData, _ := json.Marshal(urlData)

	// ACT
	request := httptest.NewRequest(http.MethodPost, storeUrl, bytes.NewBuffer(jsonData))
	tokenString, _ := jwt.GenerateToken()
	request.Header.Set("Authorization", "Bearer "+tokenString)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
}
