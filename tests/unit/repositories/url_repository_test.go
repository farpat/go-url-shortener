package repositories

import (
	"database/sql"
	"os"
	"testing"

	"github.com/farpat/go-url-shortener/internal/config"
	"github.com/farpat/go-url-shortener/internal/models"
	urlRepository "github.com/farpat/go-url-shortener/internal/repositories"
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
	allUrls, _ := urlRepository.All()

	// ASSERT
	assert.Equal(t, 2, len(allUrls))
	assert.Equal(t, 2, getUrlsCount(db))

	expectedSlugs := map[string]bool{"abc": true, "def": true}
	for _, u := range allUrls {
		assert.True(t, expectedSlugs[u.Slug])
	}
}

func TestFind(t *testing.T) {
	// ARRANGE
	teardown, db := setupTestDB()
	defer teardown()

	// ACT
	url := models.UrlShowItem{Slug: "abc", Url: "https://example.com"}
	insertUrl(db, url)
	foundUrl, _ := urlRepository.Find(url.Slug)

	// ASSERT
	assert.Equal(t, url.Slug, foundUrl.Slug)
	assert.Equal(t, url.Url, foundUrl.Url)
	assert.Equal(t, 1, getUrlsCount(db))
}

func TestFindNotFound(t *testing.T) {
	// ARRANGE
	teardown, db := setupTestDB()
	defer teardown()
	var oldUrlCounts int
	db.QueryRow("SELECT COUNT(*) FROM urls").Scan(&oldUrlCounts)

	// ACT
	_, err := urlRepository.Find("not-found")

	// ASSERT
	expectedError := urlRepository.NotFoundError{Slug: "not-found"}
	assert.EqualError(t, err, expectedError.Error())
	var newUrlCounts int
	db.QueryRow("SELECT COUNT(*) FROM urls").Scan(&newUrlCounts)
	assert.Equal(t, oldUrlCounts, newUrlCounts)
}

func TestDelete(t *testing.T) {
	// ARRANGE
	teardown, db := setupTestDB()
	defer teardown()

	// ACT
	url := models.UrlShowItem{Slug: "abc", Url: "https://example.com"}
	insertUrl(db, url)
	err := urlRepository.Delete(url.Slug)

	// ASSERT
	assert.NoError(t, err)
	assert.Equal(t, 0, getUrlsCount(db))
}

func TestDeleteNotFound(t *testing.T) {
	// ARRANGE
	teardown, db := setupTestDB()
	defer teardown()

	// ACT
	url := models.UrlShowItem{Slug: "abc", Url: "https://example.com"}
	insertUrl(db, url)
	err := urlRepository.Delete("not-found")

	// ASSERT
	expectedError := urlRepository.NotFoundError{Slug: "not-found"}
	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, 1, getUrlsCount(db))
}

func getUrlsCount(db *sql.DB) int {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM urls").Scan(&count)
	return count
}
