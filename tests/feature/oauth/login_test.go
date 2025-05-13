package oauth_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/farpat/go-url-shortener/internal/router"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// ARRANGE
	router := router.SetupRouter()

	// ACT
	request := httptest.NewRequest(http.MethodPost, "/oauth/login", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	// ASSERT
	assert.Equal(t, http.StatusOK, response.Code)
	var loginResponse map[string]string
	err := json.NewDecoder(response.Body).Decode(&loginResponse)
	assert.NoError(t, err)
	assert.NotEmpty(t, loginResponse["access_token"])
	assert.NotEmpty(t, loginResponse["expired_at"])
}
