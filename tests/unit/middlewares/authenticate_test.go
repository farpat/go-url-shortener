package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/farpat/go-url-shortener/internal/middlewares"
	"github.com/farpat/go-url-shortener/internal/utils/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticateAsSuccess(t *testing.T) {
	// Create a mock handler to pass to the middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create a request with a valid Authorization header
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	tokenString, _ := jwt.GenerateToken()
	req.Header.Set("Authorization", "Bearer "+tokenString)

	// Create a response recorder
	resp := httptest.NewRecorder()

	// Call the middleware
	middlewares.Authenticate(handler).ServeHTTP(resp, req)

	// Assert the status code is OK
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestNoAuthenticationIfNoToken(t *testing.T) {
	// Create a mock handler to pass to the middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create a request without an Authorization header
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a response recorder
	resp := httptest.NewRecorder()

	// Call the middleware
	middlewares.Authenticate(handler).ServeHTTP(resp, req)

	// Assert the status code is Unauthorized
	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}

func TestNoAuthenticationIfInvalidToken(t *testing.T) {
	// Create a mock handler to pass to the middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create a request with an invalid Authorization header
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")

	// Create a response recorder
	resp := httptest.NewRecorder()

	// Call the middleware
	middlewares.Authenticate(handler).ServeHTTP(resp, req)

	// Assert the status code is Unauthorized
	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}
