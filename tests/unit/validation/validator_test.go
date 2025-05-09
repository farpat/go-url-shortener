package validation

import (
	"database/sql"
	"os"
	"testing"

	"github.com/farpat/go-url-shortener/internal/config"
	"github.com/farpat/go-url-shortener/internal/utils/framework"
	"github.com/farpat/go-url-shortener/internal/validation"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() (teardown func(), db *sql.DB) {
	dbPath := "database_test.db"
	config.Databases["main"] = dbPath
	absoluteDbPath := framework.ProjectPath(dbPath)

	db, err := sql.Open("sqlite3", absoluteDbPath)
	if err != nil {
		panic(err)
	}

	// Crée la table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS urls (
		slug TEXT PRIMARY KEY,
		url TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		panic(err)
	}

	return func() {
		db.Close()
		os.Remove(absoluteDbPath)
	}, db
}

func TestUniqueSlugSuccess(t *testing.T) {
	// ARRANGE
	teardown, _ := setupTestDB()
	defer teardown()

	// ACT
	validator := validation.GetValidator()
	err := validator.Var("new-slug", "unique_slug")

	// ASSERT
	assert.NoError(t, err)
}

func TestUniqueSlugFailure(t *testing.T) {
	// ARRANGE
	teardown, db := setupTestDB()
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
