package repositories

import (
	"database/sql"
	"os"
	"testing"

	"github.com/farpat/go-url-shortener/internal/config"
	"github.com/farpat/go-url-shortener/internal/models"
	"github.com/farpat/go-url-shortener/internal/repositories"
	"github.com/farpat/go-url-shortener/internal/utils"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() (teardown func(), db *sql.DB) {
	dbPath := "database_test.db"
	config.Databases["main"] = dbPath
	absoluteDbPath := utils.ProjectPath(dbPath)

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

func insertUrl(db *sql.DB, url models.UrlShowItem) {
	_, err := db.Exec("INSERT INTO urls (slug, url) VALUES (?, ?)", url.Slug, url.Url)
	if err != nil {
		panic(err)
	}
}

func TestAll(t *testing.T) {
	// ARRANGE
	teardown, db := setupTestDB()
	defer teardown()

	// ACT
	urls := []models.UrlShowItem{
		{Slug: "abc", Url: "https://example.com"},
		{Slug: "def", Url: "https://google.com"},
	}
	for _, url := range urls {
		insertUrl(db, url)
	}
	allUrls, _ := repositories.All()

	// ASSERT
	assert.Equal(t, 2, len(allUrls))

	expectedSlugs := map[string]bool{"abc": true, "def": true}
	for _, u := range allUrls {
		assert.True(t, expectedSlugs[u.Slug])
	}
}
