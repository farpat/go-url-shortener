package validation

import (
	"testing"

	handlers "github.com/farpat/go-url-shortener/internal/handlers/url"
	"github.com/farpat/go-url-shortener/internal/validation"
	"github.com/farpat/go-url-shortener/tests"
	"github.com/go-playground/validator/v10"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestUniqueSlugSuccess(t *testing.T) {
	// ARRANGE
	teardown, _ := tests.SetupTestDB()
	defer teardown()

	// ACT
	err := validation.GetValidate().Var("new-slug", "unique_slug")

	// ASSERT
	assert.NoError(t, err)
}

func TestUniqueSlugFailure(t *testing.T) {
	// ARRANGE
	teardown, db := tests.SetupTestDB()
	defer teardown()

	// Insert a URL with the slug we want to test
	_, err := db.Exec("INSERT INTO urls (slug, url) VALUES (?, ?)", "existing-slug", "https://example.com")
	if err != nil {
		panic(err)
	}

	// ACT
	err = validation.GetValidate().Var("existing-slug", "unique_slug")

	// ASSERT
	assert.Error(t, err)
}

func TestFormatErrorsFunction(t *testing.T) {
	// ARRANGE
	teardown, db := tests.SetupTestDB()
	defer teardown()

	// Insert a URL with the slug we want to test
	_, err := db.Exec("INSERT INTO urls (slug, url) VALUES (?, ?)", "existing-slug", "https://example.com")
	if err != nil {
		panic(err)
	}

	// ACT
	var urlRequest handlers.StoreUrlRequest
	urlRequest.Slug = "existing-slug"
	urlRequest.Url = "https://example.com"

	err = validation.GetValidate().Struct(urlRequest)
	result := validation.FormatErrors(err.(validator.ValidationErrors))

	// ASSERT
	expected := map[string]string{"Slug": "unique_slug"}
	assert.Equal(t, expected, result)
}
