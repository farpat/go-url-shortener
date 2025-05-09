package validation

import (
	"testing"

	"github.com/farpat/go-url-shortener/internal/validation"
	"github.com/farpat/go-url-shortener/tests"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestUniqueSlugSuccess(t *testing.T) {
	// ARRANGE
	teardown, _ := tests.SetupTestDB()
	defer teardown()

	// ACT
	validator := validation.GetValidator()
	err := validator.Var("new-slug", "unique_slug")

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
	validator := validation.GetValidator()
	err = validator.Var("existing-slug", "unique_slug")

	// ASSERT
	assert.Error(t, err)
}
