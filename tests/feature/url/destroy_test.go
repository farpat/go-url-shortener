package url_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/farpat/go-url-shortener/internal/models"
	"github.com/farpat/go-url-shortener/internal/router"
	"github.com/farpat/go-url-shortener/internal/utils/jwt"
	"github.com/farpat/go-url-shortener/tests"
	"github.com/stretchr/testify/assert"
)

const destroyUrlPrefix = "/api/urls/"

func TestDestroyExistingUrl(t *testing.T) {
	// ARRANGE
	router := router.SetupRouter()
	teardown, db := tests.SetupTestDB()
	defer teardown()
	insertUrl(db, models.UrlListItem{Slug: "abc", Url: "https://example.com"})

	// ACT
	request := httptest.NewRequest(http.MethodDelete, destroyUrlPrefix+"abc", nil)
	tokenString, _ := jwt.GenerateToken()
	request.Header.Set("Authorization", "Bearer "+tokenString)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, http.StatusNoContent, response.Code)
}

func TestDestroyUnexistingUrl(t *testing.T) {
	// ARRANGE
	router := router.SetupRouter()
	teardown, _ := tests.SetupTestDB()
	defer teardown()

	// ACT
	request := httptest.NewRequest(http.MethodDelete, destroyUrlPrefix+"nonexistent", nil)
	tokenString, _ := jwt.GenerateToken()
	request.Header.Set("Authorization", "Bearer "+tokenString)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, http.StatusNotFound, response.Code)
}
