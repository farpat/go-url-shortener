package url_test

import (
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

const showUrlPrefix = "/api/urls/"

func TestShowExistingUrl(t *testing.T) {
	// ARRANGE
	router := router.SetupRouter()
	teardown, db := tests.SetupTestDB()
	defer teardown()
	insertUrl(db, models.UrlListItem{Slug: "abc", Url: "https://example.com"})

	// ACT
	request := httptest.NewRequest(http.MethodGet, showUrlPrefix+"abc", nil)
	tokenString, _ := jwt.GenerateToken()
	request.Header.Set("Authorization", "Bearer "+tokenString)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, http.StatusOK, response.Code)
	var showUrlResponse map[string]any
	err := json.NewDecoder(response.Body).Decode(&showUrlResponse)
	assert.NoError(t, err)
	showUrlResponseData := showUrlResponse["data"].(map[string]any)
	assert.Equal(t, "abc", showUrlResponseData["slug"])
	assert.Equal(t, "https://example.com", showUrlResponseData["url"])
	assert.NotEmpty(t, showUrlResponseData["created_at"])
}

func TestShowUnexistingUrl(t *testing.T) {
	// ARRANGE
	router := router.SetupRouter()
	teardown, _ := tests.SetupTestDB()
	defer teardown()

	// ACT
	request := httptest.NewRequest(http.MethodGet, showUrlPrefix+"nonexistent", nil)
	tokenString, _ := jwt.GenerateToken()
	request.Header.Set("Authorization", "Bearer "+tokenString)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, http.StatusNotFound, response.Code)
}
