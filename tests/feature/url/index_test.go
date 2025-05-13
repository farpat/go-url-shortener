package url_test

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	urlHandler "github.com/farpat/go-url-shortener/internal/handlers/url"
	"github.com/farpat/go-url-shortener/internal/models"
	"github.com/farpat/go-url-shortener/internal/router"
	"github.com/farpat/go-url-shortener/internal/utils/jwt"
	"github.com/farpat/go-url-shortener/tests"
	"github.com/stretchr/testify/assert"
)

const indexUrl = "/api/urls"

func TestIndexIsUnauthorizedIfInvalidToken(t *testing.T) {
	// ARRANGE
	router := router.SetupRouter()

	// ACT
	request := httptest.NewRequest(http.MethodGet, indexUrl, nil)
	request.Header.Set("Authorization", "Bearer invalid_token")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, http.StatusUnauthorized, response.Code)
}

func TestIndexListUrls(t *testing.T) {
	// ARRANGE
	router := router.SetupRouter()
	teardown, db := tests.SetupTestDB()
	defer teardown()
	for _, url := range []models.UrlListItem{
		{Slug: "abc", Url: "https://example.com"},
		{Slug: "def", Url: "https://google.com"},
	} {
		insertUrl(db, url)
	}

	// ACT
	request := httptest.NewRequest(http.MethodGet, indexUrl, nil)
	tokenString, _ := jwt.GenerateToken()
	request.Header.Set("Authorization", "Bearer "+tokenString)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, http.StatusOK, response.Code)
	var indexResponse urlHandler.IndexResponse
	json.NewDecoder(response.Body).Decode(&indexResponse)

	assert.Equal(t, 2, len(indexResponse.Data))
	assert.Equal(t, "abc", indexResponse.Data[0].Slug)
	assert.Equal(t, "def", indexResponse.Data[1].Slug)
}

func insertUrl(db *sql.DB, url models.UrlListItem) {
	_, err := db.Exec("INSERT INTO urls (slug, url) VALUES (?, ?)", url.Slug, url.Url)
	if err != nil {
		panic(err)
	}
}
